import hre from 'hardhat'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { HealthCheckerHelper, MockAFN } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import { constants } from 'ethers'
import { publicAbi } from '../../../test-helpers/helpers'

const { deployContract } = hre.waffle
let roles: Roles
let HealthCheckerArtifact: Artifact
let MockAFNArtifact: Artifact
let healthChecker: HealthCheckerHelper
let afn: MockAFN

let maxTimeBetweenAFNSignals: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('HealthChecker', () => {
  beforeEach(async () => {
    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    HealthCheckerArtifact = await hre.artifacts.readArtifact(
      'HealthCheckerHelper',
    )
    maxTimeBetweenAFNSignals = BigNumber.from(60).mul(60) // 1 hour

    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    healthChecker = <HealthCheckerHelper>(
      await deployContract(roles.defaultAccount, HealthCheckerArtifact, [
        afn.address,
        maxTimeBetweenAFNSignals,
      ])
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(healthChecker, [
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
      'whenHealthyFunction',
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
    it('sets the correct storage vars', async () => {
      expect(await healthChecker.getAFN()).to.equal(afn.address)
      expect(await healthChecker.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        maxTimeBetweenAFNSignals,
      )
    })

    it('fails if zero values are used', async () => {
      // Zero address afn
      await evmRevert(
        deployContract(roles.defaultAccount, HealthCheckerArtifact, [
          constants.AddressZero,
          maxTimeBetweenAFNSignals,
        ]),
        `BadHealthConfig()`,
      )
      // Zero time
      await evmRevert(
        deployContract(roles.defaultAccount, HealthCheckerArtifact, [
          afn.address,
          0,
        ]),
        `BadHealthConfig()`,
      )
    })
  })

  describe('#pause', () => {
    it('owner can pause healthChecker', async () => {
      const account = roles.defaultAccount
      await expect(healthChecker.connect(account).pause())
        .to.emit(healthChecker, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async function () {
      const account = roles.stranger
      await expect(healthChecker.connect(account).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await healthChecker.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause healthChecker', async () => {
      const account = roles.defaultAccount
      await expect(healthChecker.connect(account).unpause())
        .to.emit(healthChecker, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async function () {
      const account = roles.stranger
      await expect(healthChecker.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
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
        healthChecker.connect(roles.stranger).setAFN(newAFN.address),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('fails with zero value', async () => {
      await evmRevert(
        healthChecker
          .connect(roles.defaultAccount)
          .setAFN(constants.AddressZero),
        `BadHealthConfig()`,
      )
    })

    it('sets the new AFN', async () => {
      const tx = await healthChecker
        .connect(roles.defaultAccount)
        .setAFN(newAFN.address)
      expect(await healthChecker.getAFN()).to.equal(newAFN.address)
      await expect(tx)
        .to.emit(healthChecker, 'AFNSet')
        .withArgs(afn.address, newAFN.address)
    })
  })

  describe('#setMaxTimeWithoutAFNSignal', () => {
    let newTime: BigNumber

    beforeEach(async () => {
      newTime = maxTimeBetweenAFNSignals.mul(2)
    })

    it('only callable by owner', async () => {
      await expect(
        healthChecker
          .connect(roles.stranger)
          .setMaxSecondsWithoutAFNHeartbeat(newTime),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('fails with zero value', async () => {
      await evmRevert(
        healthChecker
          .connect(roles.defaultAccount)
          .setMaxSecondsWithoutAFNHeartbeat(0),
        `BadHealthConfig()`,
      )
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await healthChecker
        .connect(roles.defaultAccount)
        .setMaxSecondsWithoutAFNHeartbeat(newTime)
      expect(await healthChecker.getMaxSecondsWithoutAFNHeartbeat()).to.equal(
        newTime,
      )
      await expect(tx)
        .to.emit(healthChecker, 'AFNMaxHeartbeatTimeSet')
        .withArgs(maxTimeBetweenAFNSignals, newTime)
    })
  })

  describe('#whenHealthy', () => {
    // Uses HealthCheckerHelper.whenHealthyFunction() to simulate modifier

    it('fails if the afn has emitted a bad signal', async () => {
      await afn.voteBad()
      await evmRevert(healthChecker.whenHealthyFunction(), 'BadAFNSignal()')
    })

    it('fails if the heartbeat is stale', async () => {
      await afn.setTimestamp(1)
      await evmRevert(healthChecker.whenHealthyFunction(), 'StaleAFNHeartbeat')
    })

    it('it does nothing if all is well', async () => {
      await healthChecker.whenHealthyFunction()
    })
  })
})
