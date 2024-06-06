pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

contract CCIPClientBase is OwnerIsCreator {
  error InvalidRouter(address router);
  error InvalidChain(uint64 chainSelector);
  error InvalidSender(bytes sender);
  error InvalidRecipient(bytes recipient);

  address internal immutable i_ccipRouter;

  struct Chain {
    bytes extraArgsBytes;
    bytes recipient;
  }
  
  mapping(uint64 destChainSelector => Chain chainInfo) public s_chains;

  mapping(bytes sender => bool isApproved) public s_senders; // Approved addresses of which to receive messages from

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

  function updateApprovedSenders(bytes[] calldata adds, bytes[] calldata removes) external onlyOwner {
    for(uint256 i = 0; i < removes.length; ++i) {
      delete s_senders[removes[i]];
    }

    for(uint256 i = 0; i < removes.length; ++i) {
      s_senders[adds[i]] = true;
    }
  }

  /////////////////////////////////////////////////////////////////////
  // Chain Management
  /////////////////////////////////////////////////////////////////////

  function enableChain(uint64 chainSelector, Chain calldata chainInfo) external onlyOwner {
    s_chains[chainSelector] = chainInfo;
  }

  function disableChain(uint64 chainSelector) external onlyOwner {
    delete s_chains[chainSelector];
  }

  modifier validChain(uint64 chainSelector) {
    if (s_chains[chainSelector].extraArgsBytes.length == 0) revert InvalidChain(chainSelector);
    _;
  }

  modifier validSender(bytes calldata sender) {
    if (!s_senders[sender]) revert InvalidSender(sender);
    _;
  }

}
