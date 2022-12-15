// Package gethwrappers provides tools for wrapping solidity contracts with
// golang packages, using abigen.
package gethwrappers

// Make sure solidity compiler artifacts are up to date. Only output stdout on failure.
//go:generate ./generation/compile_contracts.sh

//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/FluxAggregator.abi ../../contracts/solc/v0.6/FluxAggregator.bin FluxAggregator flux_aggregator_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/VRFTestHelper.abi ../../contracts/solc/v0.6/VRFTestHelper.bin VRFTestHelper solidity_vrf_verifier_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/VRFCoordinator.abi ../../contracts/solc/v0.6/VRFCoordinator.bin VRFCoordinator solidity_vrf_coordinator_interface
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/VRFConsumer.abi ../../contracts/solc/v0.6/VRFConsumer.bin VRFConsumer solidity_vrf_consumer_interface
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/VRFRequestIDBaseTestHelper.abi ../../contracts/solc/v0.6/VRFRequestIDBaseTestHelper.bin VRFRequestIDBaseTestHelper solidity_vrf_request_id
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/Flags.abi ../../contracts/solc/v0.6/Flags.bin Flags flags_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/Oracle.abi ../../contracts/solc/v0.6/Oracle.bin Oracle oracle_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.6/BlockhashStore.abi ../../contracts/solc/v0.6/BlockhashStore.bin BlockhashStore blockhash_store
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/Consumer.abi ../../contracts/solc/v0.7/Consumer.bin Consumer consumer_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/MultiWordConsumer.abi ../../contracts/solc/v0.7/MultiWordConsumer.bin MultiWordConsumer multiwordconsumer_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/Operator.abi ../../contracts/solc/v0.7/Operator.bin Operator operator_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/OperatorFactory.abi ../../contracts/solc/v0.7/OperatorFactory.bin OperatorFactory operator_factory
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/AuthorizedForwarder.abi ../../contracts/solc/v0.7/AuthorizedForwarder.bin AuthorizedForwarder authorized_forwarder
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/AuthorizedReceiver.abi ../../contracts/solc/v0.7/AuthorizedReceiver.bin AuthorizedReceiver authorized_receiver
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/BatchBlockhashStore.abi ../../contracts/solc/v0.8.6/BatchBlockhashStore.bin BatchBlockhashStore batch_blockhash_store
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/BatchVRFCoordinatorV2.abi ../../contracts/solc/v0.8.6/BatchVRFCoordinatorV2.bin BatchVRFCoordinatorV2 batch_vrf_coordinator_v2
//go:generate go run ./generation/generate/wrap.go OffchainAggregator/OffchainAggregator.abi - OffchainAggregator offchain_aggregator_wrapper

//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/KeeperRegistry1_1.abi ../../contracts/solc/v0.7/KeeperRegistry1_1.bin KeeperRegistry keeper_registry_wrapper1_1
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/UpkeepPerformCounterRestrictive.abi ../../contracts/solc/v0.7/UpkeepPerformCounterRestrictive.bin UpkeepPerformCounterRestrictive upkeep_perform_counter_restrictive_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.7/UpkeepCounter.abi ../../contracts/solc/v0.7/UpkeepCounter.bin UpkeepCounter upkeep_counter_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/CronUpkeepFactory.abi - CronUpkeepFactory cron_upkeep_factory_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/CronUpkeep.abi - CronUpkeep cron_upkeep_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeeperRegistry1_2.abi ../../contracts/solc/v0.8.6/KeeperRegistry1_2.bin KeeperRegistry keeper_registry_wrapper1_2
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/TypeAndVersionInterface.abi ../../contracts/solc/v0.8.6/TypeAndVersionInterface.bin TypeAndVersionInterface type_and_version_interface_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeeperRegistryCheckUpkeepGasUsageWrapper1_2.abi ../../contracts/solc/v0.8.6/KeeperRegistryCheckUpkeepGasUsageWrapper1_2.bin KeeperRegistryCheckUpkeepGasUsageWrapper gas_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeeperRegistry1_3.abi ../../contracts/solc/v0.8.6/KeeperRegistry1_3.bin KeeperRegistry keeper_registry_wrapper1_3
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeeperRegistryLogic1_3.abi ../../contracts/solc/v0.8.6/KeeperRegistryLogic1_3.bin KeeperRegistryLogic keeper_registry_logic1_3
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeeperRegistry2_0.abi ../../contracts/solc/v0.8.6/KeeperRegistry2_0.bin KeeperRegistry keeper_registry_wrapper2_0
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeeperRegistryLogic2_0.abi ../../contracts/solc/v0.8.6/KeeperRegistryLogic2_0.bin KeeperRegistryLogic keeper_registry_logic2_0

