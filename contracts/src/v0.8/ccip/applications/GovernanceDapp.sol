// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";
import {Any2EVMMessageReceiverInterface} from "../interfaces/applications/Any2EVMMessageReceiverInterface.sol";
import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
import {Any2EVMOffRampRouterInterface} from "../interfaces/offRamp/Any2EVMOffRampRouterInterface.sol";
import {GERouterInterface} from "../interfaces/router/GERouterInterface.sol";
import {CCIP} from "../models/Models.sol";
import {IERC20} from "../../vendor/IERC20.sol";

contract GovernanceDapp is Any2EVMMessageReceiverInterface, TypeAndVersionInterface, OwnerIsCreator {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GovernanceDapp 1.0.0";

  using CCIP for CCIP.EVMExtraArgsV1;

  error InvalidDeliverer(address deliverer);
  event ConfigPropagated(uint256 chainId, address contractAddress);
  event ReceivedConfig(uint256 feeAmount, uint256 changedAtBlock);

  struct FeeConfig {
    uint256 feeAmount;
    uint256 changedAtBlock;
  }

  struct CrossChainClone {
    uint256 chainId;
    address contractAddress;
  }

  FeeConfig internal s_feeConfig;
  CrossChainClone[] internal s_crossChainClones;

  GERouterInterface internal s_router;

  // The fee token for CCIP billing
  address internal immutable i_feeToken;

  constructor(
    GERouterInterface sendingRouter,
    FeeConfig memory feeConfig,
    address feeToken
  ) {
    s_router = sendingRouter;
    s_feeConfig = feeConfig;
    i_feeToken = feeToken;
  }

  function voteForNewFeeConfig(FeeConfig calldata feeConfig) public onlyOwner {
    // Call for new fee config
    // Count if votes >= threshold
    // if votes passes
    if (s_router != GERouterInterface(address(0))) {
      _propagateFeeConfigChange(feeConfig);
    }
    s_feeConfig = feeConfig;
  }

  function _propagateFeeConfigChange(FeeConfig calldata feeConfig) private {
    bytes memory data = abi.encode(feeConfig);
    uint256 numberOfClones = s_crossChainClones.length;
    for (uint256 i = 0; i < numberOfClones; ++i) {
      CrossChainClone memory clone = s_crossChainClones[i];

      CCIP.EVM2AnyGEMessage memory message = CCIP.EVM2AnyGEMessage({
        receiver: abi.encode(clone.contractAddress),
        data: data,
        tokensAndAmounts: new CCIP.EVMTokenAndAmount[](0),
        feeToken: i_feeToken,
        extraArgs: CCIP.EVMExtraArgsV1({gasLimit: 3e5, strict: false})._toBytes()
      });
      s_router.ccipSend(clone.chainId, message);
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
    emit ReceivedConfig(newFeeConfig.feeAmount, newFeeConfig.changedAtBlock);
  }

  function addClone(CrossChainClone memory clone) public onlyOwner {
    s_crossChainClones.push(clone);
  }

  function setRouters(GERouterInterface router) public {
    s_router = router;
  }

  function getFeeConfig() external view returns (FeeConfig memory) {
    return s_feeConfig;
  }

  /**
   * @dev only calls from the set router are accepted.
   */
  modifier onlyRouter() {
    if (msg.sender != address(s_router)) revert InvalidDeliverer(msg.sender);
    _;
  }
}
