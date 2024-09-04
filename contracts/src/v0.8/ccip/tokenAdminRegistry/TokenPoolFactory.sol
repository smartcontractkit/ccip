pragma solidity ^0.8.24;

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";
import {DeterministicContractDeployer} from "../../shared/util/DeterministicDeployer.sol";

import {TokenAdminRegistry} from "./TokenAdminRegistry.sol";

import {RateLimiter} from "../libraries/RateLimiter.sol";
import {TokenPool} from "../pools/TokenPool.sol";

import {ITokenAdminRegistry} from "../interfaces/ITokenAdminRegistry.sol";

import {BurnMintERC677} from "../../shared/token/ERC677/BurnMintERC677.sol";
import {BurnMintTokenPool} from "../pools/BurnMintTokenPool.sol";

contract TokenPoolFactory is OwnerIsCreator {
  using DeterministicContractDeployer for bytes;

  ITokenAdminRegistry internal immutable i_tokenAdminRegistry;

  error InvalidZeroAddress();

  mapping(uint64 remoteChainSelector => address remotePoolFactory) internal s_remotePoolFactories;

  constructor(address tokenAdminRegistry) {
    if (tokenAdminRegistry == address(0)) revert InvalidZeroAddress();

    i_tokenAdminRegistry = ITokenAdminRegistry(tokenAdminRegistry);
  }

  struct ExistingTokenPool {
    uint64 remoteChainSelector;
    bytes remotePoolAddress;
    bytes remoteTokenAddress;
    RateLimiter.Config outboundRateLimiterConfig; // Outbound rate limited config, meaning the rate limits for all of the onRamps for the given chain
    RateLimiter.Config inboundRateLimiterConfig; // Inbound rate limited config, meaning the rate limits for all of the offRamps for the given chain
  }

  struct Deployment {
    BurnMintERC677 token;
    BurnMintTokenPool pool;
  }

  function createTokenPool(
    address existingToken,
    ExistingTokenPool[] memory remoteTokenPools,
    /// @notice: init code and token args have been combined into one to prevent a stack too deep error
    bytes calldata tokenPoolInitCode,
    bytes calldata tokenInitCode,
    bytes32 salt //TODO: Check that this is allowed for omni-chain deployments
  ) public returns (address tokenAddress, address poolAddress) {
    // If there is no existing ERC20-token, deploy a new one, else return the existing address
    if (existingToken == address(0)) {
      tokenAddress = tokenInitCode.deploy(salt);
    } else {
      tokenAddress = existingToken;
    }

    // Ensure a unique deployment between senders even if the same input parameter is used
    salt = keccak256(abi.encodePacked(salt, msg.sender));

    // Deploy a new token pool locally
    poolAddress = tokenPoolInitCode.deploy(salt);

    // Setup token roles
    BurnMintERC677(tokenAddress).grantMintAndBurnRoles(poolAddress);

    _setTokenPool(tokenAddress, poolAddress);

    // Stack Scoping to reduce pressure on stack too deep
    {
      TokenPool.ChainUpdate[] memory chainUpdates = new TokenPool.ChainUpdate[](remoteTokenPools.length);
      for (uint256 i = 0; i < remoteTokenPools.length; i++) {
        address remoteFactoryAddress = s_remotePoolFactories[remoteTokenPools[i].remoteChainSelector];

        chainUpdates[i] = TokenPool.ChainUpdate({
          remoteChainSelector: remoteTokenPools[i].remoteChainSelector,
          allowed: true,

          // If an address is not passed, predict the address using the remote factory address.
          remotePoolAddress: remoteTokenPools[i].remotePoolAddress.length == 0
            ? abi.encode(tokenPoolInitCode.predictAddressOfUndeployedContract(salt, remoteFactoryAddress))
            : remoteTokenPools[i].remotePoolAddress,

          remoteTokenAddress: remoteTokenPools[i].remoteTokenAddress.length == 0
            ? abi.encode(tokenInitCode.predictAddressOfUndeployedContract(salt, remoteFactoryAddress))
            : remoteTokenPools[i].remoteTokenAddress,

          outboundRateLimiterConfig: remoteTokenPools[i].outboundRateLimiterConfig,
          inboundRateLimiterConfig: remoteTokenPools[i].inboundRateLimiterConfig
        });
      }

      TokenPool(poolAddress).applyChainUpdates(chainUpdates);
    }

    _releaseOwnership(tokenAddress, poolAddress);
  }

  function _setTokenPool(address token, address pool) public {
    // propose this factory as the admin for the token in the token admin registry
    i_tokenAdminRegistry.proposeAdministrator(token, address(this));

    // Accept the admin role by the token admin registry
    i_tokenAdminRegistry.acceptAdminRole(token);

    // Set the pool address in the token admin registry
    i_tokenAdminRegistry.setPool(token, pool);
  }

  function _releaseOwnership(address token, address pool) public {
    i_tokenAdminRegistry.transferAdminRole(token, msg.sender);

    OwnerIsCreator(token).transferOwnership(address(msg.sender)); // 1 step ownership transfer
    OwnerIsCreator(pool).transferOwnership(address(msg.sender)); // 2 step ownership transfer
  }

  // TODO: Update with struct and arrays and shit. PoC for now
  function updateRemotePoolFactory(uint64 remoteChainSelector, address remotePoolFactory) public onlyOwner {
    s_remotePoolFactories[remoteChainSelector] = remotePoolFactory;
  }

  function getRemotePoolFactory(uint64 remoteChainSelector) public view returns (address) {
    return s_remotePoolFactories[remoteChainSelector];
  }
}
