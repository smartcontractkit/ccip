// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./BaseTest.t.sol";
import "./mocks/MockERC20.sol";
import "./mocks/MockPool.sol";
import "../../tests/MockV3Aggregator.sol";
import "../pools/NativeTokenPool.sol";

contract TokenSetup is BaseTest {
  IERC20[] internal s_sourceTokens;
  IERC20[] internal s_destTokens;

  PoolInterface[] internal s_sourcePools;
  PoolInterface[] internal s_destPools;

  AggregatorV2V3Interface[] internal s_sourceFeeds;

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

    PoolInterface.BucketConfig memory bucketConfig = PoolInterface.BucketConfig({rate: 1e50, capacity: 1e50});

    if (s_sourcePools.length == 0) {
      s_sourcePools.push(new NativeTokenPool(s_sourceTokens[0], bucketConfig, bucketConfig));
      s_sourcePools.push(new NativeTokenPool(s_sourceTokens[1], bucketConfig, bucketConfig));
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
