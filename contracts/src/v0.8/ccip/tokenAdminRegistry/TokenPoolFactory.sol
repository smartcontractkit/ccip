pragma solidity ^0.8.24;

import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {DeterministicContractDeployer} from "../../shared/util/DeterministicContractDeployer.sol";

import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool} from "../pools/TokenPool.sol";
import {RegistryModuleOwnerCustom} from "./RegistryModuleOwnerCustom.sol";

import {BurnMintERC677} from "../../shared/token/ERC677/BurnMintERC677.sol";
import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";

import {console2 as console} from "forge-std/console2.sol";

contract TokenPoolFactory is OwnerIsCreator {
  using DeterministicContractDeployer for bytes;

  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;
  RegistryModuleOwnerCustom internal immutable i_registryModuleOwnerCustom;
  address internal immutable i_rmnProxy;
  address internal immutable i_ccipRouter;

  event RemoteChainConfigUpdated(uint64 indexed remoteChainSelector, RemoteChainConfig remoteChainConfig);

  error InvalidZeroAddress();

  mapping(uint64 remoteChainSelector => RemoteChainConfig) internal s_remoteChainConfigs;

  bytes4 public constant EMPTY_PARAMETER_FLAG = bytes4(keccak256("EMPTY_PARAMETER"));

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
    bytes remoteTokenInitCode;
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of
      //  the onRamps for the given chain
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of
      // the offRamps for the given chain
  }

  /// @dev ORDERING IS CRITICAL IN PREDICTING ADDRESSES
  struct RemoteChainConfig {
    address remotePoolFactory;
    address remoteRouter;
    address remoteRMNProxy;
  }

  function deployTokenAndTokenPool(
    ExistingTokenPool[] memory remoteTokenPools,
    bytes memory tokenInitCode,
    bytes memory tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) external returns (address tokenAddress, address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    address token = tokenInitCode._deploy(salt);

    return _createTokenPool(token, remoteTokenPools, tokenPoolInitCode, tokenPoolInitArgs, salt, false);
  }

  function deployTokenPoolWithExistingToken(
    address token,
    ExistingTokenPool[] memory remoteTokenPools,
    bytes memory tokenInitCode,
    bytes memory tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) external returns (address tokenAddress, address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    return _createTokenPool(token, remoteTokenPools, tokenPoolInitCode, tokenPoolInitArgs, salt, true);
  }

  function _createTokenPool(
    address token,
    ExistingTokenPool[] memory remoteTokenPools,
    bytes memory tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt,
    bool isExistingToken
  ) internal returns (address tokenAddress, address poolAddress) {
    // If the user doesn't want to provide any special parameters which may be needed for the token pool
    // then use the standard burn/mint token pool params. Since the user can provide custom token pool
    // init code, they must also be able to provide custom constructor args.
    if (bytes4(tokenPoolInitArgs) == EMPTY_PARAMETER_FLAG) {
      tokenPoolInitArgs = abi.encode(token, new address[](0), i_rmnProxy, i_ccipRouter);
    }

    // Construct the code that will be depoyed from the initCode and the initArgs
    bytes memory newtokenPoolInitCode = abi.encodePacked(tokenPoolInitCode, tokenPoolInitArgs);

    // deploy the pool using the above
    poolAddress = newtokenPoolInitCode._deploy(salt);

    // Stack Scoping to reduce pressure on stack too deep
    {
      // Create an array of chain updates to apply to the token pool
      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](remoteTokenPools.length);

      // For Each remote chain in the remoteTokenPools array
      for (uint256 i = 0; i < remoteTokenPools.length; i++) {
        // The address of the remote token is needed later in the function as an address, not as
        // bytes so we store it in memory here
        address remoteTokenAddress;

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

        // If the user provides the empty parameter flag, then we need to predict the address of the token
        // otherwise we can use the address provided by the user
        if (bytes4(remoteTokenPools[i].remoteTokenAddress) == EMPTY_PARAMETER_FLAG) {
          // The user must provide the initCode for the remote token, so we can predict its address correctly. It's provided in the remoteTokenInitCode field for the remoteTokenPool
          remoteTokenAddress = remoteTokenPools[i].remoteTokenInitCode._predictAddressOfUndeployedContract(
            salt, remoteChainConfig.remotePoolFactory
          );

          // The library returns an EVM-compatible address but chainUpdate takes bytes so we encode it
          chainUpdate.remoteTokenAddress = abi.encode(remoteTokenAddress);
        } else {
          // If the user already has a remote token deployed, reuse the address. We still need it as
          // an address for later, so we store it in memory after decoding.
          // NOTE: This assumes that the provided address can be decoded into an EVM address.
          remoteTokenAddress = abi.decode(remoteTokenPools[i].remoteTokenAddress, (address));

          chainUpdate.remoteTokenAddress = remoteTokenPools[i].remoteTokenAddress;
        }

        // If the user provides the empty parameter flag, then we need to predict the address of the pool
        if (bytes4(remoteTokenPools[i].remotePoolAddress) == EMPTY_PARAMETER_FLAG) {
          // Generate the initCode that will be used on the remote chain. It is assumed that tokenInitCode
          // will be the same on all chains, so we can reuse it here.

          // Calculate the remote pool Args with an empty allowList, remote RMN, and Remote Router addresses. Since the first constructor parameter is an EVM token address, we can use the remoteTokenAddress we acquired earlier.
          bytes memory remotePoolInitArgs = abi.encode(
            remoteTokenAddress, new address[](0), remoteChainConfig.remoteRMNProxy, remoteChainConfig.remoteRouter
          );

          // Combine the initCode with the initArgs to create the full initCode
          bytes memory remotePoolInitcode = abi.encodePacked(type(BurnMintTokenPool).creationCode, remotePoolInitArgs);

          // Predict the address of the undeployed contract on the destination chain
          chainUpdate.remotePoolAddress = abi.encode(
            remotePoolInitcode._predictAddressOfUndeployedContract(salt, remoteChainConfig.remotePoolFactory)
          );
        } else {
          // If the user already has a remote pool deployed, reuse the address.
          chainUpdate.remotePoolAddress = remoteTokenPools[i].remotePoolAddress;
        }

        // Update the chainUpdate struct in the chainUpdates array
        chainUpdates[i] = chainUpdate;
      }

      // If the token already exists, then this contract will not be the owner,
      // and thus it will not be able to set the token pool or transfer ownership
      // which must be done manually by the end user.
      if (!isExistingToken) {
        _setTokenPool(token, poolAddress);
        OwnerIsCreator(token).transferOwnership(address(msg.sender)); // 2 step ownership transfer
      }

      // Apply the chain updates to the token pool
      TokenPool(poolAddress).applyChainUpdates(chainUpdates);

      OwnerIsCreator(poolAddress).transferOwnership(address(msg.sender)); // 2 step ownership transfer

      // TODO: Add more events
      return (token, poolAddress);
    }
  }

  function _setTokenPool(address token, address pool) public {
    // propose this factory as the admin for the token in the token admin registry
    i_registryModuleOwnerCustom.registerAdminViaOwner(token);

    // Accept the admin role by the token admin registry
    i_tokenAdminRegistry.acceptAdminRole(token);

    // Set the pool address in the token admin registry
    i_tokenAdminRegistry.setPool(token, pool);

    i_tokenAdminRegistry.transferAdminRole(token, msg.sender);
  }

  function updateRemoteChainConfig(
    uint64 remoteChainSelector,
    RemoteChainConfig calldata remoteConfig
  ) public onlyOwner {
    s_remoteChainConfigs[remoteChainSelector] = remoteConfig;

    emit RemoteChainConfigUpdated(remoteChainSelector, remoteConfig);
  }

  function getRemoteChainConfig(uint64 remoteChainSelector) public view returns (RemoteChainConfig memory) {
    return s_remoteChainConfigs[remoteChainSelector];
  }
}
