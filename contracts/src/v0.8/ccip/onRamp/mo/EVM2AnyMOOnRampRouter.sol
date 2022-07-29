// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../pools/PoolCollector.sol";
import "../interfaces/Any2EVMMOOnRampInterface.sol";
import "../interfaces/Any2EVMMOOnRampRouterInterface.sol";

contract EVM2AnyMOOnRampRouter is Any2EVMMOOnRampRouterInterface, TypeAndVersionInterface, OwnerIsCreator {
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "EVM2AnyMOOnRampRouter 1.0.0";

  // destination chain id => OnRamp
  mapping(uint256 => Any2EVMMOOnRampInterface) private s_onRamps;
  // A mapping to get the balance of a given subscription
  mapping(address => uint256) private s_subscriptionBalance;

  // The router configuration
  RouterConfig private s_config;

  constructor(RouterConfig memory config) {
    s_config = config;
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnyMOMessage memory message) external returns (uint64) {
    // Find and put the correct onRamp on the stack.
    Any2EVMMOOnRampInterface onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    // If fees are enabled, charge the subscription.
    if (s_config.fee > 0) {
      s_subscriptionBalance[msg.sender] -= uint256(s_config.fee);
    }

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function setOnRamp(uint256 chainId, Any2EVMMOOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function getOnRamp(uint256 chainId) external view returns (Any2EVMMOOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function setFee(uint96 newFee) external onlyFeeAdmin {
    s_config.fee = newFee;
    emit FeeSet(newFee);
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function getFee() external view returns (uint96) {
    return s_config.fee;
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function fundSubscription(uint256 amount) external {
    // TODO after spec work revisit this to improve subscriptions
    address sender = msg.sender;
    s_config.feeToken.safeTransferFrom(sender, address(this), amount);
    s_subscriptionBalance[sender] += amount;

    emit SubscriptionFunded(sender, amount);
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function unfundSubscription(uint256 amount) external {
    address sender = msg.sender;
    s_subscriptionBalance[sender] -= amount;
    s_config.feeToken.safeTransfer(sender, amount);

    emit SubscriptionUnfunded(sender, amount);
  }

  /// @inheritdoc Any2EVMMOOnRampRouterInterface
  function getBalance(address sender) external view returns (uint256 balance) {
    return s_subscriptionBalance[sender];
  }

  // Requires that the function is called by the fee admin.
  modifier onlyFeeAdmin() {
    if (msg.sender != s_config.feeAdmin) revert OnlyCallableByFeeAdmin();
    _;
  }
}
