import hre from 'hardhat'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { AFN } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { Signer } from '@ethersproject/abstract-signer'
import { expect } from 'chai'
import { evmRevert } from '../../../test-helpers/matchers'
import { ContractTransaction } from 'ethers'
import { publicAbi } from '../../../test-helpers/helpers'

const { deployContract } = hre.waffle
let roles: Roles
let AFNArtifact: Artifact
let afn: AFN
let partyAccounts: Array<Signer>
let parties: Array<string>
let weights: Array<BigNumber>
let goodQuorum: BigNumber
let badQuorum: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('AFN', () => {
  beforeEach(async () => {
    partyAccounts = [
      roles.oracleNode1,
      roles.oracleNode2,
      roles.oracleNode3,
      roles.oracleNode4,
    ]
    parties = await Promise.all(partyAccounts.map((acc) => acc.getAddress()))
    weights = [1, 2, 3, 4].map((num) => BigNumber.from(num))
    badQuorum = BigNumber.from(3)
    goodQuorum = BigNumber.from(4)

    AFNArtifact = await hre.artifacts.readArtifact('AFN')

    afn = <AFN>(
      await deployContract(roles.defaultAccount, AFNArtifact, [
        parties,
        weights,
        goodQuorum,
        badQuorum,
      ])
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(afn, [
      'hasBadSignal',
      'getLastHeartbeat',
      'voteGood',
      'voteBad',
      'recover',
      'setConfig',
      'getWeightThresholds',
      'getParticipants',
      'getWeightByParticipant',
      'getRound',
      'getCommitteeVersion',
      'getLastGoodVoteByParticipant',
      'getGoodVotes',
      'getBadVotersAndVotes',
      'hasVotedBad',
      // Ownership
      'owner',
      'transferOwnership',
      'acceptOwnership',
      // TypeAndVersionInterface
      'typeAndVersion',
    ])
  })

  describe('#constructor', () => {
    it('deploys correctly', async () => {
      const initialBadSignal = await afn.hasBadSignal()
      const initialLastHeartbeat = await afn.getLastHeartbeat()
      const initialWeightThresholds = await afn.getWeightThresholds()
      const initialParticipants = await afn.getParticipants()
      const initialRound = await afn.getRound()
      const initialCommitteeVersion = await afn.getCommitteeVersion()
      expect(initialBadSignal).to.be.false
      expect(initialLastHeartbeat.timestamp).to.equal(0)
      expect(initialWeightThresholds.good).to.equal(goodQuorum)
      expect(initialWeightThresholds.bad).to.equal(badQuorum)
      expect(initialParticipants).to.deep.equal(parties)
      expect(initialRound).to.equal(1)
      expect(initialCommitteeVersion).to.equal(1)

      for (let i = 0; i < parties.length; i++) {
        const party = parties[i]
        const initialWeight = await afn.getWeightByParticipant(party)
        expect(initialWeight).to.equal(weights[i])
      }
    })
  })

  describe('#voteGood', () => {
    describe('failure', () => {
      it('fails when the round is wrong', async () => {
        await evmRevert(
          afn.connect(partyAccounts[1]).voteGood(2),
          'IncorrectRound(1, 2)',
        )
      })
      it('fails if the signal is bad', async () => {
        await afn.connect(partyAccounts[3]).voteBad()
        await evmRevert(
          afn.connect(partyAccounts[1]).voteGood(1),
          'MustRecoverFromBadSignal',
        )
      })
      it('fails if the voter is not a registered party', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).voteGood(1),
          `InvalidVoter("${await roles.defaultAccount.getAddress()}")`,
        )
      })
      it('fails if the voter already voted in this round', async () => {
        await afn.connect(partyAccounts[1]).voteGood(1)
        await evmRevert(
          afn.connect(partyAccounts[1]).voteGood(1),
          `AlreadyVoted()`,
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      let index: number
      describe('single vote without reaching quorum', () => {
        beforeEach(async () => {
          index = 1
          tx = await afn.connect(partyAccounts[index]).voteGood(1)
        })
        it('sets the last good vote of the voter to this round', async () => {
          const lastGoodVote = await afn.getLastGoodVoteByParticipant(
            parties[index],
          )
          expect(lastGoodVote).to.equal(1)
        })
        it('adds the votes to this round', async () => {
          const votes = await afn.getGoodVotes()
          expect(votes).to.equal(weights[index])
        })
        it('emits a good vote event', async () => {
          await expect(tx).to.emit(afn, 'GoodVote').withArgs(parties[index], 1)
        })
      })

      describe('reaching good quorum', () => {
        beforeEach(async () => {
          index = 3
          tx = await afn.connect(partyAccounts[index]).voteGood(1)
        })

        it('sets the last heartbeat', async () => {
          const heartbeat = await afn.getLastHeartbeat()
          expect(heartbeat.round).to.equal(1)
          expect(heartbeat.committeeVersion).to.equal(1)
          expect(heartbeat.timestamp).to.not.equal(0)
        })
        it('increments the round', async () => {
          const roundSet = await afn.getRound()
          expect(roundSet).to.equal(2)
        })
        it('emits a heatbeat event', async () => {
          await expect(tx).to.emit(afn, 'AFNHeartbeat')
        })
      })
    })
  })

  describe('#voteBad', () => {
    describe('failure', () => {
      it('fails if the signal is already bad', async () => {
        await afn.connect(partyAccounts[3]).voteBad()
        await evmRevert(
          afn.connect(partyAccounts[2]).voteBad(),
          'MustRecoverFromBadSignal()',
        )
      })
      it('fails if the voter is not a registered party', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).voteBad(),
          `InvalidVoter("${await roles.defaultAccount.getAddress()}")`,
        )
      })
      it('fails is the voter has already voted bad', async () => {
        await afn.connect(partyAccounts[0]).voteBad()
        await evmRevert(
          afn.connect(partyAccounts[0]).voteBad(),
          'AlreadyVoted()',
        )
      })
    })

    describe('success', () => {
      it('increments votes, adds party to voters and sets s_hasVotedBad for sender', async () => {
        const index = 1
        await afn.connect(partyAccounts[index]).voteBad()
        const votersAndVotes = await afn.getBadVotersAndVotes()
        const hasVotedBad = await afn.hasVotedBad(parties[index])
        expect(votersAndVotes.voters).to.deep.equal([parties[index]])
        expect(votersAndVotes.votes).to.equal(weights[index])
        expect(hasVotedBad).to.be.true
      })
      describe('reaching bad quorum', () => {
        let tx: ContractTransaction
        beforeEach(async () => {
          tx = await afn.connect(partyAccounts[3]).voteBad()
        })

        it('sets the bad signal', async () => {
          expect(await afn.hasBadSignal()).to.be.true
        })
        it('emits an event', async () => {
          await expect(tx).to.emit(afn, 'AFNBadSignal')
        })
      })
    })
  })

  describe('#recover', () => {
    describe('failure', () => {
      it('only allows the owner to call', async () => {
        await evmRevert(
          afn.connect(roles.stranger).recover(),
          'Only callable by owner',
        )
      })
      it('fails if there is no bad signal', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).recover(),
          'RecoveryNotNecessary()',
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        await afn.connect(partyAccounts[3]).voteBad()
        tx = await afn.connect(roles.defaultAccount).recover()
      })

      it('resets s_badVoters, s_hasVotedBad and s_badVotes', async () => {
        const votersAndVotes = await afn.getBadVotersAndVotes()
        expect(votersAndVotes.voters.length).to.equal(0)
        expect(votersAndVotes.votes).to.equal(0)
      })
      it('turns off the bad signal', async () => {
        const hasBadSignal = await afn
          .connect(roles.defaultAccount)
          .hasBadSignal()
        expect(hasBadSignal).to.be.false
      })
      it('emits a Recovered event', async () => {
        await expect(tx).to.emit(afn, 'RecoveredFromBadSignal')
      })
    })
  })

  describe('#setConfig', () => {
    let newParties: Array<string>
    let newWeights: Array<BigNumber>
    let newGoodQuorum: BigNumber
    let newBadQuorum: BigNumber

    describe('failure', () => {
      beforeEach(async () => {
        newParties = [
          await roles.consumer.getAddress(),
          await roles.stranger.getAddress(),
        ]
        newWeights = [BigNumber.from(8), BigNumber.from(9)]
        newGoodQuorum = BigNumber.from(10)
        newBadQuorum = BigNumber.from(8)
      })

      it('only allows the owner to set config', async () => {
        await evmRevert(
          afn.connect(partyAccounts[0]).setConfig([], [], 1, 1),
          'Only callable by owner',
        )
      })

      it('fails if the parties length is 0', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).setConfig([], newWeights, 1, 1),
          'InvalidConfig()',
        )
      })
      it('fails if the weights length is 0', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).setConfig(newParties, [], 1, 1),
          'InvalidConfig()',
        )
      })
      it('fails if the goodQuorum is 0', async () => {
        await evmRevert(
          afn
            .connect(roles.defaultAccount)
            .setConfig(newParties, newWeights, 0, 1),
          'InvalidConfig()',
        )
      })
      it('fails if the badQuorum is 0', async () => {
        await evmRevert(
          afn
            .connect(roles.defaultAccount)
            .setConfig(newParties, newWeights, 1, 0),
          'InvalidConfig()',
        )
      })
      it('fails if a weight is 0', async () => {
        await evmRevert(
          afn.connect(roles.defaultAccount).setConfig(newParties, [0, 0], 1, 1),
          'InvalidWeight()',
        )
      })
      it('fails if the sum of weights is less that the heartbeat threshold', async () => {
        await evmRevert(
          afn
            .connect(roles.defaultAccount)
            .setConfig(newParties, [5, 5], 9, 12),
          'InvalidConfig()',
        )
      })
      it('fails if the sum of weights is less that the bad signal threshold', async () => {
        await evmRevert(
          afn
            .connect(roles.defaultAccount)
            .setConfig(newParties, [5, 5], 12, 9),
          'InvalidConfig()',
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      let initialRound: BigNumber
      let initialCommitteeVersion: BigNumber

      beforeEach(async () => {
        initialRound = await afn.getRound()
        initialCommitteeVersion = await afn.getCommitteeVersion()

        await afn.connect(partyAccounts[0]).voteBad()
        await afn.connect(partyAccounts[1]).voteGood(initialRound)

        newParties = [
          await roles.consumer.getAddress(),
          await roles.stranger.getAddress(),
        ]
        newWeights = [BigNumber.from(8), BigNumber.from(9)]
        newGoodQuorum = BigNumber.from(10)
        newBadQuorum = BigNumber.from(8)
        tx = await afn
          .connect(roles.defaultAccount)
          .setConfig(newParties, newWeights, newGoodQuorum, newBadQuorum)
      })

      it('removes the old configs', async () => {
        for (let i = 0; i < parties.length; i++) {
          const party = parties[i]
          const setWeight = await afn
            .connect(roles.defaultAccount)
            .getWeightByParticipant(party)
          expect(setWeight).to.equal(0)
        }
        const thresholds = await afn.getWeightThresholds()
        expect(thresholds.good).to.not.equal(goodQuorum)
        expect(thresholds.bad).to.not.equal(badQuorum)
        const setRound = await afn.getRound()
        const setCommitteeVersion = await afn.getCommitteeVersion()
        expect(setRound).to.not.equal(initialRound)
        expect(setCommitteeVersion).to.not.equal(initialCommitteeVersion)
      })

      it('sets the new configs', async () => {
        for (let i = 0; i < newParties.length; i++) {
          const party = newParties[i]
          const setWeight = await afn
            .connect(roles.defaultAccount)
            .getWeightByParticipant(party)
          expect(setWeight).to.equal(newWeights[i])
        }
        const thresholds = await afn.getWeightThresholds()
        expect(thresholds.good).to.not.equal(goodQuorum)
        expect(thresholds.bad).to.not.equal(badQuorum)
        const setRound = await afn.getRound()
        const setCommitteeVersion = await afn.getCommitteeVersion()
        expect(setRound).to.equal(initialRound.add(1))
        expect(setCommitteeVersion).to.equal(initialCommitteeVersion.add(1))
      })

      it('resets all votes', async () => {
        const goodVotes = await afn.getGoodVotes()
        const badVotes = await afn.getBadVotersAndVotes()
        expect(goodVotes).to.equal(0)
        expect(badVotes.votes).to.equal(0)
        expect(badVotes.voters).to.deep.equal([])
      })

      it('emits an event', async () => {
        await expect(tx)
          .to.emit(afn, 'ConfigSet')
          .withArgs(newParties, newWeights, newGoodQuorum, newBadQuorum)
      })
    })
  })
})
