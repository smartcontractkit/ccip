import hre from 'hardhat'
import {
  expectGasWithinDeviation,
  publicAbi,
  stringToBytes,
} from '../../../test-helpers/helpers'
import { BigNumber, constants, ContractTransaction } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  NativeTokenPool,
  OnRampHelper,
  MockAFN,
  MockAggregator,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  CCIPMessagePayload,
  requestEventArgsEqual,
} from '../../../test-helpers/ccip/ccip'
import { Signer } from 'ethers'
import { GAS } from '../../../test-helpers/ccip/gas-measurements'

const { deployContract } = hre.waffle

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let PriceFeedArtifact: Artifact
let RampArtifact: Artifact

let roles: Roles

let afn: MockAFN
let ramp: OnRampHelper
const numberOfTokensPoolsAndFeeds = 3
let tokens: Array<MockERC20>
let pools: Array<NativeTokenPool>
let priceFeed: MockAggregator
let priceFeedLatestAnswer: number = 100
const sourceChainId: number = 123
const destinationChainIds: Array<BigNumber> = [
  BigNumber.from(9),
  BigNumber.from(8),
  BigNumber.from(7),
]
const maxTokensLength: number = 10
const maxDataSize: number = 10 ** 3 // 1kb
const relayingFeeJuels: number = 1
let bucketRate: BigNumber
let bucketCapactiy: BigNumber
let maxTimeWithoutAFNSignal: BigNumber
let mintAmount: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('OnRamp', () => {
  beforeEach(async () => {
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    mintAmount = BigNumber.from('1000000000000000000')

    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('NativeTokenPool')
    PriceFeedArtifact = await hre.artifacts.readArtifact('MockAggregator')
    RampArtifact = await hre.artifacts.readArtifact('OnRampHelper')

    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    maxTimeWithoutAFNSignal = BigNumber.from(60).mul(60) // 1 hour

    tokens = new Array<MockERC20>()
    pools = new Array<NativeTokenPool>()
    let bucketConfig = {
      rate: bucketRate,
      capacity: bucketCapactiy,
    }
    for (let i = 0; i < numberOfTokensPoolsAndFeeds; i++) {
      tokens.push(
        <MockERC20>(
          await deployContract(roles.defaultAccount, TokenArtifact, [
            'TOKEN',
            'TOKEN',
            await roles.defaultAccount.getAddress(),
            mintAmount,
          ])
        ),
      )

      pools.push(
        <NativeTokenPool>(
          await deployContract(roles.defaultAccount, PoolArtifact, [
            tokens[i].address,
            bucketConfig,
            bucketConfig,
          ])
        ),
      )
    }

    priceFeed = <MockAggregator>(
      await deployContract(roles.defaultAccount, PriceFeedArtifact)
    )
    priceFeed
      .connect(roles.defaultAccount)
      .setLatestAnswer(priceFeedLatestAnswer)

    ramp = <OnRampHelper>(
      await deployContract(roles.defaultAccount, RampArtifact, [
        sourceChainId,
        destinationChainIds,
        tokens.map((t) => t.address),
        pools.map((p) => p.address),
        [priceFeed.address, constants.AddressZero, constants.AddressZero],
        [await roles.defaultAccount.getAddress()],
        afn.address,
        maxTimeWithoutAFNSignal,
        {
          router: hre.ethers.constants.AddressZero,
          maxTokensLength: maxTokensLength,
          maxDataSize: maxDataSize,
          relayingFeeJuels: relayingFeeJuels,
        },
      ])
    )

    for (let i = 0; i < numberOfTokensPoolsAndFeeds; i++) {
      pools[i].connect(roles.defaultAccount).setOnRamp(ramp.address, true)
    }
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(ramp, [
      // OnRampHelper
      'publicCalculateFee',
      // OnRamp
      'requestCrossChainSend',
      'withdrawAccumulatedFees',
      'CHAIN_ID',
      'setAllowlistEnabled',
      'getAllowlistEnabled',
      'setAllowlist',
      'getAllowlist',
      'getSequenceNumberOfDestinationChain',
      'getDestinationChains',
      'getConfig',
      'setConfig',
      // PriceFeedRegistry
      'addFeed',
      'removeFeed',
      'getFeed',
      'getFeedTokens',
      // TokenPoolRegistry
      'addPool',
      'removePool',
      'getPool',
      'isPool',
      'getPoolTokens',
      // HealthChecker
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
      'isHealthy',
      // TypeAndVersionInterface
      'typeAndVersion',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
      // Pausable
      'paused',
      'pause',
      'unpause',
    ])
  })

  describe('#constructor', () => {
    it('should deploy correctly', async () => {
      // Ramp
      const owner = await roles.defaultAccount.getAddress()
      expect(await ramp.CHAIN_ID()).to.equal(sourceChainId)
      expect(await ramp.owner()).to.equal(owner)
      expect(await ramp.getAllowlistEnabled()).to.be.true
      expect(await ramp.getAllowlist()).to.deep.equal([
        await roles.defaultAccount.getAddress(),
      ])
      expect(await ramp.getAFN()).to.equal(afn.address)
      expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        maxTimeWithoutAFNSignal,
      )
      const config = await ramp.getConfig()
      expect(config.maxDataSize).to.equal(maxDataSize)
      expect(config.maxTokensLength).to.equal(maxTokensLength)
      expect(config.relayingFeeJuels).to.equal(relayingFeeJuels)

      // Tokens, Pools and Price Feeds
      for (let i = 0; i < numberOfTokensPoolsAndFeeds; i++) {
        const token = tokens[i]
        const configuredPool = await ramp.getPool(token.address)
        expect(configuredPool).to.equal(pools[i].address)
        const configuredPriceFeed = await ramp.getFeed(tokens[i].address)
        if (i == 0) {
          expect(configuredPriceFeed).to.equal(priceFeed.address)
        } else {
          expect(configuredPriceFeed).to.equal(constants.AddressZero)
        }
      }

      // Sequence numbers per destination chain
      for (let i = 0; i < destinationChainIds.length; i++) {
        expect(
          await ramp.getSequenceNumberOfDestinationChain(
            destinationChainIds[i],
          ),
        ).to.equal(1)
      }
    })
  })

  describe('#_calculateFee', () => {
    let latestPrice: number

    beforeEach(async () => {
      // Ensure that the contracts are setup correctly
      expect((await ramp.getConfig()).relayingFeeJuels).to.equal(1)
      expect(await ramp.getFeed(tokens[0].address)).to.equal(priceFeed.address)

      // Generate latest price
      latestPrice = Math.ceil(Math.random() * 1000)
      await priceFeed.connect(roles.defaultAccount).setLatestAnswer(latestPrice)
    })

    it('fails if the feeToken is not a configured fee token', async () => {
      const wrongToken = pools[0].address
      await evmRevert(
        ramp.connect(roles.defaultAccount).publicCalculateFee(wrongToken),
        `UnsupportedFeeToken("${wrongToken}")`,
      )
    })

    describe('when relaying fee is 1', () => {
      it('calculates the correct fee', async () => {
        const result = await ramp
          .connect(roles.defaultAccount)
          .publicCalculateFee(tokens[0].address)
        expect(result).to.equal(latestPrice)
      })
    })

    describe('when relaying fee is different', () => {
      it('calculates the correct fee', async () => {
        const fee = Math.ceil(Math.random() * 1000)
        await ramp.connect(roles.defaultAccount).setConfig({
          router: hre.ethers.constants.AddressZero,
          relayingFeeJuels: fee,
          maxDataSize: maxDataSize,
          maxTokensLength: maxTokensLength,
        })
        const result = await ramp
          .connect(roles.defaultAccount)
          .publicCalculateFee(tokens[0].address)
        expect(result).to.equal(fee * latestPrice)
      })
    })
  })

  describe('#withdrawAccumulatedFees', () => {
    let receiver: string
    let messageData: string
    let options: string
    let amounts: Array<BigNumber>
    let payload: CCIPMessagePayload
    let recipient: Signer
    let recipientAddress: string
    let feeToken: MockERC20
    let feesTaken: BigNumber

    beforeEach(async () => {
      receiver = await roles.stranger.getAddress()
      messageData = hre.ethers.constants.HashZero
      options = hre.ethers.constants.HashZero
      amounts = [bucketRate.div(8)]
      payload = {
        receiver: receiver,
        data: messageData,
        tokens: [tokens.map((t) => t.address)[0]],
        amounts: amounts,
        destinationChainId: destinationChainIds[0],
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
      await tokens[0]
        .connect(roles.defaultAccount)
        .approve(ramp.address, amounts[0])
      await ramp
        .connect(roles.defaultAccount)
        .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
      recipient = roles.stranger
      recipientAddress = await recipient.getAddress()
      feeToken = tokens[0]
      feesTaken = (await priceFeed.latestAnswer()).mul(relayingFeeJuels)
    })

    it('success', async () => {
      const recipientBalanceBefore = await feeToken.balanceOf(recipientAddress)
      const rampBalanceBefore = await feeToken.balanceOf(ramp.address)

      const tx = await ramp
        .connect(roles.defaultAccount)
        .withdrawAccumulatedFees(feeToken.address, recipientAddress, feesTaken)

      const recipientBalanceAfter = await feeToken.balanceOf(recipientAddress)
      const rampBalanceAfter = await feeToken.balanceOf(ramp.address)

      expect(recipientBalanceAfter).to.equal(
        recipientBalanceBefore.add(feesTaken),
      )
      expect(rampBalanceAfter).to.equal(rampBalanceBefore.sub(feesTaken))

      await expect(tx)
        .to.emit(ramp, 'FeesWithdrawn')
        .withArgs(feeToken.address, recipientAddress, feesTaken)
    })

    describe('failure', () => {
      it('fails if called by a non-owner', async () => {
        await evmRevert(
          ramp
            .connect(roles.stranger)
            .withdrawAccumulatedFees(
              feeToken.address,
              recipientAddress,
              feesTaken,
            ),
          'Only callable by owner',
        )
      })
      it('fails if amount is greater than OnRamp balance', async () => {
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .withdrawAccumulatedFees(
              feeToken.address,
              recipientAddress,
              feesTaken.mul(2),
            ),
          'ERC20: transfer amount exceeds balance',
        )
      })
    })
  })

  describe('#requestCrossChainSend', async () => {
    let receiver: string
    let messageData: string
    let options: string
    let amounts: Array<BigNumber>
    let payload: CCIPMessagePayload

    beforeEach(async () => {
      receiver = await roles.stranger.getAddress()
      messageData = hre.ethers.constants.HashZero
      options = hre.ethers.constants.HashZero
      amounts = [bucketRate.div(8), bucketRate.div(4), bucketRate.div(2)]
      payload = {
        receiver: receiver,
        data: messageData,
        tokens: tokens.map((t) => t.address),
        amounts: amounts,
        destinationChainId: destinationChainIds[0],
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
    })

    describe('GASTEST', () => {
      let tx: ContractTransaction
      let gasUsed: BigNumber

      beforeEach(async () => {
        gasUsed = BigNumber.from(0)
      })

      it('GASTEST - Message only (with payment) [ @skip-coverage ]', async () => {
        payload.tokens = [tokens[0].address]
        payload.amounts = [priceFeedLatestAnswer]
        payload.data = stringToBytes('Hello World')
        tx = await tokens[0]
          .connect(roles.defaultAccount)
          .approve(ramp.address, priceFeedLatestAnswer)
        gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        tx = await ramp
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
        gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        expectGasWithinDeviation(
          gasUsed,
          GAS.OnRamp.requestCrossChainSend.MESSAGE_ONLY,
        )
      })

      it('GASTEST - Send 1 token [ @skip-coverage ]', async () => {
        payload.tokens = [tokens[0].address]
        payload.amounts = [amounts[0]]
        tx = await tokens[0]
          .connect(roles.defaultAccount)
          .approve(ramp.address, amounts[0])
        gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        tx = await ramp
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
        gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        expectGasWithinDeviation(
          gasUsed,
          GAS.OnRamp.requestCrossChainSend.ONE_TOKEN,
        )
      })

      it('GASTEST - Send 2 tokens [ @skip-coverage ]', async () => {
        payload.tokens = [tokens[0].address, tokens[1].address]
        payload.amounts = [amounts[0], amounts[1]]
        for (let i = 0; i < 2; i++) {
          tx = await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
          gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        }
        tx = await ramp
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
        gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        expectGasWithinDeviation(
          gasUsed,
          GAS.OnRamp.requestCrossChainSend.TWO_TOKENS,
        )
      })

      it('GASTEST - Send 3 tokens [ @skip-coverage ]', async () => {
        for (let i = 0; i < tokens.length; i++) {
          tx = await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
          gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        }
        tx = await ramp
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
        gasUsed = gasUsed.add((await tx.wait()).gasUsed)
        expectGasWithinDeviation(
          gasUsed,
          GAS.OnRamp.requestCrossChainSend.THREE_TOKENS,
        )
      })
    })

    describe('success (3 tokens - without fees)', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        await ramp.connect(roles.defaultAccount).setConfig({
          router: hre.ethers.constants.AddressZero,
          relayingFeeJuels: 0,
          maxDataSize: maxDataSize,
          maxTokensLength: maxTokensLength,
        })
        for (let i = 0; i < tokens.length; i++) {
          tx = await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
        }
        tx = await ramp
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
      })

      it('does not store any fees', async () => {
        expect(await tokens[0].balanceOf(ramp.address)).to.equal(0)
      })

      it('does not emit a fee taken event', async () => {
        await expect(tx).to.not.emit(ramp, 'FeeCharged')
      })

      it('locks each token into the pools', async () => {
        for (let i = 0; i < pools.length; i++) {
          const poolBalance = await tokens[i].balanceOf(pools[i].address)
          expect(poolBalance).to.equal(amounts[i])
        }
      })

      it('emits a CrossChainSendRequested event', async () => {
        const receipt = await tx.wait()
        const eventArgs = ramp.interface.parseLog(
          receipt.logs[receipt.logs.length - 1],
        ).args
        const expectedAmounts = amounts
        requestEventArgsEqual(eventArgs, {
          sequenceNumber: eventArgs?.message?.sequenceNumber,
          sourceChainId: BigNumber.from(sourceChainId),
          destinationChainId: BigNumber.from(payload.destinationChainId),
          sender: await roles.defaultAccount.getAddress(),
          receiver: receiver,
          data: messageData,
          tokens: tokens.map((t) => t.address),
          amounts: expectedAmounts,
          options: options,
        })
      })

      it('increments the sequence number per destination chain', async () => {
        expect(
          await ramp.getSequenceNumberOfDestinationChain(
            payload.destinationChainId,
          ),
        ).to.equal(2)
      })
    })

    describe('success (3 tokens)', () => {
      let tx: ContractTransaction
      let feeTaken: BigNumber
      beforeEach(async () => {
        for (let i = 0; i < tokens.length; i++) {
          tx = await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
        }
        tx = await ramp
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload, hre.ethers.constants.AddressZero)
        feeTaken = (await priceFeed.latestAnswer()).mul(relayingFeeJuels)
      })

      it('stores the fee in the OnRamp', async () => {
        expect(await tokens[0].balanceOf(ramp.address)).to.equal(feeTaken)
      })

      it('emits a fee taken event', async () => {
        await expect(tx)
          .to.emit(ramp, 'FeeCharged')
          .withArgs(
            await roles.defaultAccount.getAddress(),
            ramp.address,
            feeTaken,
          )
      })

      it('locks each token into the pools', async () => {
        for (let i = 0; i < pools.length; i++) {
          const poolBalance = await tokens[i].balanceOf(pools[i].address)
          if (i == 0) {
            // Calculate amount minus fee
            expect(poolBalance).to.equal(amounts[i].sub(feeTaken))
          } else {
            expect(poolBalance).to.equal(amounts[i])
          }
        }
      })

      it('emits a CrossChainSendRequested event', async () => {
        const receipt = await tx.wait()
        const eventArgs = ramp.interface.parseLog(
          receipt.logs[receipt.logs.length - 1],
        ).args
        const expectedAmounts = amounts
        expectedAmounts[0] = expectedAmounts[0].sub(feeTaken)
        requestEventArgsEqual(eventArgs, {
          sequenceNumber: eventArgs?.message?.sequenceNumber,
          sourceChainId: BigNumber.from(sourceChainId),
          destinationChainId: BigNumber.from(payload.destinationChainId),
          sender: await roles.defaultAccount.getAddress(),
          receiver: receiver,
          data: messageData,
          tokens: tokens.map((t) => t.address),
          amounts: expectedAmounts,
          options: options,
        })
      })

      it('increments the sequence number per destination chain', async () => {
        expect(
          await ramp.getSequenceNumberOfDestinationChain(
            payload.destinationChainId,
          ),
        ).to.equal(2)
      })
    })

    describe('failure', () => {
      it('fails when the ramp is paused', async () => {
        await ramp.pause()
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          'Pausable: paused',
        )
      })
      it('fails when the AFN signal is bad', async () => {
        await afn.voteBad()
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          'BadAFNSignal()',
        )
      })
      it('fails when the AFN signal is stale', async () => {
        await afn.setTimestamp(BigNumber.from(1))
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          'StaleAFNHeartbeat()',
        )
      })
      it('fails if the originalSender is set, but its not called by the router', async () => {
        await ramp.connect(roles.defaultAccount).setConfig({
          router: await roles.oracleNode.getAddress(),
          relayingFeeJuels: relayingFeeJuels,
          maxDataSize: maxDataSize,
          maxTokensLength: maxTokensLength,
        })
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(
              payload,
              await roles.oracleNode1.getAddress(),
            ),
          `MustBeCalledByRouter()`,
        )
      })
      it('fails when the allowlist is set and the sender is not part of it', async () => {
        await ramp.connect(roles.defaultAccount).setAllowlist([])
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `SenderNotAllowed("${await roles.defaultAccount.getAddress()}")`,
        )
      })
      it('fails if the destination chain ID is not supported by the OnRamp', async () => {
        payload.destinationChainId = BigNumber.from(999)
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `UnsupportedDestinationChain(${payload.destinationChainId})`,
        )
      })
      it('fails if the data is larger than the max data size', async () => {
        const newDataSize = 1
        await ramp.connect(roles.defaultAccount).setConfig({
          router: hre.ethers.constants.AddressZero,
          maxDataSize: newDataSize,
          maxTokensLength: maxTokensLength,
          relayingFeeJuels: relayingFeeJuels,
        })
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `MessageTooLarge(${newDataSize}, 32)`,
        )
      })
      it('fails if there are too many tokens, or the amounts are a different length to the tokens', async () => {
        await ramp.connect(roles.defaultAccount).setConfig({
          router: hre.ethers.constants.AddressZero,
          maxDataSize: maxDataSize,
          maxTokensLength: 1,
          relayingFeeJuels: relayingFeeJuels,
        })
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `UnsupportedNumberOfTokens()`,
        )
        await ramp.connect(roles.defaultAccount).setConfig({
          router: hre.ethers.constants.AddressZero,
          maxDataSize: maxDataSize,
          maxTokensLength: maxTokensLength,
          relayingFeeJuels: relayingFeeJuels,
        })
        payload.amounts = [100]
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if the fee token is not supported', async () => {
        const wrongToken = await roles.stranger.getAddress()
        payload.tokens[0] = wrongToken
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `UnsupportedFeeToken("${wrongToken}")`,
        )
      })
      it('fails if the sender does not approve the ramp for a token', async () => {
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `ERC20: transfer amount exceeds allowance`,
        )
      })
      it('fails if a token does not have a configured pool in the ramp', async () => {
        // Remove one of the 3 pools from the ramp
        await ramp
          .connect(roles.defaultAccount)
          .removePool(tokens[0].address, pools[0].address)

        // Approve so it gets past the fee taking
        await tokens[0]
          .connect(roles.defaultAccount)
          .approve(ramp.address, amounts[0])
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `UnsupportedToken("${tokens[0].address}")`,
        )
      })
      it('fails if a lock exceeds the token limit', async () => {
        await pools[0]
          .connect(roles.defaultAccount)
          .setLockOrBurnBucket(1, 1, true)

        await tokens[0]
          .connect(roles.defaultAccount)
          .approve(ramp.address, amounts[0])
        const amountToLock = amounts[0].sub(
          relayingFeeJuels * priceFeedLatestAnswer,
        )
        await evmRevert(
          ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(payload, hre.ethers.constants.AddressZero),
          `ExceedsTokenLimit(1, ${amountToLock})`,
        )
      })
    })
  })

  describe('#setConfig', () => {
    it('only allows owner to set', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setConfig({
          router: hre.ethers.constants.AddressZero,
          relayingFeeJuels: 1,
          maxDataSize: 2,
          maxTokensLength: 3,
        }),
        'Only callable by owner',
      )
    })

    it('sets the max data size correctly', async () => {
      const router = hre.ethers.constants.AddressZero
      const newDataSize = 1
      const newTokensLength = 1
      const newRelayFee = 1
      let tx = await ramp.connect(roles.defaultAccount).setConfig({
        router: hre.ethers.constants.AddressZero,
        maxDataSize: newDataSize,
        maxTokensLength: newTokensLength,
        relayingFeeJuels: newRelayFee,
      })
      const config = await await ramp.getConfig()
      expect(config.maxDataSize).to.equal(newDataSize)
      expect(config.maxTokensLength).to.equal(newTokensLength)
      expect(config.relayingFeeJuels).to.equal(newRelayFee)
      await expect(tx)
        .to.emit(ramp, 'OnRampConfigSet')
        .withArgs([router, newRelayFee, newDataSize, newTokensLength])
    })
  })

  describe('#setAllowListEnabled', () => {
    it('only allows owner to set', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setAllowlistEnabled(false),
        'Only callable by owner',
      )
    })

    it('sets the allowlistEnabled flag correctly', async () => {
      let tx = await ramp
        .connect(roles.defaultAccount)
        .setAllowlistEnabled(false)
      expect(await ramp.getAllowlistEnabled()).to.be.false
      await expect(tx).to.emit(ramp, 'AllowlistEnabledSet').withArgs(false)

      tx = await ramp.connect(roles.defaultAccount).setAllowlistEnabled(true)
      expect(await ramp.getAllowlistEnabled()).to.be.true
      await expect(tx).to.emit(ramp, 'AllowlistEnabledSet').withArgs(true)
    })
  })

  describe('#setAllowList', () => {
    let newAllowList: Array<string>

    beforeEach(async () => {
      newAllowList = [
        await roles.oracleNode1.getAddress(),
        await roles.oracleNode2.getAddress(),
      ]
    })

    it('only allows owner to set', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setAllowlist(newAllowList),
        'Only callable by owner',
      )
    })

    it('sets the correct allowlist', async () => {
      await ramp.connect(roles.defaultAccount).setAllowlist(newAllowList)
      expect(await ramp.getAllowlist()).to.deep.equal(newAllowList)
    })
  })

  describe('#setAFN', () => {
    let newAFN: MockAFN

    beforeEach(async () => {
      newAFN = <MockAFN>(
        await deployContract(roles.defaultAccount, MockAFNArtifact)
      )
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setAFN(newAFN.address),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new AFN', async () => {
      const tx = await ramp.connect(roles.defaultAccount).setAFN(newAFN.address)
      expect(await ramp.getAFN()).to.equal(newAFN.address)
      await expect(tx)
        .to.emit(ramp, 'AFNSet')
        .withArgs(afn.address, newAFN.address)
    })
  })

  describe('#setMaxSecondsWithoutAFNHeartbeat', () => {
    let newTime: BigNumber

    beforeEach(async () => {
      newTime = maxTimeWithoutAFNSignal.mul(2)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setMaxSecondsWithoutAFNHeartbeat(newTime),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setMaxSecondsWithoutAFNHeartbeat(newTime)
      expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(newTime)
      await expect(tx)
        .to.emit(ramp, 'AFNMaxHeartbeatTimeSet')
        .withArgs(maxTimeWithoutAFNSignal, newTime)
    })
  })
})
