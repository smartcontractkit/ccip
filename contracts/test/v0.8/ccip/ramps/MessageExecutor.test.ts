import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockOffRamp, MessageExecutorHelper } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { CCIPMessage, messageDeepEqual } from '../../../test-helpers/ccip/ccip'
import { BigNumber } from '@ethersproject/bignumber'
import { numToBytes32, publicAbi } from '../../../test-helpers/helpers'

interface ExecutableMessage {
  path: string[]
  index: BigNumber
  message: CCIPMessage
}

function encodeExecutableMessages(messages: ExecutableMessage[]): string {
  return ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(bytes32[] path, uint256 index, tuple(uint256 sourceChainId, uint64 sequenceNumber, address sender, tuple(address[] tokens, uint256[] amounts, uint256 destinationChainId, address receiver, address executor, bytes data, bytes options) payload) message)[] report',
    ],
    [messages],
  )
}

const { deployContract } = hre.waffle

let roles: Roles

let RampArtifact: Artifact
let ExecutorArtifact: Artifact

let ramp: MockOffRamp
let executor: MessageExecutorHelper

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('MessageExecutor', () => {
  beforeEach(async () => {
    RampArtifact = await hre.artifacts.readArtifact('MockOffRamp')
    ExecutorArtifact = await hre.artifacts.readArtifact('MessageExecutorHelper')

    ramp = <MockOffRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [])
    )
    executor = <MessageExecutorHelper>(
      await deployContract(roles.defaultAccount, ExecutorArtifact, [
        ramp.address,
        false,
      ])
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(executor, [
      'getOffRamp',
      'setNeedFee',
      'getNeedFee',
      // MessageExecutorHelper
      'report',
      'withdrawAccumulatedFees',
      // OCR2Abstract
      'setConfig',
      'latestConfigDetails',
      'latestConfigDigestAndEpoch',
      'transmit',
      // OCR2Base
      'transmitters',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
      // TypeAndVersionInterface
      'typeAndVersion',
    ])
  })

  it('deploys correctly', async () => {
    expect(await executor.getOffRamp()).to.equal(ramp.address)
  })

  it('executes 2 messages in the same tx', async () => {
    const message1: CCIPMessage = {
      sourceChainId: BigNumber.from(1),
      sequenceNumber: BigNumber.from(1),
      sender: await roles.oracleNode1.getAddress(),
      payload: {
        destinationChainId: BigNumber.from(2),
        receiver: await roles.oracleNode2.getAddress(),
        data: numToBytes32(3),
        tokens: [],
        amounts: [],
        executor: hre.ethers.constants.AddressZero,
        options: numToBytes32(4),
      },
    }
    const message2: CCIPMessage = {
      sourceChainId: BigNumber.from(1),
      sequenceNumber: BigNumber.from(2),
      sender: await roles.oracleNode3.getAddress(),
      payload: {
        destinationChainId: BigNumber.from(2),
        receiver: await roles.oracleNode4.getAddress(),
        data: numToBytes32(7),
        tokens: [],
        amounts: [],
        executor: hre.ethers.constants.AddressZero,
        options: numToBytes32(8),
      },
    }
    const path1 = [numToBytes32(9)]
    const path2 = [numToBytes32(10)]
    const index1 = BigNumber.from(0)
    const index2 = BigNumber.from(1)

    const em1: ExecutableMessage = {
      path: path1,
      index: index1,
      message: message1,
    }
    const em2: ExecutableMessage = {
      path: path2,
      index: index2,
      message: message2,
    }
    const tx = await executor
      .connect(roles.defaultAccount)
      .report(encodeExecutableMessages([em1, em2]))
    const receipt = await tx.wait()
    const event1 = ramp.interface.parseLog(receipt.logs[0])
    const event2 = ramp.interface.parseLog(receipt.logs[1])

    expect(event1.args.path).to.deep.equal(path1)
    expect(event1.args.index).to.equal(index1)
    messageDeepEqual(event1.args.message, message1)

    expect(event2.args.path).to.deep.equal(path2)
    expect(event2.args.index).to.equal(index2)
    messageDeepEqual(event2.args.message, message2)
  })
})
