// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./AFNSetup.t.sol";

contract ConfigCompare is Test {
  function assertConfigEq(AFN.Config memory actualConfig, AFN.Config memory expectedConfig) public {
    assertEq(actualConfig.voters.length, expectedConfig.voters.length);
    for (uint256 i = 0; i < expectedConfig.voters.length; ++i) {
      AFN.Voter memory expectedVoter = expectedConfig.voters[i];
      AFN.Voter memory actualVoter = actualConfig.voters[i];
      assertEq(actualVoter.blessVoteAddr, expectedVoter.blessVoteAddr);
      assertEq(actualVoter.curseVoteAddr, expectedVoter.curseVoteAddr);
      assertEq(actualVoter.blessWeight, expectedVoter.blessWeight);
      assertEq(actualVoter.curseWeight, expectedVoter.curseWeight);
    }
    assertEq(actualConfig.blessWeightThreshold, expectedConfig.blessWeightThreshold);
    assertEq(actualConfig.curseWeightThreshold, expectedConfig.curseWeightThreshold);
  }
}

contract AFN_constructor is ConfigCompare, AFNSetup {
  function testConstructorSuccess() public {
    AFN.Config memory expectedConfig = afnConstructorArgs();
    (uint32 actualVersion, , AFN.Config memory actualConfig) = s_afn.getConfigDetails();
    assertEq(actualVersion, 1);
    assertConfigEq(actualConfig, expectedConfig);
  }
}

