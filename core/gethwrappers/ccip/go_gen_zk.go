package ccip

// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/ARMProxy/ARMProxy.abi ../../../contracts/solc/v0.8.24/ARMProxy/ARMProxy.bin RMNProxyContract rmn_proxy_contract ../../../contracts/zksolc/v0.8.24/ARMProxy/ARMProxy.sol/ARMProxy.zbin
// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/RMN/RMN.abi ../../../contracts/solc/v0.8.24/RMN/RMN.bin RMNContract rmn_contract ../../../contracts/zksolc/v0.8.24/RMN/RMN.sol/RMN.zbin

// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/TokenAdminRegistry/TokenAdminRegistry.abi ../../../contracts/solc/v0.8.24/TokenAdminRegistry/TokenAdminRegistry.bin TokenAdminRegistry token_admin_registry ../../../contracts/zksolc/v0.8.24/TokenAdminRegistry/TokenAdminRegistry.sol/TokenAdminRegistry.zbin
// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/RegistryModuleOwnerCustom/RegistryModuleOwnerCustom.abi ../../../contracts/solc/v0.8.24/RegistryModuleOwnerCustom/RegistryModuleOwnerCustom.bin RegistryModuleOwnerCustom registry_module_owner_custom ../../../contracts/zksolc/v0.8.24/RegistryModuleOwnerCustom/RegistryModuleOwnerCustom.sol/RegistryModuleOwnerCustom.zbin
// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/FeeQuoter/FeeQuoter.abi ../../../contracts/solc/v0.8.24/FeeQuoter/FeeQuoter.bin FeeQuoter fee_quoter ../../../contracts/zksolc/v0.8.24/FeeQuoter/FeeQuoter.sol/FeeQuoter.zbin

// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/EVM2EVMOnRamp/EVM2EVMOnRamp.abi ../../../contracts/solc/v0.8.24/EVM2EVMOnRamp/EVM2EVMOnRamp.bin EVM2EVMOnRamp evm_2_evm_onramp ../../../contracts/zksolc/v0.8.24/EVM2EVMOnRamp/EVM2EVMOnRamp.sol/EVM2EVMOnRamp.zbin
// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/EVM2EVMOffRamp/EVM2EVMOffRamp.abi ../../../contracts/solc/v0.8.24/EVM2EVMOffRamp/EVM2EVMOffRamp.bin EVM2EVMOffRamp evm_2_evm_offramp ../../../contracts/zksolc/v0.8.24/EVM2EVMOffRamp/EVM2EVMOffRamp.sol/EVM2EVMOffRamp.zbin
// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/CommitStore/CommitStore.abi ../../../contracts/solc/v0.8.24/CommitStore/CommitStore.bin CommitStore commit_store ../../../contracts/zksolc/v0.8.24/CommitStore/CommitStore.sol/CommitStore.zbin

// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/Router/Router.abi ../../../contracts/solc/v0.8.24/Router/Router.bin Router router ../../../contracts/zksolc/v0.8.24/Router/Router.sol/Router.zbin
// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/PriceRegistry/PriceRegistry.abi ../../../contracts/solc/v0.8.24/PriceRegistry/PriceRegistry.bin PriceRegistry price_registry_1_2_0 ../../../contracts/zksolc/v0.8.24/PriceRegistry/PriceRegistry.sol/PriceRegistry.zbin

// go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/LockReleaseTokenPool/LockReleaseTokenPool.abi ../../../contracts/solc/v0.8.24/LockReleaseTokenPool/LockReleaseTokenPool.bin LockReleaseTokenPool lock_release_token_pool ../../../contracts/zksolc/v0.8.24/LockReleaseTokenPool/LockReleaseTokenPool.sol/LockReleaseTokenPool.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.19/LinkToken/LinkToken.abi ../../../contracts/solc/v0.8.19/LinkToken/LinkToken.bin LinkToken link_token ../../../contracts/zksolc/v0.8.19/LinkToken/LinkToken.sol/LinkToken.zbin
//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/WETH9/WETH9.abi ../../../contracts/solc/v0.8.24/WETH9/WETH9.bin WETH9 weth9 ../../../contracts/zksolc/v0.8.24/WETH9/WETH9.sol/WETH9.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/SelfFundedPingPong/SelfFundedPingPong.abi ../../../contracts/solc/v0.8.24/SelfFundedPingPong/SelfFundedPingPong.bin SelfFundedPingPong self_funded_ping_pong ../../../contracts/zksolc/v0.8.24/SelfFundedPingPong/SelfFundedPingPong.sol/SelfFundedPingPong.zbin

//go:generate go run ../generation/generate/wrap.go ../../../contracts/solc/v0.8.24/Storage/Storage.abi ../../../contracts/solc/v0.8.24/Storage/Storage.bin Storage storage ../../../contracts/zksolc/v0.8.24/Storage/Storage.sol/Storage.zbin
