// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";

contract OffRampTokenPoolRegistrySetup is TokenSetup {
  OffRampTokenPoolRegistry s_tokenPoolRegistry;

  function setUp() public virtual override {
    TokenSetup.setUp();
    s_tokenPoolRegistry = new OffRampTokenPoolRegistry(getCastedSourceTokens(), getCastedDestinationPools());
  }
}

contract OffRampTokenPoolRegistry_addPool is OffRampTokenPoolRegistrySetup {
  event PoolAdded(IERC20 token, IPool pool);

  function testAddPoolSuccess() public {
    IPool newPool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));

    vm.expectEmit(false, false, false, true);
    emit PoolAdded(token, newPool);
    s_tokenPoolRegistry.addPool(token, newPool);

    IPool actualPool = s_tokenPoolRegistry.getPoolBySourceToken(token);

    assertEq(address(newPool), address(actualPool));
  }

  // Reverts

  function testInvalidTokenPoolConfigReverts() public {
    address zero = address(0);

    vm.expectRevert(OffRampTokenPoolRegistry.InvalidTokenPoolConfig.selector);
    s_tokenPoolRegistry.addPool(IERC20(zero), IPool(s_destPools[0]));

    vm.expectRevert(OffRampTokenPoolRegistry.InvalidTokenPoolConfig.selector);
    s_tokenPoolRegistry.addPool(IERC20(s_destFeeToken), IPool(zero));
  }

  function testPoolAlreadyAddedReverts() public {
    IPool newPool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));
    s_tokenPoolRegistry.addPool(token, newPool);

    vm.expectRevert(OffRampTokenPoolRegistry.PoolAlreadyAdded.selector);
    s_tokenPoolRegistry.addPool(token, newPool);
  }

  function testOnlyOwnerReverts() public {
    IPool newPool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));

    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");
    s_tokenPoolRegistry.addPool(token, newPool);
  }
}

contract OffRampTokenPoolRegistry_removePool is OffRampTokenPoolRegistrySetup {
  event PoolRemoved(IERC20 token, IPool pool);

  function testRemovePoolSuccess() public {
    IPool pool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));

    s_tokenPoolRegistry.addPool(token, pool);
    IPool actualPool = s_tokenPoolRegistry.getPoolBySourceToken(token);
    assertEq(address(pool), address(actualPool));

    vm.expectEmit(false, false, false, true);
    emit PoolRemoved(token, pool);

    s_tokenPoolRegistry.removePool(token, pool);

    actualPool = s_tokenPoolRegistry.getPoolBySourceToken(token);
    assertEq(address(0), address(actualPool));
  }

  // Reverts

  function testTokenPoolMismatchReverts() public {
    IPool correctPool = new LockReleaseTokenPool(IERC20(address(2)));
    IPool wrongPool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));

    s_tokenPoolRegistry.addPool(token, correctPool);

    vm.expectRevert(OffRampTokenPoolRegistry.TokenPoolMismatch.selector);
    s_tokenPoolRegistry.removePool(token, wrongPool);
  }

  function testPoolDoesNotExistReverts() public {
    IPool newPool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));

    vm.expectRevert(OffRampTokenPoolRegistry.PoolDoesNotExist.selector);
    s_tokenPoolRegistry.removePool(token, newPool);
  }

  function testOnlyOwnerReverts() public {
    IPool newPool = new LockReleaseTokenPool(IERC20(address(2)));
    IERC20 token = IERC20(address(1));

    changePrank(STRANGER);

    vm.expectRevert("Only callable by owner");
    s_tokenPoolRegistry.removePool(token, newPool);
  }
}

contract OffRampTokenPoolRegistry_getDestinationToken is OffRampTokenPoolRegistrySetup {
  function testGetDestinationTokenSuccess() public {
    address expectedToken = address(IPool(s_destPools[0]).getToken());
    address actualToken = address(s_tokenPoolRegistry.getDestinationToken(IERC20(s_sourceTokens[0])));

    assertEq(expectedToken, actualToken);

    expectedToken = address(IPool(s_destPools[1]).getToken());
    actualToken = address(s_tokenPoolRegistry.getDestinationToken(IERC20(s_sourceTokens[1])));

    assertEq(expectedToken, actualToken);
  }
}

contract OffRampTokenPoolRegistry_getDestinationTokens is OffRampTokenPoolRegistrySetup {
  function testGetDestinationTokensSuccess() public {
    IERC20[] memory actualTokens = s_tokenPoolRegistry.getDestinationTokens();

    for (uint256 i = 0; i < actualTokens.length; ++i) {
      assertEq(address(s_destTokens[i]), address(actualTokens[i]));
    }
  }
}
