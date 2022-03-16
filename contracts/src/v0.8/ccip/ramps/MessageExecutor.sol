// SPDX-License-Identifier: MIT
pragma solidity 0.8.12;

import "../interfaces/OffRampInterface.sol";
import "../../interfaces/TypeAndVersionInterface.sol";
import "../ocr/OCR2Base.sol";
import "../utils/CCIP.sol";
import "../../vendor/SafeERC20.sol";

/**
 * @notice MessageExecutor enables OCR networks to execute multiple messages
 * in an OffRamp in a single transaction.
 */
contract MessageExecutor is TypeAndVersionInterface, OCR2Base {
  using SafeERC20 for IERC20;

  event FeesWithdrawn(IERC20 feeToken, address recipient, uint256 amount);

  /// @notice Message and its proof
  struct ExecutableMessage {
    // TODO: We have to split to MerkleProof up here into its individual parts, and also order
    // the items here to avoid a stack too deep error. This needs investigation.
    bytes32[] path;
    uint256 index;
    CCIP.Message message;
  }

  OffRampInterface private immutable s_offRamp;
  bool private s_needFee;

  constructor(OffRampInterface offRamp, bool needFee) OCR2Base(true) {
    s_offRamp = offRamp;
    s_needFee = needFee;
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
      s_offRamp.executeTransaction(em.message, CCIP.MerkleProof({path: em.path, index: em.index}), s_needFee);
    }
  }

  /**
   * @notice TODO Withraw function that will be removed once transmitter renumeration is implemented
   */
  function withdrawAccumulatedFees(
    IERC20 feeToken,
    address recipient,
    uint256 amount
  ) external onlyOwner {
    feeToken.safeTransfer(recipient, amount);
    emit FeesWithdrawn(feeToken, recipient, amount);
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

  function setNeedFee(bool flag) external onlyOwner {
    s_needFee = flag;
  }

  function getNeedFee() external view returns (bool) {
    return s_needFee;
  }

  function getOffRamp() external view returns (OffRampInterface) {
    return s_offRamp;
  }

  function typeAndVersion() external pure override returns (string memory) {
    return "MessageExecutor 1.0.0";
  }
}
