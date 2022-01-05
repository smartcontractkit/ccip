import { BigNumber, BigNumberish, BytesLike } from 'ethers'
import { ethers } from 'hardhat'
import { expect } from 'chai'
export interface RelayReport {
  merkleRoot: string
  minSequenceNumber: BigNumber
  maxSequenceNumber: BigNumber
}

export interface CCIPMessage {
  sequenceNumber: BigNumber
  sourceChainId: BigNumber
  destinationChainId: BigNumber
  sender: string
  payload: CCIPMessagePayload
}

export interface CCIPMessagePayload {
  receiver: string
  data: BytesLike
  tokens: string[]
  amounts: BigNumberish[]
  executor: string
  options: BytesLike
}

export class MerkleTree {
  public parent?: MerkleTree

  /**
   * Left subtree
   */
  public left?: MerkleTree

  /**
   * Right subtree
   */
  public right?: MerkleTree

  /**
   * Hash that is either provide or populated
   */
  public hash?: string

  constructor(hash?: string) {
    this.hash = hash
  }

  public getSiblingHash(hash?: string): string {
    if (hash == this.left?.hash) {
      return this.right?.hash!
    } else if (hash == this.right?.hash) {
      return this.left?.hash!
    } else {
      throw new Error('Hash not found')
    }
  }

  public recursiveProof(proof: string[]): string[] {
    if (this.parent != undefined) {
      proof.push(this.parent?.getSiblingHash(this.hash)!)
      this.parent.recursiveProof(proof)
    }
    return proof
  }

  /**
   * Computes the hash based on the children. If no right child
   * exists, reuse the left child's value
   */
  public computeHash(): string {
    const leftHash = this.left!.hash
    const rightHash = this.right ? this.right.hash : leftHash
    // Add the internal node domain separator.
    return ethers.utils.solidityKeccak256(
      ['bytes', 'bytes32', 'bytes32'],
      ['0x01', leftHash, rightHash],
    )
  }
}

export function generateMerkleTreeFromHashes(hashes: string[]): any {
  // Convert the initial hashes into leaf nodes. We will use these
  // leaf nodes to construuct the Merkle tree from the bottom up by
  // successively combining pairing nodes at each level to construct
  // the parent
  let nodes = hashes.map((p) => new MerkleTree(p))
  let leaves: MerkleTree[] = []

  // Loop until we reach a single node, which will be our Merkle root
  while (nodes.length > 1) {
    const parents = []

    // Successively pair up nodes at each level
    for (let i = 0; i < nodes.length; i += 2) {
      // Create the parent node, which we will add a left, try add
      // a right, then calculate the hash for the node
      const parent = new MerkleTree()
      parents.push(parent)

      // Assign the left, which will always be there
      parent.left = nodes[i]
      nodes[i].parent = parent

      // Assign the right, which won't always be there. However,
      // in JavaScript, an array overflow simply returns undefined
      // which in this context, is the same as a null pointer.
      parent.right = nodes[i + 1]
      nodes[i + 1].parent = parent

      // Finally compute the hash, which will be based on the
      // number of children.
      parent.hash = parent.computeHash()

      // Add to the leaves if we're still on the bottom level
      if (leaves.length < hashes.length) {
        leaves.push(nodes[i], nodes[i + 1])
      }
    }

    // Once all pairs have been made, the parents now become the
    // children and we start all over again
    nodes = parents
  }

  // Return the single node as our root
  return {
    root: nodes[0],
    leaves: leaves,
  }
}

export function encodeReport(report: RelayReport) {
  return ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(bytes32 merkleRoot, uint256 minSequenceNumber, uint256 maxSequenceNumber) report',
    ],
    [report],
  )
}

export function hashMessage(message: CCIPMessage) {
  const bytesMessage = ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(uint256 sequenceNumber, uint256 sourceChainId, uint256 destinationChainId, address sender, tuple(address receiver, bytes data, address[] tokens, uint256[] amounts, address executor, bytes options) payload) message',
    ],
    [message],
  )
  // Add the leaf domain separator 0x00.
  return ethers.utils.solidityKeccak256(
    ['bytes', 'bytes'],
    ['0x00', bytesMessage],
  )
}

export function messageDeepEqual(
  actualMessage: any,
  expectedMessage: CCIPMessage,
) {
  expect(actualMessage?.sequenceNumber).to.equal(expectedMessage.sequenceNumber)
  expect(actualMessage?.sourceChainId).to.equal(expectedMessage.sourceChainId)
  expect(actualMessage?.destinationChainId).to.equal(
    expectedMessage.destinationChainId,
  )
  expect(actualMessage?.sender).to.equal(expectedMessage.sender)
  const actualMessagePayload = actualMessage?.payload
  expect(actualMessagePayload?.receiver).to.equal(
    expectedMessage.payload.receiver,
  )
  expect(actualMessagePayload?.data).to.equal(expectedMessage.payload.data)
  expect(actualMessagePayload.tokens).to.deep.equal(
    expectedMessage.payload.tokens,
  )
  const expectedAmounts = actualMessagePayload.amounts
  expect(actualMessagePayload.amounts.length).to.equal(expectedAmounts.length)
  for (let i = 0; i < expectedAmounts.length; i++) {
    const expectedAmount = expectedAmounts[i].toString()
    expect(actualMessagePayload.amounts[i].toString()).to.equal(expectedAmount)
  }
  expect(actualMessagePayload?.options).to.equal(
    expectedMessage.payload.options,
  )
}

export function requestEventArgsEqual(
  actualRequestArgs: any,
  expectedRequestArgs: any,
) {
  expect(actualRequestArgs?.sequenceNumber).to.equal(
    expectedRequestArgs.sequenceNumber,
  )

  expect(actualRequestArgs?.chainId).to.equal(expectedRequestArgs.chainId)
  expect(actualRequestArgs?.sender).to.equal(expectedRequestArgs.sender)
  expect(actualRequestArgs?.payload.receiver).to.equal(
    expectedRequestArgs.receiver,
  )
  expect(actualRequestArgs?.payload.data).to.equal(expectedRequestArgs.data)
  expect(actualRequestArgs.payload.tokens).to.deep.equal(
    expectedRequestArgs.tokens,
  )
  expect(actualRequestArgs.payload.amounts.length).to.equal(
    expectedRequestArgs.amounts.length,
  )
  for (let i = 0; i < expectedRequestArgs.amounts.length; i++) {
    const expectedAmount = expectedRequestArgs?.amounts[i].toString()
    expect(actualRequestArgs?.payload.amounts[i].toString()).to.equal(
      expectedAmount,
    )
  }
  expect(actualRequestArgs?.payload.options).to.equal(
    expectedRequestArgs?.options,
  )
}
