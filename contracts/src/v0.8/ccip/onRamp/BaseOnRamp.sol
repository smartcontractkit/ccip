// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../vendor/SafeERC20.sol";
import "../pools/TokenPoolRegistry.sol";
import "../health/HealthChecker.sol";
import "../priceFeedRegistry/PriceFeedRegistry.sol";
import "../pools/PoolCollector.sol";
import "./interfaces/BaseOnRampInterface.sol";

contract BaseOnRamp is BaseOnRampInterface, HealthChecker, TokenPoolRegistry, PriceFeedRegistry, PoolCollector {
  using SafeERC20 for IERC20;

  // Chain ID of the source chain (where this contract is deployed)
  uint256 public immutable CHAIN_ID;
  // Chain ID of the destination chain (where this contract sends messages)
  uint256 public immutable DESTINATION_CHAIN_ID;

  // The current sequence number
  uint64 internal s_sequenceNumber;
  // Whether the allowlist is enabled
  bool internal s_allowlistEnabled;
  // Addresses that are allowed to send messages
  mapping(address => bool) internal s_allowed;
  // List of allowed addresses
  address[] internal s_allowList;

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal
  ) HealthChecker(afn, maxTimeWithoutAFNSignal) TokenPoolRegistry(tokens, pools) PriceFeedRegistry(tokens, feeds) {
    CHAIN_ID = chainId;
    DESTINATION_CHAIN_ID = destinationChainId;
    s_sequenceNumber = 1;

    if (allowlist.length > 0) {
      s_allowlistEnabled = true;
      s_allowList = allowlist;
    }
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
  }

  /**
   * @notice Get the pool for a specific token
   * @param token token to get the pool for
   * @return pool PoolInterface
   */
  function getTokenPool(IERC20 token) external view override returns (PoolInterface) {
    return getPool(token);
  }

  function setAllowlistEnabled(bool enabled) external onlyOwner {
    s_allowlistEnabled = enabled;
    emit AllowlistEnabledSet(enabled);
  }

  function getAllowlistEnabled() external view returns (bool) {
    return s_allowlistEnabled;
  }

  function setAllowlist(address[] calldata allowlist) external onlyOwner {
    // Remove existing allowlist
    address[] memory existingList = s_allowList;
    for (uint256 i = 0; i < existingList.length; i++) {
      s_allowed[existingList[i]] = false;
    }

    // Set the new allowlist
    s_allowList = allowlist;
    for (uint256 i = 0; i < allowlist.length; i++) {
      s_allowed[allowlist[i]] = true;
    }
    emit AllowlistSet(allowlist);
  }

  function getAllowlist() external view returns (address[] memory) {
    return s_allowList;
  }

  function getSequenceNumber() external view returns (uint64) {
    return s_sequenceNumber;
  }
}
