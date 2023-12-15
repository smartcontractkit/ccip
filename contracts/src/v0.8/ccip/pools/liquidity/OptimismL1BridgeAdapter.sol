// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IL1Bridge} from "./IBridge.sol";
import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";

import {IL1StandardBridge} from "@eth-optimism/contracts/L1/messaging/IL1StandardBridge.sol";
import {IL1CrossDomainMessenger} from "@eth-optimism/contracts/L1/messaging/IL1CrossDomainMessenger.sol";
import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/utils/SafeERC20.sol";

contract OptimismL1BridgeAdapter is IL1Bridge {
  using SafeERC20 for IERC20;

  IL1StandardBridge internal immutable i_L1Bridge;
  IL1CrossDomainMessenger internal immutable i_L1CrossDomainMessenger;
  IWrappedNative internal immutable i_wrappedNative;

  // Nonce to use for L2 deposits to allow for better tracking offchain.
  uint64 private s_nonce = 0;

  constructor(
    IL1StandardBridge l1Bridge,
    IWrappedNative wrappedNative,
    IL1CrossDomainMessenger l1CrossDomainMessenger
  ) {
    if (address(l1Bridge) == address(0) || address(wrappedNative) == address(0)) {
      revert BridgeAddressCannotBeZero();
    }
    i_L1Bridge = l1Bridge;
    i_L1CrossDomainMessenger = l1CrossDomainMessenger;
    i_wrappedNative = wrappedNative;
  }

  function depositERC20ToL2(address l1Token, address l2Token, address recipient, uint256 amount) external {
    IERC20(l1Token).safeTransferFrom(msg.sender, address(this), amount);

    // If the token is the wrapped native, we unwrap it and deposit native
    if (l1Token == address(i_wrappedNative)) {
      i_wrappedNative.withdraw(amount);
      depositNativeToL2(recipient, amount);
      return;
    }

    // Token is normal ERC20
    IERC20(l1Token).approve(address(i_L1Bridge), amount);
    i_L1Bridge.depositERC20To(l1Token, l2Token, recipient, amount, 0, abi.encode(s_nonce++));
  }

  function depositNativeToL2(address recipient, uint256 amount) public payable {
    i_L1Bridge.depositETHTo{value: amount}(recipient, 0, abi.encode(s_nonce++));
  }

  function finalizeWithdrawERC20FromL2(
    address l1Token,
    address l2Token,
    address from,
    address to,
    uint256 amount,
    bytes calldata data
  ) external {
    i_L1Bridge.finalizeERC20Withdrawal(l1Token, l2Token, from, to, amount, data);
  }

  function finalizeWithdrawNativeFromL2(address from, address to, uint256 amount, bytes calldata data) external {
    i_L1Bridge.finalizeETHWithdrawal(from, to, amount, data);
  }

  function relayMessageFromL2ToL1(
    address target,
    address sender,
    bytes memory message,
    uint256 messageNonce,
    IL1CrossDomainMessenger.L2MessageInclusionProof memory proof
  ) external {
    i_L1CrossDomainMessenger.relayMessage(target, sender, message, messageNonce, proof);
    // TODO
  }

  function getL1Bridge() external view returns (address) {
    return address(i_L1Bridge);
  }

  function getL2Bridge() external returns (address) {
    return i_L1Bridge.l2TokenBridge();
  }

  /// @notice returns the address of the
  function getWrappedNative() external view returns (address) {
    return address(i_wrappedNative);
  }
}
