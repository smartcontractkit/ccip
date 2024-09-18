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
    // The address of the remote pool to either deploy or use as is. If
    // the empty parameter flag is provided, the address will be predicted
    bytes remotePoolAddress;
    // The address of the remote token to either deploy or use as is
    // If the empty parameter flag is provided, the address will be predicted
    bytes remoteTokenAddress;
    // The init code for the remote token if it needs to be deployed
    // and includes all the constructor params already appended
    bytes remoteTokenInitCode;
    // The rate limiter config for token messages to be used in the pool.
    // The specified rate limit will also be applied to the token pool's inbound messages as well.
    RateLimiter.Config rateLimiterConfig;
  }

  // solhint-disable-next-line gas-struct-packing
  struct RemoteChainConfig {
    address remotePoolFactory; // The factory contract on the remote chain
    address remoteRouter; // The router contract on the remote chain
    address remoteRMNProxy; // The RMNProxy contract on the remote chain
  }

  struct RemoteChainConfigUpdate {
    uint64 remoteChainSelector;
    RemoteChainConfig remoteChainConfig;
  }

  bytes4 public constant EMPTY_PARAMETER_FLAG = bytes4(keccak256("EMPTY_PARAMETER_FLAG"));
  string public constant typeAndVersion = "TokenPoolFactory 1.0.0-dev";

  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;
  RegistryModuleOwnerCustom internal immutable i_registryModuleOwnerCustom;

  address private immutable i_rmnProxy;
  address private immutable i_ccipRouter;

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

  // ================================================================
  // |                   Top-Level Deployment                        |
  // ================================================================

  function deployTokenAndTokenPool(
    RemoteTokenPoolInfo[] calldata remoteTokenPools,
    bytes memory tokenInitCode,
    bytes calldata tokenPoolInitCode,
    bytes memory tokenPoolInitArgs,
    bytes32 salt
  ) external returns (address, address) {
    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // Deploy the token
    address token = Create2.deploy(0, salt, tokenInitCode);

    // Deploy the token pool
    address pool = _createTokenPool(token, remoteTokenPools, tokenPoolInitCode, tokenPoolInitArgs, salt);

    // Set the token pool for token in the token admin registry since this contract is the token and pool owner
    _setTokenPoolInTokenAdminRegistry(token, pool);

    // Transfer the ownership of the token to the msg.sender.
    // This is a 2 step process and must be accepted in a separate tx.
    IOwnable(token).transferOwnership(msg.sender);

    return (token, pool);
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

  // ================================================================
  // |                Pool Deployment/Configuration                  |
  // ================================================================

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
    // Create an array of chain updates to apply to the token pool
    TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](remoteTokenPools.length);

    for (uint256 i = 0; i < remoteTokenPools.length; i++) {
      RemoteTokenPoolInfo memory remoteTokenPool = remoteTokenPools[i];
      RemoteChainConfig memory remoteChainConfig = s_remoteChainConfigs[remoteTokenPool.remoteChainSelector];

      // If the user provides the empty parameter flag, then the address of the token needs to be predicted
      // otherwise the address provided is used.
      if (bytes4(remoteTokenPool.remoteTokenAddress) == EMPTY_PARAMETER_FLAG) {
        // The user must provide the initCode for the remote token, so we can predict its address correctly. It's
        // provided in the remoteTokenInitCode field for the remoteTokenPool
        remoteTokenPool.remoteTokenAddress = abi.encode(
          salt.computeAddress(keccak256(remoteTokenPool.remoteTokenInitCode), remoteChainConfig.remotePoolFactory)
        );
      }

      // If the user provides the empty parameter flag, the address of the pool should be predicted
      if (bytes4(remoteTokenPool.remotePoolAddress) == EMPTY_PARAMETER_FLAG) {
        // Generate the initCode that will be used on the remote chain. It is assumed that tokenInitCode
        // will be the same on all chains, so it can be reused here.

        // Combine the initCode with the initArgs to create the full initCode
        bytes32 remotePoolInitcode = keccak256(
          bytes.concat(
            type(BurnMintTokenPool).creationCode,
            // Calculate the remote pool Args with an empty allowList, remote RMN, and Remote Router addresses.
            abi.encode(
              abi.decode(remoteTokenPool.remoteTokenAddress, (address)),
              new address[](0),
              remoteChainConfig.remoteRMNProxy,
              remoteChainConfig.remoteRouter
            )
          )
        );

        // Predict the address of the undeployed contract on the destination chain
        remoteTokenPool.remotePoolAddress =
          abi.encode(salt.computeAddress(remotePoolInitcode, remoteChainConfig.remotePoolFactory));
      }

      chainUpdates[i] = TokenPool.ChainUpdate({
        remoteChainSelector: remoteTokenPool.remoteChainSelector,
        allowed: true,
        remotePoolAddress: remoteTokenPool.remotePoolAddress,
        remoteTokenAddress: remoteTokenPool.remoteTokenAddress,
        outboundRateLimiterConfig: remoteTokenPool.rateLimiterConfig,
        inboundRateLimiterConfig: remoteTokenPool.rateLimiterConfig
      });
    }

    // If the user doesn't want to provide any special parameters which may be needed for a custom the token pool then
    /// use the standard burn/mint token pool params. Since the user can provide custom token pool init code,
    // they must also provide custom constructor args.
    if (bytes4(tokenPoolInitArgs) == EMPTY_PARAMETER_FLAG) {
      tokenPoolInitArgs = abi.encode(token, new address[](0), i_rmnProxy, i_ccipRouter);
    }

    // Construct the code that will be deployed from the initCode and the initArgs
    address poolAddress = Create2.deploy(0, salt, abi.encodePacked(tokenPoolInitCode, tokenPoolInitArgs));

    // Apply the chain updates to the token pool
    TokenPool(poolAddress).applyChainUpdates(chainUpdates);

    // Begin the 2 step ownership transfer of the token pool to the msg.sender.
    IOwnable(poolAddress).transferOwnership(address(msg.sender)); // 2 step ownership transfer

    return poolAddress;
  }

  /// @notice Sets the token pool address in the token admin registry for a newly deployed token pool.
  /// @dev this function should only be called when the token is deployed by this contract as well, otherwise
  /// the token pool will not be able to be set in the token admin registry, and this function will revert.
  /// @param token The address of the token to set the pool for
  /// @param pool The address of the pool to set in the token admin registry
  function _setTokenPoolInTokenAdminRegistry(address token, address pool) internal {
    i_registryModuleOwnerCustom.registerAdminViaOwner(token);
    i_tokenAdminRegistry.acceptAdminRole(token);
    i_tokenAdminRegistry.setPool(token, pool);

    // Begin the 2 admin transfer process which must be accepted in a separate tx.
    i_tokenAdminRegistry.transferAdminRole(token, msg.sender);
  }

  // ================================================================
  // |                  Remote Chain Configuration                  |
  // ================================================================

  /// @notice Updates the remote chain config for the given remote chain selector
  /// @param remoteChainConfigs An array of remote chain configs to update
  /// @dev The function may only be called by the contract owner.
  function updateRemoteChainConfig(RemoteChainConfigUpdate[] calldata remoteChainConfigs) external onlyOwner {
    for (uint256 i = 0; i < remoteChainConfigs.length; ++i) {
      RemoteChainConfig memory remoteConfig = remoteChainConfigs[i].remoteChainConfig;

      if (
        remoteChainConfigs[i].remoteChainSelector == 0 || remoteConfig.remotePoolFactory == address(0)
          || remoteConfig.remoteRouter == address(0) || remoteConfig.remoteRMNProxy == address(0)
      ) revert InvalidZeroAddress();

      s_remoteChainConfigs[remoteChainConfigs[i].remoteChainSelector] = remoteChainConfigs[i].remoteChainConfig;
      emit RemoteChainConfigUpdated(remoteChainConfigs[i].remoteChainSelector, remoteConfig);
    }
  }

  /// @notice Get the remote chain config for a given remote chain selector
  /// @param remoteChainSelector The remote chain selector to get the config for
  /// @return remoteChainConfig The remote chain config for the given remote chain selector
  function getRemoteChainConfig(uint64 remoteChainSelector) public view returns (RemoteChainConfig memory) {
    return s_remoteChainConfigs[remoteChainSelector];
  }
}
