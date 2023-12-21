// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface IBridge {
  error BridgeAddressCannotBeZero();

  function getL1Bridge() external view returns (address);
}

interface IL1Bridge is IBridge {
  function depositERC20ToL2(address l1Token, address l2Token, address recipient, uint256 amount) external payable;

  function depositNativeToL2(address recipient, uint256 amount) external payable;

  function finalizeWithdrawERC20FromL2(
    address l1Token,
    address l2Token,
    address from,
    address to,
    uint256 amount,
    bytes calldata data
  ) external;

  function finalizeWithdrawNativeFromL2(address from, address to, uint256 amount, bytes calldata data) external;
}

interface IL2Bridge is IBridge {
  function depositERC20ToL1(address l2Token, address recipient, uint256 amount) external;

  function depositNativeToL1(address recipient, uint256 amount) external payable;
}
