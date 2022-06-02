import hre from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockERC20, MockOffRamp, OffRampHelper } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  Any2EVMTollMessage,
  encodeExecutionReport,
  executionReportDeepEqual,
  MerkleMultiTree,
} from '../../../test-helpers/ccip/ccip'
import { BigNumber } from '@ethersproject/bignumber'
import { numToBytes32, publicAbi } from '../../../test-helpers/helpers'

const { deployContract } = hre.waffle

let roles: Roles

let RampArtifact: Artifact
let ExecutorArtifact: Artifact

let ramp: MockOffRamp
let executor: OffRampHelper
let token: MockERC20

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Any2EVMTollOffRamp', () => {
  beforeEach(async () => {
    RampArtifact = await hre.artifacts.readArtifact('MockOffRamp')
    ExecutorArtifact = await hre.artifacts.readArtifact('OffRampHelper')

    const adminAddress = await roles.defaultAccount.getAddress()
    const TokenArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )

    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 1 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )

    ramp = <MockOffRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [])
    )
    executor = <OffRampHelper>(
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
      // OffRampExecutorHelper
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

  it('executes a payload of 2 messages', async () => {
    const message1: Any2EVMTollMessage = {
      sourceChainId: BigNumber.from(1),
      sequenceNumber: BigNumber.from(1),
      sender: await roles.oracleNode1.getAddress(),
      receiver: await roles.oracleNode2.getAddress(),
      data: numToBytes32(3),
      tokens: [],
      amounts: [],
      feeToken: token.address,
      feeTokenAmount: 0,
      gasLimit: 0,
    }
    const message2: Any2EVMTollMessage = {
      sourceChainId: BigNumber.from(1),
      sequenceNumber: BigNumber.from(2),
      sender: await roles.oracleNode3.getAddress(),
      receiver: await roles.oracleNode4.getAddress(),
      data: numToBytes32(7),
      tokens: [],
      amounts: [],
      feeToken: token.address,
      feeTokenAmount: 0,
      gasLimit: 0,
    }
    const tree = new MerkleMultiTree([message1, message2])
    const execReport = tree.generateExecutionReport([0, 1])
    const tx = await executor
      .connect(roles.defaultAccount)
      .report(encodeExecutionReport(execReport))
    const receipt = await tx.wait()
    const event1 = ramp.interface.parseLog(receipt.logs[0])

    executionReportDeepEqual(event1.args.report, execReport)
  })
})
