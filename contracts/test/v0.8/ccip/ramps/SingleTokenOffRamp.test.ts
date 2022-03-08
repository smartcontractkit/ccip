import hre from 'hardhat'
import {
  numToBytes32,
  publicAbi,
  stringToBytes,
} from '../../../test-helpers/helpers'
import { expect } from 'chai'
import {
  BigNumber,
  Contract,
  ContractFactory,
  ContractTransaction,
} from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  SimpleMessageReceiver,
  MockERC20,
  LockUnlockPool,
  MockAFN,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  CCIPMessage,
  CCIPMessagePayload,
  encodeReport,
  generateMerkleTreeFromHashes,
  hashMessage,
  MerkleTree,
  messageDeepEqual,
  RelayReport,
} from '../../../test-helpers/ccip'

const { deployContract } = hre.waffle

function constructReport(
  message: CCIPMessage,
  minSequenceNumber: BigNumber,
  maxSequenceNumber: BigNumber,
): RelayReport {
  const rootHash = hashMessage(message)
  let report: RelayReport = {
    merkleRoot: rootHash,
    minSequenceNumber: minSequenceNumber,
    maxSequenceNumber: maxSequenceNumber,
  }
  return report
}

let roles: Roles

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` using typechain artifacts.
let ramp: Contract
let afn: MockAFN
let token: MockERC20
let receiver: SimpleMessageReceiver
let pool: LockUnlockPool

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let rampFactory: ContractFactory

const sourceChainId: number = 123
const destinationChainId: number = 234
const initialExecutionDelay: number = 0
let bucketRate: BigNumber
let bucketCapactiy: BigNumber
let maxTimeBetweenAFNSignals: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('SingleTokenOffRamp', () => {
  beforeEach(async () => {
    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('LockUnlockPool')
    rampFactory = await hre.ethers.getContractFactory(
      'SingleTokenOffRampHelper',
    )
    const SimpleMessageReceiverArtifact: Artifact =
      await hre.artifacts.readArtifact('SimpleMessageReceiver')
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    const mintAmount = BigNumber.from('100000000000000000000')
    maxTimeBetweenAFNSignals = BigNumber.from(60).mul(60) // 1 hour
    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        mintAmount,
      ])
    )
    pool = <LockUnlockPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [token.address])
    )
    await token
      .connect(roles.defaultAccount)
      .transfer(pool.address, mintAmount.div(2))
    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    ramp = await rampFactory
      .connect(roles.defaultAccount)
      .deploy(
        sourceChainId,
        destinationChainId,
        token.address,
        pool.address,
        bucketRate,
        bucketCapactiy,
        afn.address,
        maxTimeBetweenAFNSignals,
        initialExecutionDelay,
      )
    await pool.connect(roles.defaultAccount).setOffRamp(ramp.address, true)
    receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(ramp, [
      // SingleTokenRamp
      'TOKEN',
      'POOL',
      'SOURCE_CHAIN_ID',
      'CHAIN_ID',
      'executeTransaction',
      'generateMerkleRoot',
      'getMerkleRoot',
      'getExecuted',
      'getLastReport',
      'getExecutionDelaySeconds',
      'setExecutionDelaySeconds',
      'configureTokenBucket',
      'getTokenBucket',
      // HealthChecker
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
      //SingleTokenOffRampHelper
      'report',
      // OCR2Base
      'setConfig',
      'latestConfigDetails',
      'transmitters',
      'transmit',
      'latestConfigDigestAndEpoch',
      // TypeAndVersionInterface
      'typeAndVersion',
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
    it('should deploy correctly', async () => {
      const owner = await roles.defaultAccount.getAddress()
      await expect(await ramp.TOKEN()).to.equal(token.address)
      await expect(await ramp.POOL()).to.not.equal(
        hre.ethers.constants.AddressZero,
      )
      await expect(await ramp.SOURCE_CHAIN_ID()).to.equal(sourceChainId)
      await expect(await ramp.owner()).to.equal(owner)
      await expect(await ramp.getExecutionDelaySeconds()).to.equal(0)

      await expect(await pool.owner()).to.equal(owner)
      await expect(await pool.isOffRamp(ramp.address)).to.equal(true)
      await expect(await pool.getToken()).to.equal(token.address)
    })

    it('should fail if the pool token is different from the ramp token', async () => {
      const differentToken = <MockERC20>(
        await deployContract(roles.defaultAccount, TokenArtifact, [
          'LINK Token',
          'LINK',
          await roles.defaultAccount.getAddress(),
          BigNumber.from('100000000000000000000'),
        ])
      )
      await evmRevert(
        rampFactory
          .connect(roles.defaultAccount)
          .deploy(
            sourceChainId,
            destinationChainId,
            differentToken.address,
            pool.address,
            bucketRate,
            bucketCapactiy,
            afn.address,
            maxTimeBetweenAFNSignals,
            initialExecutionDelay,
          ),
        `TokenMismatch()`,
      )
    })
  })

  describe('#generateMerkleProof', () => {
    let messages: Array<string>
    let merkle: any

    it('generates', async () => {
      messages = [
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [1]),
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [2]),
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [3]),
        hre.ethers.utils.defaultAbiCoder.encode(['uint256'], [4]),
      ]
      merkle = generateMerkleTreeFromHashes(messages)
      for (let i = 0; i < merkle.leaves.length; i++) {
        const leaf = merkle.leaves[i]
        const proof = leaf.recursiveProof([])
        const hash = leaf.hash
        expect(await ramp.generateMerkleRoot(proof, hash, i)).to.equal(
          merkle.root.hash,
        )
      }
    })
  })

  describe('#report', () => {
    describe('failure', () => {
      let report: RelayReport
      beforeEach(async () => {
        report = {
          merkleRoot: numToBytes32(1),
          minSequenceNumber: BigNumber.from(2),
          maxSequenceNumber: BigNumber.from(3),
        }
      })

      it('reverts when paused', async () => {
        await ramp.connect(roles.defaultAccount).pause()
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(stringToBytes('')),
          'Pausable: paused',
        )
      })

      it('fails whenn the AFN signal is bad', async () => {
        await afn.voteBad()
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(stringToBytes('')),
          'BadAFNSignal()',
        )
      })

      it('fails when the AFN signal is stale', async () => {
        await afn.setTimestamp(BigNumber.from(1))
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(stringToBytes('')),
          'StaleAFNHeartbeat()',
        )
      })

      it('reverts when the minSequenceNumber is greater than the maxSequenceNumber', async () => {
        report.maxSequenceNumber = BigNumber.from(1)
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(encodeReport(report)),
          'RelayReportError()',
        )
      })

      it('reverts when the minSequenceNumber is not 1 greater than the previous report maxSequenceNumber', async () => {
        await ramp.connect(roles.defaultAccount).report(encodeReport(report))
        report = {
          merkleRoot: numToBytes32(2),
          minSequenceNumber: BigNumber.from(3),
          maxSequenceNumber: BigNumber.from(4),
        }
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(encodeReport(report)),
          `SequenceError(3, 3)`,
        )
      })
    })

    describe('success', () => {
      let report: RelayReport
      let root: string
      let response: ContractTransaction
      beforeEach(async () => {
        root = numToBytes32(1)
        report = {
          merkleRoot: root,
          minSequenceNumber: BigNumber.from(1),
          maxSequenceNumber: BigNumber.from(2),
        }
        response = await ramp
          .connect(roles.defaultAccount)
          .report(encodeReport(report))
      })
      it('stores the root', async () => {
        const stored = await ramp.getMerkleRoot(root)
        expect(stored).to.not.equal(0)
      })
      it('stores the report in s_lastReport', async () => {
        const response = await ramp.getLastReport()
        expect(response.merkleRoot).to.equal(root)
        expect(response.minSequenceNumber).to.equal(report.minSequenceNumber)
        expect(response.maxSequenceNumber).to.equal(report.maxSequenceNumber)
      })
      it('emits a ReportAccepted event', async () => {
        expect(response)
          .to.emit(ramp, 'ReportAccepted')
          .withArgs([root, report.minSequenceNumber, report.maxSequenceNumber])
      })
    })
  })

  describe('#executeTransaction', () => {
    let sequenceNumber: BigNumber
    let sourceId: BigNumber
    let destinationId: BigNumber
    let sender: string
    let messagedata: string
    let amount: BigNumber
    let options: string
    let message: CCIPMessage
    let payload: CCIPMessagePayload
    let report: RelayReport
    beforeEach(async () => {
      sequenceNumber = BigNumber.from(1)
      sourceId = BigNumber.from(sourceChainId)
      destinationId = BigNumber.from(destinationChainId)
      sender = await roles.oracleNode.getAddress()
      messagedata = stringToBytes('Message')
      amount = BigNumber.from('10000000000')
      options = stringToBytes('options')
      payload = {
        receiver: receiver.address,
        data: messagedata,
        tokens: [token.address],
        amounts: [amount],
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
      message = {
        sequenceNumber: sequenceNumber,
        sourceChainId: sourceId,
        destinationChainId: destinationId,
        sender: sender,
        payload: payload,
      }
    })

    describe('failure', () => {
      describe('verifyMerkleProof failures', () => {
        let hashes: string[]
        let root: MerkleTree
        let leaves: MerkleTree[]

        beforeEach(async () => {
          const hash1 = hashMessage(message)
          const sequenceNumber2 = BigNumber.from(2)
          const payload2 = {
            receiver: receiver.address,
            data: messagedata,
            tokens: [token.address],
            amounts: [BigNumber.from('9999999')],
            executor: hre.ethers.constants.AddressZero,
            options: options,
          }
          const message2 = {
            sequenceNumber: sequenceNumber2,
            sourceChainId: sourceId,
            destinationChainId: destinationId,
            sender: sender,
            payload: payload2,
          }
          const hash2 = hashMessage(message2)
          hashes = [hash1, hash2]
          const merkle = generateMerkleTreeFromHashes(hashes)
          root = merkle.root
          leaves = merkle.leaves
          report = {
            merkleRoot: root.hash!,
            minSequenceNumber: sequenceNumber,
            maxSequenceNumber: sequenceNumber2,
          }
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
        })

        it('fails when the payload is wrong', async () => {
          const proof = leaves[0].recursiveProof([])
          message.payload.options = stringToBytes('loremipsum')
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, 0),
            `MerkleProofError(["${proof[0]}"], [${message.sequenceNumber}, ${message.sourceChainId}, ${message.destinationChainId}, "${message.sender}", ["${message.payload.receiver}", "${message.payload.data}", ["${message.payload.tokens[0]}"], [${message.payload.amounts[0]}], "${message.payload.executor}", "${message.payload.options}"]], 0)`,
          )
        })

        it('fails when the proof is wrong', async () => {
          const proof = [numToBytes32(1)]
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, 0),
            `MerkleProofError(["${proof[0]}"], [${message.sequenceNumber}, ${message.sourceChainId}, ${message.destinationChainId}, "${message.sender}", ["${message.payload.receiver}", "${message.payload.data}", ["${message.payload.tokens[0]}"], [${message.payload.amounts[0]}], "${message.payload.executor}", "${message.payload.options}"]], 0)`,
          )
        })

        it('fails when the index is wrong', async () => {
          const proof = leaves[0].recursiveProof([])
          const wrongIndex = 1

          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, wrongIndex),
            `MerkleProofError(["${proof[0]}"], [${message.sequenceNumber}, ${message.sourceChainId}, ${message.destinationChainId}, "${message.sender}", ["${message.payload.receiver}", "${message.payload.data}", ["${message.payload.tokens[0]}"], [${message.payload.amounts[0]}], "${message.payload.executor}", "${message.payload.options}"]], ${wrongIndex})`,
          )
        })

        it('fails when the execution delay has not yet passed', async () => {
          const proof = leaves[0].recursiveProof([])
          await ramp
            .connect(roles.defaultAccount)
            .setExecutionDelaySeconds(60 * 60)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(proof, message, 0),
            `ExecutionDelayError()`,
          )
        })
      })
      describe('validation fails', () => {
        it('fails if the receiver is the ramp', async () => {
          message.payload.receiver = ramp.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is the pool', async () => {
          message.payload.receiver = pool.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is the token', async () => {
          message.payload.receiver = token.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is not a contract', async () => {
          message.payload.receiver = await roles.oracleNode1.getAddress()
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails when the message executor is invalid', async () => {
          // Set the executor to a specific address, then executing with a different
          // one should revert.
          message.payload.executor = await roles.oracleNode1.getAddress()
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidExecutor(${message.sequenceNumber})`,
          )
        })
        it('fails when the message is already executed', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await ramp
            .connect(roles.defaultAccount)
            .executeTransaction([], message, 0)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `AlreadyExecuted(${message.sequenceNumber})`,
          )
        })
        it('should fail if sent from an unsupported source chain', async () => {
          message.sourceChainId = BigNumber.from(999)
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `InvalidSourceChain(${message.sourceChainId})`,
          )
        })
        it('should fail if the number of tokens sent is not 1', async () => {
          message.payload.tokens.push(await roles.oracleNode.getAddress())
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `UnsupportedNumberOfTokens()`,
          )
        })
        it('should fail if the number of amounts of tokens to send is not 1', async () => {
          message.payload.amounts.push(BigNumber.from(50000))
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `UnsupportedNumberOfTokens()`,
          )
        })
        it('should fail if sent using an unsupported token', async () => {
          message.payload.tokens[0] = await roles.oracleNode2.getAddress()
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `UnsupportedToken("${message.payload.tokens[0]}")`,
          )
        })
        it('should fail if sending more tokens than the tokenBucket allows', async () => {
          message.payload.amounts[0] = bucketCapactiy.add(1)
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `ExceedsTokenLimit(${bucketCapactiy}, ${message.payload.amounts[0]})`,
          )
        })
        it('should fail if the receiver does not support CrossChainMessageReceiverInterface', async () => {
          const nonReceiver = <MockERC20>(
            await deployContract(roles.defaultAccount, TokenArtifact, [
              'FAKE Token',
              'FAKE',
              await roles.defaultAccount.getAddress(),
              100,
            ])
          )
          message.payload.receiver = nonReceiver.address
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `ExecutionError(${message.sequenceNumber}, "0x")`,
          )
        })
        it('should fail if the contract is paused', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await ramp.connect(roles.defaultAccount).pause()
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `Pausable: paused`,
          )
        })
        it('fails whenn the AFN signal is bad', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await afn.voteBad()
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `BadAFNSignal()`,
          )
        })

        it('fails when the AFN signal is stale', async () => {
          report = constructReport(message, sequenceNumber, sequenceNumber)
          await ramp.connect(roles.defaultAccount).report(encodeReport(report))
          await afn.setTimestamp(BigNumber.from(1))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction([], message, 0),
            `StaleAFNHeartbeat()`,
          )
        })
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      beforeEach(async () => {
        await ramp
          .connect(roles.defaultAccount)
          .report(
            encodeReport(
              constructReport(message, sequenceNumber, sequenceNumber),
            ),
          )
        tx = await ramp
          .connect(roles.defaultAccount)
          .executeTransaction([], message, 0)
      })
      it('should set s_executed to true', async () => {
        expect(await ramp.getExecuted(message.sequenceNumber)).to.be.true
      })
      it('should deliver the message to the receiver', async () => {
        messageDeepEqual(await receiver.s_message(), message)
      })
      it('should send the funds to the receiver contract', async () => {
        expect(await token.balanceOf(receiver.address)).to.equal(
          message.payload.amounts[0],
        )
      })
      it('should emit a CrossChainMessageExecuted event', async () => {
        expect(tx)
          .to.emit(ramp, 'CrossChainMessageExecuted')
          .withArgs(message.sequenceNumber)
      })
      it('should execute a message specifying an executor', async () => {
        message.payload.executor = await roles.oracleNode1.getAddress()
        message.sequenceNumber = message.sequenceNumber.add(1)
        await ramp
          .connect(roles.defaultAccount)
          .report(
            encodeReport(
              constructReport(
                message,
                sequenceNumber.add(1),
                sequenceNumber.add(1),
              ),
            ),
          )
        // Should not revert
        await expect(
          ramp.connect(roles.oracleNode1).executeTransaction([], message, 0),
        )
          .to.emit(ramp, 'CrossChainMessageExecuted')
          .withArgs(message.sequenceNumber)
      })
    })
  })

  describe('#pause', () => {
    it('owner can pause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).pause())
        .to.emit(ramp, 'Paused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot pause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).pause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#unpause', () => {
    beforeEach(async () => {
      await ramp.connect(roles.defaultAccount).pause()
    })

    it('owner can unpause ramp', async () => {
      const account = roles.defaultAccount
      await expect(ramp.connect(account).unpause())
        .to.emit(ramp, 'Unpaused')
        .withArgs(await account.getAddress())
    })

    it('unknown account cannot unpause pool', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#configureTokenBucket', () => {
    let newRate: BigNumber
    let newCapacity: BigNumber

    beforeEach(async () => {
      newRate = BigNumber.from(5)
      newCapacity = bucketCapactiy.add(10)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp
          .connect(roles.stranger)
          .configureTokenBucket(newRate, newCapacity, true),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .configureTokenBucket(newRate, newCapacity, true)
      const tokenBucketParams = await ramp.getTokenBucket()
      expect(tokenBucketParams.rate).to.equal(newRate)
      expect(tokenBucketParams.capacity).to.equal(newCapacity)
      await expect(tx)
        .to.emit(ramp, 'NewTokenBucketConstructed')
        .withArgs(newRate, newCapacity, true)
    })
  })

  describe('#setExecutionDelaySeconds', () => {
    it('can only be called by the owner', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setExecutionDelaySeconds(60),
        'Only callable by owner',
      )
    })

    it('sets the execution delay', async () => {
      const delaySeconds = 60
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setExecutionDelaySeconds(delaySeconds)
      await expect(tx)
        .to.emit(ramp, 'ExecutionDelaySecondsSet')
        .withArgs(delaySeconds)
      const actualDelaySeconds = await ramp.getExecutionDelaySeconds()
      expect(actualDelaySeconds).to.equal(delaySeconds)
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
        ramp.connect(roles.stranger).setAFN(newAFN.address),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new AFN', async () => {
      const tx = await ramp.connect(roles.defaultAccount).setAFN(newAFN.address)
      expect(await ramp.getAFN()).to.equal(newAFN.address)
      await expect(tx)
        .to.emit(ramp, 'AFNSet')
        .withArgs(afn.address, newAFN.address)
    })
  })

  describe('#setMaxSecondsWithoutAFNHeartbeat', () => {
    let newTime: BigNumber

    beforeEach(async () => {
      newTime = maxTimeBetweenAFNSignals.mul(2)
    })

    it('only callable by owner', async () => {
      await expect(
        ramp.connect(roles.stranger).setMaxSecondsWithoutAFNHeartbeat(newTime),
      ).to.be.revertedWith('Only callable by owner')
    })

    it('sets the new max time without afn signal', async () => {
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setMaxSecondsWithoutAFNHeartbeat(newTime)
      expect(await ramp.getMaxSecondsWithoutAFNHeartbeat()).to.equal(newTime)
      await expect(tx)
        .to.emit(ramp, 'AFNMaxHeartbeatTimeSet')
        .withArgs(maxTimeBetweenAFNSignals, newTime)
    })
  })

  describe('#typeAndVersion', () => {
    it('should return the correct type and version', async () => {
      const response = await ramp.typeAndVersion()
      await expect(response).to.equal('SingleTokenOffRamp 1.1.0')
    })
  })
})
