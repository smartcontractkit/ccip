import hre, { ethers } from 'hardhat'
import { stringToBytes } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { BigNumber, Contract, ContractReceipt } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  NativeTokenPool,
  SimpleMessageReceiver,
  MockAFN,
  MockAggregator,
  EVM2AnyTollOnRampRouter,
  EVM2EVMTollOnRamp,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  Any2EVMTollMessage,
  encodeRelayReport,
  EVM2AnyTollMessage,
  MerkleMultiTree,
  messageDeepEqual,
} from '../../../test-helpers/ccip/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let chain1AFN: MockAFN
let chain1OnRampRouter: EVM2AnyTollOnRampRouter
let chain1OnRamp: EVM2EVMTollOnRamp
let chain1Token: MockERC20
let chain1Pool: NativeTokenPool
const chain1ID: number = 1

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` using typechain artifacts.
let chain2BlobVerifier: Contract
let chain2Router: Contract
let chain2AFN: MockAFN
let chain2Token: MockERC20
let chain2Receiver: SimpleMessageReceiver
let chain2Pool: NativeTokenPool
let priceFeed1: MockAggregator
const chain2ID: number = 2

const sendAmount = BigNumber.from('1000000000000000000')
const maxTimeBetweenAFNSignals = sendAmount
const executionDelay = 0
let bucketRate: BigNumber
let bucketCapactiy: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Contract End to End', () => {
  beforeEach(async () => {
    bucketRate = BigNumber.from('1000000000000000000')
    bucketCapactiy = BigNumber.from('10000000000000000000')
    const maxTokensLength: number = 10
    const maxDataSize: number = 10 ** 3 // 1kb
    const adminAddress = await roles.defaultAccount.getAddress()

    const MockAFNArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockAFN',
    )
    const TokenArtifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    const PoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'NativeTokenPool',
    )
    const blobVerifierHelperFactory = await ethers.getContractFactory(
      'BlobVerifierHelper',
    )
    const routerFactory = await hre.ethers.getContractFactory(
      'Any2EVMTollOffRampRouter',
    )
    const PriceFeedFactory: Artifact = await hre.artifacts.readArtifact(
      'MockAggregator',
    )
    const OnRampRouterArtifact: Artifact = await hre.artifacts.readArtifact(
      'EVM2AnyTollOnRampRouter',
    )
    const OnRampArtifact: Artifact = await hre.artifacts.readArtifact(
      'EVM2EVMTollOnRamp',
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
    let bucketConfig = {
      rate: bucketRate,
      capacity: bucketCapactiy,
    }
    chain2Pool = <NativeTokenPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain2Token.address,
        bucketConfig,
        bucketConfig,
      ])
    )
    chain2AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    priceFeed1 = <MockAggregator>(
      await deployContract(roles.defaultAccount, PriceFeedFactory)
    )

    // Deploy chain1 token
    chain1Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 1 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
    // Chain 2 OffRamp
    chain2BlobVerifier = await blobVerifierHelperFactory
      .connect(roles.defaultAccount)
      .deploy(
        chain1ID,
        chain2ID,
        [chain1Token.address],
        [chain2Pool.address],
        [priceFeed1.address],
        chain2AFN.address,
        maxTimeBetweenAFNSignals,
        executionDelay,
        maxTokensLength,
      )
    chain2Router = await routerFactory
      .connect(roles.defaultAccount)
      .deploy([chain2BlobVerifier.address])
    await chain2BlobVerifier.setRouter(chain2Router.address)
    await chain2Pool
      .connect(roles.defaultAccount)
      .setOffRamp(chain2BlobVerifier.address, true)
    await chain2Token
      .connect(roles.defaultAccount)
      .transfer(chain2Pool.address, sendAmount)
    chain2Receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
    )

    // Chain 1 contracts
    chain1Pool = <NativeTokenPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        chain1Token.address,
        bucketConfig,
        bucketConfig,
      ])
    )
    chain1AFN = <MockAFN>(
      await deployContract(roles.defaultAccount, MockAFNArtifact)
    )
    chain1OnRampRouter = <EVM2AnyTollOnRampRouter>(
      await deployContract(roles.defaultAccount, OnRampRouterArtifact)
    )
    chain1OnRamp = <EVM2EVMTollOnRamp>(
      await deployContract(roles.defaultAccount, OnRampArtifact, [
        chain1ID,
        chain2ID,
        [chain1Token.address],
        [chain1Pool.address],
        [priceFeed1.address],
        [],
        chain1AFN.address,
        maxTimeBetweenAFNSignals,
        {
          router: chain1OnRampRouter.address,
          maxTokensLength: maxTokensLength,
          maxDataSize: maxDataSize,
          relayingFeeJuels: 0,
        },
      ])
    )
    await chain1OnRampRouter.setOnRamp(chain2ID, chain1OnRamp.address)
    await chain1Pool
      .connect(roles.defaultAccount)
      .setOnRamp(chain1OnRamp.address, true)
  })

  it('should send a message and tokens from chain1 to chain2', async () => {
    const messagedata = stringToBytes('Message')
    const evmToAnyTollMessage: EVM2AnyTollMessage = {
      receiver: chain2Receiver.address,
      data: messagedata,
      tokens: [chain1Token.address],
      amounts: [sendAmount],
      feeToken: chain1Token.address,
      feeTokenAmount: 0,
      gasLimit: 0,
    }

    const initialChain1PoolBalance = await chain1Token.balanceOf(
      chain1Pool.address,
    )
    const initialChain2ReceiverBalance = await chain2Token.balanceOf(
      chain2Receiver.address,
    )
    // approve tokens and send message
    await chain1Token
      .connect(roles.defaultAccount)
      .approve(chain1OnRampRouter.address, sendAmount)
    let tx = await chain1OnRampRouter
      .connect(roles.defaultAccount)
      .ccipSend(BigNumber.from(chain2ID), evmToAnyTollMessage)

    // Check tokens are locked
    await expect(await chain1Token.balanceOf(chain1Pool.address)).to.equal(
      initialChain1PoolBalance.add(sendAmount),
    )

    // DON picks up event and reads
    let receipt: ContractReceipt = await tx.wait()
    const log = chain1OnRamp.interface.parseLog(
      receipt.logs[receipt.logs.length - 1],
    )
    const donMessage: Any2EVMTollMessage = {
      sequenceNumber: log.args.message.sequenceNumber,
      sourceChainId: BigNumber.from(chain1ID),
      sender: log.args.message.sender,
      receiver: log.args.message.receiver,
      data: log.args.message.data,
      tokens: log.args.message.tokens,
      amounts: log.args.message.amounts,
      feeToken: log.args.message.feeToken,
      feeTokenAmount: log.args.message.feeTokenAmount,
      gasLimit: log.args.message.gasLimit,
    }

    // DON encodes, reports and executes the message
    const tree = new MerkleMultiTree([donMessage])
    await chain2BlobVerifier
      .connect(roles.defaultAccount)
      .report(encodeRelayReport(tree.generateRelayReport()))
    tx = await chain2BlobVerifier
      .connect(roles.defaultAccount)
      .executeTransaction(tree.generateExecutionReport([0]), false)
    receipt = await tx.wait()

    // Check that events are emitted and receiver receives the message
    await expect(tx)
      .to.emit(chain2BlobVerifier, 'CrossChainMessageExecuted')
      .withArgs(donMessage.sequenceNumber)

    await expect(tx).to.emit(chain2Receiver, 'MessageReceived')
    const receivedPayload = await chain2Receiver.getMessage()
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
