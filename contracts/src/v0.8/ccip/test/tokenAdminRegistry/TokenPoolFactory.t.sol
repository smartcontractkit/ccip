pragma solidity ^0.8.24;

import {IOwner} from "../../interfaces/IOwner.sol";

import {BurnMintTokenPool} from "../../pools/BurnMintTokenPool.sol";
import {TokenPoolFactory} from "../../tokenAdminRegistry/TokenPoolFactory.sol";

import {TokenAdminRegistrySetup} from "./TokenAdminRegistry.t.sol";
import {RegistryModuleOwnerCustom} from "../../tokenAdminRegistry/RegistryModuleOwnerCustom.sol";


import {DeterministicContractDeployer} from "../../../shared/util/DeterministicContractDeployer.sol";
import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {BurnMintERC20} from "../../../shared/token/ERC20/BurnMintERC20.sol";

contract TokenPoolFactorySetup is TokenAdminRegistrySetup {
  TokenPoolFactory internal s_tokenPoolFactory;
  RegistryModuleOwnerCustom internal s_registryModuleOwnerCustom;

  bytes internal s_poolInitCode;
  bytes internal s_poolInitArgs;

  bytes32 internal constant s_salt = keccak256(abi.encode("FAKE_SALT"));

  address internal s_rmnProxy = address(0x1234);

  bytes internal s_tokenCreationParams;
  bytes internal s_tokenInitCode;

  function setUp() public virtual override {
    TokenAdminRegistrySetup.setUp();

    s_registryModuleOwnerCustom = new RegistryModuleOwnerCustom(address(s_tokenAdminRegistry));
    s_tokenAdminRegistry.addRegistryModule(address(s_registryModuleOwnerCustom));

    s_tokenPoolFactory = new TokenPoolFactory(address(s_tokenAdminRegistry), address(s_registryModuleOwnerCustom), s_rmnProxy, address(s_sourceRouter));

    // Create Init Code for BurnMintERC20 TestToken with 18 decimals and supply cap of max uint256 value
    s_tokenCreationParams = abi.encode("TestToken", "TT", 18, type(uint256).max);
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

    bytes memory predictedPoolCreationParams = abi.encode(predictedTokenAddress, new address[](0), s_rmnProxy, s_sourceRouter);
    bytes memory predictedPoolInitCode = abi.encodePacked(s_poolInitCode, predictedPoolCreationParams);
    address predictedPoolAddress = DeterministicContractDeployer._predictAddressOfUndeployedContract(
      predictedPoolInitCode, dynamicSalt, address(s_tokenPoolFactory)
    );


    (address tokenAddress, address poolAddress) = s_tokenPoolFactory.createTokenPool(
      address(0), new TokenPoolFactory.ExistingTokenPool[](0), s_poolInitCode, s_tokenInitCode, s_salt
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

  function test_createTokenPool_WithExistingTokenOnRemoteChain_Success() public {}

  function test_createTokenPool_predictFutureAddress_Success() public {}

  function test_createTokenPool_RemotChainNotSupported_Revert() public {}
}
