// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../../../interfaces/TypeAndVersionInterface.sol";
import "../../pools/PoolCollector.sol";
import "../../access/OwnerIsCreator.sol";
import "../interfaces/Any2EVMSubscriptionOnRampRouterInterface.sol";

contract EVM2AnySubscriptionOnRampRouter is
  Any2EVMSubscriptionOnRampRouterInterface,
  TypeAndVersionInterface,
  OwnerIsCreator,
  PoolCollector
{
  using SafeERC20 for IERC20;

  string public constant override typeAndVersion = "EVM2AnySubscriptionOnRampRouter 1.0.0";

  // destination chain id => OnRamp
  mapping(uint256 => Any2EVMSubscriptionOnRampInterface) private s_onRamps;
  // A mapping to get the balance of a given subscription
  mapping(address => uint256) private s_subscriptionBalance;

  // The router configuration
  RouterConfig private s_config;

  constructor(RouterConfig memory config) {
    s_config = config;
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnySubscriptionMessage memory message)
    external
    returns (uint64)
  {
    Any2EVMSubscriptionOnRampInterface onRamp = s_onRamps[destinationChainId];
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    if (s_config.fee > 0) {
      s_subscriptionBalance[msg.sender] -= uint256(s_config.fee);
    }

    _collectTokens(onRamp, message.tokens, message.amounts);

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function setOnRamp(uint256 chainId, Any2EVMSubscriptionOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function getOnRamp(uint256 chainId) external view returns (Any2EVMSubscriptionOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function setFee(uint96 newFee) external onlyFeeAdmin {
    s_config.fee = newFee;
    emit FeeSet(newFee);
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function getFee() external view returns (uint96) {
    return s_config.fee;
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function fundSubscription(uint256 amount) external {
    // TODO after spec work revisit this to improve subscriptions
    address sender = msg.sender;
    s_config.feeToken.safeTransferFrom(sender, address(this), amount);
    s_subscriptionBalance[sender] += amount;

    emit SubscriptionFunded(sender, amount);
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function unfundSubscription(uint256 amount) external {
    address sender = msg.sender;
    s_subscriptionBalance[sender] -= amount;
    s_config.feeToken.safeTransfer(sender, amount);

    emit SubscriptionUnfunded(sender, amount);
  }

  /// @inheritdoc Any2EVMSubscriptionOnRampRouterInterface
  function getBalance(address sender) external view returns (uint256 balance) {
    return s_subscriptionBalance[sender];
  }

  // Requires that the function is called by the fee admin.
  modifier onlyFeeAdmin() {
    if (msg.sender != s_config.feeAdmin) revert OnlyCallableByFeeAdmin();
    _;
  }
}
