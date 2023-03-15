// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "./BaseTest.t.sol";
import "./mocks/MockERC20.sol";
import "../../tests/MockV3Aggregator.sol";
import "../pools/BurnMintTokenPool.sol";
import "../pools/LockReleaseTokenPool.sol";
import "../models/Client.sol";

contract TokenSetup is BaseTest {
  address[] internal s_sourceTokens;
  address[] internal s_destTokens;

  address[] internal s_sourcePools;
  address[] internal s_destPools;

  address internal s_sourceFeeToken;
  address internal s_destFeeToken;

  IPool internal s_destFeeTokenPool;

  function setUp() public virtual override {
    BaseTest.setUp();

    // Source tokens & pools
    if (s_sourceTokens.length == 0) {
      s_sourceTokens.push(address(new MockERC20("sLINK", "sLNK", OWNER, 2**256 - 1)));
      s_sourceTokens.push(address(new MockERC20("sETH", "sETH", OWNER, 2**128)));
    }

    if (s_sourcePools.length == 0) {
      s_sourcePools.push(address(new LockReleaseTokenPool(IERC20(s_sourceTokens[0]))));
      s_sourcePools.push(address(new BurnMintTokenPool(IBurnMintERC20(s_sourceTokens[1]))));
    }

    s_sourceFeeToken = s_sourceTokens[0];

    // Destination tokens & pools
    if (s_destTokens.length == 0) {
      s_destTokens.push(address(new MockERC20("dLINK", "dLNK", OWNER, 2**256 - 1)));
      s_destTokens.push(address(new MockERC20("dETH", "dETH", OWNER, 2**128)));
    }

    if (s_destPools.length == 0) {
      s_destPools.push(address(new LockReleaseTokenPool(IERC20(s_destTokens[0]))));
      s_destPools.push(address(new BurnMintTokenPool(IBurnMintERC20(s_destTokens[1]))));

      // Float the pools with funds
      IERC20(s_destTokens[0]).transfer(address(s_destPools[0]), POOL_BALANCE);
      IERC20(s_destTokens[1]).transfer(address(s_destPools[1]), POOL_BALANCE);
    }

    s_destFeeToken = s_destTokens[0];
    s_destFeeTokenPool = IPool(s_destPools[0]);
  }

  function getCastedSourceEVMTokenAmountsWithZeroAmounts()
    internal
    view
    returns (Client.EVMTokenAmount[] memory tokenAmounts)
  {
    tokenAmounts = new Client.EVMTokenAmount[](s_sourceTokens.length);
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      tokenAmounts[i].token = s_sourceTokens[i];
    }
  }

  function getCastedDestinationEVMTokenAmountsWithZeroAmounts()
    internal
    view
    returns (Client.EVMTokenAmount[] memory tokenAmounts)
  {
    tokenAmounts = new Client.EVMTokenAmount[](s_destTokens.length);
    for (uint256 i = 0; i < tokenAmounts.length; ++i) {
      tokenAmounts[i].token = s_destTokens[i];
    }
  }

  function getCastedSourceTokens() internal view returns (IERC20[] memory sourceTokens) {
    // Convert address array into IERC20 array in one line
    sourceTokens = abi.decode(abi.encode(s_sourceTokens), (IERC20[]));
  }

  function getCastedDestinationTokens() internal view returns (IERC20[] memory destTokens) {
    // Convert address array into IERC20 array in one line
    destTokens = abi.decode(abi.encode(s_destTokens), (IERC20[]));
  }

  function getCastedSourcePools() internal view returns (IPool[] memory sourcePools) {
    // Convert address array into IPool array in one line
    sourcePools = abi.decode(abi.encode(s_sourcePools), (IPool[]));
  }

  function getCastedDestinationPools() internal view returns (IPool[] memory destPools) {
    // Convert address array into IPool array in one line
    destPools = abi.decode(abi.encode(s_destPools), (IPool[]));
  }
}
