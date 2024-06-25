
SHELL := /opt/homebrew/bin/fish

# IMPORTANT: Use solc 0.8.18 and abigen 1.14.5

.PHONY: build
build:
	rm -rf build/
	solc --abi --bin mycontract.sol -o build
	abigen --abi build/mycontract_sol_SimpleContract.abi --bin build/mycontract_sol_SimpleContract.bin --pkg=main --out=mycontract.go

.PHONY: test
test: build
	go test -v --tags "playground" ./...
