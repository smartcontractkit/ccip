import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockOffRamp, MessageExecutorHelper } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { CCIPMessage, messageDeepEqual } from '../../../test-helpers/ccip'
import { BigNumber } from '@ethersproject/bignumber'
import { numToBytes32 } from '../../../test-helpers/helpers'

interface ExecutableMessage {
  proof: string[]
  message: CCIPMessage
  index: BigNumber
}

function encodeExecutableMessages(messages: ExecutableMessage[]): string {
  return ethers.utils.defaultAbiCoder.encode(
    [
      'tuple(bytes32[] proof, tuple(uint256 sequenceNumber, uint256 sourceChainId, uint256 destinationChainId, address sender, tuple(address receiver, bytes data, address[] tokens, uint256[] amounts, address executor, bytes options) payload) message, uint256 index)[] report',
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
      ])
    )
  })

  it('deploys correctly', async () => {
    expect(await executor.s_offRamp()).to.equal(ramp.address)
  })

  it('executes 2 messages in the same tx', async () => {
    const message1: CCIPMessage = {
      sequenceNumber: BigNumber.from(1),
      sourceChainId: BigNumber.from(1),
      destinationChainId: BigNumber.from(2),
      sender: await roles.oracleNode1.getAddress(),
      payload: {
        receiver: await roles.oracleNode2.getAddress(),
        data: numToBytes32(3),
        tokens: [],
        amounts: [],
        executor: hre.ethers.constants.AddressZero,
        options: numToBytes32(4),
      },
    }
    const message2: CCIPMessage = {
      sequenceNumber: BigNumber.from(2),
      sourceChainId: BigNumber.from(1),
      destinationChainId: BigNumber.from(2),
      sender: await roles.oracleNode3.getAddress(),
      payload: {
        receiver: await roles.oracleNode4.getAddress(),
        data: numToBytes32(7),
        tokens: [],
        amounts: [],
        executor: hre.ethers.constants.AddressZero,
        options: numToBytes32(8),
      },
    }
    const proof1 = [numToBytes32(9)]
    const proof2 = [numToBytes32(10)]
    const index1 = BigNumber.from(0)
    const index2 = BigNumber.from(1)

    const em1: ExecutableMessage = {
      proof: proof1,
      message: message1,
      index: index1,
    }
    const em2: ExecutableMessage = {
      proof: proof2,
      message: message2,
      index: index2,
    }
    const tx = await executor
      .connect(roles.defaultAccount)
      .report(encodeExecutableMessages([em1, em2]))
    const receipt = await tx.wait()
    const event1 = ramp.interface.parseLog(receipt.logs[0])
    const event2 = ramp.interface.parseLog(receipt.logs[1])

    expect(event1.args.proof).to.deep.equal(proof1)
    expect(event1.args.index).to.equal(index1)
    messageDeepEqual(event1.args.message, message1)

    expect(event2.args.proof).to.deep.equal(proof2)
    expect(event2.args.index).to.equal(index2)
    messageDeepEqual(event2.args.message, message2)
  })
})
