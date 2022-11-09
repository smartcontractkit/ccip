// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../../interfaces/TypeAndVersionInterface.sol";
import {EVM2AnySubscriptionOnRampRouterInterface, EVM2EVMSubscriptionOnRampInterface, BaseOnRampRouterInterface} from "../../interfaces/onRamp/EVM2AnySubscriptionOnRampRouterInterface.sol";
import {BaseOnRampInterface} from "../../interfaces/onRamp/BaseOnRampInterface.sol";
import {PoolCollector} from "../../pools/PoolCollector.sol";
import {OwnerIsCreator} from "../../access/OwnerIsCreator.sol";
import {CCIP} from "../../models/Models.sol";
import {SafeERC20, IERC20} from "../../../vendor/SafeERC20.sol";

contract EVM2AnySubscriptionOnRampRouter is
  EVM2AnySubscriptionOnRampRouterInterface,
  TypeAndVersionInterface,
  OwnerIsCreator,
  PoolCollector
{
  using SafeERC20 for IERC20;
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "EVM2AnySubscriptionOnRampRouter 1.0.0";

  // destination chain id => OnRamp
  mapping(uint256 => EVM2EVMSubscriptionOnRampInterface) private s_onRamps;
  // A mapping to get the balance of a given subscription
  mapping(address => uint256) private s_subscriptionBalance;

  // The router configuration
  RouterConfig private s_config;

  constructor(RouterConfig memory config) {
    s_config = config;
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function ccipSend(uint256 destinationChainId, CCIP.EVM2AnySubscriptionMessage calldata message)
    external
    returns (uint64)
  {
    // Find and put the correct onRamp on the stack.
    EVM2EVMSubscriptionOnRampInterface onRamp = s_onRamps[destinationChainId];
    // Check if the onRamp is a zero address, meaning the chain is not supported.
    if (address(onRamp) == address(0)) revert UnsupportedDestinationChain(destinationChainId);
    // If fees are enabled, charge the subscription.
    if (s_config.fee > 0) {
      s_subscriptionBalance[msg.sender] -= uint256(s_config.fee);
    }

    // Transfer the tokens to the token pools.
    _collectTokens(onRamp, message.tokensAndAmounts);

    return onRamp.forwardFromRouter(message, msg.sender);
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function setOnRamp(uint256 chainId, EVM2EVMSubscriptionOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) == address(onRamp)) revert OnRampAlreadySet(chainId, onRamp);
    s_onRamps[chainId] = onRamp;
    emit OnRampSet(chainId, onRamp);
  }

  function removeOnRamp(uint256 chainId, EVM2EVMSubscriptionOnRampInterface onRamp) external onlyOwner {
    if (address(s_onRamps[chainId]) != address(onRamp))
      revert WrongOnRamp(address(onRamp), address(s_onRamps[chainId]));
    delete s_onRamps[chainId];
    emit OnRampRemoved(chainId, onRamp);
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function getOnRamp(uint256 chainId) external view returns (EVM2EVMSubscriptionOnRampInterface) {
    return s_onRamps[chainId];
  }

  /// @inheritdoc BaseOnRampRouterInterface
  function isChainSupported(uint256 chainId) external view returns (bool supported) {
    return address(s_onRamps[chainId]) != address(0);
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function setFee(uint96 newFee) external onlyFeeAdmin {
    s_config.fee = newFee;
    emit FeeSet(newFee);
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function getFee() external view returns (uint96) {
    return s_config.fee;
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function fundSubscription(uint256 amount) external {
    // TODO after spec work revisit this to improve subscriptions
    address sender = msg.sender;
    s_config.feeToken.safeTransferFrom(sender, address(this), amount);
    s_subscriptionBalance[sender] += amount;

    emit SubscriptionFunded(sender, amount);
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function unfundSubscription(uint256 amount) external {
    address sender = msg.sender;
    s_subscriptionBalance[sender] -= amount;
    s_config.feeToken.safeTransfer(sender, amount);

    emit SubscriptionUnfunded(sender, amount);
  }

  /// @inheritdoc EVM2AnySubscriptionOnRampRouterInterface
  function getBalance(address sender) external view returns (uint256 balance) {
    return s_subscriptionBalance[sender];
  }

  // Requires that the function is called by the fee admin.
  modifier onlyFeeAdmin() {
    if (msg.sender != s_config.feeAdmin) revert OnlyCallableByFeeAdmin();
    _;
  }
}
