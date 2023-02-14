// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import {TypeAndVersionInterface} from "../../interfaces/TypeAndVersionInterface.sol";

import {OwnerIsCreator} from "../access/OwnerIsCreator.sol";
// solhint-disable-next-line chainlink-solidity/explicit-imports
import {CCIPReceiver} from "./CCIPReceiver.sol";

import {IERC20} from "../../vendor/IERC20.sol";
import {Client} from "../models/Client.sol";
import {IRouterClient} from "../interfaces/router/IRouterClient.sol";
import {Client} from "../models/Client.sol";

/// @title GovernanceDapp - Example of a Governance Dapp using CCIPReceiver
contract GovernanceDapp is CCIPReceiver, TypeAndVersionInterface, OwnerIsCreator {
  // solhint-disable-next-line chainlink-solidity/all-caps-constant-storage-variables
  string public constant override typeAndVersion = "GovernanceDapp 1.0.0";

  using Client for Client.EVMExtraArgsV1;

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
  IERC20 private s_feeToken;
  CrossChainClone[] internal s_crossChainClones;

  constructor(
    address router,
    FeeConfig memory feeConfig,
    IERC20 feeToken
  ) CCIPReceiver(router) {
    s_feeConfig = feeConfig;
    s_feeToken = feeToken;
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

      Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
        receiver: abi.encode(clone.contractAddress),
        data: data,
        tokenAmounts: new Client.EVMTokenAmount[](0),
        feeToken: address(s_feeToken),
        extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 3e5, strict: false}))
      });
      IRouterClient(getRouter()).ccipSend(clone.chainId, message);
      emit ConfigPropagated(clone.chainId, clone.contractAddress);
    }
  }

  /**
   * @notice Called by the OffRamp, this function receives a message and forwards
   * the tokens sent with it to the designated EOA
   * @param message CCIP Message
   */
  function _ccipReceive(Client.Any2EVMMessage memory message) internal override {
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
    s_feeToken.transferFrom(msg.sender, address(this), amount);
    s_feeToken.approve(address(getRouter()), amount);
  }

  function addClone(CrossChainClone memory clone) public onlyOwner {
    s_crossChainClones.push(clone);
  }

  function getFeeConfig() external view returns (FeeConfig memory) {
    return s_feeConfig;
  }
}
