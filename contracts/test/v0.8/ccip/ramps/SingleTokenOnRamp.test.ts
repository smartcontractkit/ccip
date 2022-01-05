import hre from 'hardhat'
import { publicAbi } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { BigNumber, ContractReceipt, ContractTransaction, Signer } from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  MockERC20,
  LockUnlockPool,
  SingleTokenOnRamp,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  CCIPMessagePayload,
  requestEventArgsEqual,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let afn: MockAFN
let ramp: SingleTokenOnRamp
let token: MockERC20
let pool: LockUnlockPool
let destinationTokenAddress: string

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let RampArtifact: Artifact

const sourceChainId: number = 123
const destinationChainId: number = 234
let bucketRate: BigNumber
let bucketCapactiy: BigNumber
let maxTimeWithoutAFNSignal: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('SingleTokenOnRamp', () => {
  beforeEach(async () => {
    destinationTokenAddress = await roles.stranger.getAddress()

    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')

    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('LockUnlockPool')
    RampArtifact = await hre.artifacts.readArtifact('SingleTokenOnRamp')

    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    maxTimeWithoutAFNSignal = BigNumber.from(60).mul(60) // 1 hour
    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        BigNumber.from('1000000000000000000'),
      ])
    )
    pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [token.address])
    )
    ramp = <SingleTokenOnRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [
        sourceChainId,
        token.address,
        pool.address,
        destinationChainId,
        destinationTokenAddress,
        [roles.defaultAccount.getAddress()],
        true,
        bucketRate,
        bucketCapactiy,
        afn.address,
        maxTimeWithoutAFNSignal,
      ])
    )
    await pool.connect(roles.defaultAccount).setOnRamp(ramp.address, true)
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(ramp, [
      // SingleTokenRamp
      'requestCrossChainSend',
      'TOKEN',
      'DESTINATION_TOKEN',
      'POOL',
      'DESTINATION_CHAIN_ID',
      'CHAIN_ID',
      'setAllowlistEnabled',
      'getAllowlistEnabled',
      'setAllowlist',
      'getAllowlist',
      'configureTokenBucket',
      'getTokenBucket',
      // HealthChecker
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
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
      const owner = await roles.defaultAccount.getAddress()
      await expect(await ramp.TOKEN()).to.equal(token.address)
      await expect(await ramp.POOL()).to.not.equal(
        hre.ethers.constants.AddressZero,
      )
      await expect(await ramp.DESTINATION_CHAIN_ID()).to.equal(
        destinationChainId,
      )
      await expect(await ramp.owner()).to.equal(owner)
      await expect(await ramp.getAllowlistEnabled()).to.be.true
      await expect(await ramp.getAllowlist()).to.deep.equal([
        await roles.defaultAccount.getAddress(),
      ])
      await expect(await ramp.getAFN()).to.equal(afn.address)
      await expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        maxTimeWithoutAFNSignal,
      )
      const tokenBucket = await ramp.getTokenBucket()
      await expect(tokenBucket.rate).to.equal(bucketRate)
      await expect(tokenBucket.capacity).to.equal(bucketCapactiy)

      await expect(await pool.owner()).to.equal(owner)
      await expect(await pool.isOnRamp(ramp.address)).to.equal(true)
      await expect(await pool.getToken()).to.equal(token.address)
    })

    it('should fail if the pool token is different from the ramp token', async () => {
      const differentToken = <MockERC20>(
        await deployContract(roles.defaultAccount, TokenArtifact, [
          'LINK Token',
          'LINK',
          await roles.defaultAccount.getAddress(),
          BigNumber.from('1000000000000000000'),
        ])
      )
      await evmRevert(
        deployContract(roles.defaultAccount, RampArtifact, [
          sourceChainId,
          differentToken.address,
          pool.address,
          destinationChainId,
          destinationTokenAddress,
          [roles.defaultAccount.getAddress()],
          true,
          bucketRate,
          bucketCapactiy,
          afn.address,
          maxTimeWithoutAFNSignal,
        ]),
        `TokenMismatch()`,
      )
    })
  })

  describe('#requestCrossChainSend', () => {
    let receiver: string
    let messagedata: string
    let options: string
    let amount: BigNumber
    let message: CCIPMessagePayload

    beforeEach(async () => {
      receiver = await roles.stranger.getAddress()
      messagedata = hre.ethers.constants.HashZero
      options = hre.ethers.constants.HashZero
      amount = BigNumber.from('1000000000000000')
      message = {
        receiver: receiver,
        data: messagedata,
        tokens: [token.address],
        amounts: [amount],
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
    })

    describe('when contract not paused', () => {
      it('fails if there are not enough or too many tokens', async () => {
        message.tokens = []
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
        message.tokens = [token.address, token.address]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if there are not enough or too many amounts', async () => {
        message.amounts = []
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
        message.amounts = [amount, amount]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if token is not configured token', async () => {
        message.tokens = [receiver]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `UnsupportedToken("${token.address}", "${receiver}")`,
        )
      })

      it('fails if sent by a non-allowlisted address', async () => {
        await evmRevert(
          ramp.connect(roles.stranger).requestCrossChainSend(message),
          `SenderNotAllowed("${await roles.stranger.getAddress()}")`,
        )
      })

      it('fails if the send amount is greater than the bucket allows', async () => {
        message.amounts = [bucketCapactiy.add(1)]
        await evmRevert(
          ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
          `ExceedsTokenLimit(${bucketCapactiy}, ${message.amounts[0]})`,
        )
      })

      describe('sending a message', () => {
        let tx: ContractTransaction
        let owner: Signer
        let initialPoolBalance: BigNumber

        beforeEach(async () => {
          owner = roles.defaultAccount
          initialPoolBalance = await token.balanceOf(pool.address)
          await token.approve(pool.address, amount)
          tx = await ramp
            .connect(roles.defaultAccount)
            .requestCrossChainSend(message)
        })

        it('emits a Locked event in the pool', async () => {
          await expect(tx)
            .to.emit(pool, 'Locked')
            .withArgs(ramp.address, await owner.getAddress(), amount)
        })

        it('transfers the tokens to the pool', async () => {
          const expectedBalance = initialPoolBalance.add(amount)
          await expect(await token.balanceOf(pool.address)).to.equal(
            expectedBalance,
          )
        })

        it('emits a message send request', async () => {
          const receipt: ContractReceipt = await tx.wait()
          const eventArgs = receipt.events?.[3]?.args?.[0]
          requestEventArgsEqual(eventArgs, {
            sequenceNumber: eventArgs?.sequenceNumber,
            sourceChainId: BigNumber.from(sourceChainId),
            destinationChainId: BigNumber.from(destinationChainId),
            sender: await owner.getAddress(),
            receiver: receiver,
            data: messagedata,
            tokens: [destinationTokenAddress],
            amounts: [amount],
            options: options,
          })
        })
      })
    })

    it('fails when the ramp is paused', async () => {
      await ramp.pause()
      await evmRevert(
        ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
        'Pausable: paused',
      )
    })

    it('fails whenn the AFN signal is bad', async () => {
      await afn.voteBad()
      await evmRevert(
        ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
        'BadAFNSignal()',
      )
    })

    it('fails when the AFN signal is stale', async () => {
      await afn.setTimestamp(BigNumber.from(1))
      await evmRevert(
        ramp.connect(roles.defaultAccount).requestCrossChainSend(message),
        'StaleAFNHeartbeat()',
      )
    })
  })

  describe('#pause', () => {
    it('owner can pause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).pause())
        .to.emit(ramp, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await ramp.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).unpause())
        .to.emit(ramp, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#setAllowlistEnabled', () => {
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
      await expect(await ramp.getAllowlistEnabled()).to.be.false
      await expect(tx).to.emit(ramp, 'AllowlistEnabledSet').withArgs(false)

      tx = await ramp.connect(roles.defaultAccount).setAllowlistEnabled(true)
      await expect(await ramp.getAllowlistEnabled()).to.be.true
      await expect(tx).to.emit(ramp, 'AllowlistEnabledSet').withArgs(true)
    })
  })

  describe('#setAllowlist', () => {
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
      await expect(await ramp.getAllowlist()).to.deep.equal(newAllowList)
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

  describe('#configureTokenBucket', () => {
    let newRate: BigNumber
    let newCapacity: BigNumber

    beforeEach(async () => {
      newRate = BigNumber.from(5)
      newCapacity = bucketCapactiy.add(10)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp
          .connect(roles.stranger)
          .configureTokenBucket(newRate, newCapacity, true),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .configureTokenBucket(newRate, newCapacity, true)
      const tokenBucketParams = await ramp.getTokenBucket()
      expect(tokenBucketParams.rate).to.equal(newRate)
      expect(tokenBucketParams.capacity).to.equal(newCapacity)
      await expect(tx)
        .to.emit(ramp, 'NewTokenBucketConstructed')
        .withArgs(newRate, newCapacity, true)
    })
  })

  describe('#typeAndVersion', () => {
    it('should return the correct type and version', async () => {
      const response = await ramp.typeAndVersion()
      await expect(response).to.equal('SingleTokenOnRamp 1.1.0')
    })
  })
})
