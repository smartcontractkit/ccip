// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {IOwnable} from "../../../shared/interfaces/IOwnable.sol";
import {Internal} from "../../libraries/Internal.sol";
import {CCIPTestSuite} from "./CCIPTestSuite.sol";
import {ChainSelectors} from "./ChainSelectors.sol";
import {ForkedChain} from "./ForkedChain.sol";
import {MCMS} from "./MCMS.sol";

import {console2} from "forge-std/Console2.sol";
import {Test} from "forge-std/Test.sol";

/// @notice This test tests all tokens on all chains in a forked chain environment.
/// @dev Populate a .env file with the following variables:
/// # Global
/// FULL_LOGGING=false
///
/// # Has to match the network names as defined below
/// SOURCE_CHAIN=<chain>
/// DEST_CHAIN=<chain>
///
/// # Networks
/// # The RPC URL for the chain, when using older pre/post migration blocks, this will need to be an archive node.
/// <chain>_RPC_URL=''
/// # The Router contract address
/// <chain>_ROUTER=0x0BF3dE8c5D3e8A2B34D2BEeB17ABfCeBaf363A59
/// # Any block before the migration was applied.
/// <chain>_PRE_BLOCK=6797746
/// # Any block after the migration was applied. Use 0 if the migration has not been applied yet.
/// <chain>_POST_BLOCK=6904314
contract ChainLoading is Test {
  string[] public AllTestnets = [
    ChainSelectors.SEPOLIA,
    ChainSelectors.GNOSIS_TESTNET,
    ChainSelectors.BNB_TESTNET,
    ChainSelectors.MODE_TESTNET,
    ChainSelectors.OPT_SEPOLIA,
    ChainSelectors.POLYGON_AMOY,
    ChainSelectors.ARB_SEPOLIA,
    ChainSelectors.AVAX_FUJI,
    ChainSelectors.BASE_SEPOLIA
  ];
  string[] public AllMainnets = [
    ChainSelectors.BLAST,
    ChainSelectors.ETHEREUM,
    ChainSelectors.GNOSIS,
    ChainSelectors.BNB,
    ChainSelectors.MODE,
    ChainSelectors.OPTIMISM,
    ChainSelectors.POLYGON,
    ChainSelectors.ARBITRUM,
    ChainSelectors.AVAX,
    ChainSelectors.BASE
  ];

  uint256 internal constant FourHours = 4 * 60 * 60;

  mapping(string chainName => ForkedChainTestSetup) public s_chains;

  struct ForkedChainTestSetup {
    CCIPSetup ccipSetup;
    string name;
    uint256 forkId;
    uint64 postMigrationBlock;
    CCIPTestSuite testSuite;
  }

  struct CCIPSetup {
    address router;
  }

  function test_env_chain() public {
    run(vm.envString("SOURCE_CHAIN"));
  }

  function test_env_lane() public {
    string memory destChain = ChainSelectors._resolveChainSelector(uint64(vm.envUint("REMOTE_CHAIN_SELECTOR")));
    _run_lane(vm.envString("SOURCE_CHAIN"), destChain);
  }

  function test_all_chains() public {
    string[] memory chains = AllMainnets;
    for (uint256 i = 0; i < chains.length; ++i) {
      run(chains[i]);
    }
  }

  function _run_lane(string memory sourceChainName, string memory destChainName) internal {
    _loadSingleChain(sourceChainName);
    _loadSingleChain(destChainName);
    uint64 remoteChainSelector = uint64(vm.envUint("REMOTE_CHAIN_SELECTOR"));

    ForkedChainTestSetup memory sourceChain = _activateFork(sourceChainName);

    // Apply proposal on source
    _executeProposalOnTimeLock(sourceChain);

    // Send messages
    Internal.EVM2EVMMessage[] memory msgs = sourceChain.testSuite.sendTokensSingleLane(remoteChainSelector);

    ForkedChainTestSetup memory destChain = _activateFork(destChainName);

    // Apply proposal on dest
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

    if (chain.postMigrationBlock == 0) {
      console2.logString("Migration not applied yet");

      _executeProposalOnTimeLock(chain);
    } else {
      // To find the post-migration block, search the Router for the RouterOnRampSet the event sig
      // 0x1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23
      // There should be multiple events in the same block during the migration.
      vm.rollFork(chain.postMigrationBlock);
    }

    console2.logString(" +------------------------------------------------+");
    console2.logString(" |                After migration                 |");
    console2.logString(" +------------------------------------------------+");

    chain.testSuite.sendAllTokens(true);
  }

  function _executeProposalOnTimeLock(
    ForkedChainTestSetup memory chain
  ) internal {
    if (chain.postMigrationBlock != 0) {
      console2.logString("Rolling to post-migration block");
      // To find the post-migration block, search the Router for the RouterOnRampSet the event sig
      // 0x1f7d0ec248b80e5c0dde0ee531c4fc8fdb6ce9a2b3d90f560c74acd6a7202f23
      // There should be multiple events in the same block during the migration.
      vm.rollFork(chain.postMigrationBlock);

      return;
    }

    bytes memory payload = vm.envBytes(string.concat(chain.name, "_PAYLOAD"));
    (MCMS.Call[] memory calls,,,) = abi.decode(payload, (MCMS.Call[], bytes32, bytes32, uint256));

    if (calls.length == 0) {
      console2.logString("No migration to apply");
      return;
    }
    // All contracts are owned by the same contract
    changePrank(IOwnable(calls[0].target).owner());

    for (uint256 i = 0; i < calls.length; ++i) {
      MCMS.Call memory call = calls[i];

      // Exclude GHO
      if (
        call.target == 0x5756880B6a1EAba0175227bf02a7E87c1e02B28C
          || call.target == 0xF168B83598516A532a85995b52504a2Fa058C068
      ) {
        continue;
      }

      (bool success,) = call.target.call{value: call.value}(call.data);
      require(success, "RBACTimelock: underlying transaction reverted");
    }

    console2.logString("Migration not applied yet");
  }

  function _loadSingleChain(
    string memory name
  ) internal returns (ForkedChainTestSetup memory) {
    ForkedChainTestSetup memory setup;

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
