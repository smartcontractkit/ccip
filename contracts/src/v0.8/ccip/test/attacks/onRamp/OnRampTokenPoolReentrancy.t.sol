// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {Internal} from "../../../libraries/Internal.sol";
import {Client} from "../../../libraries/Client.sol";
import {FacadeClient} from "./FacadeClient.sol";
import {ReentrantMaliciousTokenPool} from "./ReentrantMaliciousTokenPool.sol";
import {EVM2EVMOnRampSetup} from "../../onRamp/EVM2EVMOnRampSetup.t.sol";

import {IERC20} from "../../../../vendor/openzeppelin-solidity/v4.8.0/contracts/token/ERC20/IERC20.sol";

/// @title OnRampTokenPoolReentrancy
/// Attempts to perform a reentrancy exploit on Onramp with a malicious TokenPool
contract OnRampTokenPoolReentrancy is EVM2EVMOnRampSetup {
  FacadeClient s_facadeClient;
  ReentrantMaliciousTokenPool s_maliciousTokenPool;
  IERC20 s_sourceToken;
  IERC20 s_feeToken;

  function setUp() public virtual override {
    EVM2EVMOnRampSetup.setUp();

    s_sourceToken = IERC20(s_sourceTokens[0]);
    s_feeToken = IERC20(s_sourceTokens[0]);

    s_facadeClient = new FacadeClient(address(s_sourceRouter), DEST_CHAIN_ID, s_sourceToken, s_feeToken);

    s_maliciousTokenPool = new ReentrantMaliciousTokenPool(address(s_facadeClient), s_sourceToken, address(s_mockARM));

    Internal.PoolUpdate[] memory removes = new Internal.PoolUpdate[](1);
    removes[0].token = address(s_sourceToken);
    removes[0].pool = address(s_sourcePools[0]);
    Internal.PoolUpdate[] memory adds = new Internal.PoolUpdate[](1);
    adds[0].token = address(s_sourceToken);
    adds[0].pool = address(s_maliciousTokenPool);

    s_onRamp.applyPoolUpdates(removes, adds);

    s_sourceToken.transfer(address(s_facadeClient), 1e18);
    s_feeToken.transfer(address(s_facadeClient), 1e18);
  }

  /// @dev This test was used to showcase a reentrancy exploit on OnRamp with malicious TokenPool.
  /// How it worked: OnRamp used to construct EVM2EVM messages after calling TokenPool's lockOrBurn.
  /// This allowed the malicious TokenPool to break message sequencing expectations as follows:
  ///   Any user -> Facade -> 1st call to ccipSend -> pool’s lockOrBurn —>
  ///   (reenter)-> Facade -> 2nd call to ccipSend
  /// In this case, Facade's second call would produce an EVM2EVM msg with a lower sequence number.
  /// The issue was fixed by moving state updates and event construction to before TokenPool calls.
  /// This test is kept to verify message sequence expectations are not broken.
  function testSuccess() public {
    uint256 amount = 1;

    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](1);
    tokenAmounts[0].token = address(s_sourceToken);
    tokenAmounts[0].amount = amount;

    Client.EVM2AnyMessage memory message1 = Client.EVM2AnyMessage({
      receiver: abi.encode(address(100)),
      data: abi.encodePacked(uint256(1)), // message 1 contains data 1
      tokenAmounts: tokenAmounts,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000})),
      feeToken: address(s_feeToken)
    });

    Client.EVM2AnyMessage memory message2 = Client.EVM2AnyMessage({
      receiver: abi.encode(address(100)),
      data: abi.encodePacked(uint256(2)), // message 2 contains data 2
      tokenAmounts: tokenAmounts,
      extraArgs: Client._argsToBytes(Client.EVMExtraArgsV1({gasLimit: 200_000})),
      feeToken: address(s_feeToken)
    });

    uint256 expectedFee = s_sourceRouter.getFee(DEST_CHAIN_ID, message1);
    assertGt(expectedFee, 0);

    // Outcome of a successful exploit:
    // Message 1 event from OnRamp contains sequence/nonce 2, message 2 contains sequence/nonce 1
    // Internal.EVM2EVMMessage memory msgEvent1 = _messageToEvent(message1, 2, 2, expectedFee, address(s_facadeClient));
    // Internal.EVM2EVMMessage memory msgEvent2 = _messageToEvent(message2, 1, 1, expectedFee, address(s_facadeClient));

    // vm.expectEmit();
    // emit CCIPSendRequested(msgEvent2);
    // vm.expectEmit();
    // emit CCIPSendRequested(msgEvent1);

    // After issue is fixed, sequence now increments as expected
    Internal.EVM2EVMMessage memory msgEvent1 = _messageToEvent(message1, 1, 1, expectedFee, address(s_facadeClient));
    Internal.EVM2EVMMessage memory msgEvent2 = _messageToEvent(message2, 2, 2, expectedFee, address(s_facadeClient));

    vm.expectEmit();
    emit CCIPSendRequested(msgEvent2);
    vm.expectEmit();
    emit CCIPSendRequested(msgEvent1);

    s_facadeClient.send(amount);
  }
}
