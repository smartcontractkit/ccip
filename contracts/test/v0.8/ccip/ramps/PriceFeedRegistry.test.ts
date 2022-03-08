import hre from 'hardhat'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { Artifact } from 'hardhat/types'
import {
  PriceFeedRegistry,
  MockERC20,
  MockAggregator,
} from '../../../../typechain'
import { publicAbi } from '../../../test-helpers/helpers'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import { constants, ContractTransaction } from 'ethers'

const { deployContract } = hre.waffle

let roles: Roles

let numberOfFeeds: number
let registry: PriceFeedRegistry
let tokens: Array<MockERC20>
let tokensAddresses: Array<string>
let feeds: Array<MockAggregator>
let feedsAddresses: Array<string>

let RegistryArtifact: Artifact
let MockERC20Artifact: Artifact
let MockFeedArtifact: Artifact

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('PriceFeedRegistry', () => {
  beforeEach(async () => {
    tokens = new Array<MockERC20>()
    feeds = new Array<MockAggregator>()
    MockERC20Artifact = await hre.artifacts.readArtifact('MockERC20')
    MockFeedArtifact = await hre.artifacts.readArtifact('MockAggregator')

    numberOfFeeds = 5
    for (let i = 0; i < numberOfFeeds; i++) {
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
      feeds.push(
        <MockAggregator>(
          await deployContract(roles.defaultAccount, MockFeedArtifact, [])
        ),
      )
    }
    tokensAddresses = tokens.map((t) => t.address)
    feedsAddresses = feeds.map((p) => p.address)

    RegistryArtifact = await hre.artifacts.readArtifact('PriceFeedRegistry')
    registry = <PriceFeedRegistry>(
      await deployContract(roles.defaultAccount, RegistryArtifact, [
        tokensAddresses,
        feedsAddresses,
      ])
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(registry, [
      'setFeeds',
      'getFeed',
      'getFeedTokens',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
    ])
  })

  describe('#constructor', () => {
    it('maps source tokens to the feeds', async () => {
      for (let i = 0; i < numberOfFeeds; i++) {
        expect(await registry.getFeed(tokensAddresses[i])).to.equal(
          feedsAddresses[i],
        )
      }
    })
    it('sets the source tokens list', async () => {
      expect(await registry.getFeedTokens()).to.deep.equal(tokensAddresses)
    })
    it('sets the owner', async () => {
      expect(await registry.owner()).to.equal(
        await roles.defaultAccount.getAddress(),
      )
    })
  })

  describe('#setFeeds', () => {
    let newTokens: Array<MockERC20>
    let newFeeds: Array<MockAggregator>
    describe('failure', () => {
      beforeEach(async () => {
        newTokens = new Array<MockERC20>()
        newFeeds = new Array<MockAggregator>()

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

        newFeeds.push(
          <MockAggregator>(
            await deployContract(roles.defaultAccount, MockFeedArtifact, [])
          ),
        )
      })
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          registry.connect(roles.stranger).setFeeds(
            newTokens.map((nt) => nt.address),
            newFeeds.map((nf) => nf.address),
          ),
          'Only callable by owner',
        )
      })
      it('fails when lengths of params are not equal', async () => {
        await evmRevert(
          registry.connect(roles.defaultAccount).setFeeds(
            newTokens.map((nt) => nt.address),
            [],
          ),
          'InvalidPriceFeedConfig()',
        )
        await evmRevert(
          registry.connect(roles.defaultAccount).setFeeds(
            [],
            newFeeds.map((nf) => nf.address),
          ),
          'InvalidPriceFeedConfig()',
        )
      })
      it('fails when length of params is zero', async () => {
        await evmRevert(
          registry.connect(roles.defaultAccount).setFeeds([], []),
          'InvalidPriceFeedConfig()',
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        newTokens = new Array<MockERC20>()
        newFeeds = new Array<MockAggregator>()

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

        newFeeds.push(
          <MockAggregator>(
            await deployContract(roles.defaultAccount, MockFeedArtifact, [])
          ),
        )

        tx = await registry.connect(roles.defaultAccount).setFeeds(
          newTokens.map((nt) => nt.address),
          newFeeds.map((nf) => nf.address),
        )
      })
      it('removes the old source tokens and feeds', async () => {
        for (let i = 0; i < numberOfFeeds; i++) {
          expect(await registry.getFeed(tokensAddresses[i])).to.equal(
            constants.AddressZero,
          )
        }
        expect((await registry.getFeedTokens()).length).to.equal(1)
      })
      it('sets the new tokens and feeds', async () => {
        expect(await registry.getFeedTokens()).to.deep.equal(
          newTokens.map((nt) => nt.address),
        )
        expect(await registry.getFeed(newTokens[0].address)).to.equal(
          newFeeds[0].address,
        )
      })
      it('emits a FeedsSet event', async () => {
        await expect(tx)
          .to.emit(registry, 'FeedsSet')
          .withArgs(
            newTokens.map((nt) => nt.address),
            newFeeds.map((nf) => nf.address),
          )
      })
    })
  })
})
