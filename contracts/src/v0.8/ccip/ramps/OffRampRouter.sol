// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "../../vendor/SafeERC20.sol";
import "../access/OwnerIsCreator.sol";
import "../interfaces/OffRampInterface.sol";
import "../interfaces/OffRampRouterInterface.sol";

contract OffRampRouter is OffRampRouterInterface, OwnerIsCreator {
  using Address for address;
  using SafeERC20 for IERC20;

  error OffRampNotConfigured(OffRampInterface offRamp);
  error AlreadyConfigured(OffRampInterface offRamp);
  error NoOffRampsConfigured();
  error MessageFailure(uint256 sequenceNumber, bytes reason);

  event OffRampAdded(OffRampInterface indexed offRamp);
  event OffRampRemoved(OffRampInterface indexed offRamp);

  struct OffRampDetails {
    uint96 listIndex;
    bool allowed;
  }

  // OffRamp allowlist mapping
  mapping(OffRampInterface => OffRampDetails) private s_offRamps;
  // OffRamp list
  OffRampInterface[] private s_offRampsList;

  constructor(OffRampInterface[] memory offRamps) {
    s_offRampsList = offRamps;
    for (uint256 i; i < offRamps.length; i++) {
      s_offRamps[offRamps[i]] = OffRampDetails({listIndex: uint96(i), allowed: true});
    }
  }

  /**
   * @notice Route the message to its intended receiver contract
   * @param receiver Receiver contract implementing CrossChainMessageReceiverInterface
   * @param message CCIP.Message struct
   */
  function routeMessage(CrossChainMessageReceiverInterface receiver, CCIP.Message calldata message)
    external
    override
    onlyOffRamp
  {
    try receiver.receiveMessage(message) {} catch (bytes memory reason) {
      revert MessageFailure(message.sequenceNumber, reason);
    }
  }

  /**
   * @notice Owner can add an offRamp from the allowlist
   * @dev Onlw callable by the owner
   * @param offRamp The offRamp to add
   */
  function addOffRamp(OffRampInterface offRamp) external onlyOwner {
    OffRampDetails memory details = s_offRamps[offRamp];
    // Check if the offramp is already allowed
    if (details.allowed) revert AlreadyConfigured(offRamp);

    // Set the s_offRamps with the new offRamp
    details.allowed = true;
    details.listIndex = uint96(s_offRampsList.length);
    s_offRamps[offRamp] = details;

    // Add to the s_offRampsList
    s_offRampsList.push(offRamp);

    emit OffRampAdded(offRamp);
  }

  /**
   * @notice Owner can remove a speicific offRamp from the allowlist
   * @dev Onlw callable by the owner
   * @param offRamp The offRamp to remove
   */
  function removeOffRamp(OffRampInterface offRamp) external onlyOwner {
    // Check that there are any feeds to remove
    uint256 listLength = s_offRampsList.length;
    if (listLength == 0) revert NoOffRampsConfigured();

    OffRampDetails memory oldDetails = s_offRamps[offRamp];
    // Check if it exists
    if (!oldDetails.allowed) revert OffRampNotConfigured(offRamp);

    // Swap the last item in the s_offRampsList with the item being removed,
    // update the index of the item moved from the end of the list to its new place,
    // then pop from the end of the list to remove.
    OffRampInterface lastItem = s_offRampsList[listLength - 1];
    // Perform swap
    s_offRampsList[listLength - 1] = s_offRampsList[oldDetails.listIndex];
    s_offRampsList[oldDetails.listIndex] = lastItem;
    // Update listIndex on moved item
    s_offRamps[lastItem].listIndex = oldDetails.listIndex;
    // Pop from list and delete from mapping
    s_offRampsList.pop();
    delete s_offRamps[offRamp];

    emit OffRampRemoved(offRamp);
  }

  function getOffRamps() external view returns (OffRampInterface[] memory offRamps) {
    offRamps = s_offRampsList;
  }

  function isOffRamp(OffRampInterface offRamp) external view returns (bool allowed) {
    return s_offRamps[offRamp].allowed;
  }

  modifier onlyOffRamp() {
    OffRampInterface offRamp = OffRampInterface(msg.sender);
    if (!s_offRamps[offRamp].allowed) revert OffRampNotConfigured(offRamp);
    _;
  }
}
