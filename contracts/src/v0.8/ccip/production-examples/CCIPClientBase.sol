pragma solidity ^0.8.0;

import {OwnerIsCreator} from "../../shared/access/OwnerIsCreator.sol";

contract CCIPClientBase is OwnerIsCreator {
  error InvalidRouter(address router);
  error InvalidChain(uint64 chainSelector);

  address internal immutable i_ccipRouter;
  mapping(uint64 destChainSelector => bytes extraArgsBytes) public s_chains;

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
  // Chain Management
  /////////////////////////////////////////////////////////////////////

  function enableChain(uint64 chainSelector, bytes memory extraArgs) external onlyOwner {
    s_chains[chainSelector] = extraArgs;
  }

  function disableChain(uint64 chainSelector) external onlyOwner {
    delete s_chains[chainSelector];
  }

  modifier validChain(uint64 chainSelector) {
    if (s_chains[chainSelector].length == 0) revert InvalidChain(chainSelector);
    _;
  }
}
