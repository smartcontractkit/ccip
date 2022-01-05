import hre, { ethers } from 'hardhat'
import { BigNumber, Contract } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { expect } from 'chai'
import {
  MockERC20,
  LockUnlockPool,
  EOASingleTokenReceiver,
  SingleTokenOnRamp,
  EOASingleTokenSender,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  CCIPMessage,
  encodeReport,
  hashMessage,
  RelayReport,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let chain1OnApp: EOASingleTokenSender
let chain1AFN: MockAFN
let chain1OnRamp: SingleTokenOnRamp
let chain1Token: MockERC20
let chain1Pool: LockUnlockPool
const chain1ID: number = 1

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` and try-catch using typechain artifacts.
let chain2OffRamp: Contract
let chain2AFN: MockAFN
let chain2OffApp: EOASingleTokenReceiver
let chain2Token: MockERC20
let chain2Pool: LockUnlockPool
const chain2ID: number = 2

const sendAmount = BigNumber.from('1000000000000000000')
const maxTimeBetweenAFNSignals = sendAmount
const executionDelay = 0

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Single Token EOA End to End', () => {
  beforeEach(async () => {
    const adminAddress = await roles.defaultAccount.getAddress()

    const MockAFNArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockAFN',
    )
    const TokenArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    const PoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'LockUnlockPool',
    )
    const offRampFactory = await ethers.getContractFactory(
      'SingleTokenOffRampHelper',
    )
    const OnRampSenderArtifact: Artifact = await hre.artifacts.readArtifact(
      'EOASingleTokenSender',
    )
    const OnRampArtifact: Artifact = await hre.artifacts.readArtifact(
      'SingleTokenOnRamp',
    )
    const OffRampReceiverArtifact: Artifact = await hre.artifacts.readArtifact(
      'EOASingleTokenReceiver',
    )

    // Deploy chain2 contracts
    chain2Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 2 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    chain2Pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain2Token.address,
      ])
    )
    chain2AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain2OffRamp = await offRampFactory.connect(roles.defaultAccount).deploy(
      chain1ID,
      chain2ID,
      chain2Token.address,
      chain2Pool.address,
      sendAmount, //bucketRate
      sendAmount, //bucketCapacity
      chain2AFN.address,
      maxTimeBetweenAFNSignals,
      executionDelay,
    )
    await chain2Pool
      .connect(roles.defaultAccount)
      .setOffRamp(chain2OffRamp.address, true)
    await chain2Token
      .connect(roles.defaultAccount)
      .approve(chain2Pool.address, sendAmount)
    await chain2Pool
      .connect(roles.defaultAccount)
      .lockOrBurn(adminAddress, sendAmount)
    chain2OffApp = <EOASingleTokenReceiver>(
      await deployContract(roles.defaultAccount, OffRampReceiverArtifact, [
        chain2OffRamp.address,
      ])
    )

    // Deploy chain1 contracts
    chain1Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 1 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    chain1Pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain1Token.address,
      ])
    )
    chain1AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain1OnRamp = <SingleTokenOnRamp>(
      await deployContract(roles.defaultAccount, OnRampArtifact, [
        chain1ID,
        chain1Token.address,
        chain1Pool.address,
        chain2ID,
        chain2Token.address,
        [],
        true,
        sendAmount, // bucketRate
        sendAmount, // bucketCapacity
        chain1AFN.address,
        maxTimeBetweenAFNSignals,
      ])
    )
    await chain1Pool
      .connect(roles.defaultAccount)
      .setOnRamp(chain1OnRamp.address, true)
    chain1OnApp = <EOASingleTokenSender>(
      await deployContract(roles.defaultAccount, OnRampSenderArtifact, [
        chain1OnRamp.address,
        chain2OffApp.address,
      ])
    )
    await chain1OnRamp.setAllowlist([chain1OnApp.address])
    await chain1Token.transfer(await roles.stranger.getAddress(), sendAmount)
  })

  it('should send tokens from chain1 to chain2 EOAs', async () => {
    // Initial balances
    const chain1StrangerInitialBalance = await chain1Token.balanceOf(
      await roles.stranger.getAddress(),
    )
    const chain2StrangerInitialBalance = await chain2Token.balanceOf(
      await roles.stranger.getAddress(),
    )

    // approve tokens and send message
    await chain1Token
      .connect(roles.stranger)
      .approve(chain1OnApp.address, sendAmount)
    let tx = await chain1OnApp
      .connect(roles.stranger)
      .sendTokens(
        await roles.stranger.getAddress(),
        sendAmount,
        ethers.constants.AddressZero,
      )

    // Parse log
    let receipt = await tx.wait()
    const log = receipt.logs[6]
    const decodedLog = chain1OnRamp.interface.parseLog(log)
    const logArgs = decodedLog.args[0]

    // Send messge to chain2
    const message: CCIPMessage = {
      sequenceNumber: logArgs.sequenceNumber,
      sourceChainId: BigNumber.from(chain1ID),
      destinationChainId: BigNumber.from(chain2ID),
      sender: logArgs.sender,
      payload: {
        receiver: logArgs.payload.receiver,
        data: logArgs.payload.data,
        tokens: logArgs.payload.tokens,
        amounts: logArgs.payload.amounts,
        executor: logArgs.payload.executor,
        options: logArgs.payload.options,
      },
    }
    // DON encodes, reports and executes the message
    let report: RelayReport = {
      merkleRoot: hashMessage(message),
      minSequenceNumber: logArgs.sequenceNumber,
      maxSequenceNumber: logArgs.sequenceNumber,
    }
    await chain2OffRamp
      .connect(roles.defaultAccount)
      .report(encodeReport(report))
    tx = await chain2OffRamp
      .connect(roles.defaultAccount)
      .executeTransaction([], message, 0)
    receipt = await tx.wait()

    const chain1StrangerBalanceAfter = await chain1Token.balanceOf(
      await roles.stranger.getAddress(),
    )
    const chain2StrangerBalanceAfter = await chain2Token.balanceOf(
      await roles.stranger.getAddress(),
    )

    expect(
      chain1StrangerInitialBalance.sub(chain1StrangerBalanceAfter),
    ).to.equal(sendAmount)
    expect(
      chain2StrangerBalanceAfter.sub(chain2StrangerInitialBalance),
    ).to.equal(sendAmount)
  })
})
