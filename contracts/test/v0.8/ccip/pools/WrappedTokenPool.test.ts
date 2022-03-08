import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { WrappedTokenPoolHelper } from '../../../../typechain'
import { BigNumber, ContractTransaction, Signer } from 'ethers'
import { getUsers, Roles } from '../../../test-helpers/setup'
import { evmRevert } from '../../../test-helpers/matchers'
import { publicAbi } from '../../../test-helpers/helpers'

const { deployContract } = hre.waffle

let roles: Roles

const DEPOSIT = BigNumber.from(123)

let onRamp: Signer
let offRamp: Signer
let pool: WrappedTokenPoolHelper

let bucketRate: BigNumber
let bucketCapactiy: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('WrappedTokenPool', () => {
  beforeEach(async () => {
    onRamp = roles.oracleNode1
    offRamp = roles.oracleNode2
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    const WrappedTokenPoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'WrappedTokenPoolHelper',
    )
    pool = <WrappedTokenPoolHelper>(
      await deployContract(roles.defaultAccount, WrappedTokenPoolArtifact, [
        'Test Wrapped Token',
        'TEST',
        bucketRate,
        bucketCapactiy,
        bucketRate,
        bucketCapactiy,
      ])
    )
    await pool
      .connect(roles.defaultAccount)
      .mint(
        await roles.defaultAccount.getAddress(),
        BigNumber.from('10000000000000000'),
      )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(pool, [
      // WrappedTokenPoolHelper
      'mint',
      // WrappedTokenPool
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
      // ERC20
      'name',
      'symbol',
      'decimals',
      'totalSupply',
      'balanceOf',
      'transfer',
      'allowance',
      'approve',
      'transferFrom',
      'increaseAllowance',
      'decreaseAllowance',
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
      expect(await pool.getToken()).to.equal(pool.address)
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

    beforeEach(async () => {
      account = roles.defaultAccount
      sender = await account.getAddress()
    })

    it('fails when called by an unknown account', async () => {
      await evmRevert(
        pool.connect(roles.stranger).lockOrBurn(sender, DEPOSIT),
        `PermissionsError()`,
      )
    })

    describe('called by the owner', () => {
      it('can burn tokens', async () => {
        const startingBalance = await pool.balanceOf(sender)
        await expect(pool.connect(account).lockOrBurn(sender, DEPOSIT))
          .to.emit(pool, 'Burned')
          .withArgs(sender, sender, DEPOSIT)

        const endingBalance = await pool.balanceOf(sender)
        await expect(startingBalance.sub(endingBalance)).to.equal(DEPOSIT)
      })

      it("can't burn when paused", async () => {
        await pool.connect(account).pause()

        await evmRevert(
          pool.connect(account).lockOrBurn(sender, DEPOSIT),
          'Pausable: paused',
        )
      })

      it('fails if the amount exceeds the token limit', async () => {
        await pool.connect(account).setLockOrBurnBucket(1, 1, true)

        await evmRevert(
          pool.connect(account).lockOrBurn(sender, DEPOSIT),
          `ExceedsTokenLimit(1, ${DEPOSIT})`,
        )
      })
    })

    describe('called by the onRamp', () => {
      describe('when the onRamp is not set yet', () => {
        it('Fails with a permissions error', async () => {
          await evmRevert(
            pool.connect(onRamp).lockOrBurn(sender, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('Once the onRamp is set', async () => {
        let onRampAddress: string

        beforeEach(async () => {
          onRampAddress = await onRamp.getAddress()
          await pool
            .connect(roles.defaultAccount)
            .setOnRamp(onRampAddress, true)
          expect(await pool.isOnRamp(onRampAddress)).to.equal(true)
        })

        it('tokens can be burned', async () => {
          await pool
            .connect(roles.defaultAccount)
            .transfer(onRampAddress, DEPOSIT)
          const balanceBefore = await pool.balanceOf(onRampAddress)

          await expect(pool.connect(onRamp).lockOrBurn(onRampAddress, DEPOSIT))
            .to.emit(pool, 'Burned')
            .withArgs(onRampAddress, onRampAddress, DEPOSIT)

          const balanceAfter = await pool.balanceOf(onRampAddress)
          expect(balanceBefore.sub(balanceAfter)).to.equal(DEPOSIT)
        })

        it("can't burn when paused", async () => {
          await pool.connect(roles.defaultAccount).pause()

          await evmRevert(
            pool.connect(onRamp).lockOrBurn(onRampAddress, DEPOSIT),
            'Pausable: paused',
          )
        })

        it('fails if the amount exceeds the token limit', async () => {
          await pool
            .connect(roles.defaultAccount)
            .setLockOrBurnBucket(1, 1, true)

          await evmRevert(
            pool.connect(onRamp).lockOrBurn(onRampAddress, DEPOSIT),
            `ExceedsTokenLimit(1, ${DEPOSIT})`,
          )
        })
      })
    })
  })

  describe('#releaseOrMint', () => {
    let recipient: string

    describe('called by the offRamp', () => {
      let offRampAddress: string
      beforeEach(async () => {
        recipient = await roles.stranger.getAddress()
        offRampAddress = await offRamp.getAddress()
      })

      describe('when the offRamp is not set yet', () => {
        it('Fails with a permissions error', async () => {
          await evmRevert(
            pool.connect(offRamp).releaseOrMint(recipient, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('once the offRamp is set', () => {
        beforeEach(async () => {
          await pool
            .connect(roles.defaultAccount)
            .setOffRamp(offRampAddress, true)
          expect(await pool.isOffRamp(offRampAddress)).to.equal(true)
        })

        it('can mint tokens', async () => {
          await expect(pool.connect(offRamp).releaseOrMint(recipient, DEPOSIT))
            .to.emit(pool, 'Minted')
            .withArgs(offRampAddress, recipient, DEPOSIT)
          const recipientBalance = await pool.balanceOf(recipient)
          await expect(recipientBalance).to.equal(DEPOSIT)
        })

        it("can't mint tokens if paused", async () => {
          await pool.connect(roles.defaultAccount).pause()

          await evmRevert(
            pool.connect(offRamp).releaseOrMint(recipient, DEPOSIT),
            'Pausable: paused',
          )
        })

        it('fails if the amount exceeds the token limit', async () => {
          await pool
            .connect(roles.defaultAccount)
            .setReleaseOrMintBucket(1, 1, true)

          await evmRevert(
            pool.connect(offRamp).releaseOrMint(recipient, DEPOSIT),
            `ExceedsTokenLimit(1, ${DEPOSIT})`,
          )
        })
      })
    })

    describe('called by the owner', () => {
      let account: string
      beforeEach(async () => {
        account = await roles.defaultAccount.getAddress()
        recipient = await roles.stranger.getAddress()
      })

      it('can mint tokens', async () => {
        await expect(
          pool.connect(roles.defaultAccount).releaseOrMint(recipient, DEPOSIT),
        )
          .to.emit(pool, 'Minted')
          .withArgs(account, recipient, DEPOSIT)
        const recipientBalance = await pool.balanceOf(recipient)
        await expect(recipientBalance).to.equal(DEPOSIT)
      })

      it("can't mint tokens if paused", async () => {
        await pool.connect(roles.defaultAccount).pause()

        await evmRevert(
          pool.connect(roles.defaultAccount).releaseOrMint(recipient, DEPOSIT),
          'Pausable: paused',
        )
      })

      it('fails if the amount exceeds the token limit', async () => {
        await pool
          .connect(roles.defaultAccount)
          .setReleaseOrMintBucket(1, 1, true)

        await evmRevert(
          pool.connect(roles.defaultAccount).releaseOrMint(recipient, DEPOSIT),
          `ExceedsTokenLimit(1, ${DEPOSIT})`,
        )
      })
    })

    it('fails when called by an unknown account', async () => {
      await evmRevert(
        pool
          .connect(roles.stranger)
          .releaseOrMint(await roles.stranger.getAddress(), DEPOSIT),
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
