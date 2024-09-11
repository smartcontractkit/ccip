pragma solidity ^0.8.24;

import {IOwner} from "../../interfaces/IOwner.sol";

import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../../pools/TokenPool.sol";

import {TokenAdminRegistry} from "../../tokenAdminRegistry/TokenAdminRegistry.sol";
import {TokenPoolFactory} from "../../tokenAdminRegistry/TokenPoolFactory.sol";

import {RegistryModuleOwnerCustom} from "../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";
import {TokenAdminRegistrySetup} from "./TokenAdminRegistry.t.sol";

import {RateLimiter} from "../../libraries/RateLimiter.sol";

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {BurnMintERC20} from "../../../shared/token/ERC20/BurnMintERC20.sol";
import {DeterministicContractDeployer} from "../../../shared/util/DeterministicContractDeployer.sol";

import {console2 as console} from "forge-std/console2.sol";

contract TokenPoolFactorySetup is TokenAdminRegistrySetup {
  TokenPoolFactory internal s_tokenPoolFactory;
  RegistryModuleOwnerCustom internal s_registryModuleOwnerCustom;

  bytes internal s_poolInitCode;
  bytes internal s_poolInitArgs;

  bytes32 internal constant s_salt = keccak256(abi.encode("FAKE_SALT"));

  address internal s_rmnProxy = address(0x1234);

  bytes internal s_tokenCreationParams;
  bytes internal s_tokenInitCode;

  uint256 public constant PREMINT_AMOUNT = 1e20; // 100 tokens in 18 decimals

  function setUp() public virtual override {
    TokenAdminRegistrySetup.setUp();

    s_registryModuleOwnerCustom = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleOwnerCustom));

    s_tokenPoolFactory = new TokenPoolFactory(
      address(s_tokenAdminRegistry), address(s_registryModuleOwnerCustom), s_rmnProxy, address(s_sourceRouter)
    );

    // Create Init Code for BurnMintERC20 TestToken with 18 decimals and supply cap of max uint256 value
    s_tokenCreationParams = abi.encode("TestToken", "TT", 18, type(uint256).max, PREMINT_AMOUNT, OWNER);

    s_tokenInitCode = abi.encodePacked(type(BurnMintERC20).creationCode, s_tokenCreationParams);

    s_poolInitCode = type(BurnMintTokenPool).creationCode;

    // Create Init Args for BurnMintTokenPool with no allowlist minus the token address
    address[] memory allowlist = new address[](1);
    allowlist[0] = OWNER;
    s_poolInitArgs = abi.encode(allowlist, address(0x1234), s_sourceRouter);
  }
}

