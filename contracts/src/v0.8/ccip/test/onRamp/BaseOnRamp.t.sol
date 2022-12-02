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

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_onRampRouter = address(50);

    s_onRamp = new BaseOnRamp(
      SOURCE_CHAIN_ID,
      DEST_CHAIN_ID,
      getCastedSourceTokens(),
      getCastedSourcePools(),
      s_allowList,
      s_afn,
      onRampConfig(),
      rateLimiterConfig(),
      TOKEN_LIMIT_ADMIN,
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
    assertEq(SOURCE_CHAIN_ID, s_onRamp.i_chainId());
    assertEq(DEST_CHAIN_ID, s_onRamp.i_destinationChainId());
    assertEq(s_onRampRouter, s_onRamp.getRouter());
    assertEq(1, s_onRamp.getExpectedNextSequenceNumber());
    assertSameConfig(onRampConfig(), s_onRamp.getConfig());
  }
}

// #getTokenPool
contract BaseOnramp_getTokenPool is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertEq(s_sourcePools[0], address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[0]))));
    assertEq(s_sourcePools[1], address(s_onRamp.getPoolBySourceToken(IERC20(s_sourceTokens[1]))));

    vm.expectRevert(abi.encodeWithSelector(BaseOnRampInterface.UnsupportedToken.selector, IERC20(s_destTokens[0])));
    s_onRamp.getPoolBySourceToken(IERC20(s_destTokens[0]));
  }
}

// #getPoolTokens
contract BaseOnramp_getPoolTokens is BaseOnrampSetup {
  // Success
  function testGetPoolTokensSuccess() public {
    IERC20[] memory supportedTokens = s_onRamp.getPoolTokens();

    assertEq(address(s_sourceTokens[0]), address(supportedTokens[0]));
    assertEq(address(s_sourceTokens[1]), address(supportedTokens[1]));
    assertEq(s_sourceTokens.length, supportedTokens.length);

    s_onRamp.removePool(IERC20(s_sourceTokens[0]), PoolInterface(s_sourcePools[0]));

    supportedTokens = s_onRamp.getPoolTokens();

    assertEq(address(s_sourceTokens[1]), address(supportedTokens[0]));
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

// #setConfig
contract BaseOnramp_setConfig is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    BaseOnRampInterface.OnRampConfig memory newConfig = BaseOnRampInterface.OnRampConfig({
      commitFeeJuels: 2400,
      maxDataSize: 400,
      maxTokensLength: 14,
      maxGasLimit: MAX_GAS_LIMIT / 2
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
    s_onRamp.setConfig(onRampConfig());
  }
}

// #getConfig
contract BaseOnramp_getConfig is BaseOnrampSetup {
  // Success
  function testSuccess() public {
    assertSameConfig(onRampConfig(), s_onRamp.getConfig());
  }
}
