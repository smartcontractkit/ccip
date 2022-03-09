// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../interfaces/OffRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";

/**
 * @notice MessageExecutor enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract MessageExecutor is TypeAndVersionInterface, OCR2Base {
  /// @notice Message and its proof
  struct ExecutableMessage {
    // TODO: We have to split to MerkleProof up here into its individual parts, and also order
    // the items here to avoid a stack too deep error. This needs investigation.
    bytes32[] path;
    uint256 index;
    CCIP.Message message;
  }

  OffRampInterface public immutable s_offRamp;

  constructor(OffRampInterface offRamp) OCR2Base(true) {
    s_offRamp = offRamp;
  }

  /**
   * @notice Entry point for execution, called by the OCR network
   * @dev Expects an encoded array of ExectableMessage tuples.
   */
  function _report(
    bytes32, /*configDigest*/
    uint40, /*epochAndRound*/
    bytes memory report
  ) internal override {
    ExecutableMessage[] memory executableMessages = abi.decode(report, (ExecutableMessage[]));
    for (uint256 i = 0; i < executableMessages.length; i++) {
      ExecutableMessage memory em = executableMessages[i];
      s_offRamp.executeTransaction(em.message, CCIP.MerkleProof({path: em.path, index: em.index}), false);
    }
  }

  function _beforeSetConfig(uint8 _threshold, bytes memory _onchainConfig) internal override {
    // TODO
  }

  function _afterSetConfig(
    uint8, /* f */
    bytes memory /* onchainConfig */
  ) internal override {
    // TODO
  }

  function _payTransmitter(uint32 initialGas, address transmitter) internal override {
    // TODO
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "MessageExecutor 1.0.0";
  }
}
