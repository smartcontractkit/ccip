import hre, { ethers } from 'hardhat'
import { BigNumber, Contract } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { expect } from 'chai'
import {
  MockERC20,
  NativeTokenPool,
  ReceiverDapp,
  SenderDapp,
  MockAFN,
  MockAggregator,
  EVM2AnyTollOnRampRouter,
  EVM2EVMTollOnRamp,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import {
  Any2EVMTollMessage,
  encodeRelayReport,
  MerkleMultiTree,
} from '../../../test-helpers/ccip/ccip'

const { deployContract } = hre.waffle

let roles: Roles
let priceFeed1: MockAggregator

let chain1OnApp: SenderDapp
let chain1AFN: MockAFN
let chain1OnRampRouter: EVM2AnyTollOnRampRouter
let chain1OnRamp: EVM2EVMTollOnRamp
let chain1Token: MockERC20
let chain1Pool: NativeTokenPool
const chain1ID: number = 1

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` and try-catch using typechain artifacts.
let chain2BlobVerifier: Contract
let chain2Router: Contract
let chain2AFN: MockAFN
let chain2OffApp: ReceiverDapp
let chain2Token: MockERC20
let chain2Pool: NativeTokenPool
const chain2ID: number = 2

const sendAmount = BigNumber.from('1000000000000000000')
const maxTimeBetweenAFNSignals = sendAmount
const executionDelay = 0

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('Single Token EOA End to End (through dapp contract)', () => {
  beforeEach(async () => {
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
    const OnRampSenderArtifact: Artifact = await hre.artifacts.readArtifact(
      'SenderDapp',
    )
    const OnRampRouterArtifact: Artifact = await hre.artifacts.readArtifact(
      'EVM2AnyTollOnRampRouter',
    )
    const OnRampArtifact: Artifact = await hre.artifacts.readArtifact(
      'EVM2EVMTollOnRamp',
    )
    const OffRampReceiverArtifact: Artifact = await hre.artifacts.readArtifact(
      'ReceiverDapp',
    )
    const PriceFeedArtifact = await hre.artifacts.readArtifact('MockAggregator')
    priceFeed1 = <MockAggregator>(
      await deployContract(roles.defaultAccount, PriceFeedArtifact)
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
    let bucketConfig = {
      rate: sendAmount,
      capacity: sendAmount,
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
    chain1Token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'Chain 1 LINK Token',
        'LINK',
        adminAddress,
        BigNumber.from('100000000000000000000'),
      ])
    )
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
        2,
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
    chain2OffApp = <ReceiverDapp>(
      await deployContract(roles.defaultAccount, OffRampReceiverArtifact, [
        chain2Router.address,
        chain2Token.address,
      ])
    )

    // Deploy chain1 contracts
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
          maxTokensLength: 2,
          maxDataSize: 10 ** 3,
          relayingFeeJuels: 0,
        },
      ])
    )
    await chain1OnRampRouter
      .connect(roles.defaultAccount)
      .setOnRamp(chain2ID, chain1OnRamp.address)
    await chain1Pool
      .connect(roles.defaultAccount)
      .setOnRamp(chain1OnRamp.address, true)
    chain1OnApp = <SenderDapp>(
      await deployContract(roles.defaultAccount, OnRampSenderArtifact, [
        chain1OnRampRouter.address,
        chain2ID,
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
        [chain1Token.address],
        [sendAmount],
        ethers.constants.AddressZero,
      )

    // Parse log
    let receipt = await tx.wait()
    const log = chain1OnRamp.interface.parseLog(
      receipt.logs[receipt.logs.length - 1],
    )

    // Send message to chain2
    const message: Any2EVMTollMessage = {
      sequenceNumber: log.args.message.sequenceNumber,
      sourceChainId: BigNumber.from(chain1ID),
      sender: log.args.message.sender,
      receiver: chain2OffApp.address,
      data: log.args.message.data,
      tokens: log.args.message.tokens,
      amounts: log.args.message.amounts,
      feeToken: log.args.message.tokens[0],
      feeTokenAmount: 0,
      gasLimit: 0,
    }
    // DON encodes, reports and executes the message
    const tree = new MerkleMultiTree([message])
    await chain2BlobVerifier
      .connect(roles.defaultAccount)
      .report(encodeRelayReport(tree.generateRelayReport()))
    tx = await chain2BlobVerifier
      .connect(roles.defaultAccount)
      .executeTransaction(tree.generateExecutionReport([0]), false)
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
