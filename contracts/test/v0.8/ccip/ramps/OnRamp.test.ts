import hre from 'hardhat'
import { publicAbi } from '../../../test-helpers/helpers'
import { BigNumber, constants, ContractTransaction } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  NativeTokenPool,
  MockAFN,
  MockAggregator,
  EVM2EVMTollOnRamp,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  EVM2AnyTollMessage,
  requestEventArgsEqual,
} from '../../../test-helpers/ccip/ccip'

const { deployContract } = hre.waffle

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let PriceFeedArtifact: Artifact
let RampArtifact: Artifact

let roles: Roles

let afn: MockAFN
let ramp: EVM2EVMTollOnRamp
const numberOfTokensPoolsAndFeeds = 3
let tokens: Array<MockERC20>
let pools: Array<NativeTokenPool>
let priceFeed: MockAggregator
let priceFeedLatestAnswer: number = 100
const sourceChainId: number = 123
const destinationChainId: BigNumber = BigNumber.from(9)
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

describe('EVM2EVMTollOnRamp', () => {
  beforeEach(async () => {
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    mintAmount = BigNumber.from('1000000000000000000')

    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('NativeTokenPool')
    PriceFeedArtifact = await hre.artifacts.readArtifact('MockAggregator')
    RampArtifact = await hre.artifacts.readArtifact('EVM2EVMTollOnRamp')

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

    ramp = <EVM2EVMTollOnRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [
        sourceChainId,
        destinationChainId,
        tokens.map((t) => t.address),
        pools.map((p) => p.address),
        [priceFeed.address, constants.AddressZero, constants.AddressZero],
        [await roles.defaultAccount.getAddress()],
        afn.address,
        maxTimeWithoutAFNSignal,
        {
          router: await roles.oracleNode.getAddress(),
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
      // EVM2EVMTollOnRamp
      'forwardFromRouter',
      'CHAIN_ID',
      'DESTINATION_CHAIN_ID',
      'getRequiredFee',
      'getTokenPool',
      'setAllowlistEnabled',
      'getAllowlistEnabled',
      'setAllowlist',
      'getAllowlist',
      'getSequenceNumber',
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
      // PoolCollector
      'withdrawAccumulatedFees',
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

      // Sequence number
      expect(await ramp.getSequenceNumber()).to.equal(1)
    })
  })

  describe('#getRequiredFee', () => {
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
        ramp.connect(roles.defaultAccount).getRequiredFee(wrongToken),
        `UnsupportedFeeToken("${wrongToken}")`,
      )
    })

    describe('when relaying fee is 1', () => {
      it('calculates the correct fee', async () => {
        const result = await ramp
          .connect(roles.defaultAccount)
          .getRequiredFee(tokens[0].address)
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
          .getRequiredFee(tokens[0].address)
        expect(result).to.equal(fee * latestPrice)
      })
    })
  })

  describe('#forwardFromRouter', async () => {
    let receiver: string
    let messageData: string
    let amounts: Array<BigNumber>
    let evmToAnyTollMessage: EVM2AnyTollMessage

    beforeEach(async () => {
      receiver = await roles.stranger.getAddress()
      messageData = hre.ethers.constants.HashZero
      amounts = [bucketRate.div(8), bucketRate.div(4), bucketRate.div(2)]
      evmToAnyTollMessage = {
        receiver: receiver,
        data: messageData,
        tokens: tokens.map((t) => t.address),
        amounts: amounts,
        feeToken: tokens.map((t) => t.address)[0],
        feeTokenAmount: 0,
        gasLimit: 0,
      }
    })

    describe('success (3 tokens)', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        for (let i = 0; i < tokens.length; i++) {
          tx = await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
        }
        tx = await ramp
          .connect(roles.oracleNode)
          .forwardFromRouter(
            evmToAnyTollMessage,
            await roles.defaultAccount.getAddress(),
          )
      })

      it('emits a CrossChainSendRequested event', async () => {
        const receipt = await tx.wait()
        const eventArgs = ramp.interface.parseLog(
          receipt.logs[receipt.logs.length - 1],
        ).args
        console.log(eventArgs)
        requestEventArgsEqual(eventArgs, {
          sequenceNumber: eventArgs?.message.sequenceNumber,
          sourceChainId: BigNumber.from(sourceChainId),
          sender: await roles.defaultAccount.getAddress(),
          receiver: receiver,
          data: messageData,
          tokens: tokens.map((t) => t.address),
          amounts: amounts,
          feeToken: tokens.map((t) => t.address)[0],
          feeTokenAmount: 0,
          gasLimit: 0,
        })
      })

      it('increments the sequence number per destination chain', async () => {
        expect(await ramp.getSequenceNumber()).to.equal(2)
      })
    })

    describe('failure', () => {
      it('fails when the ramp is paused', async () => {
        await ramp.pause()
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          'Pausable: paused',
        )
      })
      it('fails when the AFN signal is bad', async () => {
        await afn.voteBad()
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          'BadAFNSignal()',
        )
      })
      it('fails when the AFN signal is stale', async () => {
        await afn.setTimestamp(BigNumber.from(1))
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          'StaleAFNHeartbeat()',
        )
      })
      it('fails if the sender is the router but the originalSender is not set', async () => {
        for (let i = 0; i < tokens.length; i++) {
          await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
        }
        await ramp.connect(roles.defaultAccount).setConfig({
          router: await roles.oracleNode.getAddress(),
          relayingFeeJuels: relayingFeeJuels,
          maxDataSize: maxDataSize,
          maxTokensLength: maxTokensLength,
        })
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              hre.ethers.constants.AddressZero,
            ),
          `RouterMustSetOriginalSender()`,
        )
      })
      it('fails when the allowlist is set and the sender is not part of it', async () => {
        for (let i = 0; i < tokens.length; i++) {
          await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
        }
        await ramp.connect(roles.defaultAccount).setAllowlist([])
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          `SenderNotAllowed("${await roles.defaultAccount.getAddress()}")`,
        )
      })
      it('fails if the data is larger than the max data size', async () => {
        const newDataSize = 1
        await ramp.connect(roles.defaultAccount).setConfig({
          router: await roles.oracleNode.getAddress(),
          maxDataSize: newDataSize,
          maxTokensLength: maxTokensLength,
          relayingFeeJuels: relayingFeeJuels,
        })
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          `MessageTooLarge(${newDataSize}, 32)`,
        )
      })
      it('fails if there are too many tokens, or the amounts are a different length to the tokens', async () => {
        await ramp.connect(roles.defaultAccount).setConfig({
          router: await roles.oracleNode.getAddress(),
          maxDataSize: maxDataSize,
          maxTokensLength: 1,
          relayingFeeJuels: relayingFeeJuels,
        })
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          `UnsupportedNumberOfTokens()`,
        )
        await ramp.connect(roles.defaultAccount).setConfig({
          router: await roles.oracleNode.getAddress(),
          maxDataSize: maxDataSize,
          maxTokensLength: maxTokensLength,
          relayingFeeJuels: relayingFeeJuels,
        })
        evmToAnyTollMessage.amounts = [100]
        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          `UnsupportedNumberOfTokens()`,
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
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          `UnsupportedToken("${tokens[0].address}")`,
        )
      })
      it('fails if a lock exceeds the token limit', async () => {
        for (let i = 0; i < tokens.length; i++) {
          await tokens[i]
            .connect(roles.defaultAccount)
            .approve(ramp.address, amounts[i])
        }
        await pools[0]
          .connect(roles.defaultAccount)
          .setLockOrBurnBucket(1, 1, true)

        await tokens[0]
          .connect(roles.defaultAccount)
          .approve(ramp.address, amounts[0])

        await evmRevert(
          ramp
            .connect(roles.oracleNode)
            .forwardFromRouter(
              evmToAnyTollMessage,
              await roles.defaultAccount.getAddress(),
            ),
          `ExceedsTokenLimit(1, ${amounts[0]})`,
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