contract AFN_voteToBlessRoots is AFNSetup {
  event VoteToBless(uint32 indexed configVersion, address indexed voter, AFN.TaggedRoot taggedRoot, uint8 weight);

  // Success

  function _getFirstBlessVoterAndWeight() internal pure returns (address, uint8) {
    AFN.Config memory cfg = afnConstructorArgs();
    return (cfg.voters[0].blessVoteAddr, cfg.voters[0].blessWeight);
  }

  function test1RootSuccess_gas() public {
    vm.pauseGasMetering();
    (address voter, uint8 voterWeight) = _getFirstBlessVoterAndWeight();

    vm.expectEmit();
    emit VoteToBless(1, voter, makeTaggedRoot(1), voterWeight);

    changePrank(voter);
    vm.resumeGasMetering();
    s_afn.voteToBless(makeTaggedRootSingleton(1));
    vm.pauseGasMetering();

    assertFalse(s_afn.isBlessed(makeTaggedRootHash(1)));
    assertEq(voterWeight, getWeightOfVotesToBlessRoot(makeTaggedRoot(1)));
    assertTrue(hasVotedToBlessRoot(voter, makeTaggedRoot(1)));
    vm.resumeGasMetering();
  }

  function test3RootSuccess_gas() public {
    vm.pauseGasMetering();
    (address voter, uint8 voterWeight) = _getFirstBlessVoterAndWeight();

    for (uint256 i = 1; i <= 3; ++i) {
      vm.expectEmit();
      emit VoteToBless(1, voter, makeTaggedRoot(i), voterWeight);
    }

    changePrank(voter);
    vm.resumeGasMetering();
    s_afn.voteToBless(makeTaggedRootsInclusive(1, 3));
    vm.pauseGasMetering();

    for (uint256 i = 1; i <= 3; ++i) {
      assertFalse(s_afn.isBlessed(makeTaggedRootHash(i)));
      assertEq(voterWeight, getWeightOfVotesToBlessRoot(makeTaggedRoot(i)));
      assertTrue(hasVotedToBlessRoot(voter, makeTaggedRoot(i)));
    }
    vm.resumeGasMetering();
  }

  function test5RootSuccess_gas() public {
    vm.pauseGasMetering();
    (address voter, uint8 voterWeight) = _getFirstBlessVoterAndWeight();

    for (uint256 i = 1; i <= 5; ++i) {
      vm.expectEmit();
      emit VoteToBless(1, voter, makeTaggedRoot(i), voterWeight);
    }

    changePrank(voter);
    vm.resumeGasMetering();
    s_afn.voteToBless(makeTaggedRootsInclusive(1, 5));
    vm.pauseGasMetering();

    for (uint256 i = 1; i <= 5; ++i) {
      assertFalse(s_afn.isBlessed(makeTaggedRootHash(i)));
      assertEq(voterWeight, getWeightOfVotesToBlessRoot(makeTaggedRoot(i)));
      assertTrue(hasVotedToBlessRoot(voter, makeTaggedRoot(1)));
    }
    vm.resumeGasMetering();
  }

  function testIsAlreadyBlessedIgnoredSuccess() public {
    AFN.Config memory cfg = afnConstructorArgs();

    // Bless voters 2,3,4 vote to bless
    for (uint256 i = 1; i < cfg.voters.length; i++) {
      changePrank(cfg.voters[i].blessVoteAddr);
      s_afn.voteToBless(makeTaggedRootSingleton(1));
    }

    uint256 votesToBlessBefore = getWeightOfVotesToBlessRoot(makeTaggedRoot(1));
    changePrank(cfg.voters[0].blessVoteAddr);
    s_afn.voteToBless(makeTaggedRootSingleton(1));
    assertEq(votesToBlessBefore, getWeightOfVotesToBlessRoot(makeTaggedRoot(1)));
  }

  function testSenderAlreadyVotedIgnoredSuccess() public {
    (address voter, ) = _getFirstBlessVoterAndWeight();

    changePrank(voter);
    s_afn.voteToBless(makeTaggedRootSingleton(1));
    assertTrue(hasVotedToBlessRoot(voter, makeTaggedRoot(1)));

    uint256 votesToBlessBefore = getWeightOfVotesToBlessRoot(makeTaggedRoot(1));
    s_afn.voteToBless(makeTaggedRootSingleton(1));
    assertEq(votesToBlessBefore, getWeightOfVotesToBlessRoot(makeTaggedRoot(1)));
  }

  // Reverts

  function testCurseReverts() public {
    AFN.Config memory cfg = afnConstructorArgs();

    for (uint256 i = 0; i < cfg.voters.length; i++) {
      changePrank(cfg.voters[i].curseVoteAddr);
      s_afn.voteToCurse(makeCurseId(i));
    }

    changePrank(cfg.voters[0].blessVoteAddr);
    vm.expectRevert(AFN.MustRecoverFromCurse.selector);
    s_afn.voteToBless(makeTaggedRootSingleton(12903));
  }

  function testInvalidVoterReverts() public {
    changePrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(AFN.InvalidVoter.selector, STRANGER));
    s_afn.voteToBless(makeTaggedRootSingleton(12321));
  }
}

contract AFN_ownerUnbless is AFNSetup {
  function testUnblessSuccess() public {
    AFN.Config memory cfg = afnConstructorArgs();
    for (uint256 i = 0; i < cfg.voters.length; ++i) {
      changePrank(cfg.voters[i].blessVoteAddr);
      s_afn.voteToBless(makeTaggedRootSingleton(1));
    }
    assertTrue(s_afn.isBlessed(makeTaggedRootHash(1)));

    changePrank(OWNER);
    s_afn.ownerUnbless(makeTaggedRootSingleton(1));
    assertFalse(s_afn.isBlessed(makeTaggedRootHash(1)));
  }
}

