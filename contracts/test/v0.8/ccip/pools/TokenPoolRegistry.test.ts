import hre from 'hardhat'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { Artifact } from 'hardhat/types'
import { TokenPoolRegistry, MockERC20, MockPool } from '../../../../typechain'
import { publicAbi } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import { constants, ContractTransaction } from 'ethers'

const { deployContract } = hre.waffle

let roles: Roles

let numberOfPools: number
let registry: TokenPoolRegistry
let tokens: Array<MockERC20>
let tokensAddresses: Array<string>
let pools: Array<MockPool>
let poolsAddresses: Array<string>

let RegistryArtifact: Artifact
let MockERC20Artifact: Artifact
let MockPoolArtifact: Artifact

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('TokenPoolRegistry', () => {
  beforeEach(async () => {
    tokens = new Array<MockERC20>()
    pools = new Array<MockPool>()
    MockERC20Artifact = await hre.artifacts.readArtifact('MockERC20')
    MockPoolArtifact = await hre.artifacts.readArtifact('MockPool')

    numberOfPools = 5
    for (let i = 0; i < numberOfPools; i++) {
      tokens.push(
        <MockERC20>(
          await deployContract(roles.defaultAccount, MockERC20Artifact, [
            i.toString(),
            i.toString(),
            await roles.defaultAccount.getAddress(),
            100,
          ])
        ),
      )
      pools.push(
        <MockPool>(
          await deployContract(roles.defaultAccount, MockPoolArtifact, [i])
        ),
      )
    }
    tokensAddresses = tokens.map((t) => t.address)
    poolsAddresses = pools.map((p) => p.address)

    RegistryArtifact = await hre.artifacts.readArtifact('TokenPoolRegistry')
    registry = <TokenPoolRegistry>(
      await deployContract(roles.defaultAccount, RegistryArtifact, [
        tokensAddresses,
        poolsAddresses,
      ])
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(registry, [
      'setPools',
      'getPool',
      'isPool',
      'getPoolTokens',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
    ])
  })

  describe('#constructor', () => {
    it('maps the tokens to the pools', async () => {
      for (let i = 0; i < numberOfPools; i++) {
        expect(await registry.getPool(tokensAddresses[i])).to.equal(
          poolsAddresses[i],
        )
      }
    })
    it('sets the source tokens list', async () => {
      expect(await registry.getPoolTokens()).to.deep.equal(tokensAddresses)
    })
    it('sets that each token pool is configured', async () => {
      for (let i = 0; i < numberOfPools; i++) {
        expect(await registry.isPool(poolsAddresses[i])).to.be.true
      }
    })
    it('sets the owner', async () => {
      expect(await registry.owner()).to.equal(
        await roles.defaultAccount.getAddress(),
      )
    })
  })

  describe('#setPools', async () => {
    let newTokens: Array<MockERC20>
    let newPools: Array<MockPool>
    describe('failure', () => {
      beforeEach(async () => {
        newTokens = new Array<MockERC20>()
        newPools = new Array<MockPool>()

        newTokens.push(
          <MockERC20>(
            await deployContract(roles.defaultAccount, MockERC20Artifact, [
              '6',
              '6',
              await roles.defaultAccount.getAddress(),
              100,
            ])
          ),
        )

        newPools.push(
          <MockPool>(
            await deployContract(roles.defaultAccount, MockPoolArtifact, [6])
          ),
        )
      })

      it('fails when called by the non-owner', async () => {
        await evmRevert(
          registry.connect(roles.stranger).setPools(
            newTokens.map((nt) => nt.address),
            newPools.map((np) => np.address),
          ),
          'Only callable by owner',
        )
      })
      it('fails when the source tokens length and pools length is 0', async () => {
        await evmRevert(
          registry.connect(roles.defaultAccount).setPools([], []),
          'InvalidTokenPoolConfig()',
        )
      })
      it('fails if the length of source tokens is different to the length of pools', async () => {
        await evmRevert(
          registry.connect(roles.defaultAccount).setPools(
            newTokens.map((nt) => nt.address),
            [],
          ),
          'InvalidTokenPoolConfig()',
        )
        await evmRevert(
          registry.connect(roles.defaultAccount).setPools(
            [],
            newPools.map((np) => np.address),
          ),
          'InvalidTokenPoolConfig()',
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        newTokens = new Array<MockERC20>()
        newPools = new Array<MockPool>()

        newTokens.push(
          <MockERC20>(
            await deployContract(roles.defaultAccount, MockERC20Artifact, [
              '6',
              '6',
              await roles.defaultAccount.getAddress(),
              100,
            ])
          ),
        )

        newPools.push(
          <MockPool>(
            await deployContract(roles.defaultAccount, MockPoolArtifact, [6])
          ),
        )

        tx = await registry.connect(roles.defaultAccount).setPools(
          newTokens.map((nt) => nt.address),
          newPools.map((np) => np.address),
        )
      })

      it('removes the old source tokens and pools from all fields', async () => {
        for (let i = 0; i < numberOfPools; i++) {
          expect(await registry.getPool(tokensAddresses[i])).to.equal(
            constants.AddressZero,
          )
          expect(await registry.isPool(poolsAddresses[i])).to.be.false
        }
        expect((await registry.getPoolTokens()).length).to.equal(1)
      })

      it('sets the new tokens and pools in all fields', async () => {
        expect(await registry.getPoolTokens()).to.deep.equal(
          newTokens.map((nt) => nt.address),
        )
        expect(await registry.getPool(newTokens[0].address)).to.equal(
          newPools[0].address,
        )
        expect(await registry.isPool(newPools[0].address)).to.be.true
      })

      it('emits a PoolsSet event', async () => {
        await expect(tx)
          .to.emit(registry, 'PoolsSet')
          .withArgs(
            newTokens.map((nt) => nt.address),
            newPools.map((np) => np.address),
          )
      })
    })
  })
})
