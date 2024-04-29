// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import {Pool} from "../../libraries/Pool.sol";
import {BaseTest} from "../BaseTest.t.sol";
import "forge-std/console.sol";

contract Pool__generatePoolReturnDataV1 is BaseTest {
  function test__generatePoolReturnDataV1_Success() public {
    bytes memory remotePoolAddress = abi.encode(makeAddr("remotePoolAddress"));
    bytes memory destPoolData = abi.encode(makeAddr("destPoolData"));

    bytes memory generatedReturnData = Pool._encodeLockOrBurnOutV1(remotePoolAddress, destPoolData);

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = Pool._decodeLockOrBurnOutV1(generatedReturnData);

    assertEq(poolReturnDataV1.destPoolAddress, remotePoolAddress);
    assertEq(poolReturnDataV1.destPoolData, destPoolData);
  }

  function test_Fuzz__generatePoolReturnDataV1_Success(
    bytes memory destPoolData,
    bytes memory remotePoolAddress
  ) public pure {
    bytes memory generatedReturnData = Pool._encodeLockOrBurnOutV1(remotePoolAddress, destPoolData);

    Pool.LockOrBurnOutV1 memory poolReturnDataV1 = Pool._decodeLockOrBurnOutV1(generatedReturnData);

    assertEq(poolReturnDataV1.destPoolAddress, remotePoolAddress);
    assertEq(poolReturnDataV1.destPoolData, destPoolData);
  }
}
