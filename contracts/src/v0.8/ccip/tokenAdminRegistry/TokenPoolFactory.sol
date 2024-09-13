pragma solidity ^0.8.24;

import {ITypeAndVersion} from "../../shared/interfaces/ItypeAndVersion.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {DeterministicContractDeployer} from "../../shared/util/DeterministicContractDeployer.sol";

import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool} from "../pools/TokenPool.sol";
import {RegistryModuleOwnerCustom} from "./RegistryModuleOwnerCustom.sol";

import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";

contract TokenPoolFactory is OwnerIsCreator, ITypeAndVersion {
  using DeterministicContractDeployer for bytes;

  event RemoteChainConfigUpdated(uint64 indexed remoteChainSelector, RemoteChainConfig remoteChainConfig);

  error InvalidZeroAddress();

  struct RemoteTokenPoolInfo {
    uint64 remoteChainSelector;
    bytes remotePoolAddress;
    bytes remoteTokenAddress;
    bytes remoteTokenInitCode;
    RateLimiter.Config outboundRateLimiterConfig;
    RateLimiter.Config inboundRateLimiterConfig;
  }

  struct RemoteChainConfig {
    address remotePoolFactory;
    address remoteRouter;
    address remoteRMNProxy;
  }

  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;
  RegistryModuleOwnerCustom internal immutable i_registryModuleOwnerCustom;

  address internal immutable i_rmnProxy;
  address internal immutable i_ccipRouter;

  // bytes4(keccak256("EMPTY_PARAMETER_FLAG"))
  bytes4 public constant EMPTY_PARAMETER_FLAG = 0x8fc9a1e4;

  mapping(uint64 remoteChainSelector => RemoteChainConfig) internal s_remoteChainConfigs;

  constructor(address tokenAdminRegistry, address tokenAdminModule, address rmnProxy, address ccipRouter) {
    if (
      tokenAdminRegistry == address(0) || rmnProxy == address(0) || rmnProxy == address(0) || ccipRouter == address(0)
    ) revert InvalidZeroAddress();

    i_tokenAdminRegistry = ITokenAdminRegistry(tokenAdminRegistry);
    i_registryModuleOwnerCustom = RegistryModuleOwnerCustom(tokenAdminModule);
    i_rmnProxy = rmnProxy;
    i_ccipRouter = ccipRouter;
  }

  function deployTokenAndTokenPool(
    RemoteTokenPoolInfo[] memory remoteTokenPools,
    bytes memory tokenInitCode,
    bytes memory tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) external returns (address tokenAddress, address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // Deploy the token
    address token = tokenInitCode._deploy(salt);

    // Deploy the token pool
    poolAddress = _createTokenPool(token, remoteTokenPools, tokenPoolInitCode, tokenPoolInitArgs, salt);

    // Set the token pool in the token admin registry since this contract is the owner of the token and the pool
    _setTokenPool(token, poolAddress);

    // Transfer the ownership of the token to the msg.sender.
    // This is a 2 step process and must be accepted in a separate tx.
    OwnerIsCreator(token).transferOwnership(address(msg.sender)); // 2 step ownership transfer

    return (token, poolAddress);
  }

  /// @notice Deploys a token pool with an existing ERC20 token
  /// @dev Since the token already exists, this contract is not the owner and therefore cannot configure the
  /// token pool in the token admin registry in the same transaction. The user must invoke the calls to the
  /// tokenAdminRegistry manually
  /// @param token The address of the existing token to be used in the token pool
  /// @param remoteTokenPools An array of remote token pools info to be used in the pool's applyChainUpdates function
  /// @param tokenPoolInitCode The creation code for the token pool
  /// @param tokenPoolInitArgs The arguments to be passed to the token pool's constructor and concatenated with the
  /// initCode to be passed into the deployer function
  /// @param salt The salt to be used in the create2 deployment of the token pool
  /// @return poolAddress The address of the token pool that was deployed
  function deployTokenPoolWithExistingToken(
    address token,
    RemoteTokenPoolInfo[] memory remoteTokenPools,
    bytes memory tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) external returns (address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // create the token pool and return the address
    return _createTokenPool(token, remoteTokenPools, tokenPoolInitCode, tokenPoolInitArgs, salt);
  }

  /// @notice Deploys a token pool with the given token information and remote token pools
  /// @param token The token to be used in the token pool
  /// @param remoteTokenPools An array of remote token pools info to be used in the pool's applyChainUpdates function
  /// @param tokenPoolInitCode The creation code for the token pool
  /// @param tokenPoolInitArgs The arguments to be passed to the token pool's constructor and concatenated with the
  /// initCode to be passed into the deployer function
  /// @param salt The salt to be used in the create2 deployment of the token pool
  /// @return poolAddress The address of the token pool that was deployed
  function _createTokenPool(
    address token,
    RemoteTokenPoolInfo[] memory remoteTokenPools,
    bytes memory tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) internal returns (address) {
    // If the user doesn't want to provide any special parameters which may be neededfor a custom the token pool then
    /// use the standard burn/mint token pool params. Since the user can provide custom token pool init code,
    // they must also provide custom constructor args.
    if (bytes4(tokenPoolInitArgs) == EMPTY_PARAMETER_FLAG) {
      tokenPoolInitArgs = abi.encode(token, new address[](0), i_rmnProxy, i_ccipRouter);
    }


    // Stack scoping to reduce pressure on stack too deep from the concatenated initCode and initArgs
    address poolAddress;
    {
      // Construct the code that will be depoyed from the initCode and the initArgs
      bytes memory newtokenPoolInitCode = abi.encodePacked(tokenPoolInitCode, tokenPoolInitArgs);

      // deploy the pool using the above
      poolAddress = newtokenPoolInitCode._deploy(salt);
    }

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

        // If the user provides the empty parameter flag, then the address of the token needs to be predicted
        // otherwise the address provided is used.
        if (bytes4(remoteTokenPools[i].remoteTokenAddress) == EMPTY_PARAMETER_FLAG) {
          // The user must provide the initCode for the remote token, so we can predict its address correctly. It's
          // provided in the remoteTokenInitCode field for the remoteTokenPool
          remoteTokenAddress = remoteTokenPools[i].remoteTokenInitCode._predictAddressOfUndeployedContract(
            salt, remoteChainConfig.remotePoolFactory
          );

          // The library returns an EVM-compatible address but chainUpdate takes bytes so we encode it
          chainUpdate.remoteTokenAddress = abi.encode(remoteTokenAddress);
        } else {
          // If the user already has a remote token deployed, reuse the address. We still need it as
          // an address for later, so we store it in memory after decoding.
          // This assumes that the provided address can be decoded into an EVM address.
          remoteTokenAddress = abi.decode(remoteTokenPools[i].remoteTokenAddress, (address));

          chainUpdate.remoteTokenAddress = remoteTokenPools[i].remoteTokenAddress;
        }

        // If the user provides the empty parameter flag, the address of the pool should be predicted
        if (bytes4(remoteTokenPools[i].remotePoolAddress) == EMPTY_PARAMETER_FLAG) {
          // Generate the initCode that will be used on the remote chain. It is assumed that tokenInitCode
          // will be the same on all chains, so it can be reused here.

          // Calculate the remote pool Args with an empty allowList, remote RMN, and Remote Router addresses.
          // Since the first constructor parameter is an EVM token address, the remoteTokenAddress acquired earlier.
          // can be used.
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

      // Apply the chain updates to the token pool
      TokenPool(poolAddress).applyChainUpdates(chainUpdates);

      // Being the 2 step ownership transfer of the token pool to the msg.sender.
      OwnerIsCreator(poolAddress).transferOwnership(address(msg.sender)); // 2 step ownership transfer

      return poolAddress;
    }
  }

  /// @notice Sets the token pool address in the token admin registry for a newly deployed token pool.
  /// @dev this function should only be called when the token is deployed by this contract as well, otherwise
  /// the token pool will not be able to be set in the token admin registry, and this function will revert.
  /// @param token The address of the token to set the pool for
  /// @param pool The address of the pool to set in the token admin registry
  function _setTokenPool(address token, address pool) internal {
    // propose this factory as the admin for the token in the token admin registry
    i_registryModuleOwnerCustom.registerAdminViaOwner(token);

    // Accept the admin role by the token admin registry
    i_tokenAdminRegistry.acceptAdminRole(token);

    // Set the pool address in the token admin registry
    i_tokenAdminRegistry.setPool(token, pool);

    // Transfer the admin role for the token pool back to the msg.sender. This is a 2 step process
    // and must be accepted in a separate tx.
    i_tokenAdminRegistry.transferAdminRole(token, msg.sender);
  }

  function updateRemoteChainConfig(
    uint64[] calldata remoteChainSelectors,
    RemoteChainConfig[] calldata remoteConfigs
  ) external onlyOwner {
    for (uint256 i = 0; i < remoteChainSelectors.length; i++) {
      s_remoteChainConfigs[remoteChainSelectors[i]] = remoteConfigs[i];
      emit RemoteChainConfigUpdated(remoteChainSelectors[i], remoteConfigs[i]);
    }
  }

  /// @notice Get the remote chain config for a given remote chain selector
  /// @param remoteChainSelector The remote chain selector to get the config for
  /// @return remoteChainConfig The remote chain config for the given remote chain selector
  function getRemoteChainConfig(uint64 remoteChainSelector) public view returns (RemoteChainConfig memory) {
    return s_remoteChainConfigs[remoteChainSelector];
  }

  /// @notice Get the type and version of the contract
  /// @return The type and version of the contract
  function typeAndVersion() external pure returns (string memory) {
    return "TokenPoolFactory 1.0.0-dev";
  }
}
