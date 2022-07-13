// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import "../../onRamp/BaseOnRamp.sol";

contract BaseOnrampSetup is TokenSetup {
  event RouterSet(address router);
  event OnRampConfigSet(BaseOnRampInterface.OnRampConfig config);

  address[] public s_allowList;

  address public s_onRampRouter;
  BaseOnRamp public s_onRamp;
  BaseOnRampInterface.OnRampConfig public s_onRampConfig;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_onRampRouter = address(50);

    s_onRampConfig = BaseOnRampInterface.OnRampConfig({relayingFeeJuels: 0, maxDataSize: 50, maxTokensLength: 3});

    s_onRamp = new BaseOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      s_sourcePools,
      s_sourceFeeds,
      s_allowList,
      s_afn,
      HEARTBEAT,
      s_onRampConfig,
      s_onRampRouter
    );

    NativeTokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    NativeTokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);
  }

  function assertSameConfig(BaseOnRamp.OnRampConfig memory a, BaseOnRamp.OnRampConfig memory b) public {
    assertEq(a.relayingFeeJuels, b.relayingFeeJuels);
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
  }
}

// #constructor
contract BaseOnramp_constructor is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(SOURCE_CHAIN_ID, s_onRamp.CHAIN_ID());
    assertEq(DEST_CHAIN_ID, s_onRamp.DESTINATION_CHAIN_ID());
    assertEq(s_onRampRouter, s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());
    assertSameConfig(s_onRampConfig, s_onRamp.getConfig());
  }
}

// #getTokenPool
contract BaseOnramp_getTokenPool is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(address(s_sourcePools[0]), address(s_onRamp.getPool(s_sourceTokens[0])));
    assertEq(address(s_sourcePools[1]), address(s_onRamp.getPool(s_sourceTokens[1])));

    assertEq(address(0), address(s_onRamp.getPool(s_destTokens[0])));
  }
}

// #getExpectedNextSequenceNumber
contract BaseOnramp_getExpectedNextSequenceNumber is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());
  }
}

// #setRouter
contract BaseOnramp_setRouter is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(s_onRampRouter, s_onRamp.getRouter());
    address newRouter = address(100);

    vm.expectEmit(false, false, false, true);
    emit RouterSet(newRouter);

    s_onRamp.setRouter(newRouter);
    assertEq(newRouter, s_onRamp.getRouter());
  }

  // Revert
  function testSetRouterOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRamp.setRouter(address(1));
  }
}

// #getRouter
contract BaseOnramp_getRouter is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(s_onRampRouter, s_onRamp.getRouter());
  }
}

// #setConfig
contract BaseOnramp_setConfig is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    BaseOnRampInterface.OnRampConfig memory newConfig = BaseOnRampInterface.OnRampConfig({
      relayingFeeJuels: 2400,
      maxDataSize: 400,
      maxTokensLength: 14
    });

    vm.expectEmit(false, false, false, true);
    emit OnRampConfigSet(newConfig);

    s_onRamp.setConfig(newConfig);

    assertSameConfig(newConfig, s_onRamp.getConfig());
  }

  // Reverts
  function testSetConfigOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRamp.setConfig(s_onRampConfig);
  }
}

// #getConfig
contract BaseOnramp_getConfig is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertSameConfig(s_onRampConfig, s_onRamp.getConfig());
  }
}
