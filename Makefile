.PHONY: all build test clean contracts-install contracts-build contracts-test \
        resolver-build resolver-test deploy-localhost chain

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

resolver-build:
	cd resolver && go build -o bin/ ./...

resolver-test:
	cd resolver && go vet ./... && go test ./...

clean:
	rm -rf contracts/artifacts contracts/cache contracts/typechain-types resolver/bin
