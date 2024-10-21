// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {Internal} from "../../libraries/Internal.sol";
import {CCIPTestSuite} from "./CCIPTestSuite.sol";
import {ForkedChain} from "./ForkedChain.sol";
import {ManyChainMultiSig} from "./ccip-owner-contracts/ManyChainMultiSig.sol";
import {RBACTimelock} from "./ccip-owner-contracts/RBACTimelock.sol";

import {console2} from "forge-std/Console2.sol";
import {Test} from "forge-std/Test.sol";

contract ChainLoading is Test {
  string internal constant SEPOLIA = "SEPOLIA";
  string internal constant ARBITRUM_SEPOLIA = "ARB_SEPOLIA";

  uint256 internal constant FourHours = 4 * 60 * 60;

  string[] internal chainNames = [SEPOLIA];

  mapping(string chainName => ForkedChainTestSetup) public s_chains;

  struct ForkedChainTestSetup {
    MCMSSetup mcmsSetup;
    CCIPSetup ccipSetup;
    string name;
    uint256 forkId;
    uint64 postMigrationBlock;
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

  function test_Chain_Sepolia() public {
    run(SEPOLIA);
  }

  function test_Chain_Arbitrum() public {
    run(ARBITRUM_SEPOLIA);
  }

  function test_Lane_SepoliaToArbitrum() public {
    _loadSingleChain(SEPOLIA);
    _loadSingleChain(ARBITRUM_SEPOLIA);

    ForkedChainTestSetup memory chain = _activateFork(SEPOLIA);

    // Apply proposal on source
    _setProposalOnMCMS(chain);
    _executeProposalOnTimeLock(chain);

    // Send messages
    Internal.EVM2EVMMessage[] memory msgs = chain.testSuite.sendTokensSingleLane(3478487238524512106);

    ForkedChainTestSetup memory destChain = _activateFork(ARBITRUM_SEPOLIA);

    // Apply proposal on dest
    _setProposalOnMCMS(destChain);
    _executeProposalOnTimeLock(destChain);

    // Execute messages
    destChain.testSuite.ExecuteMsgs(msgs);
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

    console2.logString(" +------------------------------------------------+");
    console2.logString(" |                After migration                 |");
    console2.logString(" +------------------------------------------------+");
    _setProposalOnMCMS(chain);
    _executeProposalOnTimeLock(chain);

    chain.testSuite.sendAllTokens(true);
  }

  function _setProposalOnMCMS(
    ForkedChainTestSetup memory chain
  ) internal {
    // TODO
  }

  function _executeProposalOnTimeLock(
    ForkedChainTestSetup memory chain
  ) internal {
    // TODO actual MCMS proposal execution. For testing, we can use chains that already have the proposal executed
    // and roll to a block after the migration.
    vm.rollFork(chain.postMigrationBlock);
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
    setup.postMigrationBlock = uint64(vm.envUint(string.concat(name, "_POST_BLOCK")));

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

    console2.log("Activating chain: ", name);

    vm.selectFork(config.forkId);

    // If no test suite was deployed, deploy one
    if (address(config.testSuite) == address(0)) {
      bool fullLogging = vm.envBool("FULL_LOGGING");
      config.testSuite = new CCIPTestSuite(config.ccipSetup.router, fullLogging);
    }

    return config;
  }
}