contract AFN_unvoteToCurse is AFNSetup {
  uint256 s_curser;
  bytes32 s_cursesHash;

  function setUp() public override {
    AFN.Config memory cfg = afnConstructorArgs();
    AFNSetup.setUp();
    cfg = afnConstructorArgs();
    s_curser = 0;

    changePrank(cfg.voters[0].curseVoteAddr);
    s_afn.voteToCurse(makeCurseId(1));
    assertFalse(s_afn.badSignalReceived());
    (address[] memory cursers, uint16 weight, uint32[] memory voteCounts) = s_afn.getCurseVotersAndWeight();
    assertEq(1, cursers.length);
    assertEq(1, voteCounts.length);
    assertEq(cfg.voters[s_curser].curseVoteAddr, cursers[0]);
    assertEq(1, voteCounts[0]);
    assertEq(cfg.voters[s_curser].curseWeight, weight);

    s_cursesHash = keccak256(abi.encode(bytes32(0), block.chainid, blockhash(block.number - 1), makeCurseId(1)));
  }

  function testInvalidVoter() public {
    AFN.Config memory cfg = afnConstructorArgs();
    // Someone else cannot unvote to curse on the curser's behalf.
    address[] memory unauthorized = new address[](4);
    unauthorized[0] = cfg.voters[s_curser].blessVoteAddr;
    unauthorized[1] = cfg.voters[s_curser].curseVoteAddr;
    unauthorized[2] = OWNER;
    unauthorized[3] = cfg.voters[s_curser ^ 1].curseUnvoteAddr;

    for (uint256 i = 0; i < unauthorized.length; ++i) {
      bytes memory expectedRevert = abi.encodeWithSelector(AFN.InvalidVoter.selector, unauthorized[i]);
      changePrank(unauthorized[i]);
      // should fail when using the correct curses hash
      vm.expectRevert(expectedRevert);
      s_afn.unvoteToCurse(cfg.voters[s_curser].curseVoteAddr, s_cursesHash);
      // should fail when using garbage curses hash
      vm.expectRevert(expectedRevert);
      s_afn.unvoteToCurse(
        cfg.voters[s_curser].curseVoteAddr,
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
      );
    }
  }

  function testInvalidCursesHash() public {
    AFN.Config memory cfg = afnConstructorArgs();
    changePrank(cfg.voters[s_curser].curseUnvoteAddr);
    vm.expectRevert(
      abi.encodeWithSelector(
        AFN.InvalidCursesHash.selector,
        s_cursesHash,
        0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
      )
    );
    s_afn.unvoteToCurse(
      cfg.voters[s_curser].curseVoteAddr,
      0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
    );
  }

  function testValidCursesHash() public {
    AFN.Config memory cfg = afnConstructorArgs();
    changePrank(cfg.voters[s_curser].curseUnvoteAddr);
    s_afn.unvoteToCurse(cfg.voters[s_curser].curseVoteAddr, s_cursesHash);
  }

  function testOwnerSucceeds() public {
    AFN.Config memory cfg = afnConstructorArgs();
    changePrank(OWNER);
    AFN.UnvoteToCurseRecord[] memory records = new AFN.UnvoteToCurseRecord[](1);
    records[0] = AFN.UnvoteToCurseRecord({
      curseVoteAddr: cfg.voters[s_curser].curseUnvoteAddr,
      cursesHash: s_cursesHash,
      forceUnvote: false
    });
    s_afn.ownerUnvoteToCurse(records);
  }

  event SkippedUnvoteToCurse(address indexed voter, bytes32 expectedCursesHash, bytes32 actualCursesHash);

  function testOwnerSkips() public {
    AFN.Config memory cfg = afnConstructorArgs();
    changePrank(OWNER);
    AFN.UnvoteToCurseRecord[] memory records = new AFN.UnvoteToCurseRecord[](1);
    records[0] = AFN.UnvoteToCurseRecord({
      curseVoteAddr: cfg.voters[s_curser].curseVoteAddr,
      cursesHash: 0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff,
      forceUnvote: false
    });
    vm.expectEmit();
    emit SkippedUnvoteToCurse(
      cfg.voters[s_curser].curseVoteAddr,
      s_cursesHash,
      0xffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff
    );
    s_afn.ownerUnvoteToCurse(records);
  }
}

