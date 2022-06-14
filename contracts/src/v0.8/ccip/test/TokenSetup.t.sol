// SPDX-License-Identifier: MIT
pragma solidity 0.8.13;

import "./BaseTest.t.sol";
import "./mocks/MockERC20.sol";
import "./mocks/MockPool.sol";
import "../../tests/MockV3Aggregator.sol";

contract TokenSetup is BaseTest {
  IERC20[] public s_sourceTokens;
  IERC20[] public s_destTokens;

  PoolInterface[] s_sourcePools;
  PoolInterface[] s_destPools;

  AggregatorV2V3Interface[] public s_sourceFeeds;

  function setUp() public virtual override {
    BaseTest.setUp();
    if (s_sourceTokens.length == 0) {
      s_sourceTokens.push(new MockERC20("sLINK", "sLNK", s_owner, 2**256 - 1));
      s_sourceTokens.push(new MockERC20("sETH", "sETH", s_owner, 2**128));
    }

    if (s_destTokens.length == 0) {
      s_destTokens.push(new MockERC20("dLINK", "dLNK", s_owner, 2**256 - 1));
      s_destTokens.push(new MockERC20("dETH", "dETH", s_owner, 2**128));
    }

    if (s_sourcePools.length == 0) {
      s_sourcePools.push(new MockPool(5));
      s_sourcePools.push(new MockPool(10));
    }

    if (s_destPools.length == 0) {
      s_destPools.push(new MockPool(5));
      s_destPools.push(new MockPool(10));
    }

    if (s_sourceFeeds.length == 0) {
      s_sourceFeeds.push(new MockV3Aggregator(0, 1));
      s_sourceFeeds.push(new MockV3Aggregator(0, 2));
    }
  }
}
