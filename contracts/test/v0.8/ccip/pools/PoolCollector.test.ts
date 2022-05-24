import { expect } from 'chai'
import { BigNumber } from 'ethers'
import hre, { ethers } from 'hardhat'
import { Artifact } from 'hardhat/types'
import { MockERC20, MockOnRamp, TokenPoolHelper } from '../../../../typechain'
import { PoolCollectorHelper } from '../../../../typechain/PoolCollectorHelper'
import { CCIPMessagePayload } from '../../../test-helpers/ccip/ccip'
import { evmRevert } from '../../../test-helpers/matchers'
import { getUsers, Roles } from '../../../test-helpers/setup'

const { deployContract } = hre.waffle

let roles: Roles

let poolCollector: PoolCollectorHelper
let onRamp: MockOnRamp
let tokens: MockERC20[]
let pool: TokenPoolHelper

let PoolCollectorArtifact: Artifact
let MockOnRampArtifact: Artifact
let MockERC20Artifact: Artifact
let TokenPoolArtifact: Artifact

let payload: CCIPMessagePayload

const destinationChainId = 99
const amounts = [10, 11, 12]
const feeAmount = 1

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('PoolCollector', () => {
  beforeEach(async () => {
    const fillerNumber = BigNumber.from('10000000000000000')

    PoolCollectorArtifact = await hre.artifacts.readArtifact(
      'PoolCollectorHelper',
    )
    MockOnRampArtifact = await hre.artifacts.readArtifact('MockOnRamp')
    MockERC20Artifact = await hre.artifacts.readArtifact('MockERC20')
    TokenPoolArtifact = await hre.artifacts.readArtifact('TokenPoolHelper')

    poolCollector = <PoolCollectorHelper>(
      await deployContract(roles.defaultAccount, PoolCollectorArtifact, [])
    )

    tokens = []
    for (let i = 0; i < 3; i++) {
      const newToken = <MockERC20>(
        await deployContract(roles.defaultAccount, MockERC20Artifact, [
          `token${i}`,
          `${i}`,
          await roles.defaultAccount.getAddress(),
          fillerNumber,
        ])
      )
      tokens.push(newToken)
    }
    pool = <TokenPoolHelper>(
      await deployContract(roles.defaultAccount, TokenPoolArtifact, [
        tokens[0].address,
        fillerNumber,
        fillerNumber,
        fillerNumber,
        fillerNumber,
      ])
    )
    onRamp = <MockOnRamp>(
      await deployContract(roles.defaultAccount, MockOnRampArtifact, [
        1,
        pool.address,
        destinationChainId,
        feeAmount,
      ])
    )
  })

  describe('collectTokens', () => {
    describe('failure', () => {
      beforeEach(async () => {
        payload = {
          tokens: tokens.map((t) => t.address),
          amounts: amounts.map((a) => BigNumber.from(a)),
          destinationChainId: BigNumber.from(destinationChainId),
          receiver: await roles.consumer.getAddress(),
          executor: ethers.constants.AddressZero,
          data: ethers.constants.HashZero,
        }
        for (let i = 0; i < amounts.length; i++) {
          await tokens[i]
            .connect(roles.defaultAccount)
            .approve(poolCollector.address, BigNumber.from(amounts[i]))
        }
      })
      it('fails if a token is not supported', async () => {
        const newOnRamp = <MockOnRamp>(
          await deployContract(roles.defaultAccount, MockOnRampArtifact, [
            1,
            ethers.constants.AddressZero,
            destinationChainId,
            feeAmount,
          ])
        )
        await evmRevert(
          poolCollector.collectTokens(newOnRamp.address, payload),
          `UnsupportedToken("${tokens[0].address}")`,
        )
      })
    })
    describe('success', () => {
      beforeEach(async () => {
        payload = {
          tokens: tokens.map((t) => t.address),
          amounts: amounts.map((a) => BigNumber.from(a)),
          destinationChainId: BigNumber.from(destinationChainId),
          receiver: await roles.consumer.getAddress(),
          executor: ethers.constants.AddressZero,
          data: ethers.constants.HashZero,
        }
        for (let i = 0; i < amounts.length; i++) {
          await tokens[i]
            .connect(roles.defaultAccount)
            .approve(poolCollector.address, BigNumber.from(amounts[i]))
        }
      })

      it('calls getRequiredFee on the onRamp', async () => {
        const tx = await poolCollector.collectTokens(onRamp.address, payload)
        await expect(tx)
          .to.emit(onRamp, 'GetRequiredFee')
          .withArgs(tokens[0].address)
      })
      it('transfers the feeToken fee amount to this contract', async () => {
        await poolCollector.collectTokens(onRamp.address, payload)
        const collectorBalance = await tokens[0].balanceOf(
          poolCollector.address,
        )
        expect(collectorBalance).to.equal(feeAmount)
      })
      it('alters the payload amount of fee token correctly', async () => {
        await poolCollector.collectTokens(onRamp.address, payload)
        const amountAfterFee = BigNumber.from(payload.amounts[0]).sub(feeAmount)
        const balance = await tokens[0].balanceOf(pool.address)
        expect(balance).to.equal(amountAfterFee)
      })
      it('calls getTokenPool on the onRamp for each token', async () => {
        const tx = await poolCollector.collectTokens(onRamp.address, payload)
        for (let i = 0; i < tokens.length; i++) {
          await expect(tx)
            .to.emit(onRamp, 'GetTokenPool')
            .withArgs(tokens[i].address)
        }
      })
      it('transfers each of the payload tokens to a token pool', async () => {
        await poolCollector.collectTokens(onRamp.address, payload)
        for (let i = 0; i < tokens.length; i++) {
          const balance = await tokens[i].balanceOf(pool.address)
          let amount = BigNumber.from(amounts[i])
          if (i === 0) {
            amount = amount.sub(feeAmount)
          }
          expect(balance).to.equal(amount)
        }
      })
    })
  })
})
