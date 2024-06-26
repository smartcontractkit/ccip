pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/Address.sol";

import {ICCIPClientBase} from "../../interfaces/ICCIPClientBase.sol";

abstract contract CCIPClientBase is ICCIPClientBase, OwnerIsCreator, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using Address for address;

  address internal immutable i_ccipRouter;

  error ZeroAddressNotAllowed();

  struct ChainInfo {
    bytes recipient;
    bytes extraArgsBytes;
    mapping(bytes => bool) approvedSender;
  }

  mapping(uint64 => ChainInfo) public s_chains;

  // mapping(uint64 => mapping(bytes sender => bool)) public s_approvedSenders;
  // mapping(uint64 => bytes) public s_chains;
  // mapping(uint64 => bytes) public s_extraArgsBytes;

  constructor(address router) {
    if (router == address(0)) revert ZeroAddressNotAllowed();
    i_ccipRouter = router;
  }

  // ================================================================
  // │                      Router Management                       │
  // ================================================================

  function getRouter() public view virtual returns (address) {
    return i_ccipRouter;
  }

  /// @dev only calls from the set router are accepted.
  modifier onlyRouter() {
    if (msg.sender != getRouter()) revert InvalidRouter(msg.sender);
    _;
  }

  // ================================================================
  // │                  Sender/Receiver Management                  │
  // ================================================================

  function updateApprovedSenders(
    approvedSenderUpdate[] calldata adds,
    approvedSenderUpdate[] calldata removes
  ) external onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      // delete s_approvedSenders[removes[i].destChainSelector][removes[i].sender];
      delete s_chains[removes[i].destChainSelector].approvedSender[removes[i].sender];
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      // s_approvedSenders[adds[i].destChainSelector][adds[i].sender] = true;
      s_chains[adds[i].destChainSelector].approvedSender[adds[i].sender] = true;
    }
  }

  function isApprovedSender(uint64 sourceChainSelector, bytes calldata senderAddr) external view returns (bool) {
    return s_chains[sourceChainSelector].approvedSender[senderAddr];
  }

  // ================================================================
  // │                  Fee Token Management                       │
  // ===============================================================

  fallback() external payable {}
  receive() external payable {}

  function withdrawNativeToken(address payable to, uint256 amount) external onlyOwner {
    Address.sendValue(to, amount);
  }

  function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
    IERC20(token).safeTransfer(to, amount);
  }

  // ================================================================
  // │                      Chain Management                        │
  // ================================================================

  function enableChain(
    uint64 chainSelector,
    bytes calldata recipient,
    bytes calldata _extraArgsBytes
  ) external onlyOwner {
    s_chains[chainSelector].recipient = recipient;

    if (_extraArgsBytes.length != 0) s_chains[chainSelector].extraArgsBytes = _extraArgsBytes;
  }

  function disableChain(uint64 chainSelector) external onlyOwner {
    delete s_chains[chainSelector];
    // delete s_extraArgsBytes[chainSelector];
  }

  modifier validChain(uint64 chainSelector) {
    if (s_chains[chainSelector].recipient.length == 0) revert InvalidChain(chainSelector);
    _;
  }

  modifier validSender(uint64 chainSelector, bytes memory sender) {
    // if (!s_approvedSenders[chainSelector][sender]) revert InvalidSender(sender);
    if (!s_chains[chainSelector].approvedSender[sender]) revert InvalidSender(sender);
    _;
  }
}
