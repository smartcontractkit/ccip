// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../utils/CCIP.sol";
import "../BaseOnRamp.sol";
import "../interfaces/Any2EVMTollOnRampInterface.sol";
import "../interfaces/Any2EVMTollOnRampRouterInterface.sol";

/**
 * @notice An implementation of a toll OnRamp.
 */
contract EVM2EVMTollOnRamp is Any2EVMTollOnRampInterface, BaseOnRamp, TypeAndVersionInterface {
  using Address for address;
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "EVM2EVMTollOnRamp 1.0.0";

  constructor(
    uint256 chainId,
    uint256 destinationChainId,
    IERC20[] memory tokens,
    PoolInterface[] memory pools,
    AggregatorV2V3Interface[] memory feeds,
    address[] memory allowlist,
    AFNInterface afn,
    uint256 maxTimeWithoutAFNSignal,
    OnRampConfig memory config,
    Any2EVMTollOnRampRouterInterface router
  )
    BaseOnRamp(
      chainId,
      destinationChainId,
      tokens,
      pools,
      feeds,
      allowlist,
      afn,
      maxTimeWithoutAFNSignal,
      config,
      address(router)
    )
  {}

  /// @inheritdoc Any2EVMTollOnRampInterface
  function forwardFromRouter(CCIP.EVM2AnyTollMessage memory message, address originalSender)
    external
    override
    whenNotPaused
    whenHealthy
    returns (uint64)
  {
    if (msg.sender != address(s_router)) revert MustBeCalledByRouter();
    handleForwardFromRouter(message.data.length, message.tokens, message.amounts, originalSender);

    // Emit message request
    // we need the next available sequence number so we increment before we use the value
    CCIP.EVM2EVMTollEvent memory tollEvent = CCIP.EVM2EVMTollEvent({
      sequenceNumber: ++s_sequenceNumber,
      sourceChainId: CHAIN_ID,
      sender: originalSender,
      receiver: message.receiver,
      data: message.data,
      tokens: message.tokens,
      amounts: message.amounts,
      feeToken: message.feeToken,
      feeTokenAmount: message.feeTokenAmount,
      gasLimit: message.gasLimit
    });
    emit CCIPSendRequested(tollEvent);
    return tollEvent.sequenceNumber;
  }

  /// @inheritdoc Any2EVMTollOnRampInterface
  function getRequiredFee(IERC20 feeToken) public view override returns (uint256) {
    AggregatorV2V3Interface feed = getFeed(feeToken);
    if (address(feed) == address(0)) revert UnsupportedFeeToken(feeToken);
    return s_config.relayingFeeJuels * uint256(feed.latestAnswer());
  }
}
