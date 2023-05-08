// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import "../BaseTest.t.sol";
import {AFN} from "../../AFN.sol";
import {IAFN} from "../../interfaces/IAFN.sol";

contract AFNSetup is BaseTest {
  function makeTaggedRootsInclusive(uint256 from, uint256 to) internal pure returns (IAFN.TaggedRoot[] memory) {
    IAFN.TaggedRoot[] memory votes = new IAFN.TaggedRoot[](to - from + 1);
    for (uint256 i = from; i <= to; ++i) {
      votes[i - from] = IAFN.TaggedRoot({commitStore: address(1), root: bytes32(uint256(i))});
    }
    return votes;
  }

  function makeTaggedRootSingleton(uint256 index) internal pure returns (IAFN.TaggedRoot[] memory) {
    return makeTaggedRootsInclusive(index, index);
  }

  function makeTaggedRoot(uint256 index) internal pure returns (IAFN.TaggedRoot memory) {
    return makeTaggedRootSingleton(index)[0];
  }

  function makeTaggedRootHash(uint256 index) internal pure returns (bytes32) {
    IAFN.TaggedRoot memory taggedRoot = makeTaggedRootSingleton(index)[0];
    return keccak256(abi.encode(taggedRoot.commitStore, taggedRoot.root));
  }

  function makeCurseId(uint256 index) internal pure returns (bytes32) {
    return bytes32(index);
  }

  AFN internal s_afn;

  function setUp() public virtual override {
    BaseTest.setUp();
    s_afn = new AFN(afnConstructorArgs());
  }

  function hasVotedToBlessRoot(address voter, IAFN.TaggedRoot memory taggedRoot_) internal view returns (bool) {
    (address[] memory voters, ) = s_afn.getBlessVotersAndWeight(taggedRoot_);
    for (uint256 i = 0; i < voters.length; ++i) {
      if (voters[i] == voter) {
        return true;
      }
    }
    return false;
  }

  function getWeightOfVotesToBlessRoot(IAFN.TaggedRoot memory taggedRoot_) internal view returns (uint16) {
    (, uint16 weight) = s_afn.getBlessVotersAndWeight(taggedRoot_);
    return weight;
  }
}
