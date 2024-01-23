// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../../../pools/USDC/IMessageTransmitter.sol";

// This follows https://github.com/circlefin/evm-cctp-contracts/blob/master/src/interfaces/IMessageTransmitter.sol
interface IMessageTransmitterWithRelay is IMessageTransmitter {
  /**
   * @notice Sends an outgoing message from the source domain.
   * @dev Increment nonce, format the message, and emit `MessageSent` event with message information.
   * @param destinationDomain Domain of destination chain
   * @param recipient Address of message recipient on destination domain as bytes32
   * @param messageBody Raw bytes content of message
   * @return nonce reserved by message
   */
  function sendMessage(
    uint32 destinationDomain,
    bytes32 recipient,
    bytes calldata messageBody
  ) external returns (uint64);

  /**
   * @notice Sends an outgoing message from the source domain, with a specified caller on the
   * destination domain.
   * @dev Increment nonce, format the message, and emit `MessageSent` event with message information.
   * WARNING: if the `destinationCaller` does not represent a valid address as bytes32, then it will not be possible
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
  ) external returns (uint64);
}
