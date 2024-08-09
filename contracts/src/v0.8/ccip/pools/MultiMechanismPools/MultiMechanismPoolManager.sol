pragma solidity ^0.8.24;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";

abstract contract MultiMechanismPoolManager is OwnerIsCreator {
  mapping(uint64 => bool) internal s_shouldUseAltMech;

  event AltMechanismEnabled(uint64 indexed remoteChainSelector);
  event AltMechanismDisabled(uint64 indexed remoteChainSelector);

  /// @notice Return whether a lane should use the alternative L/R mechanism in the token pool.
  /// @param destChainSelector the source chain the message was sent from
  /// @return bool If the alternative L/R mechanism should be used
  /// @dev Function has been marked virtual and includes an unused calldata parameter in the event that
  /// more complex logic becomes necessary in the future, especially if logic changes between incoming
  /// and outgoing messages.
  function shouldUseAltMechForOutgoingMessage(uint64 destChainSelector) public view virtual returns (bool) {
    return s_shouldUseAltMech[destChainSelector];
  }

  /// @notice Return whether a lane should use the alternative L/R mechanism in the token pool.
  /// @param sourceChainSelector the source chain the message was sent from
  /// @return bool If the alternative L/R mechanism should be used
  /// @dev Function has been marked virtual and includes an unused calldata parameter in the event that
  /// more complex logic becomes necessary in the future, especially if logic changes between incoming
  /// and outgoing messages.
  function shouldUseAltMechForIncomingMessage(
    uint64 sourceChainSelector,
    bytes calldata
  ) public view virtual returns (bool) {
    return s_shouldUseAltMech[sourceChainSelector];
  }

  /// @notice Updates designations for chains on whether to use LR/BM mechanism on CCIP-messages
  /// @param removes A list of chain selectors to disable Lock-Release, and enforce BM
  /// @param adds A list of chain selectors to enable LR instead of BM
  function updateChainSelectorMechanisms(uint64[] calldata removes, uint64[] calldata adds) external onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      delete s_shouldUseAltMech[removes[i]];
      emit AltMechanismDisabled(removes[i]);
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      s_shouldUseAltMech[adds[i]] = true;
      emit AltMechanismEnabled(adds[i]);
    }
  }
}
