// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../access/OwnerIsCreator.sol";
import "../interfaces/Any2EVMTollOffRampRouterInterface.sol";
import "../BaseOffRampRouter.sol";

contract Any2EVMTollOffRampRouter is BaseOffRampRouter, Any2EVMTollOffRampRouterInterface, TypeAndVersionInterface {
  string public constant override typeAndVersion = "Any2EVMTollOffRampRouter 1.0.0";

  constructor(BaseOffRampInterface[] memory offRamps) BaseOffRampRouter(offRamps) {}

  /**
   * @notice Route the message to its intended receiver contract
   * @param receiver Receiver contract implementing CrossChainMessageReceiverInterface
   * @param message CCIP.Any2EVMTollMessage struct
   */
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Any2EVMTollMessage calldata message)
    external
    override
    onlyOffRamp
  {
    try receiver.ccipReceive(message) {} catch (bytes memory reason) {
      // TODO: use RouterResults and exact gas
      revert MessageFailure(message.sequenceNumber, reason);
    }
  }
}
