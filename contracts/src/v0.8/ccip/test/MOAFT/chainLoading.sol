// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {console2} from "forge-std/Console2.sol";
import {Test} from "forge-std/Test.sol";

contract ChainLoading is Test {
  struct ChainConfig {
    string rpcUrl;
    address mcms;
    bytes mcmcPayload;
    address callProxy;
    bytes callProxyPayload;
    address router;
    uint256 block;
    string name;
    uint256 forkId;
  }

  function test() public {
    ChainConfig memory mainnet = _loadSingleChain("MAINNET");
    _activateFork(mainnet);

    _executeMCMS(mainnet.mcms, mainnet.mcmcPayload);
  }

  function _executeMCMS(address mcms, bytes memory payload) internal {
    if (mcms.code.length == 0) {
      revert("MCMS address is empty");
    }

    (bool success, bytes memory result) = mcms.call(payload);
    console2.logBytes(result);
    assertTrue(success, "MCMS call failed");
  }

  function _loadSingleChain(
    bytes memory name
  ) internal returns (ChainConfig memory) {
    ChainConfig memory chainConfig;
    chainConfig.rpcUrl = vm.envString(string(bytes.concat(name, "_RPC_URL")));
    chainConfig.mcms = vm.envAddress(string(bytes.concat(name, "_MCMS")));
    chainConfig.mcmcPayload = vm.envBytes(string(bytes.concat(name, "_MCMS_PAYLOAD")));
    chainConfig.callProxy = vm.envAddress(string(bytes.concat(name, "_CALL_PROXY")));
    chainConfig.callProxyPayload = vm.envBytes(string(bytes.concat(name, "_CALL_PROXY_PAYLOAD")));
    chainConfig.router = vm.envAddress(string(bytes.concat(name, "_ROUTER")));
    chainConfig.block = vm.envUint(string(bytes.concat(name, "_BLOCK")));
    chainConfig.name = string(name);
    return chainConfig;
  }

  function _activateFork(
    ChainConfig memory config
  ) internal {
    console2.logString("Activating chain");
    console2.logString(config.name);

    if (config.forkId == 0) {
      config.forkId = vm.createFork(config.rpcUrl);
    }

    vm.selectFork(config.forkId);
    vm.rollFork(config.block);
  }
}
