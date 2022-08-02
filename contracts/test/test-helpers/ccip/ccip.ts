import { BigNumber, BigNumberish, BytesLike } from 'ethers'
import MerkleTree from 'merkletreejs'
import { ethers } from 'hardhat'
import { expect } from 'chai'
import { stripHexPrefix } from '../helpers'

export interface EVM2AnyTollMessage {
  receiver: string
  data: BytesLike
  tokens: string[]
  amounts: BigNumberish[]
  feeToken: string
  feeTokenAmount: BigNumberish
  gasLimit: BigNumberish
}

export interface EVM2AnyTollEvent {
  sourceChainId: BigNumber
  sequenceNumber: BigNumber
  sender: string
  receiver: string
  data: BytesLike
  tokens: string[]
  amounts: BigNumberish[]
  feeToken: string
  feeTokenAmount: BigNumberish
  gasLimit: BigNumberish
}

export interface EVM2EVMTollMessage {
  sourceChainId: BigNumber
  sequenceNumber: BigNumber
  sender: string
  receiver: string
  data: BytesLike
  tokens: string[]
  amounts: BigNumberish[]
  feeToken: string
  feeTokenAmount: BigNumberish
  gasLimit: BigNumberish
}

export const Any2EVMTollMessageTuple = `tuple(uint256 sourceChainId, uint64 sequenceNumber, address sender, address receiver, bytes data, address[] tokens, uint256[] amounts, address feeToken, uint256 feeTokenAmount, uint256 gasLimit)`

export interface ExecutionReport {
  messages: EVM2EVMTollMessage[]
  proofs: string[]
  proofFlagsBits: BigNumberish
}
export const ExecutionReportTuple = `tuple(${Any2EVMTollMessageTuple}[] messages, bytes32[] proofs, uint256 proofFlagsBits)`

export interface RelayReport {
  merkleRoot: string
  minSequenceNumber: BigNumber
  maxSequenceNumber: BigNumber
}
export const RelayReportTuple = `tuple(bytes32 merkleRoot, uint64 minSequenceNumber, uint64 maxSequenceNumber)`

/**
 * @notice MerkleMultiTree generates a merkle tree using an array CCIPMessage leaves
 * Use this to generate relay and execution reports for specific messages in tests.
 */
export class MerkleMultiTree {
  public tree?: MerkleTree
  public messages: { [hash: string]: EVM2EVMTollMessage } = {}
  public minSequenceNumber?: BigNumber
  public maxSequenceNumber?: BigNumber

  /**
   * @notice Create a new MerkleMultiTree
   * @param rawMessages EVM2EVMTollMessage[] array of messages
   */
  constructor(rawMessages: EVM2EVMTollMessage[]) {
    rawMessages.map((rm) => {
      this.messages[this.hashMessage(rm)] = rm
      if (
        !this.minSequenceNumber ||
        rm.sequenceNumber.lt(this.minSequenceNumber)
      ) {
        this.minSequenceNumber = rm.sequenceNumber
      }
      if (
        !this.maxSequenceNumber ||
        rm.sequenceNumber.gt(this.maxSequenceNumber)
      ) {
        this.maxSequenceNumber = rm.sequenceNumber
      }
    })
    this.tree = new MerkleTree(Object.keys(this.messages), this.hashInternal, {
      sort: true,
    })
  }

  /**
   * Generate a relay report for this merkle tree
   * @returns RelayReport
   */
  public generateRelayReport(): RelayReport {
    const relayReport: RelayReport = {
      merkleRoot: this.bufferToStringAddress(this.tree?.getRoot()!),
      minSequenceNumber: this.minSequenceNumber!,
      maxSequenceNumber: this.maxSequenceNumber!,
    }
    return relayReport
  }

  /**
   * @notice Generate an execution report for specific messages in this merkle tree
   * @param messageIndices indices of the messages to include in the execution report
   * @returns ExecutionReport
   */
  public generateExecutionReport(messageIndices: number[]): ExecutionReport {
    const messageHashes: string[] = messageIndices.map((i) =>
      this.bufferToStringAddress(this.tree?.getLeaf(i)!),
    )
    return this.generateExecutionReportFromHashes(messageHashes)
  }

  /**
   * @notice Generate an execution report for specific messages in this merkle tree
   * @param messageHashes hashes of the messages to include in the execution report
   * @returns ExecutionReport
   */
  public generateExecutionReportFromHashes(
    messageHashes: string[],
  ): ExecutionReport {
    const [proofsBuffer, boolFlags] = this.generateProofs(messageHashes)

    const execReport: ExecutionReport = {
      messages: messageHashes.map((mh) => this.messages[mh]),
      proofs: proofsBuffer.map((p) => this.bufferToStringAddress(p)),
      proofFlagsBits: this.generateBigNumberBitmap(boolFlags),
    }

    return execReport
  }

  /**
   * @notice Get the root hash for this merkle tree
   * @returns Root hash
   */
  public getRoot(): string {
    return this.bufferToStringAddress(this.tree?.getRoot()!)
  }

  private generateProofs(messageHashes: string[]): [Buffer[], boolean[]] {
    const bufferMessageHashes: Buffer[] = messageHashes.map((mh) =>
      this.stringAddressToBuffer(mh),
    )
    const proofs = this.tree?.getMultiProof(bufferMessageHashes)
    const proofFlags = this.tree?.getProofFlags(bufferMessageHashes, proofs!)
    expect(proofFlags!.length).to.be.lte(256)
    return [proofs!, proofFlags!]
  }

