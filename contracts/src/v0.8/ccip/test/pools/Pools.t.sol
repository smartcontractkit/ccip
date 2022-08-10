// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";

contract TokenPoolRegistrySetup is TokenSetup {
  TokenPoolRegistry tokenPoolRegistry;

  function setUp() public virtual override {
    TokenSetup.setUp();
    tokenPoolRegistry = new TokenPoolRegistry(s_sourceTokens, s_destPools);
  }
}

contract TokenPoolRegistry_getDestinationToken is TokenPoolRegistrySetup {
  function testSuccess() public {
    address expectedToken = address(s_destPools[0].getToken());

    address actualToken = address(tokenPoolRegistry.getDestinationToken(s_sourceTokens[0]));

    assertEq(expectedToken, actualToken);

    expectedToken = address(s_destPools[1].getToken());

    actualToken = address(tokenPoolRegistry.getDestinationToken(s_sourceTokens[1]));

    assertEq(expectedToken, actualToken);
  }
}

contract TokenPoolRegistry_getDestinationTokens is TokenPoolRegistrySetup {
  function testSuccess() public {
    IERC20[] memory actualTokens = tokenPoolRegistry.getDestinationTokens();

    for (uint256 i = 0; i < actualTokens.length; ++i) {
      assertEq(address(s_destTokens[i]), address(actualTokens[i]));
    }
  }
}
