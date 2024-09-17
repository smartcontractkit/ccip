pragma solidity ^0.8.24;

import {IOwnable} from "../../shared/interfaces/IOwnable.sol";
import {ITypeAndVersion} from "../../shared/interfaces/ITypeAndVersion.sol";
import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {RateLimiter} from "../libraries/RateLimiter.sol";
import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";
import {TokenPool} from "../pools/TokenPool.sol";
import {RegistryModuleOwnerCustom} from "./RegistryModuleOwnerCustom.sol";

import {Create2} from "../../vendor/openzeppelin-solidity/v5.0.2/contracts/utils/Create2.sol";

contract TokenPoolFactory is OwnerIsCreator, ITypeAndVersion {
  using Create2 for bytes32;

  event RemoteChainConfigUpdated(uint64 indexed remoteChainSelector, RemoteChainConfig remoteChainConfig);

  error InvalidZeroAddress();

  struct RemoteTokenPoolInfo {
    uint64 remoteChainSelector; // The CCIP specific selector for the remote chain
    bytes remotePoolAddress; // The address of the remote pool to either deploy or use as is. If
      // the empty parameter flag is provided, the address will be predicted
    bytes remoteTokenAddress; // The address of the remote token to either deploy or use as is
      // If the empty parameter flag is provided, the address will be predicted
    bytes remoteTokenInitCode; // The init code for the remote token if it needs to be deployed
      // and includes all the constructor params already appended
    RateLimiter.Config outboundRateLimiterConfig; // The rate limiter config for token messages to be used in the pool.
      // The specified rate limit will also be applied to the token pool's inbound messages as well.
  }

  /* solhint-disable gas-struct-packing */
  struct RemoteChainConfig {
    address remotePoolFactory;
    /// The factory contract on the remote chain
    address remoteRouter;
    /// The router contract on the remote chain
    address remoteRMNProxy;
  }
  /// The RMNProxy contract on the remote chain
  /* solhint-enable gas-struct-packing */

  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;
  RegistryModuleOwnerCustom internal immutable i_registryModuleOwnerCustom;

  address private immutable i_rmnProxy;
  address private immutable i_ccipRouter;

  bytes4 public constant EMPTY_PARAMETER_FLAG = bytes4(keccak256("EMPTY_PARAMETER_FLAG"));
  string public constant typeAndVersion = "TokenPoolFactory 1.0.0-dev";

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
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    bytes memory tokenInitCode,
    bytes calldata tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) external returns (address tokenAddress, address poolAddress) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // Deploy the token
    address token = Create2.deploy(0, salt, tokenInitCode);

    // Deploy the token pool
    poolAddress = _createTokenPool(token, remoteTokenPools, tokenPoolInitCode, tokenPoolInitArgs, salt);

    // Set the token pool for token in the token admin registry since this contract is the token and pool owner
    _setTokenPoolInTokenAdminRegistry(token, poolAddress);

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
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    bytes calldata tokenPoolInitCode,
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
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    bytes calldata tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) internal returns (address) {
    // If the user doesn't want to provide any special parameters which may be neededfor a custom the token pool then
    /// use the standard burn/mint token pool params. Since the user can provide custom token pool init code,
    // they must also provide custom constructor args.
    if (bytes4(tokenPoolInitArgs) == EMPTY_PARAMETER_FLAG) {
      tokenPoolInitArgs = abi.encode(token, new address[](0), i_rmnProxy, i_ccipRouter);
    }

    // Construct the code that will be deployed from the initCode and the initArgs
    address poolAddress = Create2.deploy(0, salt, abi.encodePacked(tokenPoolInitCode, tokenPoolInitArgs));

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
          inboundRateLimiterConfig: remoteTokenPools[i].outboundRateLimiterConfig
        });

        // Get the address of the remote factory, caching the storage value in memory
        RemoteChainConfig memory remoteChainConfig = s_remoteChainConfigs[remoteTokenPools[i].remoteChainSelector];

        // If the user provides the empty parameter flag, then the address of the token needs to be predicted
        // otherwise the address provided is used.
        if (bytes4(remoteTokenPools[i].remoteTokenAddress) == EMPTY_PARAMETER_FLAG) {
          // The user must provide the initCode for the remote token, so we can predict its address correctly. It's
          // provided in the remoteTokenInitCode field for the remoteTokenPool

          remoteTokenAddress =
            salt.computeAddress(keccak256(remoteTokenPools[i].remoteTokenInitCode), remoteChainConfig.remotePoolFactory);

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

          // Combine the initCode with the initArgs to create the full initCode
          bytes memory remotePoolInitcode = abi.encodePacked(
            type(BurnMintTokenPool).creationCode,
            abi.encode(
              remoteTokenAddress, new address[](0), remoteChainConfig.remoteRMNProxy, remoteChainConfig.remoteRouter
            )
          );

          // Predict the address of the undeployed contract on the destination chain
          chainUpdate.remotePoolAddress =
            abi.encode(salt.computeAddress(keccak256(remotePoolInitcode), remoteChainConfig.remotePoolFactory));

          chainUpdate.remotePoolAddress =
            abi.encode(salt.computeAddress(keccak256(remotePoolInitcode), remoteChainConfig.remotePoolFactory));
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
      IOwnable(poolAddress).transferOwnership(address(msg.sender)); // 2 step ownership transfer

      return poolAddress;
    }
  }

  /// @notice Sets the token pool address in the token admin registry for a newly deployed token pool.
  /// @dev this function should only be called when the token is deployed by this contract as well, otherwise
  /// the token pool will not be able to be set in the token admin registry, and this function will revert.
  /// @param token The address of the token to set the pool for
  /// @param pool The address of the pool to set in the token admin registry
  function _setTokenPoolInTokenAdminRegistry(address token, address pool) internal {
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
}
