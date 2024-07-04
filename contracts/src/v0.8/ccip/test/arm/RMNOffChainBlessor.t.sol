// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RMNBase} from "../../RMNBase.sol";

import {IAny2EVMMultiOffRamp} from "../../interfaces/IAny2EVMMultiOffRamp.sol";
import {IRMN} from "../../interfaces/IRMN.sol";

import {
  GLOBAL_CURSE_SUBJECT,
  LIFT_CURSE_VOTE_ADDR,
  OWNER_CURSE_VOTE_ADDR,
  RMNOffChainBlessor
} from "../../RMNOffChainBlessor.sol";
import {RMNBaseSetup, makeSubjects} from "./RMNSetup.t.sol";

import {Test} from "forge-std/Test.sol";

bytes28 constant GARBAGE_CURSES_HASH = bytes28(keccak256("GARBAGE_CURSES_HASH"));

contract RMNOffBlessorSetup is RMNBaseSetup {
  RMNOffChainBlessor internal s_rmn;
  IAny2EVMMultiOffRamp.MerkleRoot internal s_root;

  function setUp() public virtual override {
    super.setUp();
    s_root.sourceChainSelector = 1;
    s_root.interval = IAny2EVMMultiOffRamp.Interval({min: 1, max: 10});
    s_root.merkleRoot = bytes32("merkle root");

    vm.startPrank(OWNER);
    s_rmn = new RMNOffChainBlessor(_rmnConstructorArgs());
    vm.stopPrank();
  }

  function _signRoot(uint256 voterKey) internal view returns (bytes memory) {
    (uint8 v, bytes32 r, bytes32 s) =
      vm.sign(voterKey, keccak256(abi.encode(s_root.merkleRoot, s_root.interval.min, s_root.interval.max)));
    return abi.encodePacked(r, s, v - 27);
  }
}

contract RMNOffChainBlessor_isBlessed is RMNOffBlessorSetup {
  function _getFirstBlessVoterAndWeight() internal view returns (address, uint8) {
    RMNBase.Config memory cfg = _rmnConstructorArgs();
    return (cfg.voters[0].blessVoteAddr, cfg.voters[0].blessWeight);
  }

  // Success

  function test_VotesBelowThreshold_Sucess() public {
    s_root.rmnSignatures.push(_signRoot(s_blessVoter1Key));

    assertFalse(s_rmn.isBlessed(s_root));
  }

  function test_VotesEqThreshold_Success() public {
    s_root.rmnSignatures.push(_signRoot(s_blessVoter2Key));
    s_root.rmnSignatures.push(_signRoot(s_blessVoter3Key));
    s_root.rmnSignatures.push(_signRoot(s_blessVoter4Key));

    assertTrue(s_rmn.isBlessed(s_root));
  }

  // Reverts

  function test_MustRecoverFromCurse_Revert() public {
    RMNBase.Config memory cfg = _rmnConstructorArgs();

    for (uint256 i = 0; i < cfg.voters.length; i++) {
      vm.startPrank(cfg.voters[i].curseVoteAddr);
      s_rmn.voteToCurse(makeCurseId(i), makeSubjects(GLOBAL_CURSE_SUBJECT));
    }

    vm.startPrank(cfg.voters[0].blessVoteAddr);

    s_root.rmnSignatures.push(_signRoot(s_blessVoter1Key));
    vm.expectRevert(RMNOffChainBlessor.MustRecoverFromCurse.selector);

    s_rmn.isBlessed(s_root);
  }

  function test_UnauthorizedVoter_Revert() public {
    (address unauthorizedVoter, uint256 unauthorizedVoterKey) = makeAddrAndKey("anauthorizedVoterKey");
    s_root.rmnSignatures.push(_signRoot(unauthorizedVoterKey));
    vm.expectRevert(abi.encodeWithSelector(RMNBase.UnauthorizedVoter.selector, unauthorizedVoter));

    s_rmn.isBlessed(s_root);
  }
}
