// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";

/// @notice #_addToTokensAmounts
contract Internal__addToTokensAmounts is TokenSetup {
  function test_MatchingTokensSuccess() public {
    address matchingTokenAddress = address(235);
    uint256 totalTokens = 3;
    Client.EVMTokenAmount[] memory existingTokens = new Client.EVMTokenAmount[](totalTokens);
    existingTokens[0] = Client.EVMTokenAmount({token: address(9), amount: 565157});
    existingTokens[1] = Client.EVMTokenAmount({token: matchingTokenAddress, amount: 67124});
    existingTokens[2] = Client.EVMTokenAmount({token: address(10), amount: 8732});

    Client.EVMTokenAmount memory newToken = Client.EVMTokenAmount({token: matchingTokenAddress, amount: 89124});

    Client.EVMTokenAmount[] memory combinedTokens = Internal._addToTokensAmounts(existingTokens, newToken);

    assertEq(combinedTokens.length, totalTokens);
    assertEq(combinedTokens[0].token, existingTokens[0].token);
    assertEq(combinedTokens[0].amount, existingTokens[0].amount);
    assertEq(combinedTokens[1].token, existingTokens[1].token);
    assertEq(combinedTokens[1].amount, existingTokens[1].amount + newToken.amount);
    assertEq(combinedTokens[2].token, existingTokens[2].token);
    assertEq(combinedTokens[2].amount, existingTokens[2].amount);
  }

  function test_NonMatchingTokensSuccess() public {
    uint256 totalTokens = 3;
    Client.EVMTokenAmount[] memory existingTokens = new Client.EVMTokenAmount[](totalTokens);
    existingTokens[0] = Client.EVMTokenAmount({token: address(9), amount: 565157});
    existingTokens[1] = Client.EVMTokenAmount({token: address(10), amount: 67124});
    existingTokens[2] = Client.EVMTokenAmount({token: address(11), amount: 8732});

    Client.EVMTokenAmount memory newToken = Client.EVMTokenAmount({token: address(12), amount: 89124});

    Client.EVMTokenAmount[] memory combinedTokens = Internal._addToTokensAmounts(existingTokens, newToken);

    assertEq(combinedTokens.length, totalTokens + 1);
    assertEq(combinedTokens[0].token, existingTokens[0].token);
    assertEq(combinedTokens[0].amount, existingTokens[0].amount);
    assertEq(combinedTokens[1].token, existingTokens[1].token);
    assertEq(combinedTokens[1].amount, existingTokens[1].amount);
    assertEq(combinedTokens[2].token, existingTokens[2].token);
    assertEq(combinedTokens[2].amount, existingTokens[2].amount);
    assertEq(combinedTokens[3].token, newToken.token);
    assertEq(combinedTokens[3].amount, newToken.amount);
  }

  function testEmptyFirstArgumentSuccess() public {
    Client.EVMTokenAmount[] memory emptyExistingTokens = new Client.EVMTokenAmount[](0);
    Client.EVMTokenAmount memory newToken = Client.EVMTokenAmount({token: address(2345678), amount: 120});

    Client.EVMTokenAmount[] memory combinedTokens = Internal._addToTokensAmounts(emptyExistingTokens, newToken);
    assertEq(combinedTokens.length, 1);
    assertEq(combinedTokens[0].token, newToken.token);
    assertEq(combinedTokens[0].amount, newToken.amount);
  }
}