contract AFN_voteToCurse is AFNSetup {
  event VoteToCurse(
    uint32 indexed configVersion,
    address indexed voter,
    uint8 weight,
    uint32 voteCount,
    bytes32 curseId,
    bytes32 cursesHash,
    uint16 accumulatedWeight
  );
  event Cursed(uint32 indexed configVersion, uint256 timestamp);

  function _getFirstCurseVoterAndWeight() internal pure returns (address, uint8) {
    AFN.Config memory cfg = afnConstructorArgs();
    return (cfg.voters[0].curseVoteAddr, cfg.voters[0].curseWeight);
  }

  // Success

  function testVoteToCurseSuccess_gas() public {
    vm.pauseGasMetering();

    (address voter, uint8 weight) = _getFirstCurseVoterAndWeight();
    changePrank(voter);
    vm.expectEmit();
    emit VoteToCurse(
      1,
      voter,
      weight,
      1,
      makeCurseId(123),
      keccak256(abi.encode(bytes32(0), block.chainid, blockhash(block.number - 1), makeCurseId(123))),
      weight
    );

    vm.resumeGasMetering();
    s_afn.voteToCurse(makeCurseId(123));
    vm.pauseGasMetering();

    (address[] memory voters, uint16 votes, ) = s_afn.getCurseVotersAndWeight();
    assertEq(1, voters.length);
    assertEq(voter, voters[0]);
    assertEq(weight, votes);

    vm.resumeGasMetering();
  }

  function testEmitCurseSuccess() public {
    AFN.Config memory cfg = afnConstructorArgs();
    for (uint256 i = 0; i < cfg.voters.length - 1; ++i) {
      changePrank(cfg.voters[i].curseVoteAddr);
      s_afn.voteToCurse(makeCurseId(1));
    }

    vm.expectEmit();
    emit Cursed(1, block.timestamp);

    changePrank(cfg.voters[cfg.voters.length - 1].curseVoteAddr);
    s_afn.voteToCurse(makeCurseId(1));
  }

  function testEvenIfAlreadyCursedSuccess() public {
    AFN.Config memory cfg = afnConstructorArgs();
    for (uint256 i = 0; i < cfg.voters.length; ++i) {
      changePrank(cfg.voters[i].curseVoteAddr);
      s_afn.voteToCurse(makeCurseId(i));
    }

    // Not part of the assertion of this test but good to have as a sanity
    // check. We want a curse to be active in order for the ultimate assertion
    // to make sense.
    assert(s_afn.badSignalReceived());

    // Asserts that this call to vote with a new curse id goes through with no
    // reverts even when the AFN contract is cursed.
    s_afn.voteToCurse(makeCurseId(cfg.voters.length + 1));
  }

  // Reverts

  function testInvalidVoterReverts() public {
    changePrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(AFN.InvalidVoter.selector, STRANGER));
    s_afn.voteToCurse(makeCurseId(12312));
  }

  function testAlreadyVotedReverts() public {
    (address voter, ) = _getFirstCurseVoterAndWeight();
    changePrank(voter);
    s_afn.voteToCurse(makeCurseId(1));

    vm.expectRevert(abi.encodeWithSelector(AFN.AlreadyVotedToCurse.selector, voter, makeCurseId(1)));
    s_afn.voteToCurse(makeCurseId(1));
  }
}

