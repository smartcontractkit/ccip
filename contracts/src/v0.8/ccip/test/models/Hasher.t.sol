// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../TokenSetup.t.sol";
import {Toll} from "../../models/Toll.sol";

contract Models__hash is TokenSetup {
  function testTollHashSuccess() public {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = new Common.EVMTokenAndAmount[](1);
    tokensAndAmounts[0].token = address(0x4440000000000000000000000000000000000001);
    tokensAndAmounts[0].amount = 12345678900;

    uint64 sourceChain = 1;
    uint64 destChain = 4;

    Toll.EVM2EVMTollMessage memory message = Toll.EVM2EVMTollMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      data: "",
      tokensAndAmounts: tokensAndAmounts,
      gasLimit: 100,
      feeTokenAndAmount: Common.EVMTokenAndAmount({
        token: address(0x3330000000000000000000000000000000000001),
        amount: 987654321
      })
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        Toll.EVM_2_EVM_TOLL_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(metadataHash, 0xa5d9be067fc21429efa4a6c47a0a5d867c500cde48c7057e1342285c9f3d2f1a);
    // Note this hash must match spec
    assertEq(Toll._hash(message, metadataHash), 0x21d6ad1f79e659726a6c6b41b0f05cfd4e4d24590a67775f85b3bca4aaff4265);
  }

  function testTollHashTwoTokensSuccess() public {
    Common.EVMTokenAndAmount[] memory tokensAndAmounts = new Common.EVMTokenAndAmount[](2);
    tokensAndAmounts[0].token = address(0x4440000000000000000000000000000000000001);
    tokensAndAmounts[1].token = address(0x6660000000000000000000000000000000000001);
    tokensAndAmounts[0].amount = 12345678900;
    tokensAndAmounts[1].amount = 4204242;

    uint64 sourceChain = 1;
    uint64 destChain = 4;

    Toll.EVM2EVMTollMessage memory message = Toll.EVM2EVMTollMessage({
      sourceChainId: sourceChain,
      sequenceNumber: 1337,
      sender: address(0x1110000000000000000000000000000000000001),
      receiver: address(0x2220000000000000000000000000000000000001),
      data: "foo bar baz",
      tokensAndAmounts: tokensAndAmounts,
      gasLimit: 100,
      feeTokenAndAmount: Common.EVMTokenAndAmount({
        token: address(0x3330000000000000000000000000000000000001),
        amount: 987654321
      })
    });

    bytes32 metadataHash = keccak256(
      abi.encode(
        Toll.EVM_2_EVM_TOLL_MESSAGE_HASH,
        sourceChain,
        destChain,
        address(0x5550000000000000000000000000000000000001)
      )
    );

    assertEq(metadataHash, 0xa5d9be067fc21429efa4a6c47a0a5d867c500cde48c7057e1342285c9f3d2f1a);
    // Note this hash must match spec
    assertEq(Toll._hash(message, metadataHash), 0x26095ef772ff770beb4f2d69ec828ff194589e146dc9cd19c84711c631b3fd49);
  }
}
