// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IL2Bridge} from "./IBridge.sol";
import {IWrappedNative} from "../../interfaces/IWrappedNative.sol";

import {L2GatewayRouter} from "@arbitrum/token-bridge-contracts/contracts/tokenbridge/arbitrum/gateway/L2GatewayRouter.sol";
import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

interface IArbSys {
  function withdrawEth(address destination) external payable returns (uint256);
}

/// @notice Arbitrum L2 Bridge adapter
/// @dev Auto unwraps and re-wraps wrapped eth in the bridge.
contract ArbitrumL2BridgeAdapter is IL2Bridge {
  using SafeERC20 for IERC20;

  error InsufficientEthValue(uint256 wanted, uint256 got);

  IL2GatewayRouter internal immutable i_l2GatewayRouter;
  //  address internal immutable i_l1ERC20Gateway;
  IArbSys internal constant ARB_SYS = IArbSys(address(0x64));

  constructor(IL2GatewayRouter l2GatewayRouter) {
    if (address(l2GatewayRouter) == address(0)) {
      revert BridgeAddressCannotBeZero();
    }
    i_l2GatewayRouter = l2GatewayRouter;
  }

  function sendERC20(address l1Token, address l2Token, address recipient, uint256 amount) external payable {
    IERC20(l2Token).safeTransferFrom(msg.sender, address(this), amount);

    i_l2GatewayRouter.outboundTransfer(l1Token, recipient, amount, bytes(""));
  }

  function depositNativeToL1(address recipient) external payable {
    ARB_SYS.withdrawEth{value: msg.value}(recipient);
  }
}

interface IL2GatewayRouter {
  function outboundTransfer(address l1Token, address to, uint256 amount, bytes calldata data) external;
}
