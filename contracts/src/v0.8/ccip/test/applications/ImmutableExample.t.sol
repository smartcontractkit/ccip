pragma solidity ^0.8.0;

import "../onRamp/EVM2EVMOnRampSetup.t.sol";
import "../../applications/ImmutableExample.sol";
import {ERC165Checker} from "../../../vendor/ERC165Checker.sol";

contract ImmutableExample_sanity is EVM2EVMOnRampSetup {
  function testSetChain() public {
    ImmutableExample e = new ImmutableExample(s_sourceRouter, IERC20(s_sourceFeeToken));
    Client.EVMExtraArgsV1 memory extraArgs = Client.EVMExtraArgsV1({gasLimit: 300_000, strict: false});
    bytes memory encodedExtraArgs = Client._argsToBytes(extraArgs);
    e.enableChain(DEST_CHAIN_ID, encodedExtraArgs);
    assertEq(e.s_chains(DEST_CHAIN_ID), encodedExtraArgs);
  }

  function testCanReceive() public {
    ImmutableExample e = new ImmutableExample(s_sourceRouter, IERC20(s_sourceFeeToken));
    assertTrue(ERC165Checker.supportsInterface(address(e), type(IAny2EVMMessageReceiver).interfaceId));
  }
}