// v0.8.6 VRFConsumer
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFConsumer.abi ../../contracts/solc/v0.8.6/VRFConsumer.bin VRFConsumer solidity_vrf_consumer_interface_v08
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFRequestIDBaseTestHelper.abi ../../contracts/solc/v0.8.6/VRFRequestIDBaseTestHelper.bin VRFRequestIDBaseTestHelper solidity_vrf_request_id_v08
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFOwnerlessConsumerExample.abi ../../contracts/solc/v0.8.6/VRFOwnerlessConsumerExample.bin VRFOwnerlessConsumerExample vrf_ownerless_consumer_example
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFLoadTestOwnerlessConsumer.abi ../../contracts/solc/v0.8.6/VRFLoadTestOwnerlessConsumer.bin VRFLoadTestOwnerlessConsumer vrf_load_test_ownerless_consumer
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFLoadTestExternalSubOwner.abi ../../contracts/solc/v0.8.6/VRFLoadTestExternalSubOwner.bin VRFLoadTestExternalSubOwner vrf_load_test_external_sub_owner

//go:generate go run ./generation/generate_link/wrap_link.go

// VRF V2
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFCoordinatorV2.abi ../../contracts/solc/v0.8.6/VRFCoordinatorV2.bin VRFCoordinatorV2 vrf_coordinator_v2
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFConsumerV2.abi ../../contracts/solc/v0.8.6/VRFConsumerV2.bin VRFConsumerV2 vrf_consumer_v2
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFMaliciousConsumerV2.abi ../../contracts/solc/v0.8.6/VRFMaliciousConsumerV2.bin VRFMaliciousConsumerV2 vrf_malicious_consumer_v2
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFTestHelper.abi ../../contracts/solc/v0.8.6/VRFTestHelper.bin VRFV08TestHelper solidity_vrf_v08_verifier_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFSingleConsumerExample.abi ../../contracts/solc/v0.8.6/VRFSingleConsumerExample.bin VRFSingleConsumerExample vrf_single_consumer_example
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFExternalSubOwnerExample.abi ../../contracts/solc/v0.8.6/VRFExternalSubOwnerExample.bin VRFExternalSubOwnerExample vrf_external_sub_owner_example
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFV2RevertingExample.abi ../../contracts/solc/v0.8.6/VRFV2RevertingExample.bin VRFV2RevertingExample vrfv2_reverting_example
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFConsumerV2UpgradeableExample.abi ../../contracts/solc/v0.8.6/VRFConsumerV2UpgradeableExample.bin VRFConsumerV2UpgradeableExample vrf_consumer_v2_upgradeable_example
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFV2TransparentUpgradeableProxy.abi ../../contracts/solc/v0.8.6/VRFV2TransparentUpgradeableProxy.bin VRFV2TransparentUpgradeableProxy vrfv2_transparent_upgradeable_proxy
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFV2ProxyAdmin.abi ../../contracts/solc/v0.8.6/VRFV2ProxyAdmin.bin VRFV2ProxyAdmin vrfv2_proxy_admin
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/NoCancelVRFCoordinatorV2.abi ../../contracts/solc/v0.8.6/NoCancelVRFCoordinatorV2.bin NoCancelVRFCoordinatorV2 nocancel_vrf_coordinator_v2

