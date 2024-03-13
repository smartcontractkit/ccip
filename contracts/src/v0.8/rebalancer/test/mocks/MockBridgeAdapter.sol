// SPDX-License-Identifier: BUSL-1.1
// solhint-disable one-contract-per-file
pragma solidity ^0.8.0;

import {IBridgeAdapter} from "../../interfaces/IBridge.sol";
import {ILiquidityContainer} from "../../interfaces/ILiquidityContainer.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

/// @notice Mock multiple-stage finalization bridge adapter implementation.
/// @dev Funds are only made available after both the prove and finalization steps are completed.
/// @dev Sends the L1 tokens from the msg sender to address(this).
contract MockL1BridgeAdapter is IBridgeAdapter, ILiquidityContainer {
  using SafeERC20 for IERC20;

  error InsufficientLiquidity();
  error NonceAlreadyUsed(uint256 nonce);
  error InvalidFinalizationAction();

  IERC20 internal immutable i_token;
  uint256 internal s_nonce = 1;
  mapping(uint256 => bool) internal s_nonceProven;
  mapping(uint256 => bool) internal s_nonceFinalized;

  constructor(IERC20 token) {
    i_token = token;
  }

  /// @notice Simply transferFrom msg.sender the tokens that are to be bridged to address(this).
  function sendERC20(
    address localToken,
    address /* remoteToken */,
    address /* remoteReceiver */,
    uint256 amount,
    bytes calldata /* bridgeSpecificPayload */
  ) external payable override returns (bytes memory) {
    IERC20(localToken).transferFrom(msg.sender, address(this), amount);
    bytes memory encodedNonce = abi.encode(s_nonce++);
    return encodedNonce;
  }

  function getBridgeFeeInNative() external pure returns (uint256) {
    return 0;
  }

  function provideLiquidity(uint256 amount) external {
    i_token.safeTransferFrom(msg.sender, address(this), amount);
    emit LiquidityAdded(msg.sender, amount);
  }

  function withdrawLiquidity(uint256 amount) external {
    if (i_token.balanceOf(address(this)) < amount) revert InsufficientLiquidity();
    i_token.safeTransfer(msg.sender, amount);
    emit LiquidityRemoved(msg.sender, amount);
  }

  /// @notice Payload to "prove" the withdrawal.
  /// @dev This is just a mock setup, there's no real proving. This is so that
  /// @dev we can test the multi-step finalization code path.
  /// @param nonce the nonce emitted on the remote chain.
  struct ProvePayload {
    uint256 nonce;
  }

  /// @notice Payload to "finalize" the withdrawal.
  /// @dev This is just a mock setup, there's no real finalization. This is so that
  /// @dev we can test the multi-step finalization code path.
  /// @param nonce the nonce emitted on the remote chain.
  struct FinalizePayload {
    uint256 nonce;
    uint256 amount;
  }

  /// @notice The finalization action to take.
  /// @dev This emulates Optimism's two-step withdrawal process.
  enum FinalizationAction {
    Invalid,
    ProveWithdrawal,
    FinalizeWithdrawal
  }

  /// @notice The payload to use for the bridgeSpecificPayload in the finalizeWithdrawERC20 function.
  struct Payload {
    FinalizationAction action;
    bytes data;
  }

  /// @dev Test setup is trusted, so just transfer the tokens to the localReceiver,
  /// @dev which should be the local rebalancer.
  /// @dev Infer the amount from the bridgeSpecificPayload
  /// @dev Note that this means that this bridge adapter will need to have some tokens,
  /// @dev however this is ok in a test environment since we will have infinite tokens.
  /// @return true if the transfer was successful, revert otherwise.
  function finalizeWithdrawERC20(
    address /* remoteSender */,
    address localReceiver,
    bytes calldata bridgeSpecificPayload
  ) external override returns (bool) {
    Payload memory payload = abi.decode(bridgeSpecificPayload, (Payload));
    if (payload.action == FinalizationAction.ProveWithdrawal) {
      ProvePayload memory provePayload = abi.decode(payload.data, (ProvePayload));
      if (s_nonceProven[provePayload.nonce]) revert NonceAlreadyUsed(provePayload.nonce);
      s_nonceProven[provePayload.nonce] = true;
      return false;
    } else if (payload.action == FinalizationAction.FinalizeWithdrawal) {
      FinalizePayload memory finalizePayload = abi.decode(payload.data, (FinalizePayload));
      if (s_nonceFinalized[finalizePayload.nonce]) revert NonceAlreadyUsed(finalizePayload.nonce);
      s_nonceFinalized[finalizePayload.nonce] = true;
      i_token.safeTransfer(localReceiver, finalizePayload.amount);
      return true;
    } else {
      revert InvalidFinalizationAction();
    }
  }
}

/// @notice Mock L2 Bridge adapter
/// @dev Sends the L2 tokens from the msg sender to address(this)
contract MockL2BridgeAdapter is IBridgeAdapter {
  /// @notice Simply transferFrom msg.sender the tokens that are to be bridged.
  function sendERC20(
    address localToken,
    address /* remoteToken */,
    address /* recipient */,
    uint256 amount,
    bytes calldata /* bridgeSpecificPayload */
  ) external payable override returns (bytes memory) {
    IERC20(localToken).transferFrom(msg.sender, address(this), amount);
    return "";
  }

  function getBridgeFeeInNative() external pure returns (uint256) {
    return 0;
  }

  // No-op
  function finalizeWithdrawERC20(
    address /* remoteSender */,
    address /* localReceiver */,
    bytes calldata /* bridgeSpecificData */
  ) external override returns (bool) {
    return true;
  }
}
