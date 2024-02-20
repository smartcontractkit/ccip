// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBridgeAdapter} from "../interfaces/IBridge.sol";

import {IL1GatewayRouter} from "@arbitrum/token-bridge-contracts/contracts/tokenbridge/ethereum/gateway/IL1GatewayRouter.sol";
import {IGatewayRouter} from "@arbitrum/token-bridge-contracts/contracts/tokenbridge/libraries/gateway/IGatewayRouter.sol";
import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

interface IOutbox {
  /**
   * @notice Executes a messages in an Outbox entry.
   * @dev Reverts if dispute period hasn't expired, since the outbox entry
   *      is only created once the rollup confirms the respective assertion.
   * @dev it is not possible to execute any L2-to-L1 transaction which contains data
   *      to a contract address without any code (as enforced by the Bridge contract).
   * @param proof Merkle proof of message inclusion in send root
   * @param index Merkle path to message
   * @param l2Sender sender if original message (i.e., caller of ArbSys.sendTxToL1)
   * @param to destination address for L1 contract call
   * @param l2Block l2 block number at which sendTxToL1 call was made
   * @param l1Block l1 block number at which sendTxToL1 call was made
   * @param l2Timestamp l2 Timestamp at which sendTxToL1 call was made
   * @param value wei in L1 message
   * @param data abi-encoded L1 message data
   */
  function executeTransaction(
    bytes32[] calldata proof,
    uint256 index,
    address l2Sender,
    address to,
    uint256 l2Block,
    uint256 l1Block,
    uint256 l2Timestamp,
    uint256 value,
    bytes calldata data
  ) external;
}

