// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./AFNSetup.t.sol";

contract AFN_constructor is AFNSetup {
  function testSuccess() public {
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 basSignalThreshold
    ) = afnConstructorArgs();

    assertEq(participants, s_afn.getParticipants());
    for (uint256 i = 0; i < participants.length; ++i) {
      assertEq(weights[i], s_afn.getWeightByParticipant(participants[i]));
    }
    (uint256 blessing, uint256 badSignal) = s_afn.getWeightThresholds();
    assertEq(blessingThreshold, blessing);
    assertEq(basSignalThreshold, badSignal);
  }
}

contract AFN_voteToBlessRoots is AFNSetup {
  event VoteToBless(address indexed voter, bytes32 indexed root, uint256 weight);

  // Success

  function test1RootSuccess_gas() public {
    vm.pauseGasMetering();
    (address[] memory participants, uint256[] memory weights, , ) = afnConstructorArgs();

    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_1, weights[0]);

    bytes32[] memory roots = new bytes32[](1);
    roots[0] = ROOT_1;

    changePrank(participants[0]);

    vm.resumeGasMetering();
    s_afn.voteToBlessRoots(roots);
    vm.pauseGasMetering();

    assertFalse(s_afn.isBlessed(ROOT_1));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_1));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_1));
    vm.resumeGasMetering();
  }

  function test3RootSuccess_gas() public {
    vm.pauseGasMetering();
    (address[] memory participants, uint256[] memory weights, , ) = afnConstructorArgs();

    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_1, weights[0]);
    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_2, weights[0]);
    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_3, weights[0]);

    bytes32[] memory roots = new bytes32[](3);
    roots[0] = ROOT_1;
    roots[1] = ROOT_2;
    roots[2] = ROOT_3;

    changePrank(participants[0]);
    vm.resumeGasMetering();
    s_afn.voteToBlessRoots(roots);
    vm.pauseGasMetering();

    assertFalse(s_afn.isBlessed(ROOT_1));
    assertFalse(s_afn.isBlessed(ROOT_2));
    assertFalse(s_afn.isBlessed(ROOT_3));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_1));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_2));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_3));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_1));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_2));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_3));
    vm.resumeGasMetering();
  }

  function test5RootSuccess_gas() public {
    vm.pauseGasMetering();
    (address[] memory participants, uint256[] memory weights, , ) = afnConstructorArgs();

    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_1, weights[0]);
    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_2, weights[0]);
    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_3, weights[0]);
    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_4, weights[0]);
    vm.expectEmit();
    emit VoteToBless(participants[0], ROOT_5, weights[0]);

    bytes32[] memory roots = new bytes32[](5);
    roots[0] = ROOT_1;
    roots[1] = ROOT_2;
    roots[2] = ROOT_3;
    roots[3] = ROOT_4;
    roots[4] = ROOT_5;

    changePrank(participants[0]);
    vm.resumeGasMetering();
    s_afn.voteToBlessRoots(roots);
    vm.pauseGasMetering();

    assertFalse(s_afn.isBlessed(ROOT_1));
    assertFalse(s_afn.isBlessed(ROOT_2));
    assertFalse(s_afn.isBlessed(ROOT_3));
    assertFalse(s_afn.isBlessed(ROOT_4));
    assertFalse(s_afn.isBlessed(ROOT_5));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_1));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_2));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_3));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_4));
    assertEq(weights[0], s_afn.getVotesToBlessRoot(ROOT_5));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_1));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_2));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_3));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_4));
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_5));
    vm.resumeGasMetering();
  }

  function testIsAlreadyBlessedIgnoredSuccess() public {
    (address[] memory participants, , , ) = afnConstructorArgs();
    bytes32[] memory roots = new bytes32[](1);
    roots[0] = ROOT_1;

    // Participants 2,3,4 vote to bless
    for (uint256 i = 1; i < participants.length; i++) {
      changePrank(participants[i]);
      s_afn.voteToBlessRoots(roots);
    }

    uint256 votesToBlessBefore = s_afn.getVotesToBlessRoot(ROOT_1);
    changePrank(participants[0]);
    s_afn.voteToBlessRoots(roots);
    assertEq(votesToBlessBefore, s_afn.getVotesToBlessRoot(ROOT_1));
  }

  function testSenderAlreadyVotedIgnoredSuccess() public {
    (address[] memory participants, , , ) = afnConstructorArgs();
    bytes32[] memory roots = new bytes32[](1);
    roots[0] = ROOT_1;

    changePrank(participants[0]);
    s_afn.voteToBlessRoots(roots);
    assertTrue(s_afn.hasVotedToBlessRoot(participants[0], ROOT_1));

    uint256 votesToBlessBefore = s_afn.getVotesToBlessRoot(ROOT_1);
    s_afn.voteToBlessRoots(roots);
    assertEq(votesToBlessBefore, s_afn.getVotesToBlessRoot(ROOT_1));
  }

  // Reverts

  function testBadSignalReverts() public {
    (address[] memory participants, , , ) = afnConstructorArgs();
    bytes32[] memory roots = new bytes32[](1);
    roots[0] = ROOT_1;

    for (uint256 i = 0; i < participants.length; i++) {
      changePrank(participants[i]);
      s_afn.voteBad();
    }

    vm.expectRevert(IAFN.MustRecoverFromBadSignal.selector);
    s_afn.voteToBlessRoots(roots);
  }

  function testInvalidVoterReverts() public {
    bytes32[] memory roots = new bytes32[](1);
    roots[0] = ROOT_1;

    changePrank(STRANGER);
    vm.expectRevert(abi.encodeWithSelector(IAFN.InvalidVoter.selector, STRANGER));
    s_afn.voteToBlessRoots(roots);
  }
}

