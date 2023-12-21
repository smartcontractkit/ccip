// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IL1Bridge} from "./IBridge.sol";
import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";

import {IL1GatewayRouter} from "@arbitrum/token-bridge-contracts/contracts/tokenbridge/ethereum/gateway/IL1GatewayRouter.sol";
import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Arbitrum L1 Bridge adapter
/// @dev Auto unwraps and re-wraps wrapped eth in the bridge.
contract ArbitrumL1BridgeAdapter is IL1Bridge {
  using SafeERC20 for IERC20;

  error InsufficientEthValue(uint256 wanted, uint256 got);

  IL1GatewayRouter internal immutable i_l1GatewayRouter;
  address internal immutable i_l1ERC20Gateway;
  address internal immutable i_l1Inbox;
  address internal immutable i_l1Outbox;

  // TODO not static?
  uint256 public constant MAX_GAS = 100_000;
  uint256 public constant GAS_PRICE_BID = 300_000_000;
  uint256 public constant MAX_SUBMISSION_COST = 8e14;

  // Nonce to use for L2 deposits to allow for better tracking offchain.
  uint64 private s_nonce = 0;

  constructor(IL1GatewayRouter l1GatewayRouter, address l1Inbox, address l1Outbox, address l1ERC20Gateway) {
    if (
      address(l1GatewayRouter) == address(0) ||
      address(l1Inbox) == address(0) ||
      address(l1Outbox) == address(0) ||
      address(l1ERC20Gateway) == address(0)
    ) {
      revert BridgeAddressCannotBeZero();
    }
    i_l1GatewayRouter = l1GatewayRouter;
    i_l1Inbox = l1Inbox;
    i_l1Outbox = l1Outbox;
    i_l1ERC20Gateway = l1ERC20Gateway;
  }

  function depositERC20ToL2(address l1Token, address, address recipient, uint256 amount) external payable {
    IERC20(l1Token).safeTransferFrom(msg.sender, address(this), amount);

    IERC20(l1Token).approve(i_l1ERC20Gateway, amount);

    uint256 wantedNativeFeeCoin = MAX_SUBMISSION_COST + MAX_GAS * GAS_PRICE_BID;
    if (msg.value < wantedNativeFeeCoin) {
      revert InsufficientEthValue(wantedNativeFeeCoin, msg.value);
    }

    i_l1GatewayRouter.outboundTransferCustomRefund{value: wantedNativeFeeCoin}(
      l1Token,
      recipient,
      recipient,
      amount,
      MAX_GAS,
      GAS_PRICE_BID,
      abi.encode(MAX_SUBMISSION_COST, bytes(""))
    );
  }

  function depositNativeToL2(address recipient, uint256 amount) public payable {
    // TODO
    //    i_L1Bridge.depositETHTo{value: amount}(recipient, 0, abi.encode(s_nonce++));
  }

  function finalizeWithdrawERC20FromL2(
    address l1Token,
    address,
    address from,
    address to,
    uint256 amount,
    bytes calldata data
  ) external {
    i_l1GatewayRouter.finalizeInboundTransfer(l1Token, from, to, amount, data);
  }

  function finalizeWithdrawNativeFromL2(address from, address to, uint256 amount, bytes calldata data) external {
    // TODO
    // Outbox.executeTransaction
  }

  //  function relayMessageFromL2ToL1(
  //    address target,
  //    address sender,
  //    bytes memory message,
  //    uint256 messageNonce,
  //    IL1CrossDomainMessenger.L2MessageInclusionProof memory proof
  //  ) external {
  //    i_L1CrossDomainMessenger.relayMessage(target, sender, message, messageNonce, proof);
  //    // TODO
  //  }

  function getL1Bridge() external view returns (address) {
    return address(i_l1GatewayRouter);
  }

  function getL2Token(address l1Token) external view returns (address) {
    return i_l1GatewayRouter.calculateL2TokenAddress(l1Token);
  }
}
