pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {DeterministicContractDeployer} from "../../shared/util/DeterministicContractDeployer.sol";

import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool} from "../pools/TokenPool.sol";
import {RegistryModuleOwnerCustom} from "./RegistryModuleOwnerCustom.sol";

import {BurnMintERC677} from "../../shared/token/ERC677/BurnMintERC677.sol";
import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";

contract TokenPoolFactory is OwnerIsCreator {
  using DeterministicContractDeployer for bytes;

  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;
  RegistryModuleOwnerCustom internal immutable i_registryModuleOwnerCustom;
  address internal immutable i_rmnProxy;
  address internal immutable i_ccipRouter;

  event RemoteChainConfigUpdated(
    uint64 indexed remoteChainSelector, RemoteChainConfig remoteChainConfig
  );

  error InvalidZeroAddress();

  mapping(uint64 remoteChainSelector => RemoteChainConfig) internal s_remoteChainConfigs;

  constructor(address tokenAdminRegistry, address tokenAdminModule, address rmnProxy, address ccipRouter) {
    if (tokenAdminRegistry == address(0) || rmnProxy == address(0)) revert InvalidZeroAddress();

    i_tokenAdminRegistry = ITokenAdminRegistry(tokenAdminRegistry);
    i_registryModuleOwnerCustom = RegistryModuleOwnerCustom(tokenAdminModule);
    i_rmnProxy = rmnProxy;
    i_ccipRouter = ccipRouter;
  }

  struct ExistingTokenPool {
    uint64 remoteChainSelector;
    bytes remotePoolAddress;
    bytes remoteTokenAddress;
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of
      //  the onRamps for the given chain
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of
      // the offRamps for the given chain
  }

  struct RemoteChainConfig {
    address remotePoolFactory;
    address remoteRouter;
    address remoteRMNProxy;
  }

  struct Deployment {
    BurnMintERC677 token;
    BurnMintTokenPool pool;
  }

  function createTokenPool(
    address existingToken,
    ExistingTokenPool[] memory remoteTokenPools,

    /// @notice: init code and token args have been combined into one to prevent a stack too deep error
    bytes memory tokenPoolInitCode,
    bytes calldata tokenInitCode,
    bytes32 salt
  ) public returns (address tokenAddress, address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // If there is no existing ERC20-token, deploy a new one, else return the existing address
    if (existingToken == address(0)) {
      tokenAddress = tokenInitCode._deploy(salt);
    } else {
      tokenAddress = existingToken;
    }

    // Configure the token by granting the roles
    BurnMintERC677(tokenAddress).grantMintAndBurnRoles(poolAddress);

    bytes memory tokenPoolInitArgs = abi.encode(tokenAddress, new address[](0), i_rmnProxy, i_ccipRouter);
    tokenPoolInitCode = abi.encodePacked(tokenPoolInitCode, tokenPoolInitArgs);
    poolAddress = tokenPoolInitCode._deploy(salt);

    // Stack Scoping to reduce pressure on stack too deep
    {
      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](remoteTokenPools.length);

      // For Each remote chain in the remoteTokenPools array
      for (uint256 i = 0; i < remoteTokenPools.length; i++) {
        // Declaring the struct and updating remote addresses later in the function prevents stack too deep
        TokenPool.ChainUpdate memory chainUpdate = TokenPool.ChainUpdate({
          remoteChainSelector: remoteTokenPools[i].remoteChainSelector,
          allowed: true,
          remotePoolAddress: "",
          remoteTokenAddress: "",
          outboundRateLimiterConfig: remoteTokenPools[i].outboundRateLimiterConfig,
          inboundRateLimiterConfig: remoteTokenPools[i].inboundRateLimiterConfig
        });

        // Get the address of the remote factory, caching the storage value in memory
        RemoteChainConfig memory remoteChainConfig = s_remoteChainConfigs[remoteTokenPools[i].remoteChainSelector];

        // If the user already has a remote token deployed, reuse the address, otherwise calculate the new address
        // of the undeployed token on the destination chain
        if (remoteTokenPools[i].remoteTokenAddress.length == 0) {
          // Since the tokenInitCode doesn't require any dynamic parameters, and is already deployed, we can re-use
          // the initCode used at the beginning of the function.
          chainUpdate.remoteTokenAddress =
            abi.encode(tokenInitCode._predictAddressOfUndeployedContract(salt, remoteChainConfig.remotePoolFactory));
        } else {
          // If the user already has a remote token deployed, reuse the address
          chainUpdate.remoteTokenAddress = remoteTokenPools[i].remoteTokenAddress;
        }

        // If the user already has a remote pool deployed, reuse the address, otherwise calculate the new address
        // of the undeployed pool on the destination chain
        if (remoteTokenPools[i].remotePoolAddress.length == 0) {
          // Generate the initCode that will be used on the remote chain. It is assumed that tokenInitCode
          // will be the same on all chains.
          
          // Calculate the remote pool Args with an empty allowList, remote RMN, and Remote Router addresses
          bytes memory remotePoolInitArgs = abi.encode(
            chainUpdate.remoteTokenAddress, new address[](0), remoteChainConfig.remoteRMNProxy, remoteChainConfig.remoteRouter
          );

          // Combine the initCode with the initArgs to create the full initCode
          bytes memory remotePoolInitcode =
            abi.encodePacked(type(BurnMintTokenPool).creationCode, remotePoolInitArgs);

          // Predict the address of the undeployed contract on the destination chain
          chainUpdate.remotePoolAddress =
            abi.encode(remotePoolInitcode._predictAddressOfUndeployedContract(salt, remoteChainConfig.remotePoolFactory));
        } else {
          // If the user already has a remote pool deployed, reuse the address
          chainUpdate.remotePoolAddress = remoteTokenPools[i].remotePoolAddress;
        }

        // Update the chainUpdate struct in the chainUpdates array
        chainUpdates[i] = chainUpdate;
      }

       // Setup token roles
      _setTokenPool(tokenAddress, poolAddress);

      // Apply the chain updates to the token pool
      TokenPool(poolAddress).applyChainUpdates(chainUpdates);

      _releaseOwnership(tokenAddress, poolAddress);

      return (tokenAddress, poolAddress);
    }
  }

  function _setTokenPool(address token, address pool) public {
    // propose this factory as the admin for the token in the token admin registry
    i_registryModuleOwnerCustom.registerAdminViaOwner(token);

    // Accept the admin role by the token admin registry
    i_tokenAdminRegistry.acceptAdminRole(token);

    // Set the pool address in the token admin registry
    i_tokenAdminRegistry.setPool(token, pool);
  }

  function _releaseOwnership(address token, address pool) internal {
    i_tokenAdminRegistry.transferAdminRole(token, msg.sender);

    OwnerIsCreator(token).transferOwnership(address(msg.sender)); // 1 step ownership transfer
    OwnerIsCreator(pool).transferOwnership(address(msg.sender)); // 2 step ownership transfer
  }

  // TODO: Update Event Maybe.
  function updateRemoteChainConfig(uint64 remoteChainSelector, RemoteChainConfig calldata remoteConfig) public onlyOwner {
    s_remoteChainConfigs[remoteChainSelector] = remoteConfig;

    emit RemoteChainConfigUpdated(remoteChainSelector, remoteConfig);
  }

  function getRemoteChainConfig(uint64 remoteChainSelector) public view returns (RemoteChainConfig memory) {
    return s_remoteChainConfigs[remoteChainSelector];
  }
}
