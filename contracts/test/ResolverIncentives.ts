import { loadFixture, time } from "@nomicfoundation/hardhat-toolbox/network-helpers"
import { expect } from "chai"
import { ethers } from "hardhat"

const DAY = 24n * 60n * 60n

describe("ResolverIncentives", () => {
	async function deployFixture() {
		const [client, resolver, stranger] = await ethers.getSigners()
		const inc = await (await ethers.getContractFactory("ResolverIncentives")).deploy()
		return { inc, client, resolver, stranger }
	}

	// Opens a channel and returns its id from the emitted event.
	async function open(inc: any, client: any, resolver: any, deposit: bigint, duration = DAY) {
		const tx = await inc.connect(client).openChannel(resolver.address, duration, { value: deposit })
		const rcpt = await tx.wait()
		const ev = rcpt!.logs
			.map((l: any) => {
				try {
					return inc.interface.parseLog(l)
				} catch {
					return null
				}
			})
			.find((p: any) => p && p.name === "ChannelOpened")
		return ev!.args.id as string
	}

	// Client signs a voucher authorizing `cumulative` wei on channel `id`.
	async function voucher(inc: any, client: any, id: string, cumulative: bigint) {
		const addr = await inc.getAddress()
		const inner = ethers.keccak256(
			ethers.AbiCoder.defaultAbiCoder().encode(["address", "bytes32", "uint256"], [addr, id, cumulative]),
		)
		return client.signMessage(ethers.getBytes(inner))
	}

	it("opens a channel and records it", async () => {
		const { inc, client, resolver } = await loadFixture(deployFixture)
		const id = await open(inc, client, resolver, ethers.parseEther("1"))
		const ch = await inc.channels(id)
		expect(ch.client).to.equal(client.address)
		expect(ch.resolver).to.equal(resolver.address)
		expect(ch.deposit).to.equal(ethers.parseEther("1"))
		expect(ch.claimed).to.equal(0n)
	})

	it("rejects a zero-value or zero-resolver open", async () => {
		const { inc, client, resolver } = await loadFixture(deployFixture)
		await expect(inc.connect(client).openChannel(resolver.address, DAY, { value: 0 })).to.be.revertedWithCustomError(inc, "NoDeposit")
		await expect(inc.connect(client).openChannel(ethers.ZeroAddress, DAY, { value: 1n })).to.be.revertedWithCustomError(inc, "ZeroResolver")
	})

	it("lets the resolver claim a signed voucher and pays the delta", async () => {
		const { inc, client, resolver } = await loadFixture(deployFixture)
		const id = await open(inc, client, resolver, ethers.parseEther("1"))

		const v1 = await voucher(inc, client, id, ethers.parseEther("0.3"))
		await expect(inc.connect(resolver).claim(id, ethers.parseEther("0.3"), v1)).to.changeEtherBalance(
			resolver,
			ethers.parseEther("0.3"),
		)
		expect((await inc.channels(id)).claimed).to.equal(ethers.parseEther("0.3"))

		// A second, higher voucher pays only the increment.
		const v2 = await voucher(inc, client, id, ethers.parseEther("0.5"))
		await expect(inc.connect(resolver).claim(id, ethers.parseEther("0.5"), v2)).to.changeEtherBalance(
			resolver,
			ethers.parseEther("0.2"),
		)
	})

	it("treats a replayed (non-increasing) voucher as nothing to claim", async () => {
		const { inc, client, resolver } = await loadFixture(deployFixture)
		const id = await open(inc, client, resolver, ethers.parseEther("1"))
		const v = await voucher(inc, client, id, ethers.parseEther("0.4"))
		await inc.connect(resolver).claim(id, ethers.parseEther("0.4"), v)
		await expect(inc.connect(resolver).claim(id, ethers.parseEther("0.4"), v)).to.be.revertedWithCustomError(inc, "NothingToClaim")
	})

	it("rejects a claim by a non-resolver and a forged voucher", async () => {
		const { inc, client, resolver, stranger } = await loadFixture(deployFixture)
		const id = await open(inc, client, resolver, ethers.parseEther("1"))
		const good = await voucher(inc, client, id, ethers.parseEther("0.2"))
		await expect(inc.connect(stranger).claim(id, ethers.parseEther("0.2"), good)).to.be.revertedWithCustomError(inc, "NotResolver")

		// Voucher signed by someone other than the channel client.
		const forged = await voucher(inc, stranger, id, ethers.parseEther("0.2"))
		await expect(inc.connect(resolver).claim(id, ethers.parseEther("0.2"), forged)).to.be.revertedWithCustomError(inc, "BadVoucher")
	})

	it("caps a claim at the deposit", async () => {
		const { inc, client, resolver } = await loadFixture(deployFixture)
		const id = await open(inc, client, resolver, ethers.parseEther("1"))
		const v = await voucher(inc, client, id, ethers.parseEther("5")) // over-authorize
		await expect(inc.connect(resolver).claim(id, ethers.parseEther("5"), v)).to.changeEtherBalance(resolver, ethers.parseEther("1"))
	})

	it("refunds the remainder to the client only after expiry", async () => {
		const { inc, client, resolver } = await loadFixture(deployFixture)
		const id = await open(inc, client, resolver, ethers.parseEther("1"))
		const v = await voucher(inc, client, id, ethers.parseEther("0.6"))
		await inc.connect(resolver).claim(id, ethers.parseEther("0.6"), v)

		await expect(inc.connect(client).closeChannel(id)).to.be.revertedWithCustomError(inc, "NotExpired")
		// Past expiry but within the settlement window: the resolver's grace
		// period, so the client still cannot close.
		await time.increase(DAY + 1n)
		await expect(inc.connect(client).closeChannel(id)).to.be.revertedWithCustomError(inc, "NotExpired")
		// After the settlement window: client reclaims the unspent remainder.
		await time.increase(3601n)
		await expect(inc.connect(client).closeChannel(id)).to.changeEtherBalance(client, ethers.parseEther("0.4"))
		expect((await inc.channels(id)).client).to.equal(ethers.ZeroAddress) // deleted
	})
})
