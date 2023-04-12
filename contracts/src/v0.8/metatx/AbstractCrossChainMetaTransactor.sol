// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import {IERC2771Recipient} from "../vendor/IERC2771Recipient.sol";
import {Context} from "../vendor/Context.sol";
import {IRouterClient} from "../ccip/interfaces/IRouterClient.sol";
import {Client} from "../ccip/libraries/Client.sol";
import {OwnerIsCreator} from "../ccip/OwnerIsCreator.sol";

/// @dev AbstractCrossChainMetaTransactor extends ERC20 token to add cross chain transfer functionality
/// @dev Also, it trusts ERC2771 forwarder to forward meta-transactions
abstract contract AbstractCrossChainMetaTransactor is OwnerIsCreator, IERC2771Recipient {
  /// @dev forwarder verifies signatures for meta transactions and forwards the
  /// @dev request to this contract
  address private s_forwarder;
  address private s_ccipRouter;
  address private s_chainlinkOwner;
  /// @dev boolean only used for testing. Should be set to false in production
  /// @dev go-ethereum.simulatedBackend (used for testing) doesn't allow custom chain IDs
  /// @dev so block.chainid is hard-coded to 1337.
  bool private s_test_only_force_cross_chain_transfer;

  constructor(
    address forwarder,
    address ccipRouter,
    address chainlinkOwner,
    bool _test_only_force_cross_chain_transfer
  ) {
    s_forwarder = forwarder;
    s_ccipRouter = ccipRouter;
    s_chainlinkOwner = chainlinkOwner;
    s_test_only_force_cross_chain_transfer = _test_only_force_cross_chain_transfer;
  }

  /// @dev Transfers "amount" of this token to receiver address in destination chain.
  /// @param receiver token receiver address in destination chain. Handles distribution of tokens to recipients
  /// @param amount total token amount to be transferred
  /// @param destinationChainId destination chain ID
  function metaTransfer(
    address receiver,
    uint256 amount,
    uint64 destinationChainId
  ) external validateTrustedForwarder returns (bytes32) {
    if (!isCrossChainTransfer(destinationChainId)) {
      _transfer(_msgSender(), receiver, amount);
      return ""; // return empty bytes32 because there is no ccip message ID
    }
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: address(this), amount: amount});
    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: abi.encode(receiver),
      data: "",
      tokenAmounts: tokenAmounts,
      feeToken: address(0), // use native token instead of ERC20 tokens
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000, strict: false}))
    });
    uint256 fee = IRouterClient(s_ccipRouter).getFee(destinationChainId, message);

    _transfer(_msgSender(), address(this), amount);
    _approve(address(this), s_ccipRouter, amount);

    return IRouterClient(s_ccipRouter).ccipSend{value: fee}(destinationChainId, message);
  }

  /// @dev Sets `amount` as the allowance of `spender` over the `owner` s tokens.
  /// @param owner token owner approving allowance
  /// @param spender approved token spender
  /// @param amount total token amount to be approved
  function _approve(
    address owner,
    address spender,
    uint256 amount
  ) internal virtual;

  /// @dev Moves `amount` of tokens from `sender` to `recipient`.
  /// @param sender token sender
  /// @param recipient token recipient
  /// @param amount total token amount to be approved
  function _transfer(
    address sender,
    address recipient,
    uint256 amount
  ) internal virtual;

  function isCrossChainTransfer(uint64 chainId) private view returns (bool) {
    if (s_test_only_force_cross_chain_transfer) {
      return true;
    }
    return chainId != block.chainid;
  }

  receive() external payable {}

  error WithdrawFailure();

  function withdrawNative() external validateChainlinkOwner {
    uint256 amount = address(this).balance;
    // Owner can receive Ether since the address of owner is payable
    (bool success, ) = owner().call{value: amount}("");
    if (!success) {
      revert WithdrawFailure();
    }
  }

  /// @notice Method is not a required method to allow Recipients to trust multiple Forwarders. Not recommended yet.
  /// @notice **Warning** The Forwarder can have a full control over your Recipient. Only trust verified Forwarder.
  /// @return forwarder The address of the Forwarder contract that is being used.
  function getTrustedForwarder() public view returns (address forwarder) {
    return s_forwarder;
  }

  /// @inheritdoc IERC2771Recipient
  function isTrustedForwarder(address forwarder) public view override returns (bool) {
    return forwarder == s_forwarder;
  }

  /// @inheritdoc IERC2771Recipient
  function _msgSender() internal view override returns (address ret) {
    if (msg.data.length >= 20 && isTrustedForwarder(msg.sender)) {
      // At this point we know that the sender is a trusted forwarder,
      // so we trust that the last bytes of msg.data are the verified sender address.
      // extract sender address from the end of msg.data
      assembly {
        ret := shr(96, calldataload(sub(calldatasize(), 20)))
      }
    } else {
      ret = msg.sender;
    }
  }

  /// @inheritdoc IERC2771Recipient
  function _msgData() internal view override returns (bytes calldata ret) {
    if (msg.data.length >= 20 && isTrustedForwarder(msg.sender)) {
      return msg.data[0:msg.data.length - 20];
    } else {
      return msg.data;
    }
  }

  function getForwarder() public view returns (address) {
    return s_forwarder;
  }

  function getCCIPRouter() public view returns (address) {
    return s_ccipRouter;
  }

  function getChainlinkOwner() public view returns (address) {
    return s_chainlinkOwner;
  }

  error MustBeTrustedForwarder(address sender);
  error MustBeChainlinkOwner(address sender);

  modifier validateTrustedForwarder() {
    if (!isTrustedForwarder(msg.sender)) {
      revert MustBeTrustedForwarder(msg.sender);
    }
    _;
  }

  modifier validateChainlinkOwner() {
    if (msg.sender != s_chainlinkOwner) {
      revert MustBeChainlinkOwner(msg.sender);
    }
    _;
  }
}
