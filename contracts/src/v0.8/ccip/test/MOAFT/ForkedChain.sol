// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {console2} from "forge-std/Console2.sol";
import {Vm} from "forge-std/Vm.sol";

library ForkedChain {
  struct Chain {
    string rpcUrl;
    uint256 block;
    string name;
  }

  function LoadAndActivateFork(Vm vm, string memory name) public returns (uint256 forkId) {
    return _activateNewForkInstance(vm, _loadChainConfig(vm, name));
  }

  function _loadChainConfig(Vm vm, string memory name) internal returns (Chain memory) {
    Chain memory chain;
    chain.rpcUrl = vm.envString(string.concat(name, "_RPC_URL"));
    chain.block = vm.envUint(string.concat(name, "_BLOCK"));
    chain.name = name;
    return chain;
  }

  function _activateNewForkInstance(Vm vm, Chain memory chain) internal returns (uint256 forkId) {
    console2.logString(string.concat("Launching new fork instance for chain: ", chain.name));

    return vm.createFork(chain.rpcUrl, chain.block);
  }
}
