// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../../shared/access/OwnerIsCreator.sol";

import {IERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/IERC20.sol";
import {SafeERC20} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/token/ERC20/utils/SafeERC20.sol";
import {Address} from "../../../vendor/openzeppelin-solidity/v4.8.3/contracts/utils/Address.sol";

/// @title CCIPBase
/// @notice This contains the boilerplate code for managing chains and tokens your contract may interact with as part of CCIP. It does not send or receive messages through CCIP only manages supported chains and sources/destinations.
/// @dev This contract is abstract, but does not have any functions which must be implemented by a child.
abstract contract CCIPBase is OwnerIsCreator {
  using SafeERC20 for IERC20;
  using Address for address;

  error ZeroAddressNotAllowed();
  error InvalidRouter(address router);
  error InvalidChain(uint64 chainSelector);
  error InvalidSender(bytes sender);
  error InvalidRecipient(bytes recipient);

  event CCIPRouterModified(address indexed oldRouter, address indexed newRouter);
  event TokensWithdrawnByOwner(address indexed token, address indexed to, uint256 amount);

  // TODO 2 events, add/remove
  // TODO comment on the tradeoffs, short comment on why prefer indexed
  event ApprovedSenderModified(uint64 indexed destChainSelector, bytes indexed recipient, bool isBeingApproved);

  event ChainAdded(uint64 indexed remoteChainSelector, bytes indexed recipient, bytes extraArgsBytes);
  event ChainRemoved(uint64 indexed removeChainSelector);

  struct ApprovedSenderUpdate {
    uint64 destChainSelector; // Chainselector for a source chain that is allowed to call this dapp
    bytes sender; //             The sender address on source chain that is allowed to call, ABI encoded in the case of a remote EVM chain
  }

  struct ChainUpdate { // TODO comments
    uint64 chainSelector;
    bool allowed;
    bytes recipient;
    bytes extraArgsBytes;
  }

  struct ChainConfig {
    bytes recipient; //      The address to send messages to on the destination chain, ABI encoded in the case of a remote EVM chain.
    bytes extraArgsBytes; // Specifies extraArgs to pass into ccipSend, includes configs such as gas limit, and OOO execution.
    mapping(bytes recipient => bool isApproved) approvedSender; // Mapping is nested to support work-flows where Dapps may need to receive messages from one-or-more contracts on a source chain, or to support one-sided dapp upgrades.
  }

  address internal s_ccipRouter;

  mapping(uint64 destChainSelector => ChainConfig) public s_chainConfigs;

  constructor(address router) {
    if (router == address(0)) revert ZeroAddressNotAllowed();
    s_ccipRouter = router;
  }

  // ================================================================
  // │                      Router Management                       │
  // ================================================================

  /// @notice returns the address of the CCIP Router set at contract deployment
  function getRouter() public view virtual returns (address) {
    return s_ccipRouter;
  }

  /// @dev only calls from the set router are accepted.
  modifier onlyRouter() {
    if (msg.sender != getRouter()) revert InvalidRouter(msg.sender);
    _;
  }

  // ================================================================
  // │                  Sender/Receiver Management                  │
  // ================================================================

  /// @notice modify the list of approved source chain contracts which can send messages to this contract through CCIP
  /// @dev removes are executed before additions, so a contract present in both will be approved at the end of execution
  function updateApprovedSenders(
    ApprovedSenderUpdate[] calldata adds,
    ApprovedSenderUpdate[] calldata removes
  ) external onlyOwner {
    for (uint256 i = 0; i < removes.length; ++i) {
      delete s_chainConfigs[removes[i].destChainSelector].approvedSender[removes[i].sender];

      // Third parameter is false to indicate that the sender's previous approval is being revoked, to improve off-chain event indexing
      emit ApprovedSenderModified(removes[i].destChainSelector, removes[i].sender, false);
    }

    for (uint256 i = 0; i < adds.length; ++i) {
      s_chainConfigs[adds[i].destChainSelector].approvedSender[adds[i].sender] = true;

      // Third parameter is true to indicate that the sender is being approved, to improve off-chain event indexing
      emit ApprovedSenderModified(adds[i].destChainSelector, adds[i].sender, true);
    }
  }

  /// @notice Return whether a contract on the specified source chain is authorized to send messages to this contract through CCIP
  /// @dev This function does not revert on an unapproved-sender, and should only be used as a getter-function for querying approvals from a ChainConfig object. The isValidSender modifier should be used instead for incoming message-validation
  /// @param sourceChainSelector A unique CCIP-specific identifier for the source chain
  /// @param senderAddr The address which sent the message on the source chain, abi-encoded if evm-compatible
  /// @return bool Whether the address is approved or not to invoke functions on this contract
  function isApprovedSender(uint64 sourceChainSelector, bytes calldata senderAddr) external view returns (bool) {
    return s_chainConfigs[sourceChainSelector].approvedSender[senderAddr];
  }

  // ================================================================
  // │                  Fee Token Management                       │
  // ===============================================================

  /// @notice Accepts incoming native-tokens to support prefunding in native fee token.
  /// @dev All the example applications accept prefunding. This function should be removed if prefunding in native fee token is not required.
  receive() external payable {}

  /// @notice Allow the owner to recover any native-tokens sent to this contract out of error, or to withdraw any native-tokens which were used for pre-funding if the fee-token is switched away from native-tokens.
  /// @dev Function should not be used to recover tokens from failed messages, abandonFailedMessage() should be used instead
  /// @param to A payable address to send the recovered tokens to
  /// @param amount the amount of native tokens to recover, denominated in wei
  function withdrawNativeToken(address payable to, uint256 amount) external onlyOwner {
    Address.sendValue(to, amount);

    // Use the same withdrawal event signature as withdrawTokens() but use address(0) to denote native-tokens.
    emit TokensWithdrawnByOwner(address(0), to, amount);
  }

  /// @notice Allow the owner to recover any ERC-20 tokens sent to this contract out of error or withdraw any fee-tokens which were sent as a source of fee-token pre-funding
  /// @dev This should NOT be used for recoverying tokens from a failed message. Token recoveries can happen only if
  /// the failed message is guaranteed to not succeed upon retry, otherwise this can lead to double spend.
  /// For implementation of token recovery, see inheriting contracts.
  /// @param to A payable address to send the recovered tokens to
  /// @param amount the amount of native tokens to recover, denominated in wei  function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
  function withdrawTokens(address token, address to, uint256 amount) external onlyOwner {
    IERC20(token).safeTransfer(to, amount);

    emit TokensWithdrawnByOwner(token, to, amount);
  }

  // ================================================================
  // │                      Chain Management                        │
  // ================================================================

  // TODO comments
  // TODO name as updateRouter
  function modifyRouter(address newRouter) external onlyOwner {
    if (newRouter == address(0)) revert ZeroAddressNotAllowed();

    // Store the old router in memory to emit event
    address currentRouter = s_ccipRouter;

    s_ccipRouter = newRouter;

    emit CCIPRouterModified(currentRouter, newRouter);
  }

  // Reference similar function in TokenPool
  /// @notice Enable a remote-chain to send and receive messages to/from this contract via CCIP
  function applyChainUpdates(ChainUpdate[] calldata chains) external onlyOwner {
    for (uint256 i = 0; i < chains.length; ++i) {
      if (!chains[i].allowed) {
        delete s_chainConfigs[chains[i].chainSelector].recipient;
        emit ChainRemoved(chains[i].chainSelector);
      } else {
        // The existence of a stored recipient is used to denote a chain being enabled, so the length here cannot be zero
        if (chains[i].recipient.length == 0) revert ZeroAddressNotAllowed();

        ChainConfig storage currentConfig = s_chainConfigs[chains[i].chainSelector];

        currentConfig.recipient = chains[i].recipient;

        // Set any additional args such as enabling out-of-order execution or manual gas-limit
        if (chains[i].extraArgsBytes.length != 0) currentConfig.extraArgsBytes = chains[i].extraArgsBytes;

        emit ChainAdded(chains[i].chainSelector, chains[i].recipient, chains[i].extraArgsBytes);
      }
    }
  }

  // TODO comments
  modifier isValidChain(uint64 chainSelector) {
    // Must be storage and not memory because the struct contains a nested mapping which is not capable of being copied to memory
    ChainConfig storage currentConfig = s_chainConfigs[chainSelector];
    if (currentConfig.recipient.length == 0) revert InvalidChain(chainSelector);
    _;
  }

  modifier isValidSender(uint64 chainSelector, bytes memory sender) {
    // If the chain is disabled, then short-circuit trigger a revert because no sender should be valid
    if (s_chainConfigs[chainSelector].recipient.length == 0 || !s_chainConfigs[chainSelector].approvedSender[sender]) {
      revert InvalidSender(sender);
    }
    _;
  }
}
