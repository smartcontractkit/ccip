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
    bytes recipient; // The address to send messages to on the destination-chain, abi.encode(addr) if an EVM-compatible networks
    bytes extraArgsBytes; // Includes additional configs such as manual gas limit, and out-of-order-execution. Should not be supplied at runtime to prevent unexpected contract behavior
    mapping(bytes recipient => bool isApproved) approvedSender;
  }

  address internal immutable i_ccipRouter;

  mapping(uint64 destChainSelector => ChainConfig) public s_chainConfigs;

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

  /// @notice modify the list of approved source-chain contracts which can send messages to this contract through CCIP
  /// @dev removes are executed before additions, so a contract present in both will be approved at the end of execution
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

  /// @notice Return whether a contract on the specified source chain is authorized to send messages to this contract through CCIP
  /// @param sourceChainSelector A unique CCIP-specific identifier for the source chain
  /// @param senderAddr The address which sent the message on the source-chain, abi-encoded if evm-compatible
  /// @return bool Whether the address is approved or not to invoke functions on this contract
  function isApprovedSender(uint64 sourceChainSelector, bytes calldata senderAddr) external view returns (bool) {
    return s_chainConfigs[sourceChainSelector].approvedSender[senderAddr];
  }

  // ================================================================
  // │                  Fee Token Management                       │
  // ===============================================================

  /// @notice function support native-fee-token pre-funding in all children implementing the ccipSend function
  receive() external payable {}

  /// @notice Allow the owner to recover any native-tokens sent to this contract out of error.
  /// @dev Function should not be used to recover tokens from failed-messages, abandonFailedMessage() should be used instead
  /// @param to A payable address to send the recovered tokens to
  /// @param amount the amount of native tokens to recover, denominated in wei
  function withdrawNativeToken(address payable to, uint256 amount) external onlyOwner {
    Address.sendValue(to, amount);
  }

  /// @notice Allow the owner to recover any ERC-20 tokens sent to this contract out of error.
  /// @dev Function should not be used to recover tokens from failed-messages, abandonFailedMessage() should be used instead
  /// @param to A payable address to send the recovered tokens to
  /// @param amount the amount of native tokens to recover, denominated in wei  function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
  function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
    IERC20(token).safeTransfer(to, amount);
  }

  // ================================================================
  // │                      Chain Management                        │
  // ================================================================

  /// @notice Enable a remote-chain to send and receive messages to/from this contract via CCIP
  /// @param chainSelector A unique CCIP-specific identifier for the source chain
  /// @param recipient The address a message should be sent to on the destination chain. There should only be one per-chain, and is abi-encoded if EVM-compatible.
  /// @param _extraArgsBytes additional optional ccipSend parameters. Do not need to be set unless necessary based on the application-logic
  function enableChain(
    uint64 chainSelector,
    bytes calldata recipient,
    bytes calldata _extraArgsBytes
  ) external onlyOwner {
    ChainConfig storage currentConfig = s_chainConfigs[chainSelector];

    currentConfig.recipient = recipient;

    // Set any additional args such as enabling out-of-order execution or manual gas-limit
    if (_extraArgsBytes.length != 0) currentConfig.extraArgsBytes = _extraArgsBytes;

    // If config was previously disabled, then re-enable it;
    if (currentConfig.disabled) currentConfig.disabled = false;
  }

  /// @notice Mark a chain as not supported for sending-receiving messages to/from this contract via CCIP.
  /// @dev If a chain needs to be re-enabled after being disabled, the owner must call enableChain() to support it again.
  function disableChain(uint64 chainSelector) external onlyOwner {
    s_chainConfigs[chainSelector].disabled = true;
  }

  modifier isValidChain(uint64 chainSelector) {
    // Must be storage and not memory because the struct contains a nested mapping which is not capable of being copied to memory
    ChainConfig storage currentConfig = s_chainConfigs[chainSelector];
    if (currentConfig.recipient.length == 0 || currentConfig.disabled) revert InvalidChain(chainSelector);
    _;
  }

  modifier isValidSender(uint64 chainSelector, bytes memory sender) {
    if (!s_chainConfigs[chainSelector].approvedSender[sender]) revert InvalidSender(sender);
    _;
  }
}
