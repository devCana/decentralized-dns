import { ethers, network } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * Seeds a freshly deployed local chain with demo data used by resolver
 * integration tests and the demo script: registers "example" and writes an
 * A record plus two SVC records with different selectors.
 *
 * Reads contract addresses from deployments/<network>.json (run deploy.ts
 * first). Signs records with the second hardhat account (alice).
 */
async function main() {
  const file = path.join(__dirname, "..", "deployments", `${network.name}.json`);
  const { contracts } = JSON.parse(fs.readFileSync(file, "utf8"));
  const [, alice] = await ethers.getSigners();

  const dapp = await ethers.getContractAt("NamespaceDApp", contracts.NamespaceDApp);

  const price = await dapp.priceOf("example");
  // alice's uncompressed secp256k1 public key as the on-chain pubKey
  const pubKey = ethers.SigningKey.computePublicKey(
    "0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d", // hardhat account #1
    false
  );
  await (await dapp.connect(alice).register("example", pubKey, { value: price })).wait();
  console.log("registered 'example' for", alice.address);

  const sig = "0x" + "ab".repeat(65); // placeholder until the PKI pipeline (issue #7)
  const zero = ethers.ZeroHash;

  await (
    await dapp
      .connect(alice)
      .setRecord("example", "A", "", ["address"], ["93.184.216.34"], 3600, sig, zero)
  ).wait();
  await (
    await dapp
      .connect(alice)
      .setRecord(
        "example",
        "SVC",
        "port=25&service=SMTP&transport=TCP",
        ["target", "service", "transport", "port"],
        ["mail.example", "SMTP", "TCP", "25"],
        300,
        sig,
        zero
      )
  ).wait();
  await (
    await dapp
      .connect(alice)
      .setRecord(
        "example",
        "SVC",
        "port=443&service=HTTP&transport=QUIC",
        ["target", "service", "transport", "port"],
        ["web.example", "HTTP", "QUIC", "443"],
        300,
        sig,
        zero
      )
  ).wait();
  console.log("seeded A + 2 SVC records for 'example'");
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