contract AFN_ownerUnvoteToCurse is AFNSetup {
  event RecoveredFromCurse();

  // These cursers are going to curse in setUp curseCount times.
  function getCursersAndCurseCounts() internal pure returns (address[] memory cursers, uint32[] memory curseCounts) {
    // NOTE: Change this when changing setUp or afnConstructorArgs.
    // This is a bit ugly and error prone but if we read from storage we would
    // not get an accurate gas reading for ownerUnvoteToCurse when we need it.
    cursers = new address[](4);
    cursers[0] = CURSE_VOTER_1;
    cursers[1] = CURSE_VOTER_2;
    cursers[2] = CURSE_VOTER_3;
    cursers[3] = CURSE_VOTER_4;
    curseCounts = new uint32[](cursers.length);
    for (uint256 i = 0; i < cursers.length; ++i) {
      curseCounts[i] = 1;
    }
  }

  function setUp() public virtual override {
    AFNSetup.setUp();
    (address[] memory cursers, uint32[] memory curseCounts) = getCursersAndCurseCounts();
    for (uint256 i = 0; i < cursers.length; ++i) {
      changePrank(cursers[i]);
      for (uint256 j = 0; j < curseCounts[i]; ++j) {
        s_afn.voteToCurse(makeCurseId(j));
      }
    }
  }

  function ownerUnvoteToCurse() internal {
    s_afn.ownerUnvoteToCurse(makeUnvoteToCurseRecords());
  }

  function makeUnvoteToCurseRecords() internal pure returns (AFN.UnvoteToCurseRecord[] memory) {
    (address[] memory cursers, ) = getCursersAndCurseCounts();
    AFN.UnvoteToCurseRecord[] memory records = new AFN.UnvoteToCurseRecord[](cursers.length);
    for (uint256 i = 0; i < cursers.length; ++i) {
      records[i] = AFN.UnvoteToCurseRecord({
        curseVoteAddr: cursers[i],
        cursesHash: bytes32(uint256(0)),
        forceUnvote: true
      });
    }
    return records;
  }

  // Success

  function testOwnerUnvoteToCurseSuccess_gas() public {
    vm.pauseGasMetering();
    changePrank(OWNER);

    vm.expectEmit();
    emit RecoveredFromCurse();

    vm.resumeGasMetering();
    ownerUnvoteToCurse();
    vm.pauseGasMetering();

    assertFalse(s_afn.badSignalReceived());
    (address[] memory voters, uint256 weight, ) = s_afn.getCurseVotersAndWeight();
    assertEq(voters.length, 0);
    assertEq(weight, 0);
    vm.resumeGasMetering();
  }

  // Reverts

  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    ownerUnvoteToCurse();
  }
}

