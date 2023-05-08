// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../../../interfaces/IAny2EVMMessageReceiver.sol";

import "../../../../vendor/IERC165.sol";

contract MaybeRevertMessageReceiver is IAny2EVMMessageReceiver, IERC165 {
  address private s_manager;
  bool public s_toRevert;
  bytes private s_err;
  event MessageReceived();
  error CustomError(bytes err);

  constructor(bool toRevert) {
    s_manager = msg.sender;
    s_toRevert = toRevert;
  }

  function setRevert(bool toRevert) external {
    s_toRevert = toRevert;
  }

  function setErr(bytes memory err) external {
    s_err = err;
  }

  /// @notice IERC165 supports an interfaceId
  /// @param interfaceId The interfaceId to check
  /// @return true if the interfaceId is supported
  function supportsInterface(bytes4 interfaceId) public pure override returns (bool) {
    return interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  function ccipReceive(Client.Any2EVMMessage calldata) external override {
    if (s_toRevert) {
      revert CustomError(s_err);
    }
    emit MessageReceived();
  }
}
