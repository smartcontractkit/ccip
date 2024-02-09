// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {IBridgeAdapter} from "../interfaces/IBridge.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

interface IArbSys {
  function withdrawEth(address destination) external payable returns (uint256);
}

interface IL2GatewayRouter {
  function outboundTransfer(
    address l1Token,
    address to,
    uint256 amount,
    bytes calldata data
  ) external payable returns (bytes memory);
}

/// @notice Arbitrum L2 Bridge adapter
/// @dev Auto unwraps and re-wraps wrapped eth in the bridge.
contract ArbitrumL2BridgeAdapter is IBridgeAdapter {
  using SafeERC20 for IERC20;

  IL2GatewayRouter internal immutable i_l2GatewayRouter;
  //  address internal immutable i_l1ERC20Gateway;
  IArbSys internal constant ARB_SYS = IArbSys(address(0x64));

  /// @notice event to track the L2 to L1 transfer offchain
  /// @dev while this bridge adapter is trustless and anyone can use it,
  /// @dev its highly unlikely that anyone would prefer it over the official bridge
  /// @dev contracts. And since the official bridge contracts probably have lots of
  /// @dev events and logs, we can use this event to track the L2 to L1 transfers
  /// @dev with less load on the rebalancer oracles.
  /// @param localToken the token address on L2
  /// @param remoteToken the token address on L1
  /// @param recipient the recipient of the tokens on L1
  /// @param amount the amount of tokens transferred
  /// @param outboundTransferResult the result of the outbound transfer, which is the unique id used to identify the L2 to L1 tx
  event ArbitrumL2ToL1ERC20Sent(
    address indexed localToken,
    address indexed remoteToken,
    address indexed recipient,
    uint256 amount,
    bytes outboundTransferResult
  );
  /// @notice event to track the finalization of the L2 to L1 transfer
  /// @dev no data to emit since there isn't typically a finalization step for L1 to L2 transfers
  event ArbitrumL1ToL2ERC20Finalized();

  constructor(IL2GatewayRouter l2GatewayRouter) {
    if (address(l2GatewayRouter) == address(0)) {
      revert BridgeAddressCannotBeZero();
    }
    i_l2GatewayRouter = l2GatewayRouter;
  }

  /// @inheritdoc IBridgeAdapter
  function sendERC20(
    address localToken,
    address remoteToken,
    address recipient,
    uint256 amount,
    bytes calldata /* bridgeSpecificPayload */
  ) external payable override returns (bytes memory) {
    if (msg.value != 0) {
      revert MsgShouldNotContainValue(msg.value);
    }

    IERC20(localToken).safeTransferFrom(msg.sender, address(this), amount);

    // TODO: handle return data bombs?
    // NOTE: the data returned is the unique id of the L2 to L1 transfer
    // see https://github.com/OffchainLabs/token-bridge-contracts/blob/bf9ad3d7f25c0eaf0a5f89eec7a0a370833cea16/contracts/tokenbridge/arbitrum/gateway/L2ArbitrumGateway.sol#L169-L191
    // so we can probably cap the return data to 256 bits, e.g using https://github.com/nomad-xyz/ExcessivelySafeCall
    // TODO: should we decode res here or offchain?
    // No approval needed, the bridge will burn the tokens from this contract.
    bytes memory res = i_l2GatewayRouter.outboundTransfer(remoteToken, recipient, amount, bytes(""));

    emit ArbitrumL2ToL1ERC20Sent(localToken, remoteToken, recipient, amount, res);

    return res;
  }

  /// @notice No-op since L1 -> L2 transfers do not need finalization.
  function finalizeWithdrawERC20(
    address /* remoteSender */,
    address /* localReceiver */,
    bytes calldata /* bridgeSpecificPayload */
  ) external {
    emit ArbitrumL1ToL2ERC20Finalized();
  }

  /// @notice There are no fees to bridge back to L1
  function getBridgeFeeInNative() external pure returns (uint256) {
    return 0;
  }

  function depositNativeToL1(address recipient) external payable {
    ARB_SYS.withdrawEth{value: msg.value}(recipient);
  }
}
