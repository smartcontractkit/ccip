import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { MockERC20, NativeTokenPool } from '../../../../typechain'
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
let pool: NativeTokenPool

let bucketRate: BigNumber
let bucketCapactiy: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('NativeTokenPool', () => {
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
    const NativeTokenPoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'NativeTokenPool',
    )
    pool = <NativeTokenPool>(
      await deployContract(roles.defaultAccount, NativeTokenPoolArtifact, [
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
      // NativeTokenPool
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

  describe('#lockOrBurn', () => {
    let account: Signer
    let sender: string
    let depositor: string

    describe('called by the owner', () => {
      beforeEach(async () => {
        account = roles.defaultAccount
        sender = await account.getAddress()
        depositor = await account.getAddress()

        await token.connect(account).approve(pool.address, DEPOSIT)
        const allowance = await token.allowance(depositor, pool.address)
        await expect(allowance).to.equal(DEPOSIT)
      })

      it('can lock tokens', async () => {
        await expect(pool.connect(account).lockOrBurn(depositor, DEPOSIT))
          .to.emit(pool, 'Locked')
          .withArgs(sender, depositor, DEPOSIT)

        const poolBalance = await token.balanceOf(pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      it("can't store when paused", async () => {
        await pool.connect(account).pause()

        await evmRevert(
          pool.connect(account).lockOrBurn(depositor, DEPOSIT),
          'Pausable: paused',
        )
      })

      it('fails if the amount exceeds the token limit', async () => {
        await pool.connect(account).setLockOrBurnBucket(1, 1, true)

        await evmRevert(
          pool.connect(account).lockOrBurn(depositor, DEPOSIT),
          `ExceedsTokenLimit(1, ${DEPOSIT})`,
        )
      })
    })

    describe('called by the onRamp', () => {
      beforeEach(async () => {
        account = onRamp
        sender = await account.getAddress()
        depositor = await account.getAddress()

        await token.connect(account).approve(pool.address, DEPOSIT)
        const allowance = await token.allowance(depositor, pool.address)
        await expect(allowance).to.equal(DEPOSIT)
      })

      describe('when the onRamp is not set yet', () => {
        it('Fails with a permissions error', async () => {
          await evmRevert(
            pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('Once the onRamp is set', async () => {
        beforeEach(async () => {
          await pool
            .connect(roles.defaultAccount)
            .setOnRamp(await onRamp.getAddress(), true)
          expect(await pool.isOnRamp(await onRamp.getAddress())).to.equal(true)
        })

        it('tokens can be locked', async () => {
          await expect(pool.connect(account).lockOrBurn(depositor, DEPOSIT))
            .to.emit(pool, 'Locked')
            .withArgs(sender, depositor, DEPOSIT)

          const poolBalance = await token.balanceOf(pool.address)
          await expect(poolBalance).to.equal(DEPOSIT)
        })

        it("can't store when paused", async () => {
          await pool.connect(roles.defaultAccount).pause()

          await evmRevert(
            pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            'Pausable: paused',
          )
        })

        it('fails if the amount exceeds the token limit', async () => {
          await pool
            .connect(roles.defaultAccount)
            .setLockOrBurnBucket(1, 1, true)

          await evmRevert(
            pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            `ExceedsTokenLimit(1, ${DEPOSIT})`,
          )
        })
      })
    })

    it('fails when called by an unknown account', async () => {
      const account = roles.stranger
      const depositor = await account.getAddress()
      await token.connect(account).approve(pool.address, DEPOSIT)
      await evmRevert(
        pool.connect(account).lockOrBurn(depositor, DEPOSIT),
        `PermissionsError()`,
      )
    })
  })

  describe('#releaseOrMint', () => {
    let account: Signer
    let sender: string
    let depositor: string
    let recipient: string

    describe('called by the offRamp', () => {
      beforeEach(async () => {
        account = offRamp
        sender = await account.getAddress()
        depositor = await account.getAddress()
        recipient = await roles.stranger.getAddress()

        // Store using the onRamp first
        await pool
          .connect(roles.defaultAccount)
          .setOnRamp(await onRamp.getAddress(), true)
        expect(await pool.isOnRamp(await onRamp.getAddress())).to.equal(true)

        await token.connect(onRamp).approve(pool.address, DEPOSIT)
        await pool
          .connect(onRamp)
          .lockOrBurn(await onRamp.getAddress(), DEPOSIT)
        // // Check pool balance
        const poolBalance = await token.balanceOf(pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      describe('when the offRamp is not set yet', () => {
        it('Fails with a permissions error', async () => {
          await evmRevert(
            pool.connect(account).releaseOrMint(depositor, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('once the offRamp is set', () => {
        beforeEach(async () => {
          await pool
            .connect(roles.defaultAccount)
            .setOffRamp(await offRamp.getAddress(), true)
          expect(await pool.isOffRamp(await offRamp.getAddress())).to.equal(
            true,
          )
        })

        it('can release tokens', async () => {
          await expect(pool.connect(account).releaseOrMint(recipient, DEPOSIT))
            .to.emit(pool, 'Released')
            .withArgs(sender, recipient, DEPOSIT)
          const recipientBalance = await token.balanceOf(recipient)
          await expect(recipientBalance).to.equal(DEPOSIT)
        })

        it("can't release tokens if paused", async () => {
          await pool.connect(roles.defaultAccount).pause()

          await evmRevert(
            pool.connect(account).releaseOrMint(recipient, DEPOSIT),
            'Pausable: paused',
          )
        })

        it('fails if the amount exceeds the token limit', async () => {
          await pool
            .connect(roles.defaultAccount)
            .setReleaseOrMintBucket(1, 1, true)

          await evmRevert(
            pool.connect(account).releaseOrMint(recipient, DEPOSIT),
            `ExceedsTokenLimit(1, ${DEPOSIT})`,
          )
        })
      })
    })

    describe('called by the owner', () => {
      beforeEach(async () => {
        account = roles.defaultAccount
        sender = await account.getAddress()
        depositor = await account.getAddress()
        recipient = await roles.stranger.getAddress()

        await token.connect(account).approve(pool.address, DEPOSIT)
        await pool.connect(account).lockOrBurn(depositor, DEPOSIT)
        const poolBalance = await token.balanceOf(pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      it('can release tokens', async () => {
        await expect(pool.connect(account).releaseOrMint(recipient, DEPOSIT))
          .to.emit(pool, 'Released')
          .withArgs(sender, recipient, DEPOSIT)
        const recipientBalance = await token.balanceOf(recipient)
        await expect(recipientBalance).to.equal(DEPOSIT)
      })

      it("can't release tokens if paused", async () => {
        await pool.connect(roles.defaultAccount).pause()

        await evmRevert(
          pool.connect(account).releaseOrMint(recipient, DEPOSIT),
          'Pausable: paused',
        )
      })

      it('fails if the amount exceeds the token limit', async () => {
        await pool
          .connect(roles.defaultAccount)
          .setReleaseOrMintBucket(1, 1, true)

        await evmRevert(
          pool.connect(account).releaseOrMint(recipient, DEPOSIT),
          `ExceedsTokenLimit(1, ${DEPOSIT})`,
        )
      })
    })

    it('fails when called by an unknown account', async () => {
      const account = roles.stranger
      const depositor = await account.getAddress()
      await evmRevert(
        pool.connect(account).releaseOrMint(depositor, DEPOSIT),
        `PermissionsError()`,
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
})
