// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../access/OwnerIsCreator.sol";
import "./interfaces/BaseOffRampRouterInterface.sol";

contract BaseOffRampRouter is BaseOffRampRouterInterface, OwnerIsCreator {
  // Mapping from offRamp to allowed status
  mapping(BaseOffRampInterface => OffRampDetails) internal s_offRamps;
  // List of all offRamps that have  OffRampDetails
  BaseOffRampInterface[] internal s_offRampsList;

  constructor(BaseOffRampInterface[] memory offRamps) {
    s_offRampsList = offRamps;
    for (uint256 i; i < offRamps.length; ++i) {
      s_offRamps[offRamps[i]] = OffRampDetails({listIndex: uint96(i), allowed: true});
    }
  }

  /// @inheritdoc BaseOffRampRouterInterface
  function addOffRamp(BaseOffRampInterface offRamp) external onlyOwner {
    if (address(offRamp) == address(0)) revert InvalidAddress();
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

  /// @inheritdoc BaseOffRampRouterInterface
  function removeOffRamp(BaseOffRampInterface offRamp) external onlyOwner {
    // Check that there are any feeds to remove
    uint256 listLength = s_offRampsList.length;
    if (listLength == 0) revert NoOffRampsConfigured();

    OffRampDetails memory oldDetails = s_offRamps[offRamp];
    // Check if it exists
    if (!oldDetails.allowed) revert OffRampNotAllowed(offRamp);

    // Swap the last item in the s_offRampsList with the item being removed,
    // update the index of the item moved from the end of the list to its new place,
    // then pop from the end of the list to remove.
    BaseOffRampInterface lastItem = s_offRampsList[listLength - 1];
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

  /// @inheritdoc BaseOffRampRouterInterface
  function getOffRamps() external view returns (BaseOffRampInterface[] memory offRamps) {
    offRamps = s_offRampsList;
  }

  /// @inheritdoc BaseOffRampRouterInterface
  function isOffRamp(BaseOffRampInterface offRamp) external view returns (bool allowed) {
    return s_offRamps[offRamp].allowed;
  }

  // @notice only lets allowed offRamps execute
  modifier onlyOffRamp() {
    BaseOffRampInterface offRamp = BaseOffRampInterface(msg.sender);
    if (!s_offRamps[offRamp].allowed) revert MustCallFromOffRamp(msg.sender);
    _;
  }
}
