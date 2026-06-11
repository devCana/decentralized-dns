.PHONY: all build test clean contracts-install contracts-build contracts-test \
        resolver-build resolver-test deploy-localhost seed-localhost chain \
        bindings zk-setup

all: build

build: contracts-build resolver-build

test: contracts-test resolver-test

contracts-install:
	cd contracts && npm install

contracts-build:
	cd contracts && npx hardhat compile

contracts-test:
	cd contracts && npx hardhat test

chain:
	cd contracts && npx hardhat node

deploy-localhost:
	cd contracts && npx hardhat run scripts/deploy.ts --network localhost

seed-localhost:
	cd contracts && npx hardhat run scripts/seed.ts --network localhost

bindings:
	scripts/gen-bindings.sh

# Dev-only Groth16 ceremony: regenerates resolver/internal/zk/artifacts,
# contracts/contracts/ZKVerifier.sol and the hardhat proof fixture.
zk-setup:
	cd resolver && go run ./cmd/zkgen

resolver-build:
	cd resolver && go build -o bin/ ./...

resolver-test:
	cd resolver && go vet ./... && go test ./...

clean:
	rm -rf contracts/artifacts contracts/cache contracts/typechain-types resolver/bin
