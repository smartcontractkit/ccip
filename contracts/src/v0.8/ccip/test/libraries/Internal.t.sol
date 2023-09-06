// SPDX-License-Identifier: BUSL-1.1
pragma solidity 0.8.19;

import {TokenSetup} from "../TokenSetup.t.sol";
import {Internal} from "../../libraries/Internal.sol";
import {Client} from "../../libraries/Client.sol";

contract InternalHelperSetup is TokenSetup {
  function setUp() public virtual override {
    TokenSetup.setUp();
  }
}

contract Internal_hash is InternalHelperSetup {
  function test_SameHashResultSimple() public {
    Client.EVMTokenAmount[] memory tokenAmounts = new Client.EVMTokenAmount[](2);
    tokenAmounts[0].token = s_sourceTokens[0];
    tokenAmounts[0].amount = 100e18;
    tokenAmounts[1].token = s_sourceTokens[1];
    tokenAmounts[1].amount = 200e18;

    Internal.EVM2EVMMessage memory message = Internal.EVM2EVMMessage({
      sequenceNumber: 1,
      sender: OWNER,
      nonce: 1,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainSelector: SOURCE_CHAIN_ID,
      receiver: address(123),
      data: abi.encode("Simple Test Message"),
      tokenAmounts: tokenAmounts,
      sourceTokenData: new bytes[](tokenAmounts.length),
      feeToken: s_sourceFeeToken,
      feeTokenAmount: uint256(100),
      messageId: ""
    });
    bytes32 betterHash = Internal._hash(
      message,
      keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
    );
    bytes32 currentHash = Internal._hashCurrent(
      message,
      keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
    );
    bytes32 legacyHash = Internal._hashLegacy(
      message,
      keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
    );

    assertEq(betterHash, currentHash);
    assertEq(betterHash, legacyHash);
  }

  function testFuzz_SameHashResultSimple(
    address sender,
    address receiver,
    uint64 sequenceNumber,
    uint64 nonce,
    address feeToken,
    uint256 feeTokenAmount,
    bytes calldata data,
    Client.EVMTokenAmount[] calldata tokenAmounts,
    bytes[] calldata sourceTokenData
  ) public {
    Internal.EVM2EVMMessage memory message = Internal.EVM2EVMMessage({
      sequenceNumber: sequenceNumber,
      sender: sender,
      nonce: nonce,
      gasLimit: GAS_LIMIT,
      strict: false,
      sourceChainSelector: SOURCE_CHAIN_ID,
      receiver: receiver,
      data: data,
      tokenAmounts: tokenAmounts,
      sourceTokenData: sourceTokenData,
      feeToken: feeToken,
      feeTokenAmount: feeTokenAmount,
      messageId: ""
    });

    assertEq(
      Internal._hash(
        message,
        keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
      ),
      Internal._hashCurrent(
        message,
        keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
      )
    );

    assertEq(
      Internal._hash(
        message,
        keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
      ),
      Internal._hashLegacy(
        message,
        keccak256(abi.encode(Internal.EVM_2_EVM_MESSAGE_HASH, SOURCE_CHAIN_ID, DEST_CHAIN_ID, ON_RAMP_ADDRESS))
      )
    );
  }
}