/// @notice Arbitrum L1 Bridge adapter
/// @dev Auto unwraps and re-wraps wrapped eth in the bridge.
contract ArbitrumL1BridgeAdapter is IBridgeAdapter {
  using SafeERC20 for IERC20;

  IL1GatewayRouter internal immutable i_l1GatewayRouter;
  IOutbox internal immutable i_l1Outbox;

  // Nonce to use for L2 deposits to allow for better tracking offchain.
  // TODO: increment and emit event w/ nonce
  uint64 private s_nonce = 0;

  error NoGatewayForToken(address token);
  error Unimplemented();

  /// @notice event to track the finalization of the L2 to L1 transfer offchain
  /// @dev while this bridge adapter is trustless and anyone can use it,
  /// @dev its highly unlikely that anyone would prefer it over the official bridge
  /// @dev contracts. And since the official bridge contracts probably have lots of
  /// @dev events and logs, we can use this event to track the L2 to L1 transfers
  /// @dev with less load on the rebalancer oracles.
  event ArbitrumL2ToL1ERC20Finalized(
    address indexed remoteSender,
    address indexed localReceiver,
    uint256 amount,
    ArbitrumFinalizationPayload payload
  );
  /// @notice event to track the an L1 to L2 transfer offchain
  /// @param localToken the token address on L1
  /// @param remoteToken the token address on L2
  /// @param recipient the recipient of the tokens on L1
  /// @param nonce the nonce of the transfer
  /// @param amount the amount of tokens transferred
  /// @param outboundTransferResult the result of the outbound transfer, which is the unique id used to identify the L1 to L2 tx
  event ArbitrumL1ToL2ERC20Sent(
    address indexed localToken,
    address indexed remoteToken,
    address indexed recipient,
    uint256 nonce,
    uint256 amount,
    bytes outboundTransferResult
  );

  constructor(IL1GatewayRouter l1GatewayRouter, IOutbox l1Outbox) {
    if (address(l1GatewayRouter) == address(0) || address(l1Outbox) == address(0)) {
      revert BridgeAddressCannotBeZero();
    }
    i_l1GatewayRouter = l1GatewayRouter;
    i_l1Outbox = l1Outbox;
  }

  /// @dev these are parameters provided by the caller of the sendERC20 function
  /// @dev these must be determined offchain.
  struct SendERC20Params {
    uint256 gasLimit;
    uint256 maxSubmissionCost;
    uint256 maxFeePerGas;
  }

  /// @inheritdoc IBridgeAdapter
  function sendERC20(
    address localToken,
    address /* remoteToken */,
    address recipient,
    uint256 amount,
    bytes calldata bridgeSpecificPayload
  ) external payable override returns (bytes memory) {
    // receive the token transfer from the msg.sender
    IERC20(localToken).safeTransferFrom(msg.sender, address(this), amount);

    // Note: the gateway router could return 0x0 for the gateway address
    // if that token is not yet registered
    address gateway = IGatewayRouter(address(i_l1GatewayRouter)).getGateway(localToken);
    if (gateway == address(0)) {
      revert NoGatewayForToken(localToken);
    }

    // approve the gateway to transfer the token amount sent to the adapter
    IERC20(localToken).safeApprove(gateway, amount);

    SendERC20Params memory params = abi.decode(bridgeSpecificPayload, (SendERC20Params));

    uint256 expectedMsgValue = (params.gasLimit * params.maxFeePerGas) + params.maxSubmissionCost;
    if (msg.value < expectedMsgValue) {
      revert MsgValueDoesNotMatchAmount(msg.value, expectedMsgValue);
    }

    // The router will route the call to the gateway that we approved
    // above. The gateway will then transfer the tokens to the L2.
    // outboundTransferCustomRefund will return the abi encoded inbox sequence number
    // which is 256 bits, so we can cap the return data to 256 bits.
    bytes memory inboxSequenceNumber = i_l1GatewayRouter.outboundTransferCustomRefund{value: msg.value}(
      localToken,
      recipient,
      recipient,
      amount,
      params.gasLimit,
      params.maxFeePerGas,
      abi.encode(params.maxSubmissionCost, bytes(""))
    );

    emit ArbitrumL1ToL2ERC20Sent(localToken, localToken, recipient, ++s_nonce, amount, inboxSequenceNumber);

    return inboxSequenceNumber;
  }

  /// @dev This function is so that we can easily abi-encode the arbitrum-specific
  /// @dev payload for the sendERC20 function.
  function exposeSendERC20Params(SendERC20Params memory params) public pure {}

  /// @dev fees have to be determined offchain for arbitrum
  /// @dev therefore revert here to discourage usage
  function getBridgeFeeInNative() public pure override returns (uint256) {
    revert Unimplemented();
  }

  /// @param proof Merkle proof of message inclusion in send root
  /// @param index Merkle path to message
  /// @param l2Sender sender if original message (i.e., caller of ArbSys.sendTxToL1)
  /// @param to destination address for L1 contract call
  /// @param l2Block l2 block number at which sendTxToL1 call was made
  /// @param l1Block l1 block number at which sendTxToL1 call was made
  /// @param l2Timestamp l2 Timestamp at which sendTxToL1 call was made
  /// @param value wei in L1 message
  /// @param data abi-encoded L1 message data
  struct ArbitrumFinalizationPayload {
    bytes32[] proof;
    uint256 index;
    address l2Sender;
    address to;
    uint256 l2Block;
    uint256 l1Block;
    uint256 l2Timestamp;
    uint256 value;
    bytes data;
  }

  /// @dev This function is so that we can easily abi-encode the arbitrum-specific
  /// @dev payload for the finalizeWithdrawERC20 function.
  function exposeArbitrumFinalizationPayload(ArbitrumFinalizationPayload memory payload) public pure {}

  /// @notice Finalize an L2 -> L1 transfer.
  function finalizeWithdrawERC20(
    address /* remoteSender */,
    address /* localReceiver */,
    bytes calldata arbitrumFinalizationPayload
  ) external {
    ArbitrumFinalizationPayload memory payload = abi.decode(arbitrumFinalizationPayload, (ArbitrumFinalizationPayload));
    i_l1Outbox.executeTransaction(
      payload.proof,
      payload.index,
      payload.l2Sender,
      payload.to,
      payload.l2Block,
      payload.l1Block,
      payload.l2Timestamp,
      payload.value,
      payload.data
    );
    emit ArbitrumL2ToL1ERC20Finalized(payload.l2Sender, payload.to, payload.value, payload);
  }

  function getL2Token(address l1Token) external view returns (address) {
    return i_l1GatewayRouter.calculateL2TokenAddress(l1Token);
  }
}
