// SPDX-License-Identifier: MIT
pragma solidity ^0.8.15;

import "../../../onRamp/toll/EVM2EVMTollOnRamp.sol";
import "../../../onRamp/toll/EVM2AnyTollOnRampRouter.sol";
import "./OnRampSetup.t.sol";

contract EVM2EVMTollOnRampTest is OnRampSetup {
  address[] public s_allowList;

  TollOnRampRouterInterface public s_onRampRouter;
  TollOnRampInterface public s_onRamp;

  TollOnRampInterface.OnRampConfig public s_onRampConfig;

  function setUp() public virtual override {
    OnRampSetup.setUp();
    s_onRampRouter = new EVM2AnyTollOnRampRouter();

    s_onRampConfig = TollOnRampInterface.OnRampConfig({
      router: address(s_onRampRouter),
      relayingFeeJuels: 0,
      maxDataSize: 50,
      maxTokensLength: 3
    });

    s_onRamp = new EVM2EVMTollOnRamp(
      s_sourceChainId,
      s_destChainId,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      s_allowList,
      s_afn,
      1e18,
      s_onRampConfig
    );

    s_onRampRouter.setOnRamp(s_destChainId, s_onRamp);

    // Pre approve the first token so the gas estimates of the tests
    // only cover actual gas usage from the ramps
    s_sourceTokens[0].approve(address(s_onRampRouter), 2**128);
  }

  function testGetRequiredFee() public {
    s_onRamp.getRequiredFee(s_sourceTokens[0]);
  }

  // Asserts that the correct event is emitted after a tx
  // is processed by the onramp. Each property of the event
  // is checked with this testing syntax.
  function testShouldEmitCCIPSendRequested() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    CCIP.EVM2EVMTollEvent memory tollEvent = messageToEvent(message, 1);

    vm.expectEmit(true, false, false, true);
    emit CCIPSendRequested(tollEvent);

    s_onRampRouter.ccipSend(s_destChainId, message);
  }

  function testRequestXChainSendsExactApprove() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**128;
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = s_sourceTokens[0];
    message.amounts = amounts;
    message.tokens = tokens;
    uint64 seqNum = s_onRampRouter.ccipSend(s_destChainId, message);
    assertEq(seqNum, 1);
  }

  function testRequestXChainSends() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 2**64;
    IERC20[] memory tokens = new IERC20[](1);
    tokens[0] = s_sourceTokens[0];
    message.amounts = amounts;
    message.tokens = tokens;
    uint64 seqNum = s_onRampRouter.ccipSend(s_destChainId, message);
    assertEq(seqNum, 1);
  }

  function testShouldIncrementSeqNum() public {
    uint64 seqNum = s_onRampRouter.ccipSend(s_destChainId, getEmptyMessage());
    assertEq(seqNum, 1);
    seqNum = s_onRampRouter.ccipSend(s_destChainId, getEmptyMessage());
    assertEq(seqNum, 2);
    seqNum = s_onRampRouter.ccipSend(s_destChainId, getEmptyMessage());
    assertEq(seqNum, 3);
  }

  function testOnlyKnownOnRamps() public {
    uint256 wrongChain = s_destChainId + 1;
    vm.expectRevert(abi.encodeWithSelector(BaseOnRampRouterInterface.UnsupportedDestinationChain.selector, wrongChain));
    s_onRampRouter.ccipSend(wrongChain, getEmptyMessage());
  }

  function testMustBeCalledByRouter() public {
    vm.expectRevert(BaseOnRampInterface.MustBeCalledByRouter.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), s_owner);
  }

  function testRouterMustSetOriginalSender() public {
    vm.stopPrank();
    vm.startPrank(address(s_onRampRouter));
    vm.expectRevert(BaseOnRampInterface.RouterMustSetOriginalSender.selector);
    s_onRamp.forwardFromRouter(getEmptyMessage(), address(0));
  }

  function testMessageTooLarge() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.data = "000000000000000000000000000000000000000000000000000";
    vm.expectRevert(
      abi.encodeWithSelector(
        BaseOnRampInterface.MessageTooLarge.selector,
        s_onRampConfig.maxDataSize,
        message.data.length
      )
    );

    s_onRampRouter.ccipSend(s_destChainId, message);
  }

  function testUnsupportedNumberOfTokens() public {
    CCIP.EVM2AnyTollMessage memory message = getEmptyMessage();
    message.tokens = new IERC20[](s_onRampConfig.maxTokensLength + 1);
    vm.expectRevert(BaseOnRampInterface.UnsupportedNumberOfTokens.selector);
    s_onRampRouter.ccipSend(s_destChainId, message);
  }
}
