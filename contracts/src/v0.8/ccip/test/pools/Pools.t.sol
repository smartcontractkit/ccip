// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";

contract OffRampTokenPoolRegistrySetup is TokenSetup {
  OffRampTokenPoolRegistry tokenPoolRegistry;

  function setUp() public virtual override {
    TokenSetup.setUp();
    tokenPoolRegistry = new OffRampTokenPoolRegistry(getCastedSourceTokens(), getCastedDestinationPools());
  }
}

contract OffRampTokenPoolRegistry_getDestinationToken is OffRampTokenPoolRegistrySetup {
  function testSuccess() public {
    address expectedToken = address(PoolInterface(s_destPools[0]).getToken());
    address actualToken = address(tokenPoolRegistry.getDestinationToken(IERC20(s_sourceTokens[0])));

    assertEq(expectedToken, actualToken);

    expectedToken = address(PoolInterface(s_destPools[1]).getToken());
    actualToken = address(tokenPoolRegistry.getDestinationToken(IERC20(s_sourceTokens[1])));

    assertEq(expectedToken, actualToken);
  }
}

contract OffRampTokenPoolRegistry_getDestinationTokens is OffRampTokenPoolRegistrySetup {
  function testSuccess() public {
    IERC20[] memory actualTokens = tokenPoolRegistry.getDestinationTokens();

    for (uint256 i = 0; i < actualTokens.length; ++i) {
      assertEq(address(s_destTokens[i]), address(actualTokens[i]));
    }
  }
}
