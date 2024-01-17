// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import {IMessageTransmitterWithRelay} from "./interfaces/IMessageTransmitterWithRelay.sol";

contract MockUSDCTransmitter is IMessageTransmitterWithRelay {
  // Indicated whether the receiveMessage() call should succeed.
  bool public s_shouldSucceed;
  uint32 private immutable i_version;
  uint32 private immutable i_localDomain;
  // Next available nonce from this source domain
  uint64 public nextAvailableNonce;

  constructor(uint32 version, uint32 localDomain) {
    i_version = version;
    i_localDomain = localDomain;
    s_shouldSucceed = true;
  }

  function receiveMessage(bytes calldata, bytes calldata) external view returns (bool success) {
    return s_shouldSucceed;
  }

  function setShouldSucceed(bool shouldSucceed) external {
    s_shouldSucceed = shouldSucceed;
  }

  function version() external view returns (uint32) {
    return i_version;
  }

  function localDomain() external view returns (uint32) {
    return i_localDomain;
  }

  /**
 * @notice Send the message to the destination domain and recipient
     * @dev Increment nonce, format the message, and emit `MessageSent` event with message information.
     * @param destinationDomain Domain of destination chain
     * @param recipient Address of message recipient on destination chain as bytes32
     * @param messageBody Raw bytes content of message
     * @return nonce reserved by message
     */
  function sendMessage(
    uint32 destinationDomain,
    bytes32 recipient,
    bytes calldata messageBody
  ) external override whenNotPaused returns (uint64) {
    bytes32 _emptyDestinationCaller = bytes32(0);
    uint64 _nonce = _reserveAndIncrementNonce();
    bytes32 _messageSender = bytes32(uint256(uint160((msg.sender))));

    _sendMessage(
      destinationDomain,
      recipient,
      _emptyDestinationCaller,
      _messageSender,
      _nonce,
      messageBody
    );

    return _nonce;
  }

  /**
     * @notice Send the message to the destination domain and recipient, for a specified `destinationCaller` on the
     * destination domain.
     * @dev Increment nonce, format the message, and emit `MessageSent` event with message information.
     * WARNING: if the `destinationCaller` does not represent a valid address, then it will not be possible
     * to broadcast the message on the destination domain. This is an advanced feature, and the standard
     * sendMessage() should be preferred for use cases where a specific destination caller is not required.
     * @param destinationDomain Domain of destination chain
     * @param recipient Address of message recipient on destination domain as bytes32
     * @param destinationCaller caller on the destination domain, as bytes32
     * @param messageBody Raw bytes content of message
     * @return nonce reserved by message
     */
  function sendMessageWithCaller(
    uint32 destinationDomain,
    bytes32 recipient,
    bytes32 destinationCaller,
    bytes calldata messageBody
  ) external override whenNotPaused returns (uint64) {
    require(
      destinationCaller != bytes32(0),
      "Destination caller must be nonzero"
    );

    uint64 _nonce = _reserveAndIncrementNonce();
    bytes32 _messageSender = bytes32(uint256(uint160((msg.sender))));

    _sendMessage(
      destinationDomain,
      recipient,
      destinationCaller,
      _messageSender,
      _nonce,
      messageBody
    );

    return _nonce;
  }

  /**
    * Reserve and increment next available nonce
    * @return nonce reserved
     */
  function _reserveAndIncrementNonce() internal returns (uint64) {
    uint64 _nonceReserved = nextAvailableNonce;
    nextAvailableNonce = nextAvailableNonce + 1;
    return _nonceReserved;
  }
}
