import { loadFixture } from "@nomicfoundation/hardhat-toolbox/network-helpers"
import { expect } from "chai"
import { ethers } from "hardhat"

const PUBKEY_A = "0x" + "aa".repeat(32) // dummy 32-byte ed25519 key
const PUBKEY_B = "0x" + "bb".repeat(32)
const ZERO32 = "0x" + "00".repeat(32)

describe("ResolverRegistry", () => {
	async function deployFixture() {
		const [deployer, alice, bob] = await ethers.getSigners()
		const reg = await (await ethers.getContractFactory("ResolverRegistry")).deploy()
		return { reg, deployer, alice, bob }
	}

	describe("announce", () => {
		it("registers a resolver and emits the event", async () => {
			const { reg, alice } = await loadFixture(deployFixture)
			await expect(reg.connect(alice).announce(PUBKEY_A, "https://r.alice:8080"))
				.to.emit(reg, "ResolverAnnounced")
				.withArgs(alice.address, PUBKEY_A, "https://r.alice:8080")

			const r = await reg.getResolver(alice.address)
			expect(r.pubKey).to.equal(PUBKEY_A)
			expect(r.endpoint).to.equal("https://r.alice:8080")
			expect(r.active).to.equal(true)
			expect(await reg.operatorCount()).to.equal(1)
			expect(await reg.operatorAt(0)).to.equal(alice.address)
		})

		it("updates in place without duplicating the enumeration entry", async () => {
			const { reg, alice } = await loadFixture(deployFixture)
			await reg.connect(alice).announce(PUBKEY_A, "https://old:8080")
			await reg.connect(alice).announce(PUBKEY_B, "https://new:8080")
			expect(await reg.operatorCount()).to.equal(1)
			const r = await reg.getResolver(alice.address)
			expect(r.pubKey).to.equal(PUBKEY_B)
			expect(r.endpoint).to.equal("https://new:8080")
		})

		it("rejects an empty pubkey, empty endpoint, or over-long endpoint", async () => {
			const { reg, alice } = await loadFixture(deployFixture)
			await expect(reg.connect(alice).announce(ZERO32, "https://x")).to.be.revertedWithCustomError(reg, "EmptyPubKey")
			await expect(reg.connect(alice).announce(PUBKEY_A, "")).to.be.revertedWithCustomError(reg, "EmptyEndpoint")
			await expect(reg.connect(alice).announce(PUBKEY_A, "x".repeat(257))).to.be.revertedWithCustomError(reg, "EndpointTooLong")
		})
	})

	describe("revoke", () => {
		it("marks a resolver inactive and drops it from the active set", async () => {
			const { reg, alice } = await loadFixture(deployFixture)
			await reg.connect(alice).announce(PUBKEY_A, "https://r.alice:8080")
			await expect(reg.connect(alice).revoke()).to.emit(reg, "ResolverRevoked").withArgs(alice.address)
			expect((await reg.getResolver(alice.address)).active).to.equal(false)
			const [ops] = await reg.activeResolvers()
			expect(ops.length).to.equal(0)
		})

		it("reverts when the caller has no active resolver", async () => {
			const { reg, bob } = await loadFixture(deployFixture)
			await expect(reg.connect(bob).revoke()).to.be.revertedWithCustomError(reg, "NotRegistered")
		})

		it("can be re-announced after revoke (no double enumeration)", async () => {
			const { reg, alice } = await loadFixture(deployFixture)
			await reg.connect(alice).announce(PUBKEY_A, "https://r.alice:8080")
			await reg.connect(alice).revoke()
			await reg.connect(alice).announce(PUBKEY_A, "https://r.alice:8080")
			expect(await reg.operatorCount()).to.equal(1)
			expect((await reg.getResolver(alice.address)).active).to.equal(true)
		})
	})

	describe("activeResolvers (discovery)", () => {
		it("returns only active resolvers with their keys and endpoints", async () => {
			const { reg, alice, bob } = await loadFixture(deployFixture)
			await reg.connect(alice).announce(PUBKEY_A, "https://r.alice:8080")
			await reg.connect(bob).announce(PUBKEY_B, "https://r.bob:8080")
			await reg.connect(alice).revoke()

			const [ops, keys, endpoints] = await reg.activeResolvers()
			expect(ops).to.deep.equal([bob.address])
			expect(keys).to.deep.equal([PUBKEY_B])
			expect(endpoints).to.deep.equal(["https://r.bob:8080"])
		})
	})
})