// VRF V2 Wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFV2Wrapper.abi ../../contracts/solc/v0.8.6/VRFV2Wrapper.bin VRFV2Wrapper vrfv2_wrapper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFV2WrapperInterface.abi ../../contracts/solc/v0.8.6/VRFV2WrapperInterface.bin VRFV2WrapperInterface vrfv2_wrapper_interface
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/VRFV2WrapperConsumerExample.abi ../../contracts/solc/v0.8.6/VRFV2WrapperConsumerExample.bin VRFV2WrapperConsumerExample vrfv2_wrapper_consumer_example

// Keepers X VRF v2
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/KeepersVRFConsumer.abi ../../contracts/solc/v0.8.6/KeepersVRFConsumer.bin KeepersVRFConsumer keepers_vrf_consumer

// Aggregators
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/AggregatorV2V3Interface.abi ../../contracts/solc/v0.8.6/AggregatorV2V3Interface.bin AggregatorV2V3Interface aggregator_v2v3_interface
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/AggregatorV3Interface.abi ../../contracts/solc/v0.8.6/AggregatorV3Interface.bin AggregatorV3Interface aggregator_v3_interface
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/DerivedPriceFeed.abi ../../contracts/solc/v0.8.6/DerivedPriceFeed.bin DerivedPriceFeed derived_price_feed_wrapper

// Log tester
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/LogEmitter.abi ../../contracts/solc/v0.8.6/LogEmitter.bin LogEmitter log_emitter

// Direct Request OCR
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/OCR2DR.abi ../../contracts/solc/v0.8.6/OCR2DR.bin OCR2DR ocr2dr
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/OCR2DRClient.abi ../../contracts/solc/v0.8.6/OCR2DRClient.bin OCR2DRClient ocr2dr_client
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/OCR2DRClientExample.abi ../../contracts/solc/v0.8.6/OCR2DRClientExample.bin OCR2DRClientExample ocr2dr_client_example
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/OCR2DROracle.abi ../../contracts/solc/v0.8.6/OCR2DROracle.bin OCR2DROracle ocr2dr_oracle
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.6/OCR2DRRegistry.abi ../../contracts/solc/v0.8.6/OCR2DRRegistry.bin OCR2DRRegistry ocr2dr_registry

// CCIP
// Generic
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/CommitStore.abi ../../contracts/solc/v0.8.15/CommitStore.bin CommitStore commit_store
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/CommitStoreHelper.abi ../../contracts/solc/v0.8.15/CommitStoreHelper.bin CommitStoreHelper commit_store_helper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/BurnMintTokenPool.abi ../../contracts/solc/v0.8.15/BurnMintTokenPool.bin BurnMintTokenPool burn_mint_token_pool
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/NativeTokenPool.abi ../../contracts/solc/v0.8.15/NativeTokenPool.bin NativeTokenPool native_token_pool
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/CustomTokenPool.abi ../../contracts/solc/v0.8.15/CustomTokenPool.bin CustomTokenPool custom_token_pool
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/AFN.abi ../../contracts/solc/v0.8.15/AFN.bin AFNContract afn_contract
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/MockAFN.abi ../../contracts/solc/v0.8.15/MockAFN.bin MockAFNContract mock_afn_contract

//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/ReceiverDapp.abi ../../contracts/solc/v0.8.15/ReceiverDapp.bin ReceiverDapp receiver_dapp
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/SimpleMessageReceiver.abi ../../contracts/solc/v0.8.15/SimpleMessageReceiver.bin SimpleMessageReceiver simple_message_receiver
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/MaybeRevertMessageReceiver.abi ../../contracts/solc/v0.8.15/MaybeRevertMessageReceiver.bin MaybeRevertMessageReceiver maybe_revert_message_receiver

