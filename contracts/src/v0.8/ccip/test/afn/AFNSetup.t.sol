// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import "../../AFN.sol";

contract AFNSetup is BaseTest {
  bytes32 internal constant ROOT_1 = bytes32("1");
  bytes32 internal constant ROOT_2 = bytes32("2");
  bytes32 internal constant ROOT_3 = bytes32("3");
  bytes32 internal constant ROOT_4 = bytes32("4");
  bytes32 internal constant ROOT_5 = bytes32("5");

  function setUp() public virtual override {
    BaseTest.setUp();
    (
      address[] memory participants,
      uint256[] memory weights,
      uint256 blessingThreshold,
      uint256 basSignalThreshold
    ) = afnConstructorArgs();
    s_afn = new AFN(participants, weights, blessingThreshold, basSignalThreshold);
  }
}
