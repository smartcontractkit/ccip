// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./BaseTest.t.sol";
import "./mocks/MockERC20.sol";
import "./mocks/MockPool.sol";
import "../../tests/MockV3Aggregator.sol";
import "../pools/NativeTokenPool.sol";
import "../health/HealthChecker.sol";
import "../pools/TokenPoolRegistry.sol";

contract TokenSetup is BaseTest {
  IERC20[] internal s_sourceTokens;
  IERC20[] internal s_destTokens;

  PoolInterface[] internal s_sourcePools;
  PoolInterface[] internal s_destPools;

  uint256 internal constant TOKENS_PER_FEE_COIN = 2e20;

  function setUp() public virtual override {
    BaseTest.setUp();
    if (s_sourceTokens.length == 0) {
      s_sourceTokens.push(new MockERC20("sLINK", "sLNK", OWNER, 2**256 - 1));
      s_sourceTokens.push(new MockERC20("sETH", "sETH", OWNER, 2**128));
    }

    if (s_destTokens.length == 0) {
      s_destTokens.push(new MockERC20("dLINK", "dLNK", OWNER, 2**256 - 1));
      s_destTokens.push(new MockERC20("dETH", "dETH", OWNER, 2**128));
    }

    if (s_sourcePools.length == 0) {
      s_sourcePools.push(new NativeTokenPool(s_sourceTokens[0]));
      s_sourcePools.push(new NativeTokenPool(s_sourceTokens[1]));
    }

    if (s_destPools.length == 0) {
      s_destPools.push(new NativeTokenPool(s_destTokens[0]));
      s_destPools.push(new NativeTokenPool(s_destTokens[1]));

      // Float the pools with funds
      s_destTokens[0].transfer(address(s_destPools[0]), POOL_BALANCE);
      s_destTokens[1].transfer(address(s_destPools[1]), POOL_BALANCE);
    }
  }
}
