// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

import {IARM} from "../../interfaces/IARM.sol";
import {ARM} from "../../ARM.sol";

contract MockARM is IARM {
  bool private s_curse;

  function isCursed() external view override returns (bool) {
    return s_curse;
  }

  function voteToCurse(bytes32) external {
    s_curse = true;
  }

  function ownerUnvoteToCurse(ARM.UnvoteToCurseRecord[] memory) external {
    s_curse = false;
  }

  function isBlessed(IARM.TaggedRoot calldata) external view override returns (bool) {
    return !s_curse;
  }
}
