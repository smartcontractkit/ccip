// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {CCIPTestSuite} from "./CCIPTestSuite.sol";
import {ForkedChain} from "./ForkedChain.sol";
import {ManyChainMultiSig} from "./ccip-owner-contracts/ManyChainMultiSig.sol";
import {RBACTimelock} from "./ccip-owner-contracts/RBACTimelock.sol";

import {console2} from "forge-std/Console2.sol";
import {Test} from "forge-std/Test.sol";

contract ChainLoading is Test {
  string internal constant MAINNET = "MAINNET";
  string internal constant ARBITRUM = "ARBITRUM";

  uint256 internal constant TwentyFiveHours = 25 * 60 * 60;

  string[] internal chainNames = [MAINNET];

  mapping(string chainName => ForkedChainTestSetup) public s_chains;

  struct ForkedChainTestSetup {
    MCMSSetup mcmsSetup;
    CCIPSetup ccipSetup;
    string name;
    uint256 forkId;
    CCIPTestSuite testSuite;
  }

  struct MCMSSetup {
    ManyChainMultiSig mcms;
    bytes mcmcPayload;
    RBACTimelock callProxy;
    bytes callProxyPayload;
  }

  struct CCIPSetup {
    address router;
  }

  function test_Sepolia() public {
    run(MAINNET);
  }

  function run(
    string memory chainName
  ) public {
    _loadSingleChain(chainName);

    // activate one
    ForkedChainTestSetup memory chain = _activateFork(chainName);

    console2.logString(" +------------------------------------------------+");
    console2.logString(" |                Before migration                |");
    console2.logString(" +------------------------------------------------+");

    chain.testSuite.sendAllTokens(false);

    _setProposalOnMCMS(chain.mcmsSetup);
    vm.warp(block.timestamp + TwentyFiveHours);
    _executeProposalOnTimeLock(chain.mcmsSetup);

    //    console2.logString(" +------------------------------------------------+");
    //    console2.logString(" |          Immediately after migration           |");
    //    console2.logString(" +------------------------------------------------+");
    //    vm.rollFork(6807978);
    //    chain.testSuite.sendAllTokens(true);

    console2.logString(" +------------------------------------------------+");
    console2.logString(" |                  Latest block                  |");
    console2.logString(" +------------------------------------------------+");
    vm.rollFork(6904314);
    chain.testSuite.sendAllTokens(true);
  }

  function _setProposalOnMCMS(
    MCMSSetup memory chain
  ) internal {
    // TODO
  }

  function _executeProposalOnTimeLock(
    MCMSSetup memory chain
  ) internal {
    // TODO
    //vm.rollFork(6807978);
    //    vm.rollFork(6904314);
  }

  function _loadAllChains() internal {
    for (uint256 i = 0; i < chainNames.length; ++i) {
      _loadSingleChain(chainNames[i]);
    }
  }

  function _loadSingleChain(
    string memory name
  ) internal returns (ForkedChainTestSetup memory) {
    ForkedChainTestSetup memory setup;
    //    setup.mcmsSetup.mcms = ManyChainMultiSig(payable(vm.envAddress(string.concat(name, "_MCMS"))));
    //    setup.mcmsSetup.mcmcPayload = vm.envBytes(string.concat(name, "_MCMS_PAYLOAD"));
    //    setup.mcmsSetup.callProxy = RBACTimelock(payable(vm.envAddress(string.concat(name, "_CALL_PROXY"))));
    //    setup.mcmsSetup.callProxyPayload = vm.envBytes(string.concat(name, "_CALL_PROXY_PAYLOAD"));

    setup.ccipSetup.router = vm.envAddress(string.concat(name, "_ROUTER"));

    setup.name = string(name);

    s_chains[name] = setup;
    return setup;
  }

  function _activateFork(
    string memory name
  ) internal returns (ForkedChainTestSetup memory) {
    ForkedChainTestSetup storage config = s_chains[name];
    if (config.forkId == 0) {
      config.forkId = ForkedChain.LoadAndActivateFork(vm, config.name);
    }

    console2.logString(string.concat("Activating chain: ", name));

    vm.selectFork(config.forkId);

    // If no test suite was deployed, deploy one
    if (address(config.testSuite) == address(0)) {
      config.testSuite = new CCIPTestSuite(config.ccipSetup.router);
    }

    return config;
  }
}
