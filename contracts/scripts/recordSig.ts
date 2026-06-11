import { ethers } from "ethers"

/**
 * Canonical record message mirrored byte-for-byte with the Go side
 * (resolver/internal/pki/owner.go RecordMessage):
 *
 *   "ddns-record-v1" u16(len)name u16(len)type u16(len)selector
 *   u32(ttl) u16(nFields) { u16(len)name u16(len)value }...
 *
 * Field pairs sorted by name byte order. The message is then signed with
 * EIP-191 personal-sign (ethers signMessage).
 */
export function recordMessage(
	name: string,
	recordType: string,
	selector: string,
	ttl: number,
	fieldNames: string[],
	fieldValues: string[],
): Uint8Array {
	const enc = new TextEncoder()
	const chunks: Uint8Array[] = [enc.encode("ddns-record-v1")]
	const u16 = (n: number) => new Uint8Array([(n >> 8) & 0xff, n & 0xff])
	const pushStr = (s: string) => {
		const b = enc.encode(s)
		if (b.length > 0xffff) throw new Error(`string too long: ${s.slice(0, 32)}…`)
		chunks.push(u16(b.length), b)
	}
	pushStr(name)
	pushStr(recordType)
	pushStr(selector)
	chunks.push(new Uint8Array([(ttl >>> 24) & 0xff, (ttl >>> 16) & 0xff, (ttl >>> 8) & 0xff, ttl & 0xff]))
	const pairs = fieldNames.map((k, i) => [k, fieldValues[i]] as const)
	pairs.sort((a, b) => (a[0] < b[0] ? -1 : a[0] > b[0] ? 1 : 0))
	chunks.push(u16(pairs.length))
	for (const [k, v] of pairs) {
		pushStr(k)
		pushStr(v)
	}
	const total = chunks.reduce((n, c) => n + c.length, 0)
	const out = new Uint8Array(total)
	let off = 0
	for (const c of chunks) {
		out.set(c, off)
		off += c.length
	}
	return out
}

/** Signs the canonical record message with an owner wallet (EIP-191). */
export async function signRecord(
	wallet: ethers.Wallet | ethers.HDNodeWallet,
	name: string,
	recordType: string,
	selector: string,
	ttl: number,
	fieldNames: string[],
	fieldValues: string[],
): Promise<string> {
	return wallet.signMessage(recordMessage(name, recordType, selector, ttl, fieldNames, fieldValues))
}
