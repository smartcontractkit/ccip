// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IL2Bridge} from "./IBridge.sol";
import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";

import {L2StandardBridge} from "@eth-optimism/contracts/L2/messaging/L2StandardBridge.sol";
import {Lib_PredeployAddresses} from "@eth-optimism/contracts/libraries/constants/Lib_PredeployAddresses.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

contract OptimismL2BridgeAdapter is IL2Bridge {
  using SafeERC20 for IERC20;

  L2StandardBridge internal i_L2Bridge = L2StandardBridge(Lib_PredeployAddresses.L2_STANDARD_BRIDGE);
  IWrappedNative internal immutable i_wrappedNative;

  // Nonce to use for L1 withdrawals to allow for better tracking offchain.
  uint64 private s_nonce = 0;

  constructor(IWrappedNative wrappedNative) {
    // Wrapped native can be address zero, this means that auto-wrapping is disabled.
    i_wrappedNative = wrappedNative;
  }

  function depositERC20ToL1(address l2Token, address recipient, uint256 amount) external {
    IERC20(l2Token).safeTransferFrom(msg.sender, address(this), amount);

    // If the token is the wrapped native, we unwrap it and deposit native
    if (l2Token == address(i_wrappedNative)) {
      i_wrappedNative.withdraw(amount);
      depositNativeToL1(recipient, amount);
      return;
    }

    // Token is normal ERC20
    IERC20(l2Token).approve(address(i_L2Bridge), amount);
    i_L2Bridge.withdrawTo(l2Token, recipient, amount, 0, abi.encode(s_nonce++));
  }

  function depositNativeToL1(address recipient, uint256 amount) public payable {
    i_L2Bridge.withdrawTo(Lib_PredeployAddresses.OVM_ETH, recipient, amount, 0, abi.encode(s_nonce++));
  }

  function getL1Bridge() external view returns (address) {
    return i_L2Bridge.l1TokenBridge();
  }

  function getL2Bridge() external view returns (address) {
    return address(i_L2Bridge);
  }

  /// @notice returns the address of the
  function getWrappedNative() external view returns (address) {
    return address(i_wrappedNative);
  }
}
