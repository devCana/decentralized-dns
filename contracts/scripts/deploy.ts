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

  // Contracts are added here as they are implemented (issues #2, #3, #8).
  const contractNames: string[] = [];

  for (const name of contractNames) {
    const factory = await ethers.getContractFactory(name);
    const instance = await factory.deploy();
    await instance.waitForDeployment();
    deployed[name] = await instance.getAddress();
    console.log(`${name} deployed at ${deployed[name]}`);
  }

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
