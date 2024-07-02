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

  error ZeroAddressNotAllowed();
  error InvalidRouter(address router);
  error InvalidChain(uint64 chainSelector);
  error InvalidSender(bytes sender);
  error InvalidRecipient(bytes recipient);

  struct ApprovedSenderUpdate {
    uint64 destChainSelector;
    bytes sender; // The address which initiated source-chain message, abi-encoded in case of a non-EVM-compatible network
  }

  struct ChainConfig {
    bool disabled;
    bytes recipient;
    // Includes additional configs such as manual gas limit, and OOO-execution.
    // Should not be supplied at runtime to prevent unexpected contract behavior
    bytes extraArgsBytes;
    mapping(bytes => bool) approvedSender;
  }

  address internal immutable i_ccipRouter;

  mapping(uint64 => ChainConfig) public s_chainConfigs;

  constructor(address router) {
    if (router == address(0)) revert ZeroAddressNotAllowed();
    i_ccipRouter = router;
  }

  // ================================================================
  // │                      Router Management                       │
  // ================================================================

  /// @notice returns the address of the CCIP Router set at contract-deployment
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
    ApprovedSenderUpdate[] calldata adds,
    ApprovedSenderUpdate[] calldata removes
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

  /// @notice function is set in client base to support native-fee-token pre-funding in all children implemending CCIPSend
  receive() external payable {}

  function withdrawNativeToken(address payable to, uint256 amount) external onlyOwner {
    Address.sendValue(to, amount);
  }

  /// @notice Function should NEVER be used for transfering tokens from a failed message, only for recovering tokens sent in error
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
    if (currentConfig.disabled) currentConfig.disabled = false;
  }

  function disableChain(uint64 chainSelector) external onlyOwner {
    s_chainConfigs[chainSelector].disabled = true;
  }

  modifier isValidChain(uint64 chainSelector) {
    // Must be storage and not memory because the struct contains a nested mapping
    ChainConfig storage currentConfig = s_chainConfigs[chainSelector];
    if (currentConfig.recipient.length == 0 || currentConfig.disabled) revert InvalidChain(chainSelector);
    _;
  }

  modifier isValidSender(uint64 chainSelector, bytes memory sender) {
    if (!s_chainConfigs[chainSelector].approvedSender[sender]) revert InvalidSender(sender);
    _;
  }
}
