// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import "../../onRamp/BaseOnRamp.sol";

contract BaseOnrampSetup is TokenSetup {
  event RouterSet(address router);
  event OnRampConfigSet(IBaseOnRamp.OnRampConfig config);

  address[] public s_allowList;

  address public s_onRampRouter;
  BaseOnRamp public s_onRamp;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_onRampRouter = address(50);

    s_onRamp = new BaseOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      s_sourceTokens,
      getCastedSourcePools(),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      s_onRampRouter
    );

    s_onRamp.setPrices(getCastedSourceTokens(), getTokenPrices());

    TokenPool(address(s_sourcePools[0])).setOnRamp(s_onRamp, true);
    TokenPool(address(s_sourcePools[1])).setOnRamp(s_onRamp, true);
  }

  function assertSameConfig(BaseOnRamp.OnRampConfig memory a, BaseOnRamp.OnRampConfig memory b) public {
    assertEq(a.commitFeeJuels, b.commitFeeJuels);
    assertEq(a.maxDataSize, b.maxDataSize);
    assertEq(a.maxTokensLength, b.maxTokensLength);
    assertEq(a.maxGasLimit, b.maxGasLimit);
  }
}

// #constructor
contract BaseOnramp_constructor is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(SOURCE_CHAIN_ID, s_onRamp.getChainId());
    assertEq(DEST_CHAIN_ID, s_onRamp.getDestinationChainId());
    assertEq(s_onRampRouter, s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());
    assertSameConfig(onRampConfig(), s_onRamp.getOnRampConfig());
  }
}

// #getTokenPool
contract BaseOnramp_getTokenPool is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(s_sourcePools[0], address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[0]))));
    assertEq(s_sourcePools[1], address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[1]))));

    vm.expectRevert(abi.encodeWithSelector(IBaseOnRamp.UnsupportedToken.selector, IERC20(s_destTokens[0])));
    s_onRamp.getPoolBySourceToken(IERC20(s_destTokens[0]));
  }
}

// #getSupportedTokens
contract BaseOnramp_getSupportedTokens is BaseOnrampSetup {
  // Success
  function testGetSupportedTokensSuccess() public {
    address[] memory supportedTokens = s_onRamp.getSupportedTokens();

    assertEq(s_sourceTokens, supportedTokens);

    s_onRamp.removePool(IERC20(s_sourceTokens[0]), IPool(s_sourcePools[0]));

    supportedTokens = s_onRamp.getSupportedTokens();

    assertEq(address(s_sourceTokens[1]), supportedTokens[0]);
    assertEq(s_sourceTokens.length - 1, supportedTokens.length);
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

// #setOnRampConfig
contract BaseOnramp_setOnRampConfig is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    IBaseOnRamp.OnRampConfig memory newConfig = IBaseOnRamp.OnRampConfig({
      commitFeeJuels: 2400,
      maxDataSize: 400,
      maxTokensLength: 14,
      maxGasLimit: MAX_GAS_LIMIT / 2
    });

    vm.expectEmit(false, false, false, true);
    emit OnRampConfigSet(newConfig);

    s_onRamp.setOnRampConfig(newConfig);

    assertSameConfig(newConfig, s_onRamp.getOnRampConfig());
  }

  // Reverts
  function testSetConfigOnlyOwnerReverts() public {
    vm.stopPrank();
    vm.expectRevert("Only callable by owner");
    s_onRamp.setOnRampConfig(onRampConfig());
  }
}

// #getConfig
contract BaseOnramp_getConfig is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertSameConfig(onRampConfig(), s_onRamp.getOnRampConfig());
  }
}