contract TokenPoolFactoryTests is TokenPoolFactorySetup {
  function test_createTokenPool_WithNoExistingTokenOnRemoteChain_Success() public {
    vm.startPrank(OWNER);

    bytes32 dynamicSalt = keccak256(abi.encodePacked(s_salt, OWNER));

    address predictedTokenAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      s_tokenInitCode, dynamicSalt, address(s_tokenPoolFactory)
    );

    // Create the constructor params for the predicted pool
    bytes memory poolCreationParams = abi.encode(predictedTokenAddress, new address[](0), s_rmnProxy, s_sourceRouter);

    // Predict the address of the pool before we make the tx by using the init code and the params
    bytes memory predictedPoolInitCode = abi.encodePacked(s_poolInitCode, poolCreationParams);
    address predictedPoolAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      predictedPoolInitCode, dynamicSalt, address(s_tokenPoolFactory)
    );

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      new TokenPoolFactory.ExistingTokenPool[](0), s_tokenInitCode, s_poolInitCode, poolCreationParams, s_salt
    );

    assertNotEq(address(0), tokenAddress, "Token Address should not be 0");
    assertNotEq(address(0), poolAddress, "Pool Address should not be 0");

    assertEq(predictedTokenAddress, tokenAddress, "Token Address should have been predicted");
    assertEq(predictedPoolAddress, poolAddress, "Pool Address should have been predicted");

    s_tokenAdminRegistry.acceptAdminRole(tokenAddress);
    OwnerIsCreator(tokenAddress).acceptOwnership();
    OwnerIsCreator(poolAddress).acceptOwnership();

    assertEq(poolAddress, s_tokenAdminRegistry.getPool(tokenAddress), "Token Pool should be set");
    assertEq(IOwner(tokenAddress).owner(), OWNER, "Token should be owned by the owner");
    assertEq(IOwner(poolAddress).owner(), OWNER, "Token should be owned by the owner");
  }

  function test_createTokenPool_WithNoExistingRemoteContracts_predict_Success() public {
    vm.startPrank(OWNER);
    bytes32 dynamicSalt = keccak256(abi.encodePacked(s_salt, OWNER));

    // We have to create a new factory, registry module, and token admin registry to simulate the other chain
    TokenAdminRegistry newTokenAdminRegistry = new TokenAdminRegistry();
    RegistryModuleOwnerCustom newRegistryModule = new RegistryModuleOwnerCustom(address(newTokenAdminRegistry));

    // We want to deploy a new factory and Owner Module.
    TokenPoolFactory newTokenPoolFactory = new TokenPoolFactory(
      address(newTokenAdminRegistry), address(newRegistryModule), s_rmnProxy, address(s_destRouter)
    );

    newTokenAdminRegistry.addRegistryModule(address(newRegistryModule));

    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig =
      TokenPoolFactory.RemoteChainConfig(address(newTokenPoolFactory), address(s_destRouter), address(s_rmnProxy));

    // Add the new token Factory to the remote chain config and set it for the simulated destination chain
    s_tokenPoolFactory.updateRemoteChainConfig(DEST_CHAIN_SELECTOR, remoteChainConfig);

    // Create an array of remote pools where nothing exists yet, but we want to predict the address for
    // the new pool and token on DEST_CHAIN_SELECTOR
    TokenPoolFactory.ExistingTokenPool[] memory remoteTokenPools = new TokenPoolFactory.ExistingTokenPool[](1);

    // The only field that matters is DEST_CHAIN_SELECTOR because we dont want any existing token pool or token
    // on the remote chain
    remoteTokenPools[0] = TokenPoolFactory.ExistingTokenPool(
      DEST_CHAIN_SELECTOR,
      abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()),
      abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()),
      s_tokenInitCode,
      RateLimiter.Config(false, 0, 0),
      RateLimiter.Config(false, 0, 0)
    );

    // Predict the address of the token and pool on the DESTINATION chain
    address predictedTokenAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      s_tokenInitCode, dynamicSalt, address(newTokenPoolFactory)
    );

    // Since the remote chain information was provided, we should be able to get the information from the newly
    // deployed token pool using the available getter functions

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools, s_tokenInitCode, s_poolInitCode, abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()), s_salt
    );

    // Ensure that the remote Token was set to the one we predicted
    assertEq(
      abi.encode(predictedTokenAddress),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Token Address should have been predicted"
    );

    // Create the constructor params for the predicted pool
    // The predictedTokenAddress is NOT abi-encoded since the raw evm-address
    // is used in the constructor params
    bytes memory predictedPoolCreationParams =
      abi.encode(predictedTokenAddress, new address[](0), s_rmnProxy, address(s_destRouter));

    // Take the init code and concat the destination params to it, the initCode shouldn't change
    bytes memory predictedPoolInitCode = abi.encodePacked(s_poolInitCode, predictedPoolCreationParams);

    // Predict the address of the pool on the DESTINATION chain
    address predictedPoolAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      predictedPoolInitCode, dynamicSalt, address(newTokenPoolFactory)
    );

    // Assert that the address set for the remote pool is the same as the predicted address
    assertEq(
      abi.encode(predictedPoolAddress),
      TokenPool(poolAddress).getRemotePool(DEST_CHAIN_SELECTOR),
      "Pool Address should have been predicted"
    );

    // On the new token pool factory, representing a destination chain,
    // deploy a new token and a new pool
    (address newTokenAddress, address newPoolAddress) = newTokenPoolFactory.deployTokenAndTokenPool(
      new TokenPoolFactory.ExistingTokenPool[](0),
      s_tokenInitCode,
      s_poolInitCode,
      abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()),
      s_salt
    );

    assertEq(
      TokenPool(poolAddress).getRemotePool(DEST_CHAIN_SELECTOR),
      abi.encode(newPoolAddress),
      "New Pool Address should have been deployed correctly"
    );

    assertEq(
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      abi.encode(newTokenAddress),
      "New Token Address should have been deployed correctly"
    );
  }

  function test_createTokenPool_ExistingRemoteToken_AndPredictPool_Success() public {
    vm.startPrank(OWNER);
    bytes32 dynamicSalt = keccak256(abi.encodePacked(s_salt, OWNER));

    BurnMintERC20 newRemoteToken = new BurnMintERC20("TestToken", "TT", 18, type(uint256).max, PREMINT_AMOUNT, OWNER);

    // We have to create a new factory, registry module, and token admin registry to simulate the other chain

    TokenAdminRegistry newTokenAdminRegistry = new TokenAdminRegistry();
    RegistryModuleOwnerCustom newRegistryModule = new RegistryModuleOwnerCustom(address(newTokenAdminRegistry));

    // We want to deploy a new factory and Owner Module.
    TokenPoolFactory newTokenPoolFactory = new TokenPoolFactory(
      address(newTokenAdminRegistry), address(newRegistryModule), s_rmnProxy, address(s_destRouter)
    );

    newTokenAdminRegistry.addRegistryModule(address(newRegistryModule));

    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig =
      TokenPoolFactory.RemoteChainConfig(address(newTokenPoolFactory), address(s_destRouter), address(s_rmnProxy));

    // Add the new token Factory to the remote chain config and set it for the simulated destination chain
    s_tokenPoolFactory.updateRemoteChainConfig(DEST_CHAIN_SELECTOR, remoteChainConfig);

    // Create an array of remote pools where nothing exists yet, but we want to predict the address for
    // the new pool and token on DEST_CHAIN_SELECTOR
    TokenPoolFactory.ExistingTokenPool[] memory remoteTokenPools = new TokenPoolFactory.ExistingTokenPool[](1);

    // The only field that matters is DEST_CHAIN_SELECTOR because we dont want any existing token pool or token
    // on the remote chain
    remoteTokenPools[0] = TokenPoolFactory.ExistingTokenPool(
      DEST_CHAIN_SELECTOR,
      abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()),
      abi.encode(address(newRemoteToken)),
      s_tokenInitCode,
      RateLimiter.Config(false, 0, 0),
      RateLimiter.Config(false, 0, 0)
    );

    // Since the remote chain information was provided, we should be able to get the information from the newly
    // deployed token pool using the available getter functions
    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools, s_tokenInitCode, s_poolInitCode, abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()), s_salt
    );

    // Ensure that the remote Token was set to the one we predicted
    assertEq(
      abi.encode(address(newRemoteToken)),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Token Address should have been predicted"
    );

    // Create the constructor params for the predicted pool
    // The predictedTokenAddress is NOT abi-encoded since the raw evm-address
    // is used in the constructor params
    bytes memory predictedPoolCreationParams =
      abi.encode(address(newRemoteToken), new address[](0), s_rmnProxy, address(s_destRouter));

    // Take the init code and concat the destination params to it, the initCode shouldn't change
    bytes memory predictedPoolInitCode = abi.encodePacked(s_poolInitCode, predictedPoolCreationParams);

    // Predict the address of the pool on the DESTINATION chain
    address predictedPoolAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      predictedPoolInitCode, dynamicSalt, address(newTokenPoolFactory)
    );

    // Assert that the address set for the remote pool is the same as the predicted address
    assertEq(
      abi.encode(predictedPoolAddress),
      TokenPool(poolAddress).getRemotePool(DEST_CHAIN_SELECTOR),
      "Pool Address should have been predicted"
    );

    // On the new token pool factory, representing a destination chain,
    // deploy a new token and a new pool
    (address newTokenAddress, address newPoolAddress) = newTokenPoolFactory.deployTokenPoolWithExistingToken(
      address(newRemoteToken),
      new TokenPoolFactory.ExistingTokenPool[](0),
      s_tokenInitCode,
      s_poolInitCode,
      abi.encode(s_tokenPoolFactory.EMPTY_PARAMETER_FLAG()),
      s_salt
    );

    assertEq(
      abi.encode(newTokenAddress),
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      "Remote Token Address should have been set correctly"
    );

    assertEq(
      newTokenAddress,
      address(newRemoteToken),
      "Remote Token Address returned should be the same as the one we deployed"
    );

    assertEq(
      TokenPool(poolAddress).getRemotePool(DEST_CHAIN_SELECTOR),
      abi.encode(newPoolAddress),
      "New Pool Address should have been deployed correctly"
    );
  }

  function test_createTokenPool_WithRemoteTokenAndRemotePool_Success() public {
    vm.startPrank(OWNER);

    bytes32 dynamicSalt = keccak256(abi.encodePacked(s_salt, OWNER));

    bytes memory RANDOM_TOKEN_ADDRESS = abi.encode(makeAddr("RANDOM_TOKEN"));
    bytes memory RANDOM_POOL_ADDRESS = abi.encode(makeAddr("RANDOM_POOL"));

    address predictedTokenAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      s_tokenInitCode, dynamicSalt, address(s_tokenPoolFactory)
    );

    // Create the constructor params for the predicted pool
    bytes memory poolCreationParams = abi.encode(predictedTokenAddress, new address[](0), s_rmnProxy, s_sourceRouter);

    // Predict the address of the pool before we make the tx by using the init code and the params
    bytes memory predictedPoolInitCode = abi.encodePacked(s_poolInitCode, poolCreationParams);
    address predictedPoolAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      predictedPoolInitCode, dynamicSalt, address(s_tokenPoolFactory)
    );

    // Create an array of remote pools with some fake addresses
    TokenPoolFactory.ExistingTokenPool[] memory remoteTokenPools = new TokenPoolFactory.ExistingTokenPool[](1);

    remoteTokenPools[0] = TokenPoolFactory.ExistingTokenPool(
      DEST_CHAIN_SELECTOR,
      RANDOM_POOL_ADDRESS,
      RANDOM_TOKEN_ADDRESS,
      "",
      RateLimiter.Config(false, 0, 0),
      RateLimiter.Config(false, 0, 0)
    );

    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.deployTokenAndTokenPool(
      remoteTokenPools, s_tokenInitCode, s_poolInitCode, poolCreationParams, s_salt
    );

    assertNotEq(address(0), tokenAddress, "Token Address should not be 0");
    assertNotEq(address(0), poolAddress, "Pool Address should not be 0");

    s_tokenAdminRegistry.acceptAdminRole(tokenAddress);
    OwnerIsCreator(tokenAddress).acceptOwnership();
    OwnerIsCreator(poolAddress).acceptOwnership();

    assertEq(
      TokenPool(poolAddress).getRemoteToken(DEST_CHAIN_SELECTOR),
      RANDOM_TOKEN_ADDRESS,
      "Remote Token Address should have been set"
    );

    assertEq(
      TokenPool(poolAddress).getRemotePool(DEST_CHAIN_SELECTOR),
      RANDOM_POOL_ADDRESS,
      "Remote Pool Address should have been set"
    );

    assertEq(poolAddress, s_tokenAdminRegistry.getPool(tokenAddress), "Token Pool should be set");

    assertEq(IOwner(tokenAddress).owner(), OWNER, "Token should be owned by the owner");

    assertEq(IOwner(poolAddress).owner(), OWNER, "Token should be owned by the owner");
  }

  function test_updateRemoteChainConfig_Success() public {
    TokenPoolFactory.RemoteChainConfig memory remoteChainConfig = TokenPoolFactory.RemoteChainConfig({
      remotePoolFactory: address(0x1234),
      remoteRouter: address(0x5678),
      remoteRMNProxy: address(0x9abc)
    });

    s_tokenPoolFactory.updateRemoteChainConfig(DEST_CHAIN_SELECTOR, remoteChainConfig);

    TokenPoolFactory.RemoteChainConfig memory updatedRemoteChainConfig =
      s_tokenPoolFactory.getRemoteChainConfig(DEST_CHAIN_SELECTOR);

    assertEq(
      remoteChainConfig.remotePoolFactory,
      updatedRemoteChainConfig.remotePoolFactory,
      "Token Pool Factory should be set"
    );

    assertEq(remoteChainConfig.remoteRouter, updatedRemoteChainConfig.remoteRouter, "Router should be set");

    assertEq(remoteChainConfig.remoteRMNProxy, updatedRemoteChainConfig.remoteRMNProxy, "RMN Proxy should be set");
  }
}