contract AFN_voteBad is AFNSetup {
  event VoteBad(address indexed voter, uint256 weight);
  event AFNBadSignal(uint256 timestamp);

  // Success

  function testVoteBadSuccess_gas() public {
    vm.pauseGasMetering();
    (address[] memory participants, uint256[] memory weights, , ) = afnConstructorArgs();
    address voter = participants[0];
    uint256 weight = weights[0];
    changePrank(voter);
    vm.expectEmit();
    emit VoteBad(voter, weight);

    vm.resumeGasMetering();
    s_afn.voteBad();
    vm.pauseGasMetering();

    assertTrue(s_afn.hasVotedBad(voter));
    (address[] memory voters, uint256 votes) = s_afn.getBadVotersAndVotes();
    assertEq(1, voters.length);
    assertEq(voter, voters[0]);
    assertEq(weight, votes);
    vm.resumeGasMetering();
  }

  function testEmitBadSignalSuccess() public {
    (address[] memory participants, , , ) = afnConstructorArgs();
    for (uint256 i = 0; i < participants.length - 1; ++i) {
      changePrank(participants[i]);
      s_afn.voteBad();
    }

    vm.expectEmit();
    emit AFNBadSignal(block.timestamp);

    changePrank(participants[participants.length - 1]);
    s_afn.voteBad();
  }

  // Reverts

  function testIsBadReverts() public {
    (address[] memory participants, , , ) = afnConstructorArgs();
    for (uint256 i = 0; i < participants.length; ++i) {
      changePrank(participants[i]);
      s_afn.voteBad();
    }

    vm.expectRevert(IAFN.MustRecoverFromBadSignal.selector);
    s_afn.voteBad();
  }

  function testInvalidVoterReverts() public {
    changePrank(STRANGER);

    vm.expectRevert(abi.encodeWithSelector(IAFN.InvalidVoter.selector, STRANGER));
    s_afn.voteBad();
  }

  function testAlreadyVotedReverts() public {
    (address[] memory participants, , , ) = afnConstructorArgs();
    changePrank(participants[0]);
    s_afn.voteBad();

    vm.expectRevert(IAFN.AlreadyVoted.selector);
    s_afn.voteBad();
  }
}

contract AFN_recover is AFNSetup {
  event RecoveredFromBadSignal();

  function setUp() public virtual override {
    AFNSetup.setUp();
    (address[] memory participants, , , ) = afnConstructorArgs();
    for (uint256 i = 0; i < participants.length; ++i) {
      changePrank(participants[i]);
      s_afn.voteBad();
    }
  }

  // Success

  function testRecoverSuccess_gas() public {
    vm.pauseGasMetering();
    changePrank(OWNER);
    vm.expectEmit();
    emit RecoveredFromBadSignal();

    vm.resumeGasMetering();
    s_afn.recoverFromBadSignal();
    vm.pauseGasMetering();

    assertFalse(s_afn.badSignalReceived());
    (address[] memory voters, uint256 votes) = s_afn.getBadVotersAndVotes();
    assertEq(0, voters.length);
    assertEq(0, votes);
    vm.resumeGasMetering();
  }

  // Reverts

  function testNonOwnerReverts() public {
    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_afn.recoverFromBadSignal();
  }

  function testNotBadReverts() public {
    changePrank(OWNER);
    s_afn.recoverFromBadSignal();

    vm.expectRevert(IAFN.RecoveryNotNecessary.selector);
    s_afn.recoverFromBadSignal();
  }
}