// Toll
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/EVM2AnyTollOnRampRouter.abi ../../contracts/solc/v0.8.15/EVM2AnyTollOnRampRouter.bin EVM2AnyTollOnRampRouter evm_2_any_toll_onramp_router
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/EVM2EVMTollOnRamp.abi ../../contracts/solc/v0.8.15/EVM2EVMTollOnRamp.bin EVM2EVMTollOnRamp evm_2_evm_toll_onramp
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/EVM2EVMTollOffRamp.abi ../../contracts/solc/v0.8.15/EVM2EVMTollOffRamp.bin EVM2EVMTollOffRamp evm_2_evm_toll_offramp
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/Any2EVMTollOffRampRouter.abi ../../contracts/solc/v0.8.15/Any2EVMTollOffRampRouter.bin Any2EVMTollOffRampRouter any_2_evm_toll_offramp_router
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/EVM2EVMTollOffRampHelper.abi ../../contracts/solc/v0.8.15/EVM2EVMTollOffRampHelper.bin EVM2EVMTollOffRampHelper any_2_evm_toll_offramp_helper
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/TollSenderDapp.abi ../../contracts/solc/v0.8.15/TollSenderDapp.bin TollSenderDapp toll_sender_dapp

// GE
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/EVM2EVMGEOnRamp.abi ../../contracts/solc/v0.8.15/EVM2EVMGEOnRamp.bin EVM2EVMGEOnRamp evm_2_evm_ge_onramp
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/EVM2EVMGEOffRamp.abi ../../contracts/solc/v0.8.15/EVM2EVMGEOffRamp.bin EVM2EVMGEOffRamp evm_2_evm_ge_offramp
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/GERouter.abi ../../contracts/solc/v0.8.15/GERouter.bin GERouter ge_router
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/GasFeeCache.abi ../../contracts/solc/v0.8.15/GasFeeCache.bin GasFeeCache gas_fee_cache

//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/GovernanceDapp.abi ../../contracts/solc/v0.8.15/GovernanceDapp.bin GovernanceDapp governance_dapp
//go:generate go run ./generation/generate/wrap.go ../../contracts/solc/v0.8.15/PingPongDemo.abi ../../contracts/solc/v0.8.15/PingPongDemo.bin PingPongDemo ping_pong_demo

// To run these commands, you must either install docker, or the correct version
// of abigen. The latter can be installed with these commands, at least on linux:
//
//   git clone https://github.com/ethereum/go-ethereum
//   cd go-ethereum/cmd/abigen
//   git checkout v<version-needed>
//   go install
//
// Here, <version-needed> is the version of go-ethereum specified in chainlink's
// go.mod. This will install abigen in "$GOPATH/bin", which you should add to
// your $PATH.
//
// To reduce explicit dependencies, and in case the system does not have the
// correct version of abigen installed , the above commands spin up docker
// containers. In my hands, total running time including compilation is about
// 13s. If you're modifying solidity code and testing against go code a lot, it
// might be worthwhile to generate the the wrappers using a static container
// with abigen and solc, which will complete much faster. E.g.
//
//   abigen -sol ../../contracts/src/v0.6/VRFAll.sol -pkg vrf -out solidity_interfaces.go
//
// where VRFAll.sol simply contains `import "contract_path";` instructions for
// all the contracts you wish to target. This runs in about 0.25 seconds in my
// hands.
//
// If you're on linux, you can copy the correct version of solc out of the
// appropriate docker container. At least, the following works on ubuntu:
//
//   $ docker run --name solc ethereum/solc:0.6.2
//   $ sudo docker cp solc:/usr/bin/solc /usr/bin
//   $ docker rm solc
//
// If you need to point abigen at your solc executable, you can specify the path
// with the abigen --solc <path-to-executable> option.
