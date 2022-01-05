import hre from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { TokenLimitsHelper } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { increaseTime5Minutes } from '../../../test-helpers/helpers'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle
let roles: Roles
let TokenLimitsHelperArtifact: Artifact
let helper: TokenLimitsHelper

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('TokenLimits library', () => {
  beforeEach(async () => {
    TokenLimitsHelperArtifact = await hre.artifacts.readArtifact(
      'TokenLimitsHelper',
    )

    helper = <TokenLimitsHelper>(
      await deployContract(roles.defaultAccount, TokenLimitsHelperArtifact)
    )
  })

  describe('#constructTokenBucket', () => {
    let rate: BigNumber
    let capacity: BigNumber

    beforeEach(async () => {
      rate = BigNumber.from('10')
      capacity = BigNumber.from('100')
    })

    it('returns an empty bucket', async () => {
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, false)
      const bucket = await helper.s_bucket()

      expect(bucket.rate).to.equal(rate)
      expect(bucket.capacity).to.equal(capacity)
      expect(bucket.tokens).to.equal(BigNumber.from('0'))
    })

    it('returns a full bucket', async () => {
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, true)
      const bucket = await helper.s_bucket()

      expect(bucket.rate).to.equal(rate)
      expect(bucket.capacity).to.equal(capacity)
      expect(bucket.tokens).to.equal(capacity)
    })
  })

  describe('#update', () => {
    let rate: BigNumber
    let capacity: BigNumber

    beforeEach(async () => {
      rate = BigNumber.from('1')
      capacity = BigNumber.from('1000000')
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, false)
    })

    it('increases the tokens in the bucket over time at expected rate', async () => {
      await increaseTime5Minutes(hre.ethers.provider)
      await helper.connect(roles.defaultAccount).update()
      const bucket = await helper.s_bucket()
      const threehundred = BigNumber.from(300)
      expect(bucket.tokens)
        .to.be.at.least(threehundred)
        .but.be.below(threehundred.add(5))
    })

    it('does not update the timestamp if the bucket is already full', async () => {
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, true)
      const initialBucket = await helper.s_bucket()
      await helper.connect(roles.defaultAccount).update()
      const updatedBucket = await helper.s_bucket()
      expect(initialBucket.lastUpdated).to.equal(updatedBucket.lastUpdated)
    })

    describe('altering the capacity', () => {
      let newCapacity: BigNumber

      beforeEach(async () => {
        await helper
          .connect(roles.defaultAccount)
          .constructTokenBucket(rate, capacity, true)
      })

      it('reverts if the capacity is reduced to below the current tokens amount', async () => {
        newCapacity = capacity.div(2)
        await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        await evmRevert(
          helper.connect(roles.defaultAccount).update(),
          'BucketOverfilled()',
        )
      })

      it('works correctly when capacity increased', async () => {
        newCapacity = capacity.mul(2)
        const initialBucket = await helper.s_bucket()
        await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        await helper.connect(roles.defaultAccount).update()
        const updatedBucket = await helper.s_bucket()
        expect(updatedBucket.lastUpdated).to.be.gt(initialBucket.lastUpdated)
        expect(updatedBucket.tokens).to.be.gt(capacity)
      })
    })
  })

  describe('#remove', () => {
    let rate: BigNumber
    let capacity: BigNumber

    beforeEach(async () => {
      rate = BigNumber.from('1')
      capacity = BigNumber.from('1000000')
      await helper
        .connect(roles.defaultAccount)
        .constructTokenBucket(rate, capacity, true)
    })

    it('removes from bucket', async () => {
      const removeAmount = BigNumber.from('100')
      const tx = await helper.connect(roles.defaultAccount).remove(removeAmount)
      await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(true)
      const bucket = await helper.s_bucket()
      expect(bucket.tokens).to.equal(bucket.capacity.sub(removeAmount))
    })

    it('does not remove if token amount greater than capacity', async () => {
      const removeAmount = capacity.add(1)
      const tx = await helper.connect(roles.defaultAccount).remove(removeAmount)
      await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(false)
      const bucket = await helper.s_bucket()
      expect(bucket.tokens).to.equal(bucket.capacity)
    })

    it('does not remove if token amount greater than tokens in bucket', async () => {
      const removeAmount = capacity.sub(1)
      await helper.connect(roles.defaultAccount).remove(removeAmount)
      const tx = await helper.connect(roles.defaultAccount).remove(removeAmount)
      await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(false)
    })

    describe('altering the capacity', () => {
      let newCapacity: BigNumber

      beforeEach(async () => {
        await helper
          .connect(roles.defaultAccount)
          .constructTokenBucket(rate, capacity, true)
      })

      describe('when the capacity is reduced', () => {
        beforeEach(async () => {
          newCapacity = capacity.div(2)
          await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        })

        it('fails when the tokens amount is greater than the capacity', async () => {
          const removeAmount = BigNumber.from('100')
          await evmRevert(
            helper.connect(roles.defaultAccount).remove(removeAmount),
            'BucketOverfilled()',
          )
        })
      })

      describe('when the capacity is increased', () => {
        let initialBucket: any

        beforeEach(async () => {
          initialBucket = await helper.s_bucket()
          newCapacity = capacity.mul(2)
          await helper.connect(roles.defaultAccount).alterCapacity(newCapacity)
        })

        it('removes from the bucket', async () => {
          const removeAmount = BigNumber.from('100')
          const tx = await helper
            .connect(roles.defaultAccount)
            .remove(removeAmount)
          await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(true)
          const bucket = await helper.s_bucket()
          // Depending on provider time, the bucket might update itself
          expect(bucket.tokens)
            .to.be.at.least(initialBucket.tokens.sub(removeAmount))
            .but.be.below(initialBucket.tokens.sub(removeAmount).add(5))
        })

        it('does not remove if the token amount greater than tokens in the bucket', async () => {
          const removeAmount = newCapacity.sub(1)
          await helper.connect(roles.defaultAccount).remove(removeAmount)
          const tx = await helper
            .connect(roles.defaultAccount)
            .remove(removeAmount)
          await expect(tx).to.emit(helper, 'RemovalSuccess').withArgs(false)
        })
      })
    })
  })
})
