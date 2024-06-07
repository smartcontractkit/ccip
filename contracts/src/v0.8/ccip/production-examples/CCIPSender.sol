// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {IRouterClient} from "../interfaces/IRouterClient.sol";

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {Client} from "../libraries/Client.sol";

import {ICCIPClientBase} from "./interfaces/ICCIPClientBase.sol";
import {CCIPClientBase} from "./CCIPClientBase.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

// @notice Example of a client which supports EVM/non-EVM chains
// @dev If chain specific logic is required for different chain families (e.g. particular
// decoding the bytes sender for authorization checks), it may be required to point to a helper
// authorization contract unless all chain families are known up front.
// @dev If contract does not implement IAny2EVMMessageReceiver and IERC165,
// and tokens are sent to it, ccipReceive will not be called but tokens will be transferred.
// @dev If the client is upgradeable you have significantly more flexibility and
// can avoid storage based options like the below contract uses. However it's
// worth carefully considering how the trust assumptions of your client dapp will
// change if you introduce upgradeability. An immutable dapp building on top of CCIP
// like the example below will inherit the trust properties of CCIP (i.e. the oracle network).
// @dev The receiver's are encoded offchain and passed as direct arguments to permit supporting
// new chain family receivers (e.g. a solana encoded receiver address) without upgrading.
contract CCIPSender is CCIPClientBase {
  using SafeERC20 for IERC20;
  
  error InvalidConfig();
  error InsufficientNativeFeeTokenAmount();

  event MessageSent(bytes32 messageId);
  event MessageReceived(bytes32 messageId);

 
  constructor(address router) CCIPClientBase(router) {}

  function ccipSend(
    uint64 destChainSelector,
    Client.EVMTokenAmount[] memory tokenAmounts,
    bytes calldata data,
    address feeToken
  ) public payable validChain(destChainSelector) {

    // TODO: Decide whether workflow should assume contract is funded with tokens to send already
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      IERC20(tokenAmounts[i].token).transferFrom(msg.sender, address(this), tokenAmounts[i].amount);
      IERC20(tokenAmounts[i].token).approve(i_ccipRouter, tokenAmounts[i].amount);
    }

    ICCIPClientBase.ChainInfo memory chainInfo = s_chains[destChainSelector];

    Client.EVM2AnyMessage memory message = Client.EVM2AnyMessage({
      receiver: chainInfo.recipient,
      data: data,
      tokenAmounts: tokenAmounts,
      extraArgs: chainInfo.extraArgsBytes,
      feeToken: feeToken
    });

    uint256 fee = IRouterClient(i_ccipRouter).getFee(destChainSelector, message);

    // Transfer fee token from sender and approve router to pay for message
    if (feeToken != address(0)) {
      IERC20(feeToken).safeTransferFrom(msg.sender, address(this), fee);
      IERC20(feeToken).safeApprove(i_ccipRouter, fee);
    }

    bytes32 messageId = IRouterClient(i_ccipRouter).ccipSend{
      value: feeToken == address(0) ? fee : 0
    } (destChainSelector, message);

    emit MessageSent(messageId);
  }

}
