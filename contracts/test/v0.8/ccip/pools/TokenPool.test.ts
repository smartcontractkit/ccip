import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { MockERC20, TokenPoolHelper } from '../../../../typechain'
import { BigNumber, ContractTransaction, Signer } from 'ethers'
import { getUsers, Roles } from '../../../test-helpers/setup'
import { evmRevert } from '../../../test-helpers/matchers'
import { publicAbi } from '../../../test-helpers/helpers'

const { deployContract } = hre.waffle

let roles: Roles

const DEPOSIT = BigNumber.from(123)

let token: MockERC20
let onRamp: Signer
let offRamp: Signer
let pool: TokenPoolHelper

let bucketRate: BigNumber
let bucketCapactiy: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('TokenPool', () => {
  beforeEach(async () => {
    const MockERC20Artifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    token = <MockERC20>(
      await deployContract(roles.defaultAccount, MockERC20Artifact, [
        'Test Token',
        'TEST',
        await roles.defaultAccount.getAddress(),
        BigNumber.from(1000000),
      ])
    )

    // For testing onlyRamp
    onRamp = roles.oracleNode1
    offRamp = roles.oracleNode2
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    token
      .connect(roles.defaultAccount)
      .mint(await onRamp.getAddress(), BigNumber.from(1000000))
    token
      .connect(roles.defaultAccount)
      .mint(await offRamp.getAddress(), BigNumber.from(1000000))
    const TokenPoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'TokenPoolHelper',
    )
    pool = <TokenPoolHelper>(
      await deployContract(roles.defaultAccount, TokenPoolArtifact, [
        token.address,
        bucketRate,
        bucketCapactiy,
        bucketRate,
        bucketCapactiy,
      ])
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(pool, [
      // TokenPoolHelper
      'assertLockOrBurnModifier',
      'assertMintOrReleaseModifier',
      // TokenPool
      'lockOrBurn',
      'releaseOrMint',
      'setOnRamp',
      'setOffRamp',
      'isOnRamp',
      'isOffRamp',
      'getToken',
      'setLockOrBurnBucket',
      'setReleaseOrMintBucket',
      'getLockOrBurnBucket',
      'getReleaseOrMintBucket',
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
    it('is initialized correctly', async () => {
      expect(await pool.getToken()).to.equal(token.address)
      expect(await pool.owner()).to.equal(
        await roles.defaultAccount.getAddress(),
      )
      const lockBucket = await pool.getLockOrBurnBucket()
      expect(lockBucket.rate).to.equal(bucketRate)
      expect(lockBucket.capacity).to.equal(bucketCapactiy)
      const releaseBucket = await pool.getReleaseOrMintBucket()
      expect(releaseBucket.rate).to.equal(bucketRate)
      expect(releaseBucket.capacity).to.equal(bucketCapactiy)
    })
  })

  describe('#pause', () => {
    it('owner can pause pool', async () => {
      const account = roles.defaultAccount
      await expect(pool.connect(account).pause())
        .to.emit(pool, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async () => {
      await expect(pool.connect(onRamp).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(pool.connect(offRamp).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(pool.connect(roles.stranger).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await pool.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause pool', async () => {
      const account = roles.defaultAccount
      await expect(pool.connect(account).unpause())
        .to.emit(pool, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async () => {
      await expect(pool.connect(onRamp).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(pool.connect(offRamp).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(pool.connect(roles.stranger).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#setOnRamp', () => {
    let expectedRamp: string
    beforeEach(async () => {
      expectedRamp = await roles.oracleNode3.getAddress()
    })

    it('sets the on ramp when called by the owner', async () => {
      await pool.connect(roles.defaultAccount).setOnRamp(expectedRamp, true)
      expect(await pool.isOnRamp(expectedRamp)).to.equal(true)
    })

    it('reverts when called by any other account', async () => {
      await evmRevert(
        pool.connect(roles.stranger).setOnRamp(expectedRamp, true),
        'Only callable by owner',
      )
    })
  })

  describe('#setOffRamp', () => {
    let expectedRamp: string
    beforeEach(async () => {
      expectedRamp = await roles.oracleNode3.getAddress()
    })

    it('sets the on ramp when called by the owner', async () => {
      await pool.connect(roles.defaultAccount).setOffRamp(expectedRamp, true)
      expect(await pool.isOffRamp(expectedRamp)).to.equal(true)
    })

    it('reverts when called by any other account', async () => {
      await evmRevert(
        pool.connect(roles.stranger).setOffRamp(expectedRamp, true),
        'Only callable by owner',
      )
    })
  })

  describe('#setLockOrBurnBucket', () => {
    it('fails when caller is not owner', async () => {
      await evmRevert(
        pool.connect(roles.stranger).setLockOrBurnBucket(1, 1, true),
        'Only callable by owner',
      )
    })

    describe('success', () => {
      const newRate = 500
      const newCapacity = 5000000
      let tx: ContractTransaction
      beforeEach(async () => {
        tx = await pool
          .connect(roles.defaultAccount)
          .setLockOrBurnBucket(newRate, newCapacity, true)
      })

      it('sets the new bucket', async () => {
        const newBucket = await pool.getLockOrBurnBucket()
        expect(newBucket.rate).to.equal(newRate)
        expect(newBucket.capacity).to.equal(newCapacity)
        expect(newBucket.tokens).to.equal(newCapacity)
      })
      it('emits an event', async () => {
        await expect(tx)
          .to.emit(pool, 'NewLockBurnBucketConstructed')
          .withArgs(newRate, newCapacity, true)
      })
    })
  })

  describe('#setReleaseOrMintBucket', () => {
    it('fails when caller is not owner', async () => {
      await evmRevert(
        pool.connect(roles.stranger).setReleaseOrMintBucket(1, 1, true),
        'Only callable by owner',
      )
    })

    describe('success', () => {
      const newRate = 500
      const newCapacity = 5000000
      let tx: ContractTransaction
      beforeEach(async () => {
        tx = await pool
          .connect(roles.defaultAccount)
          .setReleaseOrMintBucket(newRate, newCapacity, true)
      })

      it('sets the new bucket', async () => {
        const newBucket = await pool.getReleaseOrMintBucket()
        expect(newBucket.rate).to.equal(newRate)
        expect(newBucket.capacity).to.equal(newCapacity)
        expect(newBucket.tokens).to.equal(newCapacity)
      })
      it('emits an event', async () => {
        await expect(tx)
          .to.emit(pool, 'NewReleaseMintBucketConstructed')
          .withArgs(newRate, newCapacity, true)
      })
    })
  })

  describe('modifier #assertLockOrBurn', () => {
    describe('failure', () => {
      it('fails when called by neither the owner or onramp', async () => {
        await evmRevert(
          pool.connect(roles.stranger).assertLockOrBurnModifier(DEPOSIT),
          'PermissionsError()',
        )
      })
      it('fails when amount exceeds token limit', async () => {
        const tooMuch = bucketCapactiy.add(1)
        await evmRevert(
          pool.connect(roles.defaultAccount).assertLockOrBurnModifier(tooMuch),
          `ExceedsTokenLimit(${bucketCapactiy}, ${tooMuch})`,
        )
      })
    })
    describe('success', () => {
      it('passes when called by the owner within limit', async () => {
        expect(
          await pool
            .connect(roles.defaultAccount)
            .assertLockOrBurnModifier(DEPOSIT),
        ).to.emit(pool, 'AssertionPassed')
      })
      it('passes when called by the onramp within limit', async () => {
        await pool
          .connect(roles.defaultAccount)
          .setOnRamp(await roles.oracleNode.getAddress(), true)
        expect(
          await pool
            .connect(roles.oracleNode)
            .assertLockOrBurnModifier(DEPOSIT),
        ).to.emit(pool, 'AssertionPassed')
      })
    })
  })
  describe('modifier #assertMintOrRelease', () => {
    describe('failure', () => {
      it('fails when called by neither the owner or offramp', async () => {
        await evmRevert(
          pool.connect(roles.stranger).assertMintOrReleaseModifier(DEPOSIT),
          'PermissionsError()',
        )
      })
      it('fails when amount exceeds token limit', async () => {
        const tooMuch = bucketCapactiy.add(1)
        await evmRevert(
          pool
            .connect(roles.defaultAccount)
            .assertMintOrReleaseModifier(tooMuch),
          `ExceedsTokenLimit(${bucketCapactiy}, ${tooMuch})`,
        )
      })
    })
    describe('success', () => {
      it('passes when called by the owner within limit', async () => {
        expect(
          await pool
            .connect(roles.defaultAccount)
            .assertMintOrReleaseModifier(DEPOSIT),
        ).to.emit(pool, 'AssertionPassed')
      })
      it('passes when called by the offramp within limit', async () => {
        await pool
          .connect(roles.defaultAccount)
          .setOffRamp(await roles.oracleNode.getAddress(), true)
        expect(
          await pool
            .connect(roles.oracleNode)
            .assertMintOrReleaseModifier(DEPOSIT),
        ).to.emit(pool, 'AssertionPassed')
      })
    })
  })
})
