// SPDX-License-Identifier: MIT
pragma solidity ^0.8.13;

import "./TollOffRampSetup.t.sol";
import "../../../offRamp/toll/Any2EVMTollOffRampRouter.sol";
import "../../mocks/MockTollOffRampRouter.sol";
import "../../helpers/RevertingReceiver.sol";
import "../../helpers/Any2EVMTollOffRampHelper.sol";

contract Any2EVMTollOffRampTest is TollOffRampSetup {
  Any2EVMTollOffRampHelper s_offRamp;
  TollOffRampRouterInterface s_router;

  function setUp() public virtual override {
    TollOffRampSetup.setUp();
    s_offRamp = new Any2EVMTollOffRampHelper(
      s_sourceChainId,
      s_offRampConfig,
      s_mockBlobVerifier,
      s_onRampAddress,
      s_afn,
      s_sourceTokens,
      s_destPools,
      2**20
    );
    TollOffRampInterface[] memory offRamps = new TollOffRampInterface[](1);
    offRamps[0] = s_offRamp;
    s_router = new Any2EVMTollOffRampRouter(offRamps);
    s_offRamp.setRouter(s_router);
  }

  // Assert that setRouter will set the router to the given router argument.
  function testSetRouter() public {
    TollOffRampRouterInterface newRouter = new MockTollOffRampRouter();
    assertTrue(address(newRouter) != address(s_offRamp.s_router()));
    s_offRamp.setRouter(newRouter);
    assertEq(address(newRouter), address(s_offRamp.s_router()));
  }

  // Asserts that any call to executeSingleMessage will revert when not
  // it's not a self call.
  function testExecuteSingleMessageNoSelfCall() public {
    vm.expectRevert(TollOffRampInterface.CanOnlySelfCall.selector);
    s_offRamp.executeSingleMessage(getAny2EVMTollMessageNoTokens(1));
  }

  // Assert that any call to executeSingleMessage with an invalid receiver
  // will revert.
  function testExecuteSingleMessageInvalidReceiver() public {
    vm.stopPrank();
    vm.prank(address(s_offRamp));
    CCIP.Any2EVMTollMessage memory message = getAny2EVMTollMessageNoTokens(1);
    address wrongAddress = 0xDEaD7E64e1Fb0c487f25Dd6D3601fF6AF8D32e4E;
    message.receiver = wrongAddress;
    vm.expectRevert(abi.encodeWithSelector(TollOffRampInterface.InvalidReceiver.selector, wrongAddress));
    s_offRamp.executeSingleMessage(message);
  }

  // Assert that a self call to executeSingleMessage with a valid receiver
  // will succeed.
  function testExecuteSingleMessageSuccessNoTokens() public {
    vm.stopPrank();
    vm.prank(address(s_offRamp));
    s_offRamp.executeSingleMessage(getAny2EVMTollMessageNoTokens(1));
  }

  // Asserts that a call to execute will revert when the router is unset.
  function testExecuteNoRouterSet() public {
    TollOffRampRouterInterface newRouter = MockTollOffRampRouter(address(0));
    s_offRamp.setRouter(newRouter);
    vm.expectRevert(TollOffRampInterface.RouterNotSet.selector);
    s_offRamp.execute(createReportFromMessages(getBasicMessages()), false);
  }

  // Asserts that a properly formed call to execute will succeed.
  function testExecuteNoTokensSingleMessageSuccess() public {
    CCIP.ExecutionReport memory executionReport = createReportFromMessages(getBasicMessages());
    CCIP.ExecutionResult[] memory report = s_offRamp.execute(executionReport, false);
    assertEq(report.length, executionReport.encodedMessages.length);
    assertEq(report[0].sequenceNumber, executionReport.sequenceNumbers[0]);
    assertEq(report[0].timestampRelayed, 1);
    assertTrue(report[0].state == CCIP.MessageExecutionState.Success);
  }

  // Asserts that a call to execute will revert if a message in the execution report
  // is already executed.
  function testExecuteAlreadyExecuted() public {
    CCIP.ExecutionReport memory executionReport = createReportFromMessages(getBasicMessages());
    CCIP.ExecutionResult[] memory report = s_offRamp.execute(executionReport, false);
    vm.expectRevert(
      abi.encodeWithSelector(TollOffRampInterface.AlreadyExecuted.selector, executionReport.sequenceNumbers[0])
    );
    report = s_offRamp.execute(executionReport, false);
  }

  // Asserts that a call to execute will revert if the tokens and amounts
  // properties are not of the same length.
  function testExecuteUnsupportedNumberOfTokens() public {
    CCIP.Any2EVMTollMessage[] memory messages = getBasicMessages();
    IERC20[] memory newTokens = new IERC20[](1);
    newTokens[0] = s_sourceTokens[0];
    messages[0].tokens = newTokens;
    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);

    vm.expectRevert(
      abi.encodeWithSelector(
        TollOffRampInterface.UnsupportedNumberOfTokens.selector,
        executionReport.sequenceNumbers[0]
      )
    );
    s_offRamp.execute(executionReport, false);
  }

  // Asserts that a call to execute will revert when a message has the wrong
  // source chain id.
  function testExecuteInvalidSourceChain() public {
    CCIP.Any2EVMTollMessage[] memory messages = getBasicMessages();
    messages[0].sourceChainId = s_sourceChainId + 1;

    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);
    vm.expectRevert(abi.encodeWithSelector(TollOffRampInterface.InvalidSourceChain.selector, s_sourceChainId + 1));
    s_offRamp.execute(executionReport, false);
  }

  // Asserts that a call to execute will revert when a message has data that
  // exceeds the maximum data length.
  function testExecuteMessageDataTooLarge() public {
    CCIP.Any2EVMTollMessage[] memory messages = getBasicMessages();
    messages[0]
      .data = "3.1415926535897932384626433832795028841971693993751058209749445923078164062862089986280348253421170679821480865132823066470938446095505822317253594081284811174502841027019385211055596446229489549303819644288109756659334461284756482337867831652712019091456485669234603486104543266482133936072602491412737245870066063155881748815209209628292540917153643678925903600113305305488204665213841469519415116094330572703657595919530921861173819326117931051185480744623799627495673518857527248912279381830119491";

    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);
    vm.expectRevert(
      abi.encodeWithSelector(
        TollOffRampInterface.MessageTooLarge.selector,
        s_offRampConfig.maxDataSize,
        messages[0].data.length
      )
    );
    s_offRamp.execute(executionReport, false);
  }

  // Asserts that a call to execute succeeds even though the call
  // to execute the tx fails. The resulting tx state is set to Failed.
  function testExecuteNoTokensSingleMessageFailedCallSuccess() public {
    CCIP.Any2EVMTollMessage[] memory messages = getBasicMessages();
    RevertingReceiver newReceiver = new RevertingReceiver();
    messages[0].receiver = address(newReceiver);

    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);
    CCIP.ExecutionResult[] memory report = s_offRamp.execute(executionReport, false);
    assertEq(report.length, executionReport.encodedMessages.length);
    assertEq(report[0].sequenceNumber, executionReport.sequenceNumbers[0]);
    assertTrue(report[0].state == CCIP.MessageExecutionState.Failure);
  }

  function testPoolsProperlySet() public {
    IERC20[] memory pools = s_offRamp.getPoolTokens();
    assertEq(pools.length, s_sourceTokens.length);
    assertTrue(address(pools[0]) == address(s_sourceTokens[0]));
    assertTrue(address(pools[1]) == address(s_sourceTokens[1]));
  }

  function testExecuteWithTokensSuccess() public {
    CCIP.Any2EVMTollMessage[] memory messages = getMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);

    CCIP.ExecutionResult[] memory report = s_offRamp.execute(executionReport, false);

    assertEq(report.length, messages.length);
    assertEq(report[0].sequenceNumber, messages[0].sequenceNumber);
    assertTrue(report[0].state == CCIP.MessageExecutionState.Success);
    assertEq(report[1].sequenceNumber, messages[1].sequenceNumber);
    assertTrue(report[1].state == CCIP.MessageExecutionState.Success);
  }

  function testExecuteWithTokensSuccessWithToll() public {
    CCIP.Any2EVMTollMessage[] memory messages = getMessagesWithTokens();
    // Set message 1 to use another receiver to simulate more fair gas costs
    messages[1].receiver = address(s_secondary_receiver);

    CCIP.ExecutionReport memory executionReport = createReportFromMessages(messages);

    CCIP.ExecutionResult[] memory report = s_offRamp.execute(executionReport, true);

    assertEq(report.length, messages.length);
    assertEq(report[0].sequenceNumber, messages[0].sequenceNumber);
    assertTrue(report[0].state == CCIP.MessageExecutionState.Success);
    assertEq(report[1].sequenceNumber, messages[1].sequenceNumber);
    assertTrue(report[1].state == CCIP.MessageExecutionState.Success);
  }
}
