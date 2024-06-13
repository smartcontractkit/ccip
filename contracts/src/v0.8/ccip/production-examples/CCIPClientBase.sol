pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

import {IERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";

import {ICCIPClientBase} from "./interfaces/ICCIPClientBase.sol";

abstract contract CCIPClientBase is ICCIPClientBase, OwnerIsCreator {
  using SafeERC20 for IERC20;

  address internal immutable i_ccipRouter;

  mapping(uint64 destChainSelector => mapping(bytes sender => bool approved)) public s_approvedSenders;
  mapping(uint64 destChainSelector => bytes recipient) public s_chains;
  mapping(uint64 destChainselector => bytes extraArgsBytes) public s_extraArgsBytes;

  constructor(address router) {
    if (router == address(0)) revert InvalidRouter(address(0));
    i_ccipRouter = router;
  }

  /////////////////////////////////////////////////////////////////////
  // Router Management
  /////////////////////////////////////////////////////////////////////

  function getRouter() public view returns (address) {
    return i_ccipRouter;
  }

  /// @dev only calls from the set router are accepted.
  modifier onlyRouter() {
    if (msg.sender != getRouter()) revert InvalidRouter(msg.sender);
    _;
  }

  /////////////////////////////////////////////////////////////////////
  // Sender/Receiver Management
  /////////////////////////////////////////////////////////////////////

  function updateApprovedSenders(approvedSenderUpdate[] calldata adds, approvedSenderUpdate[] calldata removes) external onlyOwner {
    for(uint256 i = 0; i < removes.length; ++i) {
      delete s_approvedSenders[removes[i].destChainSelector][removes[i].sender];
    }

    for(uint256 i = 0; i < adds.length; ++i) {
      s_approvedSenders[adds[i].destChainSelector][adds[i].sender] = true;
    }
  }

  /////////////////////////////////////////////////////////////////////
  // Fee Token Management
  /////////////////////////////////////////////////////////////////////

  fallback() external payable {}
  receive() external payable {}

  function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
    IERC20(token).safeTransfer(to, amount);
  }

  /////////////////////////////////////////////////////////////////////
  // Chain Management
  /////////////////////////////////////////////////////////////////////

  function enableChain(uint64 chainSelector, bytes calldata recipient, bytes calldata extraArgsBytes) external onlyOwner {
    s_chains[chainSelector] = recipient;

    if (extraArgsBytes.length != 0) s_extraArgsBytes[chainSelector] = extraArgsBytes;
  }

  function disableChain(uint64 chainSelector) external onlyOwner {
    delete s_chains[chainSelector];
    delete s_extraArgsBytes[chainSelector];
  }

  modifier validChain(uint64 chainSelector) {
    if (s_chains[chainSelector].length == 0) revert InvalidChain(chainSelector);
    _;
  }

  modifier validSender(uint64 chainSelector, bytes memory sender) {
    if (!s_approvedSenders[chainSelector][sender]) revert InvalidSender(sender);
    _;
  }

}
