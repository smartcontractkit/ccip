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
      'addPool',
      'removePool',
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

  describe('#removePool', () => {
    let newToken: MockERC20
    let newPool: MockPool
    beforeEach(async () => {
      newToken = <MockERC20>(
        await deployContract(roles.defaultAccount, MockERC20Artifact, [
          '6',
          '6',
          await roles.defaultAccount.getAddress(),
          100,
        ])
      )
      newPool = <MockPool>(
        await deployContract(roles.defaultAccount, MockPoolArtifact, [6])
      )
    })
    describe('failure', () => {
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          registry
            .connect(roles.stranger)
            .removePool(tokens[2].address, pools[2].address),
          'Only callable by owner',
        )
      })
      it('fails when there are no pools', async () => {
        let newRegistry: TokenPoolRegistry = <TokenPoolRegistry>(
          await deployContract(roles.defaultAccount, RegistryArtifact, [[], []])
        )
        await evmRevert(
          newRegistry
            .connect(roles.defaultAccount)
            .removePool(newPool.address, newToken.address),
          'NoPools()',
        )
      })
      it('fails when the pool does not exist', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .removePool(newToken.address, newPool.address),
          'PoolDoesNotExist()',
        )
      })
      it('fails when the token doesnt match the configuration', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .removePool(tokens[2].address, pools[3].address),
          'TokenPoolMismatch()',
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      let theToken: string
      let thePool: string

      beforeEach(async () => {
        theToken = tokens[2].address
        thePool = pools[2].address
        tx = await registry
          .connect(roles.defaultAccount)
          .removePool(theToken, thePool)
      })
      it('removes the token from the mapping', async () => {
        const response = await registry.getPool(theToken)
        expect(response).to.equal(constants.AddressZero)
      })
      it('removes the token from the list', async () => {
        const response = await registry.getPoolTokens()
        expect(response).to.not.contain(theToken)
      })
      it('emits an event', async () => {
        await expect(tx)
          .to.emit(registry, 'PoolRemoved')
          .withArgs(theToken, thePool)
      })
      it('sets the isPool flag to false', async () => {
        const response = await registry.isPool(thePool)
        expect(response).to.be.false
      })
    })
  })

  describe('#addPool', () => {
    let newToken: MockERC20
    let newPool: MockPool
    beforeEach(async () => {
      newToken = <MockERC20>(
        await deployContract(roles.defaultAccount, MockERC20Artifact, [
          '6',
          '6',
          await roles.defaultAccount.getAddress(),
          100,
        ])
      )
      newPool = <MockPool>(
        await deployContract(roles.defaultAccount, MockPoolArtifact, [6])
      )
    })

    describe('failure', () => {
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          registry
            .connect(roles.stranger)
            .addPool(newToken.address, newPool.address),
          'Only callable by owner',
        )
      })
      it('fails when the pool already exists', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .addPool(tokens[1].address, pools[1].address),
          `PoolAlreadyAdded()`,
        )
      })
      it('fails when the token is zero address', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .addPool(constants.AddressZero, newPool.address),
          `InvalidTokenPoolConfig()`,
        )
      })
      it('fails when the token is a zer address', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .addPool(newToken.address, constants.AddressZero),
          `InvalidTokenPoolConfig()`,
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        tx = await registry
          .connect(roles.defaultAccount)
          .addPool(newToken.address, newPool.address)
      })
      it('adds a new pool to the mapping', async () => {
        const configuredPool = await registry.getPool(newToken.address)
        expect(configuredPool).to.equal(newPool.address)
      })
      it('adds the token to the s_tokenList', async () => {
        const tokenList = await registry.getPoolTokens()
        expect(tokenList).to.contain(newToken.address)
      })
      it('emits an event', async () => {
        await expect(tx)
          .to.emit(registry, 'PoolAdded')
          .withArgs(newToken.address, newPool.address)
      })
      it('sets the configured flag to true', async () => {
        const response = await registry.isPool(newPool.address)
        expect(response).to.be.true
      })
    })
  })
})
