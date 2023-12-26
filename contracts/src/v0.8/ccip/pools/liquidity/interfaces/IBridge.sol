// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.0;

interface IBridge {
  error BridgeAddressCannotBeZero();
  error MsgValueDoesNotMatchAmount(uint256 msgValue, uint256 amount);

  function sendERC20(address l1Token, address l2Token, address recipient, uint256 amount) external payable;
}

interface IL1Bridge is IBridge {
  function finalizeWithdrawERC20FromL2(
    address l2Sender,
    address l1Receiver,
    bytes calldata bridgeSpecificPayload
  ) external;
}

interface IL2Bridge is IBridge {}