contract AFN_setConfig is ConfigCompare, AFNSetup {
  /// @notice Test-specific function to use only in setConfig tests
  function getDifferentConfigArgs() private pure returns (AFN.Config memory) {
    AFN.Voter[] memory voters = new AFN.Voter[](2);
    voters[0] = AFN.Voter({
      blessVoteAddr: BLESS_VOTER_1,
      curseVoteAddr: CURSE_VOTER_1,
      curseUnvoteAddr: CURSE_UNVOTER_1,
      blessWeight: WEIGHT_1,
      curseWeight: WEIGHT_1
    });
    voters[1] = AFN.Voter({
      blessVoteAddr: BLESS_VOTER_2,
      curseVoteAddr: CURSE_VOTER_2,
      curseUnvoteAddr: CURSE_UNVOTER_2,
      blessWeight: WEIGHT_10,
      curseWeight: WEIGHT_10
    });
    return
      AFN.Config({
        voters: voters,
        blessWeightThreshold: WEIGHT_1 + WEIGHT_10,
        curseWeightThreshold: WEIGHT_1 + WEIGHT_10
      });
  }

  function setUp() public virtual override {
    AFNSetup.setUp();
    AFN.Config memory cfg = afnConstructorArgs();

    // Setup some partial state
    changePrank(cfg.voters[0].blessVoteAddr);
    s_afn.voteToBless(makeTaggedRootSingleton(1));
    changePrank(cfg.voters[1].blessVoteAddr);
    s_afn.voteToBless(makeTaggedRootSingleton(1));
    changePrank(cfg.voters[1].curseVoteAddr);
    s_afn.voteToCurse(makeCurseId(1));
  }

  // Success

  event ConfigSet(uint32 indexed configVersion, AFN.Config config);

  function testVoteToBlessByEjectedVoterReverts() public {
    // Previous config included BLESS_VOTER_4. Change to new config that doesn't.
    AFN.Config memory cfg = getDifferentConfigArgs();
    changePrank(OWNER);
    s_afn.setConfig(cfg);

    // BLESS_VOTER_4 is not part of cfg anymore, vote to bless should revert.
    changePrank(BLESS_VOTER_4);
    vm.expectRevert(abi.encodeWithSelector(AFN.InvalidVoter.selector, BLESS_VOTER_4));
    s_afn.voteToBless(makeTaggedRootSingleton(2));
  }

  function testSetConfigSuccess_gas() public {
    vm.pauseGasMetering();
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(OWNER);
    vm.expectEmit();
    emit ConfigSet(2, cfg);

    (uint32 configVersionBefore, , ) = s_afn.getConfigDetails();
    vm.resumeGasMetering();
    s_afn.setConfig(cfg);
    vm.pauseGasMetering();
    // Assert VersionedConfig has changed correctly
    (uint32 configVersionAfter, , AFN.Config memory configAfter) = s_afn.getConfigDetails();
    assertEq(configVersionBefore + 1, configVersionAfter);
    assertConfigEq(configAfter, cfg);

    // Assert that curse votes have been cleared, except for CURSE_VOTER_2 who
    // has already voted and is also part of the new config
    (address[] memory curseVoters, uint256 curseWeight, ) = s_afn.getCurseVotersAndWeight();
    assertEq(1, curseVoters.length);
    assertEq(WEIGHT_10, curseWeight);

    // Assert that good votes have been cleared
    uint256 votesToBlessRoot = getWeightOfVotesToBlessRoot(makeTaggedRoot(1));
    assertEq(ZERO, votesToBlessRoot);
    assertFalse(hasVotedToBlessRoot(cfg.voters[0].blessVoteAddr, makeTaggedRoot(1)));
    assertFalse(hasVotedToBlessRoot(cfg.voters[1].blessVoteAddr, makeTaggedRoot(1)));
    vm.resumeGasMetering();
  }

  // Reverts

  function testNonOwnerReverts() public {
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_afn.setConfig(cfg);
  }

  function testVotersLengthIsZeroReverts() public {
    changePrank(OWNER);
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(AFN.Config({voters: new AFN.Voter[](0), blessWeightThreshold: 1, curseWeightThreshold: 1}));
  }

  function testEitherThresholdIsZeroReverts() public {
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(OWNER);
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(
      AFN.Config({voters: cfg.voters, blessWeightThreshold: ZERO, curseWeightThreshold: cfg.curseWeightThreshold})
    );
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(
      AFN.Config({voters: cfg.voters, blessWeightThreshold: cfg.blessWeightThreshold, curseWeightThreshold: ZERO})
    );
  }

  function testBlessVoterIsZeroAddressReverts() public {
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(OWNER);
    cfg.voters[0].blessVoteAddr = ZERO_ADDRESS;
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(cfg);
  }

  function testWeightIsZeroAddressReverts() public {
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(OWNER);
    cfg.voters[0].blessWeight = ZERO;
    cfg.voters[0].curseWeight = ZERO;
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(cfg);
  }

  function testTotalWeightsSmallerThanEachThresholdReverts() public {
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(OWNER);
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(
      AFN.Config({voters: cfg.voters, blessWeightThreshold: WEIGHT_40, curseWeightThreshold: cfg.curseWeightThreshold})
    );
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(
      AFN.Config({voters: cfg.voters, blessWeightThreshold: cfg.blessWeightThreshold, curseWeightThreshold: WEIGHT_40})
    );
  }

  function testRepeatedAddressReverts() public {
    AFN.Config memory cfg = getDifferentConfigArgs();

    changePrank(OWNER);
    cfg.voters[0].blessVoteAddr = cfg.voters[1].curseVoteAddr;
    vm.expectRevert(AFN.InvalidConfig.selector);
    s_afn.setConfig(cfg);
  }
}