  private generateBigNumberBitmap(boolArray: Array<boolean>): BigNumber {
    let bitmap = BigNumber.from(0)
    for (let i = 0; i < boolArray.length; i++) {
      const zeroOrOne: BigNumber = boolArray[i]
        ? BigNumber.from(1)
        : BigNumber.from(0)
      bitmap = bitmap.or(zeroOrOne.shl(i))
    }
    return bitmap
  }

  private hashMessage(message: EVM2EVMTollMessage): string {
    const bytesMessage = ethers.utils.defaultAbiCoder.encode(
      [Any2EVMTollMessageTuple],
      [message],
    )
    return this.hashLeaf(bytesMessage)
  }

  private hashLeaf(value: string): string {
    // Add the leaf domain separator 0x00.
    return ethers.utils.solidityKeccak256(['bytes', 'bytes'], ['0x00', value])
  }

  private hashInternal(value: string): string {
    // Add the internal domain separator 0x01.
    return ethers.utils.solidityKeccak256(['bytes', 'bytes'], ['0x01', value])
  }

  private bufferToStringAddress(buf: Buffer): string {
    return '0x' + buf.toString('hex')
  }

  private stringAddressToBuffer(addr: string): Buffer {
    return Buffer.from(stripHexPrefix(addr), 'hex')
  }
}

/**
 * @notice Encode a RelayReport
 * @param report RelayReport
 * @returns encoded bytes string
 */
export function encodeRelayReport(report: RelayReport): string {
  return ethers.utils.defaultAbiCoder.encode([RelayReportTuple], [report])
}

/**
 * @notice Encode an ExecutionReport
 * @param report ExecutionReport
 * @returns encoded bytes string
 */
export function encodeExecutionReport(report: ExecutionReport): string {
  return ethers.utils.defaultAbiCoder.encode([ExecutionReportTuple], [report])
}

export function executionReportDeepEqual(
  actualReport: any,
  expectedReport: ExecutionReport,
) {
  expect(actualReport?.proofs).to.deep.equal(expectedReport.proofs)
  expect(actualReport?.proofFlagsBits).to.equal(expectedReport.proofFlagsBits)
  for (let i = 0; i < expectedReport.messages.length; i++) {
    const expectedMsg = expectedReport.messages[i]
    messageDeepEqual(actualReport?.messages?.[i], expectedMsg)
  }
}

export function messageDeepEqual(
  actualMessage: any,
  expectedMessage: EVM2EVMTollMessage,
) {
  expect(actualMessage?.sequenceNumber).to.equal(expectedMessage.sequenceNumber)
  expect(actualMessage?.sourceChainId).to.equal(expectedMessage.sourceChainId)
  expect(actualMessage?.sender).to.equal(expectedMessage.sender)
  expect(actualMessage?.receiver).to.equal(expectedMessage.receiver)
  expect(actualMessage?.data).to.equal(expectedMessage.data)
  expect(actualMessage?.tokens.length).to.equal(expectedMessage.tokens.length)
  for (let i = 0; i < expectedMessage.tokens.length; i++) {
    const expectedAmount = expectedMessage.tokens[i].toString()
    expect(actualMessage.tokens[i].toString()).to.equal(expectedAmount)
  }
  expect(actualMessage?.amounts.length).to.equal(expectedMessage.amounts.length)
  for (let i = 0; i < expectedMessage.amounts.length; i++) {
    const expectedAmount = expectedMessage.amounts[i].toString()
    expect(actualMessage.amounts[i].toString()).to.equal(expectedAmount)
  }
  expect(actualMessage?.feeToken).to.equal(expectedMessage.feeToken)
  expect(actualMessage?.feeTokenAmount).to.equal(expectedMessage.feeTokenAmount)
  expect(actualMessage?.gasLimit).to.equal(expectedMessage.gasLimit)
}

export function requestEventArgsEqual(
  actualRequestArgs: any,
  expectedRequestArgs: any,
) {
  expect(actualRequestArgs.message.sequenceNumber).to.equal(
    expectedRequestArgs.sequenceNumber,
  )
  expect(actualRequestArgs.message.sourceChainId).to.equal(
    expectedRequestArgs.sourceChainId,
  )
  expect(actualRequestArgs.message.sender).to.equal(expectedRequestArgs.sender)
  expect(actualRequestArgs.message.receiver).to.equal(
    expectedRequestArgs.receiver,
  )
  expect(actualRequestArgs.message.data).to.equal(expectedRequestArgs.data)
  expect(actualRequestArgs.message.tokens).to.deep.equal(
    expectedRequestArgs.tokens,
  )
  expect(actualRequestArgs.message.amounts.length).to.equal(
    expectedRequestArgs.amounts.length,
  )
  for (let i = 0; i < expectedRequestArgs.amounts.length; i++) {
    const expectedAmount = expectedRequestArgs.amounts[i].toString()
    expect(actualRequestArgs.message.amounts[i].toString()).to.equal(
      expectedAmount,
    )
  }
  expect(actualRequestArgs.message.feeToken).to.equal(
    expectedRequestArgs.feeToken,
  )
  expect(actualRequestArgs.message.feeTokenAmount).to.equal(
    expectedRequestArgs.feeTokenAmount,
  )
  expect(actualRequestArgs.message.gasLimit).to.equal(
    expectedRequestArgs.gasLimit,
  )
}
