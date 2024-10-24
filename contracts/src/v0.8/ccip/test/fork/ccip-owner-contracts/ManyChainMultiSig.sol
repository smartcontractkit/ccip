// SPDX-License-Identifier: BUSL-1.1
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/access/Ownable2Step.sol";
import "@openzeppelin/contracts/utils/cryptography/ECDSA.sol";
import "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";

// Should be used as the first 32 bytes of the pre-image of the leaf that holds a
// op. This value is for domain separation of the different values stored in the
// Merkle tree.
bytes32 constant MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP = keccak256("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP");

// Should be used as the first 32 bytes of the pre-image of the leaf that holds the
// root metadata. This value is for domain separation of the different values stored in the
// Merkle tree.
bytes32 constant MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA =
  keccak256("MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA");

/// @notice This is a multi-sig contract that supports signing many transactions (called "ops" in
/// the context of this contract to prevent confusion with transactions on the underlying chain)
/// targeting many chains with a single set of signatures. Authorized ops along with some metadata
/// are stored in a Merkle tree, which is generated offchain. Each op has an associated chain id,
/// ManyChainMultiSig contract address and nonce. The nonce enables us to enforce the
/// (per-ManyChainMultiSig contract instance) ordering of ops.
///
/// At any time, this contract stores at most one Merkle root. In the typical case, all ops
/// in the Merkle tree are expected to be executed before another root is set. Since the Merkle root
/// itself reveals ~ no information about the tree's contents, we take two measures to improve
/// transparency. First, we attach an expiration time to each Merkle root after which it cannot
/// be used any more. Second, we embed metadata in the tree itself that has to be proven/revealed
/// to the contract when a new root is set; the metadata contains the range of nonces (and thus
/// number of ops) in the tree intended for the ManyChainMultiSig contract instance on which the
/// root is being set.
///
/// Once a root is registered, *anyone* is allowed to furnish proofs of op inclusion in the Merkle
/// tree and execute the corresponding op. The contract enforces that ops are executed in the
/// correct order and with the correct arguments. A notable exception to this is the gas limit of
/// the call, which can be freely determined by the executor. We expect (transitive) callees to
/// implement standard behavior of simply reverting if insufficient gas is provided. In particular,
/// this means callees should not have non-reverting gas-dependent branches.
///
/// Note: In the typical case, we expect the time from a root being set to all of the ops
/// therein having been executed to be on the order of a couple of minutes.
contract ManyChainMultiSig is Ownable2Step {
  receive() external payable {}

  uint8 public constant NUM_GROUPS = 32;
  uint8 public constant MAX_NUM_SIGNERS = 200;

  struct Signer {
    address addr;
    uint8 index; // index of signer in s_config.signers
    uint8 group; // 0 <= group < NUM_GROUPS. Each signer can only be in one group.
  }

  // s_signers is used to easily validate the existence of the signer by its address. We still
  // have signers stored in s_config in order to easily deactivate them when a new config is set.
  mapping(address => Signer) s_signers;

  // Signing groups are arranged in a tree. Each group is an interior node and has its own quorum.
  // Signers are the leaves of the tree. A signer/leaf node is successful iff it furnishes a valid
  // signature. A group/interior node is successful iff a quorum of its children are successful.
  // setRoot succeeds only if the root group is successful.
  // Here is an example:
  //
  //                    ┌──────┐
  //                 ┌─►│2-of-3│◄───────┐
  //                 │  └──────┘        │
  //                 │        ▲         │
  //                 │        │         │
  //              ┌──┴───┐ ┌──┴───┐ ┌───┴────┐
  //          ┌──►│1-of-2│ │2-of-2│ │signer A│
  //          │   └──────┘ └──────┘ └────────┘
  //          │       ▲      ▲  ▲
  //          │       │      │  │     ┌──────┐
  //          │       │      │  └─────┤1-of-2│◄─┐
  //          │       │      │        └──────┘  │
  //  ┌───────┴┐ ┌────┴───┐ ┌┴───────┐ ▲        │
  //  │signer B│ │signer C│ │signer D│ │        │
  //  └────────┘ └────────┘ └────────┘ │        │
  //                                   │        │
  //                            ┌──────┴─┐ ┌────┴───┐
  //                            │signer E│ │signer F│
  //                            └────────┘ └────────┘
  //
  // - If signers [A, B] sign, they can set a root.
  // - If signers [B, D, E] sign, they can set a root.
  // - If signers [B, D, E, F] sign, they can set a root. (Either E's or F's signature was
  //   superfluous.)
  // - If signers [B, C, D] sign, they cannot set a root, because the 2-of-2 group on the second
  //   level isn't successful and therefore the root group isn't successful either.
  //
  // To map this tree to a Config, we:
  // - create an entry in signers for each signer (sorted by address in ascending order)
  // - assign the root group to index 0 and have it be its own parent
  // - assign an index to each non-root group, such that each group's parent has a lower index
  //   than the group itself
  // For example, we could transform the above tree structure into:
  // groupQuorums = [2, 1, 2, 1] + [0, 0, ...] (rightpad with 0s to NUM_GROUPS)
  // groupParents = [0, 0, 0, 2] + [0, 0, ...] (rightpad with 0s to NUM_GROUPS)
  // and assuming that address(A) < address(C) < address(E) < address(F) < address(D) < address(B)
  // signers = [
  //    {addr: address(A), index: 0, group: 0}, {addr: address(C), index: 1, group: 1},
  //    {addr: address(E), index: 2, group: 3}, {addr: address(F), index: 3, group: 3},
  //    {addr: address(D), index: 4, group: 2}, {addr: address(B), index: 5, group: 1},
  //  ]
  struct Config {
    Signer[] signers;
    // groupQuorums[i] stores the quorum for the i-th signer group. Any group with
    // groupQuorums[i] = 0 is considered disabled. The i-th group is successful if
    // it is enabled and at least groupQuorums[i] of its children are successful.
    uint8[NUM_GROUPS] groupQuorums;
    // groupParents[i] stores the parent group of the i-th signer group. We ensure that the
    // groups form a tree structure (where the root/0-th signer group points to itself as
    // parent) by enforcing
    // - (i != 0) implies (groupParents[i] < i)
    // - groupParents[0] == 0
    uint8[NUM_GROUPS] groupParents;
  }

  Config s_config;

  // Remember signedHashes that this contract has seen. Each signedHash can only be set once.
  mapping(bytes32 => bool) s_seenSignedHashes;

  // MerkleRoots are a bit tricky since they reveal almost no information about the contents of
  // the tree they authenticate. To mitigate this, we enforce that this contract can only execute
  // ops from a single root at any given point in time. We further associate an expiry
  // with each root to ensure that messages are executed in a timely manner. setRoot and various
  // execute calls are expected to happen in quick succession. We put the expiring root and
  // opCount in same struct in order to reduce gas costs of reading and writing.
  struct ExpiringRootAndOpCount {
    bytes32 root;
    // We prefer using block.timestamp instead of block.number, as a single
    // root may target many chains. We assume that block.timestamp can
    // be manipulated by block producers but only within relatively tight
    // bounds (a few minutes at most).
    uint32 validUntil;
    // each ManyChainMultiSig instance has it own independent opCount.
    uint40 opCount;
  }

  ExpiringRootAndOpCount s_expiringRootAndOpCount;

  /// @notice Each root also authenticates metadata about itself (stored as one of the leaves)
  /// which must be revealed when the root is set.
  ///
  /// @dev We need to be careful that abi.encode(MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA, RootMetadata)
  /// is greater than 64 bytes to prevent collisions with internal nodes in the Merkle tree. See
  /// openzeppelin-contracts/contracts/utils/cryptography/MerkleProof.sol:15 for details.
  struct RootMetadata {
    // chainId and multiSig uniquely identify a ManyChainMultiSig contract instance that the
    // root is destined for.
    // uint256 since it is unclear if we can represent chainId as uint64. There is a proposal (
    // https://ethereum-magicians.org/t/eip-2294-explicit-bound-to-chain-id/11090) to
    // bound chainid to 64 bits, but it is still unresolved.
    uint256 chainId;
    address multiSig;
    // opCount before adding this root
    uint40 preOpCount;
    // opCount after executing all ops in this root
    uint40 postOpCount;
    // override whatever root was already stored in this contract even if some of its
    // ops weren't executed.
    // Important: it is strongly recommended that offchain code set this to false by default.
    // Be careful setting this to true as it may break assumptions about what transactions from
    // the previous root have already been executed.
    bool overridePreviousRoot;
  }

  RootMetadata s_rootMetadata;

  /// @notice An ECDSA signature.
  struct Signature {
    uint8 v;
    bytes32 r;
    bytes32 s;
  }

  /// @notice setRoot Sets a new expiring root.
  ///
  /// @param root is the new expiring root.
  /// @param validUntil is the time by which root is valid
  /// @param metadata is the authenticated metadata about the root, which is stored as one of
  /// the leaves.
  /// @param metadataProof is the MerkleProof of inclusion of the metadata in the Merkle tree.
  /// @param signatures the ECDSA signatures on (root, validUntil).
  ///
  /// @dev the message (root, validUntil) should be signed by a sufficient set of signers.
  /// This signature authenticates also the metadata.
  ///
  /// @dev this method can be executed by anyone who has the root and valid signatures.
  /// as we validate the correctness of signatures, this imposes no risk.
  function setRoot(
    bytes32 root,
    uint32 validUntil,
    RootMetadata calldata metadata,
    bytes32[] calldata metadataProof,
    Signature[] calldata signatures
  ) external {
    bytes32 signedHash = ECDSA.toEthSignedMessageHash(keccak256(abi.encode(root, validUntil)));

    // Each (root, validUntil) tuple can only bet set once. For example, this prevents a
    // scenario where there are two signed roots with overridePreviousRoot = true and
    // an adversary keeps alternatively calling setRoot(root1), setRoot(root2),
    // setRoot(root1), ...
    if (s_seenSignedHashes[signedHash]) {
      revert SignedHashAlreadySeen();
    }

    // verify ECDSA signatures on (root, validUntil) and ensure that the root group is successful
    {
      // verify sigs and count number of signers in each group
      Signer memory signer;
      address prevAddress = address(0x0);
      uint8[NUM_GROUPS] memory groupVoteCounts; // number of votes per group
      for (uint256 i = 0; i < signatures.length; i++) {
        Signature calldata sig = signatures[i];
        address signerAddress = ECDSA.recover(signedHash, sig.v, sig.r, sig.s);
        // the off-chain system is required to sort the signatures by the
        // signer address in an increasing order
        if (prevAddress >= signerAddress) {
          revert SignersAddressesMustBeStrictlyIncreasing();
        }
        prevAddress = signerAddress;

        signer = s_signers[signerAddress];
        if (signer.addr != signerAddress) {
          revert InvalidSigner();
        }
        uint8 group = signer.group;
        while (true) {
          groupVoteCounts[group]++;
          if (groupVoteCounts[group] != s_config.groupQuorums[group]) {
            // bail out unless we just hit the quorum. we only hit each quorum once,
            // so we never move on to the parent of a group more than once.
            break;
          }
          if (group == 0) {
            // reached root
            break;
          }

          group = s_config.groupParents[group];
        }
      }
      // the group at the root of the tree (with index 0) determines whether the vote passed,
      // we cannot proceed if it isn't configured with a valid (non-zero) quorum
      if (s_config.groupQuorums[0] == 0) {
        revert MissingConfig();
      }
      // did the root group reach its quorum?
      if (groupVoteCounts[0] < s_config.groupQuorums[0]) {
        revert InsufficientSigners();
      }
    }

    if (validUntil < block.timestamp) {
      revert ValidUntilHasAlreadyPassed();
    }

    {
      // verify metadataProof
      bytes32 hashedLeaf = keccak256(abi.encode(MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_METADATA, metadata));
      if (!MerkleProof.verify(metadataProof, root, hashedLeaf)) {
        revert ProofCannotBeVerified();
      }
    }

    if (block.chainid != metadata.chainId) {
      revert WrongChainId();
    }

    if (address(this) != metadata.multiSig) {
      revert WrongMultiSig();
    }

    uint40 opCount = s_expiringRootAndOpCount.opCount;

    // don't allow a new root to be set if there are still outstanding ops that have not been
    // executed, unless overridePreviousRoot is set
    if (opCount != s_rootMetadata.postOpCount && !metadata.overridePreviousRoot) {
      revert PendingOps();
    }

    // the signers are responsible for tracking opCount offchain and ensuring that
    // preOpCount equals to opCount
    if (opCount != metadata.preOpCount) {
      revert WrongPreOpCount();
    }

    if (metadata.preOpCount > metadata.postOpCount) {
      revert WrongPostOpCount();
    }

    // done with validation, persist in in contract state
    s_seenSignedHashes[signedHash] = true;
    s_expiringRootAndOpCount =
      ExpiringRootAndOpCount({root: root, validUntil: validUntil, opCount: metadata.preOpCount});
    s_rootMetadata = metadata;
    emit NewRoot(root, validUntil, metadata);
  }

  /// @notice an op to be executed by the ManyChainMultiSig contract
  ///
  /// @dev We need to be careful that abi.encode(LEAF_OP_DOMAIN_SEPARATOR, RootMetadata)
  /// is greater than 64 bytes to prevent collisions with internal nodes in the Merkle tree. See
  /// openzeppelin-contracts/contracts/utils/cryptography/MerkleProof.sol:15 for details.
  struct Op {
    uint256 chainId;
    address multiSig;
    uint40 nonce;
    address to;
    uint256 value;
    bytes data;
  }

  /// @notice Execute the received op after verifying the proof of its inclusion in the
  /// current Merkle tree. The op should be the next op according to the order
  /// enforced by the merkle tree whose root is stored in s_expiringRootAndOpCount, i.e., the
  /// nonce of the op should be equal to s_expiringRootAndOpCount.opCount.
  ///
  /// @param op is Op to be executed
  /// @param proof is the MerkleProof for the op's inclusion in the MerkleTree which its
  /// root is the s_expiringRootAndOpCount.root.
  ///
  /// @dev ANYONE can call this function! That's intentional. Callers can only execute verified,
  /// ordered ops in the Merkle tree.
  ///
  /// @dev we perform a raw call to each target. Raw calls to targets that don't have associated
  /// contract code will always succeed regardless of data.
  ///
  /// @dev the gas limit of the call can be freely determined by the caller of this function.
  /// We expect callees to revert if they run out of gas.
  function execute(Op calldata op, bytes32[] calldata proof) external payable {
    ExpiringRootAndOpCount memory currentExpiringRootAndOpCount = s_expiringRootAndOpCount;

    if (s_rootMetadata.postOpCount <= currentExpiringRootAndOpCount.opCount) {
      revert PostOpCountReached();
    }

    if (op.chainId != block.chainid) {
      revert WrongChainId();
    }

    if (op.multiSig != address(this)) {
      revert WrongMultiSig();
    }

    if (block.timestamp > currentExpiringRootAndOpCount.validUntil) {
      revert RootExpired();
    }

    if (op.nonce != currentExpiringRootAndOpCount.opCount) {
      revert WrongNonce();
    }

    // verify that the op exists in the merkle tree
    bytes32 hashedLeaf = keccak256(abi.encode(MANY_CHAIN_MULTI_SIG_DOMAIN_SEPARATOR_OP, op));
    if (!MerkleProof.verify(proof, currentExpiringRootAndOpCount.root, hashedLeaf)) {
      revert ProofCannotBeVerified();
    }

    // increase the counter *before* execution to prevent reentrancy issues
    s_expiringRootAndOpCount.opCount = currentExpiringRootAndOpCount.opCount + 1;

    _execute(op.to, op.value, op.data);
    emit OpExecuted(op.nonce, op.to, op.data, op.value);
  }

  /// @notice sets a new s_config. If clearRoot is true, then it also invalidates
  /// s_expiringRootAndOpCount.root.
  ///
  /// @param signerAddresses holds the addresses of the active signers. The addresses must be in
  /// ascending order.
  /// @param signerGroups maps each signer to its group
  /// @param groupQuorums holds the required number of valid signatures in each group.
  /// A group i is called successful group if at least groupQuorum[i] distinct signers provide a
  /// valid signature.
  /// @param groupParents holds each group's parent. The groups must be arranged in a tree s.t.
  /// group 0 is the root of the tree and the i-th group's parent has index j less than i.
  /// Iff setRoot is called with a set of signatures that causes the root group to be successful,
  /// setRoot allows a root to be set.
  /// @param clearRoot, if set to true, invalidates the current root. This option is needed to
  /// invalidate the current root, so to prevent further ops from being executed. This
  /// might be used when the current root was signed under a loser group configuration or when
  /// some previous signers aren't trusted any more.
  function setConfig(
    address[] calldata signerAddresses,
    uint8[] calldata signerGroups,
    uint8[NUM_GROUPS] calldata groupQuorums,
    uint8[NUM_GROUPS] calldata groupParents,
    bool clearRoot
  ) external onlyOwner {
    if (signerAddresses.length == 0 || signerAddresses.length > MAX_NUM_SIGNERS) {
      revert OutOfBoundsNumOfSigners();
    }

    if (signerAddresses.length != signerGroups.length) {
      revert SignerGroupsLengthMismatch();
    }

    {
      // validate group structure
      // counts the number of children of each group
      uint8[NUM_GROUPS] memory groupChildrenCounts;
      // first, we count the signers as children
      for (uint256 i = 0; i < signerGroups.length; i++) {
        if (signerGroups[i] >= NUM_GROUPS) {
          revert OutOfBoundsGroup();
        }
        groupChildrenCounts[signerGroups[i]]++;
      }
      // second, we iterate backwards so as to check each group and propagate counts from
      // child group to parent groups up the tree to the root
      for (uint256 j = 0; j < NUM_GROUPS; j++) {
        uint256 i = NUM_GROUPS - 1 - j;
        // ensure we have a well-formed group tree. the root should have itself as parent
        if ((i != 0 && groupParents[i] >= i) || (i == 0 && groupParents[i] != 0)) {
          revert GroupTreeNotWellFormed();
        }
        bool disabled = groupQuorums[i] == 0;
        if (disabled) {
          // a disabled group shouldn't have any children
          if (0 < groupChildrenCounts[i]) {
            revert SignerInDisabledGroup();
          }
        } else {
          // ensure that the group quorum can be met
          if (groupChildrenCounts[i] < groupQuorums[i]) {
            revert OutOfBoundsGroupQuorum();
          }
          groupChildrenCounts[groupParents[i]]++;
          // the above line clobbers groupChildrenCounts[0] in last iteration, don't use it after the loop ends
        }
      }
    }

    Signer[] memory oldSigners = s_config.signers;
    // remove any old signer addresses
    for (uint256 i = 0; i < oldSigners.length; i++) {
      address oldSignerAddress = oldSigners[i].addr;
      delete s_signers[oldSignerAddress];
      s_config.signers.pop();
    }

    // we cannot just write s_config = Config({...}) because solc doesn't support that
    assert(s_config.signers.length == 0);
    s_config.groupQuorums = groupQuorums;
    s_config.groupParents = groupParents;

    // add new signers' addresses, we require that the signers' list be a strictly monotone
    // increasing sequence
    address prevSigner = address(0x0);
    for (uint256 i = 0; i < signerAddresses.length; i++) {
      if (prevSigner >= signerAddresses[i]) {
        revert SignersAddressesMustBeStrictlyIncreasing();
      }
      Signer memory signer = Signer({addr: signerAddresses[i], index: uint8(i), group: signerGroups[i]});
      s_signers[signerAddresses[i]] = signer;
      s_config.signers.push(signer);
      prevSigner = signerAddresses[i];
    }

    if (clearRoot) {
      // clearRoot is equivalent to overriding with a completely empty root
      uint40 opCount = s_expiringRootAndOpCount.opCount;
      s_expiringRootAndOpCount = ExpiringRootAndOpCount({root: 0, validUntil: 0, opCount: opCount});
      s_rootMetadata = RootMetadata({
        chainId: block.chainid,
        multiSig: address(this),
        preOpCount: opCount,
        postOpCount: opCount,
        overridePreviousRoot: true
      });
    }
    emit ConfigSet(s_config, clearRoot);
  }

  /// @notice Execute an op's call. Performs a raw call that always succeeds if the
  /// target isn't a contract.
  function _execute(address target, uint256 value, bytes calldata data) internal virtual {
    (bool success, bytes memory ret) = target.call{value: value}(data);
    if (!success) {
      revert CallReverted(ret);
    }
  }

  /*
   * Getters
   */

  function getConfig() public view returns (Config memory) {
    return s_config;
  }

  function getOpCount() public view returns (uint40) {
    return s_expiringRootAndOpCount.opCount;
  }

  function getRoot() public view returns (bytes32 root, uint32 validUntil) {
    ExpiringRootAndOpCount memory currentRootAndOpCount = s_expiringRootAndOpCount;
    return (currentRootAndOpCount.root, currentRootAndOpCount.validUntil);
  }

  function getRootMetadata() public view returns (RootMetadata memory) {
    return s_rootMetadata;
  }

  /*
   * Events and Errors
   */

  /// @notice Emitted when a new root is set.
  event NewRoot(bytes32 indexed root, uint32 validUntil, RootMetadata metadata);

  /// @notice Emitted when a new config is set.
  event ConfigSet(Config config, bool isRootCleared);

  /// @notice Emitted when an op gets successfully executed.
  event OpExecuted(uint40 indexed nonce, address to, bytes data, uint256 value);

  /// @notice Thrown when number of signers is 0 or greater than MAX_NUM_SIGNERS.
  error OutOfBoundsNumOfSigners();

  /// @notice Thrown when signerAddresses and signerGroups have different lengths.
  error SignerGroupsLengthMismatch();

  /// @notice Thrown when number of some signer's group is greater than (NUM_GROUPS-1).
  error OutOfBoundsGroup();

  /// @notice Thrown when the group tree isn't well-formed.
  error GroupTreeNotWellFormed();

  /// @notice Thrown when the quorum of some group is larger than the number of signers in it.
  error OutOfBoundsGroupQuorum();

  /// @notice Thrown when a disabled group contains a signer.
  error SignerInDisabledGroup();

  /// @notice Thrown when the signers' addresses are not a strictly increasing monotone sequence.
  /// Prevents signers from including more than one signature.
  error SignersAddressesMustBeStrictlyIncreasing();

  /// @notice Thrown when the signature corresponds to invalid signer.
  error InvalidSigner();

  /// @notice Thrown when there is no sufficient set of valid signatures provided to make the
  /// root group successful.
  error InsufficientSigners();

  /// @notice Thrown when attempt to set metadata or execute op for another chain.
  error WrongChainId();

  /// @notice Thrown when the multiSig address in metadata or op is
  /// incompatible with the address of this contract.
  error WrongMultiSig();

  /// @notice Thrown when the preOpCount <= postOpCount invariant is violated.
  error WrongPostOpCount();

  /// @notice Thrown when attempting to set a new root while there are still pending ops
  /// from the previous root without explicitly overriding it.
  error PendingOps();

  /// @notice Thrown when preOpCount in metadata is incompatible with the current opCount.
  error WrongPreOpCount();

  /// @notice Thrown when the provided merkle proof cannot be verified.
  error ProofCannotBeVerified();

  /// @notice Thrown when attempt to execute an op after
  /// s_expiringRootAndOpCount.validUntil has passed.
  error RootExpired();

  /// @notice Thrown when attempt to bypass the enforced ops' order in the merkle tree or
  /// re-execute an op.
  error WrongNonce();

  /// @notice Thrown when attempting to execute an op even though opCount equals
  /// metadata.postOpCount.
  error PostOpCountReached();

  /// @notice Thrown when the underlying call in _execute() reverts.
  error CallReverted(bytes error);

  /// @notice Thrown when attempt to set past validUntil for the root.
  error ValidUntilHasAlreadyPassed();

  /// @notice Thrown when setRoot() is called before setting a config.
  error MissingConfig();

  /// @notice Thrown when attempt to set the same (root, validUntil) in setRoot().
  error SignedHashAlreadySeen();
}
