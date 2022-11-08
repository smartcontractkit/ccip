// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../BaseTest.t.sol";
import "../../models/Models.sol";

contract ExtraArgs__toBytes is BaseTest {
  using CCIP for CCIP.EVMExtraArgsV1;

  function testSuccess() public {
    CCIP.EVMExtraArgsV1 memory extraArgs = CCIP.EVMExtraArgsV1({
      gasLimit : 100
    });

    bytes memory expected = bytes.concat(CCIP.EVM_EXTRA_ARGS_V1_TAG, abi.encode(extraArgs.gasLimit));
    assertEq(expected, extraArgs._toBytes());
  }

  function testEVMExtraArgsV1TagSuccess() public {
    bytes4 tag = bytes4(keccak256("CCIP EVMExtraArgsV1"));
    assertEq(tag, CCIP.EVM_EXTRA_ARGS_V1_TAG);
  }
}

contract ExtraArgs_fromBytes is BaseTest {
  using CCIP for CCIP.EVMExtraArgsV1;

  function testSuccess() public {
    CCIP.EVMExtraArgsV1 memory extraArgs = CCIP.EVMExtraArgsV1({
      gasLimit : 1234567
    });

    CCIP.EVMExtraArgsV1 memory extraArgs2 = this.fromBytesHelper(extraArgs._toBytes());

    assertEq(extraArgs.gasLimit, extraArgs2.gasLimit);
  }


  function testDefaultsSuccess() public {
    bytes memory empty;
    CCIP.EVMExtraArgsV1 memory extraArgs = this.fromBytesHelper(empty);

    assertEq(extraArgs.gasLimit, CCIP.EVM_DEFAULT_GAS_LIMIT);
  }

  // Reverts
  function testInvalidExtraArgsTagReverts() public {
    bytes4 wrongTag = 0x9879e47a;

    bytes memory bts = bytes.concat(wrongTag,  abi.encode(100));

    vm.expectRevert(abi.encodeWithSelector(CCIP.InvalidExtraArgsTag.selector, CCIP.EVM_EXTRA_ARGS_V1_TAG, wrongTag));
    this.fromBytesHelper(bts);
  }
}