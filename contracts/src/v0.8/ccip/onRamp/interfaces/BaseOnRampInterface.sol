// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "../../interfaces/PoolInterface.sol";

interface BaseOnRampInterface {
  error MessageTooLarge(uint256 maxSize, uint256 actualSize);
  error UnsupportedNumberOfTokens();
  error UnsupportedToken(IERC20 token);
  error SenderNotAllowed(address sender);
  error MustBeCalledByRouter();
  error RouterMustSetOriginalSender();

  event AllowlistSet(address[] allowlist);
  event AllowlistEnabledSet(bool enabled);
  event RouterSet(address router);

  function getTokenPool(IERC20 token) external returns (PoolInterface);

  function getSequenceNumber() external view returns (uint64);
}
