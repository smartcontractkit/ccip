// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";

contract Models__hash is TokenSetup {
  using CCIP for CCIP.EVM2EVMSubscriptionMessage;
  using CCIP for CCIP.EVM2EVMTollMessage;

  function testSubscriptionHashSuccess() public {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    tokensAndAmounts[0].token = address(0x4440000000000000000000000000000000000001);
    tokensAndAmounts[0].amount = 12345678900;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMSubscriptionMessage memory message = CCIP.EVM2EVMSubscriptionMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      nonce: 666,
      data: "",
      tokensAndAmounts: tokensAndAmounts,
      gasLimit: 100
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        CCIP.EVM_2_EVM_SUBSCRIPTION_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(metadataHash, 0xe8b93c9d01a7a72ec6c7235e238701cf1511b267a31fdb78dd342649ee58c08d);
    // Note this hash must match spec
    assertEq(message._hash(metadataHash), 0x6b4d88effbfa2121b6e1c16918d5a0003bb68437473117daaf631925e949bd02);
  }

  function testSubscriptionHashTwoTokensSuccess() public {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](2);
    tokensAndAmounts[0].token = address(0x4440000000000000000000000000000000000001);
    tokensAndAmounts[1].token = address(0x6660000000000000000000000000000000000001);
    tokensAndAmounts[0].amount = 12345678900;
    tokensAndAmounts[1].amount = 4204242;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMSubscriptionMessage memory message = CCIP.EVM2EVMSubscriptionMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      nonce: 210,
      data: "foo bar baz",
      tokensAndAmounts: tokensAndAmounts,
      gasLimit: 100
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        CCIP.EVM_2_EVM_SUBSCRIPTION_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(metadataHash, 0xe8b93c9d01a7a72ec6c7235e238701cf1511b267a31fdb78dd342649ee58c08d);
    // Note this hash must match spec
    assertEq(message._hash(metadataHash), 0xfaa461863c42b1548d5687f535d81be27ec84db8b6f60d904beeac07abcad71a);
  }

  function testTollHashSuccess() public {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](1);
    tokensAndAmounts[0].token = address(0x4440000000000000000000000000000000000001);
    tokensAndAmounts[0].amount = 12345678900;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMTollMessage memory message = CCIP.EVM2EVMTollMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      data: "",
      tokensAndAmounts: tokensAndAmounts,
      gasLimit: 100,
      feeTokenAndAmount: CCIP.EVMTokenAndAmount({
        token: address(0x3330000000000000000000000000000000000001),
        amount: 987654321
      })
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        CCIP.EVM_2_EVM_TOLL_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(metadataHash, 0x73ba062fc2abb9b2d37ef43de292ddd56a89f10ee0e344d500e63a0474073b03);
    // Note this hash must match spec
    assertEq(message._hash(metadataHash), 0xd4504baca27221b294969ab5a2989e2121cb3577a209b85d7d83371b3429df4d);
  }

  function testTollHashTwoTokensSuccess() public {
    CCIP.EVMTokenAndAmount[] memory tokensAndAmounts = new CCIP.EVMTokenAndAmount[](2);
    tokensAndAmounts[0].token = address(0x4440000000000000000000000000000000000001);
    tokensAndAmounts[1].token = address(0x6660000000000000000000000000000000000001);
    tokensAndAmounts[0].amount = 12345678900;
    tokensAndAmounts[1].amount = 4204242;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMTollMessage memory message = CCIP.EVM2EVMTollMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      data: "foo bar baz",
      tokensAndAmounts: tokensAndAmounts,
      gasLimit: 100,
      feeTokenAndAmount: CCIP.EVMTokenAndAmount({
        token: address(0x3330000000000000000000000000000000000001),
        amount: 987654321
      })
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        CCIP.EVM_2_EVM_TOLL_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(metadataHash, 0x73ba062fc2abb9b2d37ef43de292ddd56a89f10ee0e344d500e63a0474073b03);
    // Note this hash must match spec
    assertEq(message._hash(metadataHash), 0xc0b2bac538afab5af9c654028ff27f3a3cc5aa9e1082efc70656b8467dd41fb2);
  }
}
