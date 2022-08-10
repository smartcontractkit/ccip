// SPDX-License-Identifier: MIT
pragma solidity 0.8.15;

import "../helpers/MerkleHelper.sol";
import "./BlobVerifier.t.sol";

contract BlobVerifier_merkleRoot is BlobVerifierSetup {
  MerkleHelper s_merkleHelper;

  function setUp() public virtual override {
    BlobVerifierSetup.setUp();
    s_merkleHelper = new MerkleHelper();
  }

  function testMerkleRoot2() public {
    bytes32 expectedRoot = 0x82bb4cfba0e54ea8cf3e7928b2afeb764e964163cff346fe612f9ba0b4154295;
    bytes32[] memory leaves = new bytes32[](2);
    leaves[0] = keccak256("a");
    leaves[1] = keccak256("b");
    bytes32[] memory proofs = new bytes32[](0);

    bytes32 root = s_blobVerifier.merkleRoot(leaves, proofs, 2**100 - 1);

    assertEq(root, expectedRoot);
  }

  function testMerkleRoot4() public {
    bytes32 expectedRoot = 0xabc9a9e8c822c572596afbd6ccc83e7f504efad22c2757eb6ded56a2428a6927;
    bytes32[] memory leaves = new bytes32[](4);
    leaves[0] = keccak256("a");
    leaves[1] = keccak256("b");
    leaves[2] = keccak256("c");
    leaves[3] = keccak256("d");
    bytes32[] memory proofs = new bytes32[](0);

    bytes32 root = s_blobVerifier.merkleRoot(leaves, proofs, 2**100 - 1);

    assertEq(root, expectedRoot);
  }

  function testMerkleRoot6() public {
    bytes32 expectedRoot = 0xc8a718ada8222645ec770e7e86686f7276e0c4a77f7979947c50a390879449c2;
    bytes32[] memory leaves = new bytes32[](6);
    leaves[0] = keccak256("a");
    leaves[1] = keccak256("b");
    leaves[2] = keccak256("c");
    leaves[3] = keccak256("d");
    leaves[4] = keccak256("e");
    leaves[5] = keccak256("f");
    bytes32[] memory proofs = new bytes32[](0);

    bytes32 root = s_blobVerifier.merkleRoot(leaves, proofs, 2**100 - 1);

    assertEq(root, expectedRoot);
  }

  function testMerkleRoot256() public {
    bytes32 expectedRoot = 0x37fea55ec5229f7d71bfc23c1975dc5c33e5178e1ddfc2bb0ff254fb3913d80d;
    bytes32[] memory leaves = new bytes32[](256);
    for (uint256 i = 0; i < leaves.length; ++i) {
      leaves[i] = keccak256("a");
    }
    bytes32[] memory proofs = new bytes32[](0);

    bytes32 root = s_blobVerifier.merkleRoot(leaves, proofs, 2**256 - 1);

    assertEq(root, expectedRoot);
  }

  function testMerkleMulti1of4(
    bytes32 leaf1,
    bytes32 proof1,
    bytes32 proof2
  ) public {
    bytes32[] memory leaves = new bytes32[](1);
    leaves[0] = leaf1;
    bytes32[] memory proofs = new bytes32[](2);
    proofs[0] = proof1;
    proofs[1] = proof2;

    // Proof flag = false
    bytes32 result = s_merkleHelper.hashPair(leaves[0], proofs[0]);
    // Proof flag = false
    result = s_merkleHelper.hashPair(result, proofs[1]);

    assertEq(s_blobVerifier.merkleRoot(leaves, proofs, 0), result);
  }

  function testMerkleMulti2of4(
    bytes32 leaf1,
    bytes32 leaf2,
    bytes32 proof1,
    bytes32 proof2
  ) public {
    bytes32[] memory leaves = new bytes32[](2);
    leaves[0] = leaf1;
    leaves[1] = leaf2;
    bytes32[] memory proofs = new bytes32[](2);
    proofs[0] = proof1;
    proofs[1] = proof2;

    // Proof flag = false
    bytes32 result1 = s_merkleHelper.hashPair(leaves[0], proofs[0]);
    // Proof flag = false
    bytes32 result2 = s_merkleHelper.hashPair(leaves[1], proofs[1]);
    // Proof flag = true
    bytes32 finalResult = s_merkleHelper.hashPair(result1, result2);

    assertEq(s_blobVerifier.merkleRoot(leaves, proofs, 4), finalResult);
  }

  function testMerkleMulti3of4(
    bytes32 leaf1,
    bytes32 leaf2,
    bytes32 leaf3,
    bytes32 proof
  ) public {
    bytes32[] memory leaves = new bytes32[](3);
    leaves[0] = leaf1;
    leaves[1] = leaf2;
    leaves[2] = leaf3;
    bytes32[] memory proofs = new bytes32[](1);
    proofs[0] = proof;

    // Proof flag = true
    bytes32 result1 = s_merkleHelper.hashPair(leaves[0], leaves[1]);
    // Proof flag = false
    bytes32 result2 = s_merkleHelper.hashPair(leaves[2], proofs[0]);
    // Proof flag = true
    bytes32 finalResult = s_merkleHelper.hashPair(result1, result2);

    assertEq(s_blobVerifier.merkleRoot(leaves, proofs, 5), finalResult);
  }

  function testMerkleMulti4of4(
    bytes32 leaf1,
    bytes32 leaf2,
    bytes32 leaf3,
    bytes32 leaf4
  ) public {
    bytes32[] memory leaves = new bytes32[](4);
    leaves[0] = leaf1;
    leaves[1] = leaf2;
    leaves[2] = leaf3;
    leaves[3] = leaf4;
    bytes32[] memory proofs = new bytes32[](0);

    // Proof flag = true
    bytes32 result1 = s_merkleHelper.hashPair(leaves[0], leaves[1]);
    // Proof flag = true
    bytes32 result2 = s_merkleHelper.hashPair(leaves[2], leaves[3]);
    // Proof flag = true
    bytes32 finalResult = s_merkleHelper.hashPair(result1, result2);

    assertEq(s_blobVerifier.merkleRoot(leaves, proofs, 7), finalResult);
  }
}
