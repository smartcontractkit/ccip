import hre, { ethers } from 'hardhat'
import { stringToBytes } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { BigNumber, Contract, ContractReceipt } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  LockUnlockPool,
  SimpleMessageReceiver,
  SingleTokenOnRamp,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  CCIPMessage,
  CCIPMessagePayload,
  encodeReport,
  hashMessage,
  messageDeepEqual,
  RelayReport,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let chain1AFN: MockAFN
let chain1OnRamp: SingleTokenOnRamp
let chain1Token: MockERC20
let chain1Pool: LockUnlockPool
const chain1ID: number = 1

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` using typechain artifacts.
let chain2OffRamp: Contract
let chain2AFN: MockAFN
let chain2Token: MockERC20
let chain2Receiver: SimpleMessageReceiver
let chain2Pool: LockUnlockPool
const chain2ID: number = 2

const sendAmount = BigNumber.from('1000000000000000000')
const maxTimeBetweenAFNSignals = sendAmount
const executionDelay = 0

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Single Token Contract End to End', () => {
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
    const OnRampArtifact: Artifact = await hre.artifacts.readArtifact(
      'SingleTokenOnRamp',
    )
    const SimpleMessageReceiverArtifact: Artifact =
      await hre.artifacts.readArtifact('SimpleMessageReceiver')

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
      sendAmount, // bucketRate
      sendAmount, // bucketCapacity
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
    chain2Receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
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
        [roles.defaultAccount.getAddress()],
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
  })

  it('should send a message and tokens from chain1 to chain2', async () => {
    const messagedata = stringToBytes('Message')
    const options = hre.ethers.constants.HashZero
    const payload: CCIPMessagePayload = {
      receiver: chain2Receiver.address,
      data: messagedata,
      tokens: [chain1Token.address],
      amounts: [sendAmount],
      executor: hre.ethers.constants.AddressZero,
      options: options,
    }

    const initialChain1PoolBalance = await chain1Token.balanceOf(
      chain1Pool.address,
    )
    const initialChain2ReceiverBalance = await chain2Token.balanceOf(
      chain2Receiver.address,
    )
    // approve tokens and send message
    const chain1PoolAddress = await chain1OnRamp.POOL()
    await chain1Token.approve(chain1PoolAddress, sendAmount)
    let tx = await chain1OnRamp.requestCrossChainSend(payload)

    // Check tokens are locked
    await expect(await chain1Token.balanceOf(chain1Pool.address)).to.equal(
      initialChain1PoolBalance.add(sendAmount),
    )

    // DON picks up event and reads
    let receipt: ContractReceipt = await tx.wait()
    let eventArgs = receipt.events?.[3]?.args?.[0]
    const sequenceNumber = eventArgs?.sequenceNumber
    const donPayload: CCIPMessagePayload = {
      receiver: eventArgs?.payload.receiver,
      data: eventArgs?.payload.data,
      tokens: eventArgs?.payload.tokens,
      amounts: eventArgs?.payload.amounts,
      executor: eventArgs?.payload.executor,
      options: eventArgs?.payload.options,
    }
    const donMessage: CCIPMessage = {
      sequenceNumber: sequenceNumber,
      sourceChainId: BigNumber.from(chain1ID),
      destinationChainId: BigNumber.from(chain2ID),
      sender: eventArgs?.sender,
      payload: donPayload,
    }

    // DON encodes, reports and executes the message
    let report: RelayReport = {
      merkleRoot: hashMessage(donMessage),
      minSequenceNumber: sequenceNumber,
      maxSequenceNumber: sequenceNumber,
    }
    await chain2OffRamp
      .connect(roles.defaultAccount)
      .report(encodeReport(report))
    tx = await chain2OffRamp
      .connect(roles.defaultAccount)
      .executeTransaction([], donMessage, 0)
    receipt = await tx.wait()

    // Check that events are emitted and receiver receives the message
    await expect(tx)
      .to.emit(chain2OffRamp, 'CrossChainMessageExecuted')
      .withArgs(donMessage.sequenceNumber)

    await expect(tx).to.emit(chain2Receiver, 'MessageReceived')
    const receivedPayload = await chain2Receiver.s_message()
    messageDeepEqual(receivedPayload, donMessage)

    // Check balance of contract
    const afterChain2ReceiverBalance = await chain2Token.balanceOf(
      chain2Receiver.address,
    )
    expect(afterChain2ReceiverBalance).to.equal(
      initialChain2ReceiverBalance.add(sendAmount),
    )
  })
})
