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
import { constants } from 'ethers'
import { ContractTransaction } from 'ethers'

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
      'addFeed',
      'removeFeed',
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

  describe('#removeFeed', () => {
    let newToken: MockERC20
    let newFeed: MockAggregator
    beforeEach(async () => {
      newToken = <MockERC20>(
        await deployContract(roles.defaultAccount, MockERC20Artifact, [
          '6',
          '6',
          await roles.defaultAccount.getAddress(),
          100,
        ])
      )
      newFeed = <MockAggregator>(
        await deployContract(roles.defaultAccount, MockFeedArtifact, [])
      )
    })
    describe('failure', () => {
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          registry
            .connect(roles.stranger)
            .removeFeed(tokens[2].address, feeds[2].address),
          'Only callable by owner',
        )
      })
      it('fails when there are no feeds', async () => {
        let newRegistry: PriceFeedRegistry = <PriceFeedRegistry>(
          await deployContract(roles.defaultAccount, RegistryArtifact, [[], []])
        )
        await evmRevert(
          newRegistry
            .connect(roles.defaultAccount)
            .removeFeed(newFeed.address, newToken.address),
          'NoFeeds()',
        )
      })
      it('fails when the feed does not exist', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .removeFeed(newToken.address, newFeed.address),
          'FeedDoesNotExist()',
        )
      })
      it('fails when the token doesnt match the configuration', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .removeFeed(tokens[2].address, feeds[3].address),
          'TokenFeedMistmatch()',
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      let theToken: string
      let theFeed: string

      beforeEach(async () => {
        theToken = tokens[2].address
        theFeed = feeds[2].address
        tx = await registry
          .connect(roles.defaultAccount)
          .removeFeed(theToken, theFeed)
      })
      it('removes the token from the mapping', async () => {
        const response = await registry.getFeed(theToken)
        expect(response).to.equal(constants.AddressZero)
      })
      it('removes the token from the list', async () => {
        const response = await registry.getFeedTokens()
        expect(response).to.not.contain(theToken)
      })
      it('emits an event', async () => {
        await expect(tx)
          .to.emit(registry, 'FeedRemoved')
          .withArgs(theToken, theFeed)
      })
    })
  })

  describe('#addFeed', () => {
    let newToken: MockERC20
    let newFeed: MockAggregator
    beforeEach(async () => {
      newToken = <MockERC20>(
        await deployContract(roles.defaultAccount, MockERC20Artifact, [
          '6',
          '6',
          await roles.defaultAccount.getAddress(),
          100,
        ])
      )
      newFeed = <MockAggregator>(
        await deployContract(roles.defaultAccount, MockFeedArtifact, [])
      )
    })

    describe('failure', () => {
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          registry
            .connect(roles.stranger)
            .addFeed(newToken.address, newFeed.address),
          'Only callable by owner',
        )
      })
      it('fails when the feed already exists', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .addFeed(tokens[1].address, feeds[1].address),
          `FeedAlreadyAdded()`,
        )
      })
      it('fails when the token is zero address', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .addFeed(constants.AddressZero, newFeed.address),
          `InvalidPriceFeedConfig()`,
        )
      })
      it('fails when the token is a zer address', async () => {
        await evmRevert(
          registry
            .connect(roles.defaultAccount)
            .addFeed(newToken.address, constants.AddressZero),
          `InvalidPriceFeedConfig()`,
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        tx = await registry
          .connect(roles.defaultAccount)
          .addFeed(newToken.address, newFeed.address)
      })
      it('adds a new feed to the mapping', async () => {
        const configuredFeed = await registry.getFeed(newToken.address)
        expect(configuredFeed).to.equal(newFeed.address)
      })
      it('adds the token to the s_tokenList', async () => {
        const tokenList = await registry.getFeedTokens()
        expect(tokenList).to.contain(newToken.address)
      })
      it('emits an event', async () => {
        await expect(tx)
          .to.emit(registry, 'FeedAdded')
          .withArgs(newToken.address, newFeed.address)
      })
    })
  })
})
