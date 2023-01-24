// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
// solhint-disable-next-line chainlink-solidity/explicit-imports
import "./CCIPConsumer.sol";

import {IERC20} from "../../vendor/IERC20.sol";

/// @title GovernanceDapp - Example of a Governance Dapp using CCIPConsumer
contract GovernanceDapp is CCIPConsumer, TypeAndVersionInterface, OwnerIsCreator {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GovernanceDapp 1.0.0";

  using GEConsumer for GEConsumer.EVMExtraArgsV1;

  error InvalidDeliverer(address deliverer);
  event ConfigPropagated(uint64 chainId, address contractAddress);
  event ReceivedConfig(uint256 feeAmount, uint256 changedAtBlock);

  struct FeeConfig {
    uint256 feeAmount;
    uint256 changedAtBlock;
  }

  struct CrossChainClone {
    uint64 chainId;
    address contractAddress;
  }

  FeeConfig internal s_feeConfig;
  CrossChainClone[] internal s_crossChainClones;

  constructor(
    address router,
    FeeConfig memory feeConfig,
    address feeToken
  ) CCIPConsumer(router, feeToken) {
    s_feeConfig = feeConfig;
  }

  function voteForNewFeeConfig(FeeConfig calldata feeConfig) public onlyOwner {
    // Call for new fee config
    // Count if votes >= threshold
    // if votes passes
    if (getRouter() != address(0)) {
      _propagateFeeConfigChange(feeConfig);
    }
    s_feeConfig = feeConfig;
  }

  function _propagateFeeConfigChange(FeeConfig calldata feeConfig) private {
    bytes memory data = abi.encode(feeConfig);
    uint256 numberOfClones = s_crossChainClones.length;
    for (uint256 i = 0; i < numberOfClones; ++i) {
      CrossChainClone memory clone = s_crossChainClones[i];

      GEConsumer.EVM2AnyGEMessage memory message = GEConsumer.EVM2AnyGEMessage({
        receiver: abi.encode(clone.contractAddress),
        data: data,
        tokensAndAmounts: new Common.EVMTokenAndAmount[](0),
        feeToken: getFeeToken(),
        extraArgs: GEConsumer._argsToBytes(GEConsumer.EVMExtraArgsV1({gasLimit: 3e5, strict: false}))
      });
      _ccipSend(clone.chainId, message);
      emit ConfigPropagated(clone.chainId, clone.contractAddress);
    }
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function _ccipReceive(Common.Any2EVMMessage memory message) internal override {
    FeeConfig memory newFeeConfig = abi.decode(message.data, (FeeConfig));

    s_feeConfig = newFeeConfig;
    emit ReceivedConfig(newFeeConfig.feeAmount, newFeeConfig.changedAtBlock);
  }

  /**
   * @notice Fund this contract with configured feeToken and approve tokens to the router
   * @dev Requires prior approval from the msg.sender
   * @param amount The amount of feeToken to be funded
   */
  function fund(uint256 amount) external {
    IERC20 token = IERC20(getFeeToken());
    token.transferFrom(msg.sender, address(this), amount);
    token.approve(address(getRouter()), amount);
  }

  function addClone(CrossChainClone memory clone) public onlyOwner {
    s_crossChainClones.push(clone);
  }

  function getFeeConfig() external view returns (FeeConfig memory) {
    return s_feeConfig;
  }
}
