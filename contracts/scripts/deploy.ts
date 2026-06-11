import { artifacts, ethers, network } from "hardhat";
import * as fs from "fs";
import * as path from "path";

/**
 * Deploys all Decentralized DNS contracts and writes their addresses to
 * deployments/<network>.json so the resolver and CLIs can pick them up.
 *
 * Contract list grows as the project evolves; deployment is idempotent per
 * run (fresh addresses every time on a dev chain).
 */
async function main() {
  const [deployer] = await ethers.getSigners();
  console.log(`Deploying to '${network.name}' as ${deployer.address}`);

  const deployed: Record<string, string> = {};

  const registry = await (
    await ethers.getContractFactory("RecordSchemaRegistry")
  ).deploy();
  await registry.waitForDeployment();
  deployed.RecordSchemaRegistry = await registry.getAddress();
  console.log(`RecordSchemaRegistry deployed at ${deployed.RecordSchemaRegistry}`);

  // 0.01 ether yearly base fee (HLD pricing decision: length-based multiplier).
  const basePrice = ethers.parseEther(process.env.BASE_PRICE_ETH || "0.01");
  const namespace = await (
    await ethers.getContractFactory("NamespaceDApp")
  ).deploy(basePrice, deployed.RecordSchemaRegistry);
  await namespace.waitForDeployment();
  deployed.NamespaceDApp = await namespace.getAddress();
  console.log(`NamespaceDApp deployed at ${deployed.NamespaceDApp}`);

  // gnark-exported Groth16 verifier for record-commitment proofs.
  const verifier = await (await ethers.getContractFactory("Verifier")).deploy();
  await verifier.waitForDeployment();
  deployed.ZKVerifier = await verifier.getAddress();
  console.log(`ZKVerifier deployed at ${deployed.ZKVerifier}`);

  const outDir = path.join(__dirname, "..", "deployments");
  fs.mkdirSync(outDir, { recursive: true });
  const outFile = path.join(outDir, `${network.name}.json`);
  fs.writeFileSync(
    outFile,
    JSON.stringify(
      {
        network: network.name,
        chainId: (await ethers.provider.getNetwork()).chainId.toString(),
        deployer: deployer.address,
        contracts: deployed,
        timestamp: new Date().toISOString(),
      },
      null,
      2
    )
  );
  console.log(`Wrote ${outFile}`);

  // Keep typechain/artifacts honest in CI.
  void artifacts;
}

main().catch((err) => {
  console.error(err);
  process.exitCode = 1;
});
