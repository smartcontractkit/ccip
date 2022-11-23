// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {Any2EVMMessageReceiverInterface} from "../interfaces/applications/Any2EVMMessageReceiverInterface.sol";
import {EVM2AnySubscriptionOnRampRouterInterface, CCIP, IERC20} from "../interfaces/onRamp/EVM2AnySubscriptionOnRampRouterInterface.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Any2EVMOffRampRouterInterface} from "../interfaces/offRamp/Any2EVMOffRampRouterInterface.sol";

contract GovernanceDapp is Any2EVMMessageReceiverInterface, TypeAndVersionInterface, OwnerIsCreator {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GovernanceDapp 1.0.0";

  using CCIP for CCIP.EVMExtraArgsV1;

  error InvalidDeliverer(address deliverer);
  event ConfigPropagated(uint256 chainId, address contractAddress);
  event ReceivedConfig(uint256 feeAmount, address subscriptionManager, uint256 changedAtBlock);

  struct FeeConfig {
    uint256 feeAmount;
    address subscriptionManager;
    uint256 changedAtBlock;
  }

  struct CrossChainClone {
    uint256 chainId;
    address contractAddress;
  }

  FeeConfig internal s_feeConfig;
  CrossChainClone[] internal s_crossChainClones;

  Any2EVMOffRampRouterInterface internal s_receivingRouter;
  EVM2AnySubscriptionOnRampRouterInterface internal s_sendingRouter;

  constructor(
    Any2EVMOffRampRouterInterface receivingRouter,
    EVM2AnySubscriptionOnRampRouterInterface sendingRouter,
    FeeConfig memory feeConfig
  ) {
    s_receivingRouter = receivingRouter;
    s_sendingRouter = sendingRouter;
    s_feeConfig = feeConfig;
  }

  function voteForNewFeeConfig(FeeConfig calldata feeConfig) public onlyOwner {
    // Call for new fee config
    // Count if votes >= threshold
    // if votes passes
    if (s_sendingRouter != EVM2AnySubscriptionOnRampRouterInterface(address(0))) {
      _propagateFeeConfigChange(feeConfig);
    }
    s_feeConfig = feeConfig;
  }

  function _propagateFeeConfigChange(FeeConfig calldata feeConfig) private {
    bytes memory data = abi.encode(feeConfig);
    uint256 numberOfClones = s_crossChainClones.length;
    for (uint256 i = 0; i < numberOfClones; ++i) {
      CrossChainClone memory clone = s_crossChainClones[i];

      CCIP.EVM2AnySubscriptionMessage memory message = CCIP.EVM2AnySubscriptionMessage({
        receiver: abi.encode(clone.contractAddress),
        data: data,
        tokensAndAmounts: new CCIP.EVMTokenAndAmount[](0),
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: 3e5, strict: false})._toBytes()
      });
      s_sendingRouter.ccipSend(clone.chainId, message);
      emit ConfigPropagated(clone.chainId, clone.contractAddress);
    }
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function ccipReceive(CCIP.Any2EVMMessage memory message) external override onlyRouter {
    FeeConfig memory newFeeConfig = abi.decode(message.data, (FeeConfig));

    s_feeConfig = newFeeConfig;
    emit ReceivedConfig(newFeeConfig.feeAmount, newFeeConfig.subscriptionManager, newFeeConfig.changedAtBlock);
  }

  function addClone(CrossChainClone memory clone) public onlyOwner {
    s_crossChainClones.push(clone);
  }

  function setRouters(
    Any2EVMOffRampRouterInterface receivingRouter,
    EVM2AnySubscriptionOnRampRouterInterface sendingRouter
  ) public {
    s_receivingRouter = receivingRouter;
    s_sendingRouter = sendingRouter;
  }

  function getSubscriptionManager() external view returns (address) {
    return s_feeConfig.subscriptionManager;
  }

  function getFeeConfig() external view returns (FeeConfig memory) {
    return s_feeConfig;
  }

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_receivingRouter)) revert InvalidDeliverer(msg.sender);
    _;
  }
}
