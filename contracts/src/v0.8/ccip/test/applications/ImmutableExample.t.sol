pragma solidity ^0.8.0;

import "../onRamp/EVM2EVMOnRampSetup.t.sol";
import "../../applications/ImmutableExample.sol";
import {ERC165Checker} from "../../../vendor/openzeppelin-solidity/v4.8.0/utils/introspection/ERC165Checker.sol";

contract ImmutableExample_sanity is EVM2EVMOnRampSetup {
  function testExamples() public {
    ImmutableExample e = new ImmutableExample(s_sourceRouter, IERC20(s_sourceFeeToken));
    deal(address(e), 100 ether);
    deal(s_sourceFeeToken, address(e), 100 ether);

    // feeToken approval works
    assertEq(IERC20(s_sourceFeeToken).allowance(address(e), address(s_sourceRouter)), 2 ** 256 - 1);

    // Can set chain
    Client.EVMExtraArgsV1 memory extraArgs = Client.EVMExtraArgsV1({gasLimit: 300_000, strict: false});
    bytes memory encodedExtraArgs = Client._argsToBytes(extraArgs);
    e.enableChain(DEST_CHAIN_ID, encodedExtraArgs);
    assertEq(e.s_chains(DEST_CHAIN_ID), encodedExtraArgs);

    // Can send data pay native
    e.sendDataPayNative(DEST_CHAIN_ID, abi.encode(address(1)), bytes("hello"));

    // Can send data pay feeToken
    e.sendDataPayFeeToken(DEST_CHAIN_ID, abi.encode(address(1)), bytes("hello"));

    // Can send data tokens
    assertEq(address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[1]))), address(s_sourcePools[1]));
    deal(s_sourceTokens[1], OWNER, 100 ether);
    IERC20(s_sourceTokens[1]).approve(address(e), 1 ether);
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0] = Client.EVMTokenAmount({token: s_sourceTokens[1], amount: 1 ether});
    e.sendDataAndTokens(DEST_CHAIN_ID, abi.encode(address(1)), bytes("hello"), tokenAmounts);
    // Tokens transferred from owner to router then burned in pool.
    assertEq(IERC20(s_sourceTokens[1]).balanceOf(OWNER), 99 ether);
    assertEq(IERC20(s_sourceTokens[1]).balanceOf(address(s_sourceRouter)), 0);

    // Can send just tokens
    IERC20(s_sourceTokens[1]).approve(address(e), 1 ether);
    e.sendTokens(DEST_CHAIN_ID, abi.encode(address(1)), tokenAmounts);

    // Can receive
    assertTrue(ERC165Checker.supportsInterface(address(e), type(IAny2EVMMessageReceiver).interfaceId));

    // Can disable chain
    e.disableChain(DEST_CHAIN_ID);
  }
}
