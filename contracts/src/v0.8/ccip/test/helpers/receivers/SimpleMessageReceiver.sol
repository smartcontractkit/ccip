// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../../../interfaces/IAny2EVMMessageReceiver.sol";

import "../../../../vendor/openzeppelin-solidity/v4.8.0/utils/introspection/IERC165.sol";

contract SimpleMessageReceiver is IAny2EVMMessageReceiver, IERC165 {
  event MessageReceived();

  address private immutable i_manager;

  constructor() {
    i_manager = msg.sender;
  }

  /// @notice IERC165 supports an interfaceId
  /// @param interfaceId The interfaceId to check
  /// @return true if the interfaceId is supported
  function supportsInterface(bytes4 interfaceId) public pure override returns (bool) {
    return interfaceId == type(IAny2EVMMessageReceiver).interfaceId || interfaceId == type(IERC165).interfaceId;
  }

  function ccipReceive(Client.Any2EVMMessage calldata) external override {
    emit MessageReceived();
  }

  function getManager() external view returns (address) {
    return i_manager;
  }
}
