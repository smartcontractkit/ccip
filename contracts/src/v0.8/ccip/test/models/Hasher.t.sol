// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";

contract Models__hash is TokenSetup {
  using CCIP for CCIP.EVM2EVMSubscriptionMessage;
  using CCIP for CCIP.EVM2EVMTollMessage;

  function testSubscriptionHashSuccess() public {
    address[] memory tokens = new address[](1);
    tokens[0] = address(0x4440000000000000000000000000000000000001);
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 12345678900;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMSubscriptionMessage memory message = CCIP.EVM2EVMSubscriptionMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      nonce: 666,
      data: "",
      tokens: tokens,
      amounts: amounts,
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

    assertEq(0xe8b93c9d01a7a72ec6c7235e238701cf1511b267a31fdb78dd342649ee58c08d, metadataHash);
    // Note this hash must match spec
    assertEq(0xcae032f60dc29a4d98e135908afa3f562674954c9d3378606e8b0473d27e94c9, message._hash(metadataHash));
  }

  function testSubscriptionHashTwoTokensSuccess() public {
    address[] memory tokens = new address[](2);
    tokens[0] = address(0x4440000000000000000000000000000000000001);
    tokens[1] = address(0x6660000000000000000000000000000000000001);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 12345678900;
    amounts[1] = 4204242;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMSubscriptionMessage memory message = CCIP.EVM2EVMSubscriptionMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      nonce: 210,
      data: "foo bar baz",
      tokens: tokens,
      amounts: amounts,
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

    assertEq(0xe8b93c9d01a7a72ec6c7235e238701cf1511b267a31fdb78dd342649ee58c08d, metadataHash);
    // Note this hash must match spec
    assertEq(0xaef2f373966c54aec50e619bacd6e66275f660c5b5ff3bd53b00386d345bcfa9, message._hash(metadataHash));
  }

  function testTollHashSuccess() public {
    address[] memory tokens = new address[](1);
    tokens[0] = address(0x4440000000000000000000000000000000000001);
    uint256[] memory amounts = new uint256[](1);
    amounts[0] = 12345678900;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMTollMessage memory message = CCIP.EVM2EVMTollMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      data: "",
      tokens: tokens,
      amounts: amounts,
      gasLimit: 100,
      feeToken: address(0x3330000000000000000000000000000000000001),
      feeTokenAmount: 987654321
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        CCIP.EVM_2_EVM_TOLL_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(0x73ba062fc2abb9b2d37ef43de292ddd56a89f10ee0e344d500e63a0474073b03, metadataHash);
    // Note this hash must match spec
    assertEq(0x9c014cce73a389409d5dbc863cb4d0054e61698bafb21eb88cafd670ee45ed12, message._hash(metadataHash));
  }

  function testTollHashTwoTokensSuccess() public {
    address[] memory tokens = new address[](2);
    tokens[0] = address(0x4440000000000000000000000000000000000001);
    tokens[1] = address(0x6660000000000000000000000000000000000001);
    uint256[] memory amounts = new uint256[](2);
    amounts[0] = 12345678900;
    amounts[1] = 4204242;

    uint256 sourceChain = 1;
    uint256 destChain = 4;

    CCIP.EVM2EVMTollMessage memory message = CCIP.EVM2EVMTollMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      data: "foo bar baz",
      tokens: tokens,
      amounts: amounts,
      gasLimit: 100,
      feeToken: address(0x3330000000000000000000000000000000000001),
      feeTokenAmount: 987654321
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        CCIP.EVM_2_EVM_TOLL_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(0x73ba062fc2abb9b2d37ef43de292ddd56a89f10ee0e344d500e63a0474073b03, metadataHash);
    // Note this hash must match spec
    assertEq(0xb70e53658377bb46b430d3ca5bbfed10c1e97d82dd8feb0af896224b4bf890c8, message._hash(metadataHash));
  }
}
