// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";
import {ITypeAndVersion} from "../../../shared/interfaces/ITypeAndVersion.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/Address.sol";

abstract contract CCIPClientBase is OwnerIsCreator, ITypeAndVersion {
  using SafeERC20 for IERC20;
  using Address for address;

  address internal immutable i_ccipRouter;

  error ZeroAddressNotAllowed();
  error InvalidRouter(address router);
  error InvalidChain(uint64 chainSelector);
  error InvalidSender(bytes sender);
  error InvalidRecipient(bytes recipient);

  struct approvedSenderUpdate { // TODO capitalize
    uint64 destChainSelector;
    bytes sender; // TODO: add spack comments on what this field is
  }

  struct ChainConfig {
    bool isDisabled; // TODO rename: disabled
    bytes recipient;
    bytes extraArgsBytes; // TODO add explanas on why extraArgs is in ChainConfig, and not supposed to be supplied by sender at runtime
    mapping(bytes => bool) approvedSender;
  }

  mapping(uint64 => ChainConfig) public s_chainConfigs;

  constructor(address router) {
    if (router == address(0)) revert ZeroAddressNotAllowed();
    i_ccipRouter = router;
  }

  // ================================================================
  // │                      Router Management                       │
  // ================================================================
  // TODO natspec this contract
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
      delete s_chainConfigs[removes[i].destChainSelector].approvedSender[removes[i].sender];
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      s_chainConfigs[adds[i].destChainSelector].approvedSender[adds[i].sender] = true;
    }
  }

  function isApprovedSender(uint64 sourceChainSelector, bytes calldata senderAddr) external view returns (bool) {
    return s_chainConfigs[sourceChainSelector].approvedSender[senderAddr];
  }

  // ================================================================
  // │                  Fee Token Management                       │
  // ===============================================================

  fallback() external {} // TODO confirm with Devrel why they want this; we would prefer to not have a fallback to avoid silent failures

  // TODO add a comment why payable receive is in client base
  receive() external payable {}

  function withdrawNativeToken(address payable to, uint256 amount) external onlyOwner {
    Address.sendValue(to, amount);
  }

  // TODO add a warning message, this should only be used for things like withdrawing tokens that werent sent in error
  // this should never be used for recovering tokens from a message
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
    ChainConfig storage currentConfig = s_chainConfigs[chainSelector];

    currentConfig.recipient = recipient;

    if (_extraArgsBytes.length != 0) currentConfig.extraArgsBytes = _extraArgsBytes;

    // If config was previously disabled, then re-enable it;
    if (currentConfig.isDisabled) currentConfig.isDisabled = false;
  }

  function disableChain(uint64 chainSelector) external onlyOwner {
    s_chainConfigs[chainSelector].isDisabled = true;
  }

  // TODO change to supportedChain or isValidChain
  modifier validChain(uint64 chainSelector) { 
    // Must be storage and not memory because the struct contains a nested mapping
    ChainConfig storage currentConfig = s_chainConfigs[chainSelector];
    if (currentConfig.recipient.length == 0 || currentConfig.isDisabled) revert InvalidChain(chainSelector);
    _;
  }

  // TODO ditto
  modifier validSender(uint64 chainSelector, bytes memory sender) {
    if (!s_chainConfigs[chainSelector].approvedSender[sender]) revert InvalidSender(sender);
    _;
  }
}