contract AFN_setAFNConfig is AFNSetup {
  event AFNConfigSet(address[] parties, uint256[] weights, uint256 goodQuorum, uint256 badQuorum);

  /// @notice Test-specific function to use only in setAFNConfig tests
  function getDifferentConfigArgs()
    private
    pure
    returns (
      address[] memory,
      uint256[] memory,
      uint256,
      uint256
    )
  {
    address[] memory participants = new address[](2);
    participants[0] = USER_1;
    participants[1] = USER_2;
    uint256[] memory weights = new uint256[](2);
    weights[0] = WEIGHT_1;
    weights[1] = WEIGHT_10;
    uint256 blessingThreshold = WEIGHT_1 + WEIGHT_10;
    uint256 badSignalThreshold = WEIGHT_1 + WEIGHT_10;
    return (participants, weights, blessingThreshold, badSignalThreshold);
  }

  function setUp() public virtual override {
    AFNSetup.setUp();
    (address[] memory participants, , , ) = afnConstructorArgs();

    // Setup some partial state
    changePrank(participants[0]);
    bytes32[] memory roots = new bytes32[](1);
    roots[0] = ROOT_1;
    s_afn.voteToBlessRoots(roots);
    changePrank(participants[1]);
    s_afn.voteToBlessRoots(roots);
    s_afn.voteBad();
  }

  // Success

  function testSetAFNConfigSuccess_gas() public {
    vm.pauseGasMetering();
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 badSignalThreshold
    ) = getDifferentConfigArgs();

    changePrank(OWNER);
    vm.expectEmit();
    emit AFNConfigSet(participants, weights, blessingThreshold, badSignalThreshold);

    uint256 configVersionBefore = s_afn.getConfigVersion();

    vm.resumeGasMetering();
    s_afn.setAFNConfig(participants, weights, blessingThreshold, badSignalThreshold);
    vm.pauseGasMetering();

    // Assert Config has changed correctly
    assertEq(configVersionBefore + 1, s_afn.getConfigVersion());
    assertEq(participants, s_afn.getParticipants());
    assertEq(weights[0], s_afn.getWeightByParticipant(participants[0]));
    (uint256 blessing, uint256 badSignal) = s_afn.getWeightThresholds();
    assertEq(blessing, blessingThreshold);
    assertEq(badSignal, badSignalThreshold);
    // Assert that bad votes have been cleared
    (address[] memory badVoters, uint256 badVotes) = s_afn.getBadVotersAndVotes();
    assertEq(ZERO, badVoters.length);
    assertEq(ZERO, badVotes);
    // Assert that good votes have been cleared
    uint256 votesToBlessRoot = s_afn.getVotesToBlessRoot(ROOT_1);
    assertEq(ZERO, votesToBlessRoot);
    assertFalse(s_afn.hasVotedToBlessRoot(participants[0], ROOT_1));
    assertFalse(s_afn.hasVotedToBlessRoot(participants[1], ROOT_1));
    vm.resumeGasMetering();
  }

  // Reverts

  function testNonOwnerReverts() public {
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 badSignalThreshold
    ) = getDifferentConfigArgs();

    changePrank(STRANGER);
    vm.expectRevert("Only callable by owner");
    s_afn.setAFNConfig(participants, weights, blessingThreshold, badSignalThreshold);
  }

  function testParticipantsLengthIsZeroReverts() public {
    (, uint256[] memory weights, uint256 blessingThreshold, uint256 badSignalThreshold) = getDifferentConfigArgs();
    address[] memory participants = new address[](0);

    changePrank(OWNER);
    vm.expectRevert(IAFN.InvalidConfig.selector);
    s_afn.setAFNConfig(participants, weights, blessingThreshold, badSignalThreshold);
  }

  function testEitherThresholdIsZeroReverts() public {
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 badSignalThreshold
    ) = getDifferentConfigArgs();

    changePrank(OWNER);
    vm.expectRevert(IAFN.InvalidConfig.selector);
    s_afn.setAFNConfig(participants, weights, ZERO, badSignalThreshold);
    vm.expectRevert(IAFN.InvalidConfig.selector);
    s_afn.setAFNConfig(participants, weights, blessingThreshold, ZERO);
  }

  function testParticipantIsZeroAddressReverts() public {
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 badSignalThreshold
    ) = getDifferentConfigArgs();

    changePrank(OWNER);
    participants[0] = ZERO_ADDRESS;
    vm.expectRevert(IAFN.InvalidConfig.selector);
    s_afn.setAFNConfig(participants, weights, blessingThreshold, badSignalThreshold);
  }

  function testWeightIsZeroAddressReverts() public {
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 badSignalThreshold
    ) = getDifferentConfigArgs();

    changePrank(OWNER);
    weights[0] = ZERO;
    vm.expectRevert(IAFN.InvalidWeight.selector);
    s_afn.setAFNConfig(participants, weights, blessingThreshold, badSignalThreshold);
  }

  function testTotalWeightsSmallerThanEachThresholdReverts() public {
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 badSignalThreshold
    ) = getDifferentConfigArgs();

    changePrank(OWNER);
    vm.expectRevert(IAFN.InvalidConfig.selector);
    s_afn.setAFNConfig(participants, weights, WEIGHT_40, badSignalThreshold);
    vm.expectRevert(IAFN.InvalidConfig.selector);
    s_afn.setAFNConfig(participants, weights, blessingThreshold, WEIGHT_40);
  }
}
