// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.24;

import {RMN} from "../../RMN.sol";

import {RMNBase} from "../../RMNBase.sol";
import {IRMN} from "../../interfaces/IRMN.sol";

import {Test} from "forge-std/Test.sol";

function makeSubjects(bytes16 a) pure returns (bytes16[] memory) {
  bytes16[] memory subjects = new bytes16[](1);
  subjects[0] = a;
  return subjects;
}

function makeSubjects(bytes16 a, bytes16 b) pure returns (bytes16[] memory) {
  bytes16[] memory subjects = new bytes16[](2);
  subjects[0] = a;
  subjects[1] = b;
  return subjects;
}

// in order from earliest to latest curse ids
function makeCursesHashFromList(bytes32[] memory curseIds) pure returns (bytes28 cursesHash) {
  for (uint256 i = 0; i < curseIds.length; ++i) {
    cursesHash = bytes28(keccak256(abi.encode(cursesHash, curseIds[i])));
  }
}

// hides the ugliness from tests
function makeCursesHash(bytes32 a) pure returns (bytes28) {
  bytes32[] memory curseIds = new bytes32[](1);
  curseIds[0] = a;
  return makeCursesHashFromList(curseIds);
}

function makeCursesHash(bytes32 a, bytes32 b) pure returns (bytes28) {
  bytes32[] memory curseIds = new bytes32[](2);
  curseIds[0] = a;
  curseIds[1] = b;
  return makeCursesHashFromList(curseIds);
}

contract RMNBaseSetup is Test {
  // Addresses
  address internal constant OWNER = 0x00007e64E1fB0C487F25dd6D3601ff6aF8d32e4e;
  address internal constant STRANGER = address(999999);
  address internal constant ZERO_ADDRESS = address(0);
  address internal s_blessVoter1;
  address internal s_blessVoter2;
  address internal s_blessVoter3;
  address internal s_blessVoter4;
  uint256 internal s_blessVoter1Key;
  uint256 internal s_blessVoter2Key;
  uint256 internal s_blessVoter3Key;
  uint256 internal s_blessVoter4Key;
  address internal constant CURSE_VOTER_1 = address(10);
  address internal constant CURSE_VOTER_2 = address(12);
  address internal constant CURSE_VOTER_3 = address(13);
  address internal constant CURSE_VOTER_4 = address(14);

  function setUp() public virtual {
    (s_blessVoter1, s_blessVoter1Key) = makeAddrAndKey("voter1");
    (s_blessVoter2, s_blessVoter2Key) = makeAddrAndKey("voter2");
    (s_blessVoter3, s_blessVoter3Key) = makeAddrAndKey("voter3");
    (s_blessVoter4, s_blessVoter4Key) = makeAddrAndKey("voter4");
  }

  // Arm
  function _rmnConstructorArgs() internal view returns (RMNBase.Config memory) {
    RMNBase.Voter[] memory voters = new RMNBase.Voter[](4);
    voters[0] = RMNBase.Voter({
      blessVoteAddr: s_blessVoter1,
      curseVoteAddr: CURSE_VOTER_1,
      blessWeight: WEIGHT_1,
      curseWeight: WEIGHT_1
    });
    voters[1] = RMNBase.Voter({
      blessVoteAddr: s_blessVoter2,
      curseVoteAddr: CURSE_VOTER_2,
      blessWeight: WEIGHT_10,
      curseWeight: WEIGHT_10
    });
    voters[2] = RMNBase.Voter({
      blessVoteAddr: s_blessVoter3,
      curseVoteAddr: CURSE_VOTER_3,
      blessWeight: WEIGHT_20,
      curseWeight: WEIGHT_20
    });
    voters[3] = RMNBase.Voter({
      blessVoteAddr: s_blessVoter4,
      curseVoteAddr: CURSE_VOTER_4,
      blessWeight: WEIGHT_40,
      curseWeight: WEIGHT_40
    });
    return RMNBase.Config({
      voters: voters,
      blessWeightThreshold: WEIGHT_10 + WEIGHT_20 + WEIGHT_40,
      curseWeightThreshold: WEIGHT_1 + WEIGHT_10 + WEIGHT_20 + WEIGHT_40
    });
  }

  uint8 internal constant ZERO = 0;
  uint8 internal constant WEIGHT_1 = 1;
  uint8 internal constant WEIGHT_10 = 10;
  uint8 internal constant WEIGHT_20 = 20;
  uint8 internal constant WEIGHT_40 = 40;

  function makeCurseId(uint256 index) internal pure returns (bytes16) {
    return bytes16(uint128(index));
  }
}

contract RMNSetup is RMNBaseSetup {
  RMN internal s_rmn;

  function setUp() public virtual override {
    super.setUp();
    vm.startPrank(OWNER);
    s_rmn = new RMN(_rmnConstructorArgs());
    vm.stopPrank();
  }

  function makeTaggedRootsInclusive(uint256 from, uint256 to) internal pure returns (IRMN.TaggedRoot[] memory) {
    IRMN.TaggedRoot[] memory votes = new IRMN.TaggedRoot[](to - from + 1);
    for (uint256 i = from; i <= to; ++i) {
      votes[i - from] = IRMN.TaggedRoot({commitStore: address(1), root: bytes32(uint256(i))});
    }
    return votes;
  }

  function makeTaggedRootSingleton(uint256 index) internal pure returns (IRMN.TaggedRoot[] memory) {
    return makeTaggedRootsInclusive(index, index);
  }

  function makeTaggedRoot(uint256 index) internal pure returns (IRMN.TaggedRoot memory) {
    return makeTaggedRootSingleton(index)[0];
  }

  function makeTaggedRootHash(uint256 index) internal pure returns (bytes32) {
    IRMN.TaggedRoot memory taggedRoot = makeTaggedRootSingleton(index)[0];
    return keccak256(abi.encode(taggedRoot.commitStore, taggedRoot.root));
  }

  function hasVotedToBlessRoot(address voter, IRMN.TaggedRoot memory taggedRoot_) internal view returns (bool) {
    (address[] memory voters,,) = s_rmn.getBlessProgress(taggedRoot_);
    for (uint256 i = 0; i < voters.length; ++i) {
      if (voters[i] == voter) {
        return true;
      }
    }
    return false;
  }

  function getWeightOfVotesToBlessRoot(IRMN.TaggedRoot memory taggedRoot_) internal view returns (uint16) {
    (, uint16 weight,) = s_rmn.getBlessProgress(taggedRoot_);
    return weight;
  }
}
