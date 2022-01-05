import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { expect } from 'chai'
import { MockERC20, LockUnlockPool } from '../../../../typechain'
import { BigNumber, Signer } from 'ethers'
import { getUsers } from '../../../test-helpers/setup'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle

const DEPOSIT = BigNumber.from(123)

describe('LockUnlockPool', function () {
  before(async function () {
    const users = await getUsers()
    this.roles = users.roles
  })

  beforeEach(async function () {
    const MockERC20Artifact: Artifact = await hre.artifacts.readArtifact(
      'MockERC20',
    )
    this.token = <MockERC20>(
      await deployContract(this.roles.defaultAccount, MockERC20Artifact, [
        'Test Token',
        'TEST',
        this.roles.defaultAccount.address,
        BigNumber.from(1000000),
      ])
    )

    // For testing onlyRamp
    this.onRamp = this.roles.oracleNode1
    this.offRamp = this.roles.oracleNode2
    this.token
      .connect(this.roles.defaultAccount)
      .mint(this.onRamp.address, BigNumber.from(1000000))
    this.token
      .connect(this.roles.defaultAccount)
      .mint(this.offRamp.address, BigNumber.from(1000000))
    const LockUnlockPoolArtifact: Artifact = await hre.artifacts.readArtifact(
      'LockUnlockPool',
    )
    this.pool = <LockUnlockPool>(
      await deployContract(this.roles.defaultAccount, LockUnlockPoolArtifact, [
        this.token.address,
      ])
    )
  })

  describe('#constructor', () => {
    it('is initialized correctly', async function () {
      expect(await this.pool.getToken()).to.equal(this.token.address)
      expect(await this.pool.owner()).to.equal(
        this.roles.defaultAccount.address,
      )
    })
  })

  describe('#pause', () => {
    it('owner can pause pool', async function () {
      const account = this.roles.defaultAccount
      await expect(this.pool.connect(account).pause())
        .to.emit(this.pool, 'Paused')
        .withArgs(account.address)
    })

    it('unknown account cannot pause pool', async function () {
      await expect(this.pool.connect(this.onRamp).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(this.pool.connect(this.offRamp).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(
        this.pool.connect(this.roles.stranger).pause(),
      ).to.be.revertedWith('Only callable by owner')
    })
  })

  describe('#unpause', () => {
    beforeEach(async function () {
      await this.pool.connect(this.roles.defaultAccount).pause()
    })

    it('owner can unpause pool', async function () {
      const account = this.roles.defaultAccount
      await expect(this.pool.connect(account).unpause())
        .to.emit(this.pool, 'Unpaused')
        .withArgs(account.address)
    })

    it('unknown account cannot unpause pool', async function () {
      await expect(this.pool.connect(this.onRamp).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
      await expect(
        this.pool.connect(this.offRamp).unpause(),
      ).to.be.revertedWith('Only callable by owner')
      await expect(
        this.pool.connect(this.roles.stranger).unpause(),
      ).to.be.revertedWith('Only callable by owner')
    })
  })

  describe('#lockOrBurn', () => {
    let account: Signer
    let sender: string
    let depositor: string

    describe('called by the owner', () => {
      beforeEach(async function () {
        account = this.roles.defaultAccount
        sender = await account.getAddress()
        depositor = await account.getAddress()

        await this.token.connect(account).approve(this.pool.address, DEPOSIT)
        const allowance = await this.token.allowance(
          depositor,
          this.pool.address,
        )
        await expect(allowance).to.equal(DEPOSIT)
      })

      it('can lock tokens', async function () {
        await expect(this.pool.connect(account).lockOrBurn(depositor, DEPOSIT))
          .to.emit(this.pool, 'Locked')
          .withArgs(sender, depositor, DEPOSIT)

        const poolBalance = await this.token.balanceOf(this.pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      it("can't store when paused", async function () {
        await this.pool.connect(this.roles.defaultAccount).pause()

        await evmRevert(
          this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
          'Pausable: paused',
        )
      })
    })

    describe('called by the onRamp', () => {
      beforeEach(async function () {
        account = this.onRamp
        sender = await account.getAddress()
        depositor = await account.getAddress()

        await this.token.connect(account).approve(this.pool.address, DEPOSIT)
        const allowance = await this.token.allowance(
          depositor,
          this.pool.address,
        )
        await expect(allowance).to.equal(DEPOSIT)
      })

      describe('when the onRamp is not set yet', () => {
        it('Fails with a permissions error', async function () {
          await evmRevert(
            this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('Once the onRamp is set', async function () {
        beforeEach(async function () {
          await this.pool
            .connect(this.roles.defaultAccount)
            .setOnRamp(this.onRamp.address, true)
          expect(await this.pool.isOnRamp(this.onRamp.address)).to.equal(true)
        })

        it('tokens can be locked', async function () {
          await expect(
            this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
          )
            .to.emit(this.pool, 'Locked')
            .withArgs(sender, depositor, DEPOSIT)

          const poolBalance = await this.token.balanceOf(this.pool.address)
          await expect(poolBalance).to.equal(DEPOSIT)
        })

        it("can't store when paused", async function () {
          await this.pool.connect(this.roles.defaultAccount).pause()

          await evmRevert(
            this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
            'Pausable: paused',
          )
        })
      })
    })

    it('fails when called by an unknown account', async function () {
      const account = this.roles.stranger
      const depositor = account.address
      await this.token.connect(account).approve(this.pool.address, DEPOSIT)
      await evmRevert(
        this.pool.connect(account).lockOrBurn(depositor, DEPOSIT),
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
      beforeEach(async function () {
        account = this.offRamp
        sender = await account.getAddress()
        depositor = await account.getAddress()
        recipient = this.roles.stranger.address

        // Store using the onRamp first
        await this.pool
          .connect(this.roles.defaultAccount)
          .setOnRamp(this.onRamp.address, true)
        expect(await this.pool.isOnRamp(this.onRamp.address)).to.equal(true)

        await this.token
          .connect(this.onRamp)
          .approve(this.pool.address, DEPOSIT)
        await this.pool
          .connect(this.onRamp)
          .lockOrBurn(this.onRamp.address, DEPOSIT)
        // // Check pool balance
        const poolBalance = await this.token.balanceOf(this.pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      describe('when the offRamp is not set yet', () => {
        it('Fails with a permissions error', async function () {
          await evmRevert(
            this.pool.connect(account).releaseOrMint(depositor, DEPOSIT),
            'PermissionsError()',
          )
        })
      })

      describe('once the offRamp is set', () => {
        beforeEach(async function () {
          await this.pool
            .connect(this.roles.defaultAccount)
            .setOffRamp(this.offRamp.address, true)
          expect(await this.pool.isOffRamp(this.offRamp.address)).to.equal(true)
        })

        it('can release tokens', async function () {
          await expect(
            this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
          )
            .to.emit(this.pool, 'Released')
            .withArgs(sender, recipient, DEPOSIT)
          const recipientBalance = await this.token.balanceOf(recipient)
          await expect(recipientBalance).to.equal(DEPOSIT)
        })

        it("can't release tokens if paused", async function () {
          await this.pool.connect(this.roles.defaultAccount).pause()

          await evmRevert(
            this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
            'Pausable: paused',
          )
        })
      })
    })

    describe('called by the owner', () => {
      beforeEach(async function () {
        account = this.roles.defaultAccount
        sender = await account.getAddress()
        depositor = await account.getAddress()
        recipient = this.roles.stranger.address

        await this.token.connect(account).approve(this.pool.address, DEPOSIT)
        await this.pool.connect(account).lockOrBurn(depositor, DEPOSIT)
        const poolBalance = await this.token.balanceOf(this.pool.address)
        await expect(poolBalance).to.equal(DEPOSIT)
      })

      it('can release tokens', async function () {
        await expect(
          this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
        )
          .to.emit(this.pool, 'Released')
          .withArgs(sender, recipient, DEPOSIT)
        const recipientBalance = await this.token.balanceOf(recipient)
        await expect(recipientBalance).to.equal(DEPOSIT)
      })

      it("can't release tokens if paused", async function () {
        await this.pool.connect(this.roles.defaultAccount).pause()

        await evmRevert(
          this.pool.connect(account).releaseOrMint(recipient, DEPOSIT),
          'Pausable: paused',
        )
      })
    })

    it('fails when called by an unknown account', async function () {
      const account = this.roles.stranger
      const depositor = account.address
      await evmRevert(
        this.pool.connect(account).releaseOrMint(depositor, DEPOSIT),
        `PermissionsError()`,
      )
    })
  })

  describe('#setOnRamp', () => {
    let expectedRamp: string
    beforeEach(async function () {
      expectedRamp = this.roles.oracleNode3.address
    })

    it('sets the on ramp when called by the owner', async function () {
      await this.pool
        .connect(this.roles.defaultAccount)
        .setOnRamp(expectedRamp, true)
      expect(await this.pool.isOnRamp(expectedRamp)).to.equal(true)
    })

    it('reverts when called by any other account', async function () {
      await evmRevert(
        this.pool.connect(this.roles.stranger).setOnRamp(expectedRamp, true),
        'Only callable by owner',
      )
    })
  })

  describe('#setOffRamp', () => {
    let expectedRamp: string
    beforeEach(async function () {
      expectedRamp = this.roles.oracleNode3.address
    })

    it('sets the on ramp when called by the owner', async function () {
      await this.pool
        .connect(this.roles.defaultAccount)
        .setOffRamp(expectedRamp, true)
      expect(await this.pool.isOffRamp(expectedRamp)).to.equal(true)
    })

    it('reverts when called by any other account', async function () {
      await evmRevert(
        this.pool.connect(this.roles.stranger).setOffRamp(expectedRamp, true),
        'Only callable by owner',
      )
    })
  })
})
