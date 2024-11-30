package ccip

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/Storage/Storage.abi ../../../contracts/solc/v0.8.24/Storage/Storage.bin Storage storage ../../../contracts/zksolc/v0.8.24/Storage/Storage.sol/Storage.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/EVMCustom/EVMCustomTest.abi ../../../contracts/solc/v0.8.24/EVMCustom/EVMCustomTest.bin EvmcustomContract evmcustom_contract ../../../contracts/zksolc/v0.8.24/EVMCustom/EVMCustom.sol/EVMCustomTest.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.5.0/opCodes/OpCodes.abi ../../../contracts/solc/v0.5.0/opCodes/OpCodes.bin OpcodesContract opcodes_contract ../../../contracts/zksolc/v0.5.0/opCodes/opCodes.sol/OpCodes.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.19/manyChainMultisig/ManyChainMultiSig.abi ../../../contracts/solc/v0.8.19/manyChainMultisig/ManyChainMultiSig.bin ManychainmultisigContract manychainmultisig_contract ../../../contracts/zksolc/v0.8.19/manyChainMultisig/ManyChainMultiSig.sol/ManyChainMultiSig.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.19/RBACTimelock/RBACTimelock.abi ../../../contracts/solc/v0.8.19/RBACTimelock/RBACTimelock.bin RbactimelockContract rbactimelock_contract ../../../contracts/zksolc/v0.8.19/RBACTimelock/RBACTimelock.sol/RBACTimelock.zbin
