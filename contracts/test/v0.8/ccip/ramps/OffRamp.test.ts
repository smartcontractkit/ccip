import hre, { ethers } from 'hardhat'
import {
  expectGasWithinDeviation,
  numToBytes32,
  publicAbi,
  stringToBytes,
} from '../../../test-helpers/helpers'
import { expect } from 'chai'
import {
  BigNumber,
  Contract,
  ContractFactory,
  ContractReceipt,
  ContractTransaction,
} from 'ethers'
import { Roles, getUsers } from '../../../test-helpers/setup'
import {
  SimpleMessageReceiver,
  MockERC20,
  NativeTokenPool,
  MockAFN,
  MockAggregator,
} from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import {
  CCIPMessage,
  CCIPMessagePayload,
  encodeRelayReport,
  ExecutionReport,
  MerkleMultiTree,
  messageDeepEqual,
  RelayReport,
} from '../../../test-helpers/ccip/ccip'
import { constants } from 'ethers'
import { GAS } from '../../../test-helpers/ccip/gas-measurements'
const { deployContract } = hre.waffle

let roles: Roles

// This has to be ethers.Contract because of an issue with
// `address.call(abi.encodeWithSelector(...))` using typechain artifacts.
let ramp: Contract
let router: Contract
let afn: MockAFN
let sourceToken1: MockERC20
let destinationToken1: MockERC20
let sourceToken2: MockERC20
let destinationToken2: MockERC20
let receiver: SimpleMessageReceiver
let pool1: NativeTokenPool
let pool2: NativeTokenPool
let priceFeed1: MockAggregator

let MockAFNArtifact: Artifact
let TokenArtifact: Artifact
let PoolArtifact: Artifact
let PriceFeedArtifact: Artifact
let SimpleMessageReceiverArtifact: Artifact
let rampFactory: ContractFactory
let routerFactory: ContractFactory

const priceFeed1LatestAnswer: number = 100
const sourceChainId: number = 123
const destinationChainId: number = 234
const initialExecutionDelay: number = 0
const maxTokenLength: number = 10
const initialConfig = {
  executionFeeJuels: 1,
  executionDelaySeconds: initialExecutionDelay,
  maxDataSize: 1000,
  maxTokensLength: maxTokenLength,
}
let bucketRate: BigNumber
let bucketCapactiy: BigNumber
let maxTimeBetweenAFNSignals: BigNumber

async function executionValidationFail(
  ramp: Contract,
  messages: CCIPMessage[],
  revertReason: string,
  takeFees: boolean = false,
) {
  const tree = new MerkleMultiTree(messages)
  await ramp
    .connect(roles.defaultAccount)
    .report(encodeRelayReport(tree.generateRelayReport()))
  await evmRevert(
    ramp
      .connect(roles.defaultAccount)
      .executeTransaction(tree.generateExecutionReport([0]), takeFees),
    revertReason,
  )
}

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('OffRamp', () => {
  beforeEach(async () => {
    MockAFNArtifact = await hre.artifacts.readArtifact('MockAFN')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    PoolArtifact = await hre.artifacts.readArtifact('NativeTokenPool')
    PriceFeedArtifact = await hre.artifacts.readArtifact('MockAggregator')
    rampFactory = await hre.ethers.getContractFactory('OffRampHelper')
    routerFactory = await hre.ethers.getContractFactory('OffRampRouter')

    SimpleMessageReceiverArtifact = await hre.artifacts.readArtifact(
      'SimpleMessageReceiver',
    )
    bucketRate = BigNumber.from('10000000000000000')
    bucketCapactiy = BigNumber.from('100000000000000000')
    const mintAmount = BigNumber.from('100000000000000000000')
    maxTimeBetweenAFNSignals = BigNumber.from(60).mul(60) // 1 hour
    sourceToken1 = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK sourceToken1',
        'LINK',
        await roles.defaultAccount.getAddress(),
        mintAmount,
      ])
    )
    sourceToken2 = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK sourceToken2',
        'LINK',
        await roles.defaultAccount.getAddress(),
        mintAmount,
      ])
    )
    destinationToken1 = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK destinationToken1',
        'LINK',
        await roles.defaultAccount.getAddress(),
        mintAmount,
      ])
    )
    destinationToken2 = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK destinationToken2',
        'LINK',
        await roles.defaultAccount.getAddress(),
        mintAmount,
      ])
    )
    let bucketConfig = {
      rate: bucketRate,
      capacity: bucketCapactiy,
    }
    pool1 = <NativeTokenPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        destinationToken1.address,
        bucketConfig,
        bucketConfig,
      ])
    )
    pool2 = <NativeTokenPool>(
      await deployContract(roles.defaultAccount, PoolArtifact, [
        destinationToken2.address,
        bucketConfig,
        bucketConfig,
      ])
    )
    await destinationToken1
      .connect(roles.defaultAccount)
      .transfer(pool1.address, mintAmount.div(2))
    await destinationToken2
      .connect(roles.defaultAccount)
      .transfer(pool2.address, mintAmount.div(2))
    priceFeed1 = <MockAggregator>(
      await deployContract(roles.defaultAccount, PriceFeedArtifact)
    )
    await priceFeed1
      .connect(roles.defaultAccount)
      .setLatestAnswer(priceFeed1LatestAnswer)
    afn = <MockAFN>await deployContract(roles.defaultAccount, MockAFNArtifact)
    ramp = await rampFactory
      .connect(roles.defaultAccount)
      .deploy(
        sourceChainId,
        destinationChainId,
        [sourceToken1.address, sourceToken2.address],
        [pool1.address, pool2.address],
        [priceFeed1.address, constants.AddressZero],
        afn.address,
        maxTimeBetweenAFNSignals,
        initialExecutionDelay,
        maxTokenLength,
      )
    router = await routerFactory
      .connect(roles.defaultAccount)
      .deploy([ramp.address])
    await ramp.connect(roles.defaultAccount).setRouter(router.address)
    await pool1.connect(roles.defaultAccount).setOffRamp(ramp.address, true)
    await pool2.connect(roles.defaultAccount).setOffRamp(ramp.address, true)
    receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
    )
  })

  it('has a limited public interface [ @skip-coverage ]', async () => {
    publicAbi(ramp, [
      // Ramp
      'SOURCE_CHAIN_ID',
      'CHAIN_ID',
      'executeTransaction',
      'withdrawAccumulatedFees',
      'merkleRoot',
      'getMerkleRoot',
      'getExecuted',
      'getLastReport',
      'getOffRampConfig',
      'setOffRampConfig',
      'getRouter',
      'setRouter',
      // HealthChecker
      'setAFN',
      'getAFN',
      'setMaxSecondsWithoutAFNHeartbeat',
      'getMaxSecondsWithoutAFNHeartbeat',
      'isHealthy',
      // TokenPoolRegistry
      'addPool',
      'removePool',
      'getPool',
      'isPool',
      'getPoolTokens',
      // PriceFeedRegistry
      'addFeed',
      'removeFeed',
      'getFeed',
      'getFeedTokens',
      // OffRampHelper
      'report',
      // OCR2Abstract
      'setConfig',
      'latestConfigDetails',
      'latestConfigDigestAndEpoch',
      'transmit',
      // OCR2Base
      'transmitters',
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
      await expect(await ramp.SOURCE_CHAIN_ID()).to.equal(sourceChainId)
      await expect(await ramp.owner()).to.equal(owner)
      await expect(await ramp.getOffRampConfig()).to.deep.equal([
        BigNumber.from(initialConfig.executionFeeJuels),
        BigNumber.from(initialConfig.executionDelaySeconds),
        BigNumber.from(initialConfig.maxDataSize),
        BigNumber.from(initialConfig.maxTokensLength),
      ])
      await expect(await pool1.owner()).to.equal(owner)
      await expect(await pool1.isOffRamp(ramp.address)).to.equal(true)
      await expect(await pool1.getToken()).to.equal(destinationToken1.address)
    })
  })

  describe('#merkleRoot', () => {
    let messages: Array<CCIPMessage>
    let tree: MerkleMultiTree

    beforeEach(async () => {
      const receiver = await roles.oracleNode1.getAddress()
      messages = [
        {
          sourceChainId: BigNumber.from(1),
          sequenceNumber: BigNumber.from(1),
          sender: receiver,
          payload: {
            destinationChainId: BigNumber.from(2),
            tokens: [],
            amounts: [],
            receiver: receiver,
            executor: ethers.constants.AddressZero,
            data: ethers.constants.HashZero,
          },
        },
        {
          sourceChainId: BigNumber.from(1),
          sequenceNumber: BigNumber.from(2),
          sender: receiver,
          payload: {
            destinationChainId: BigNumber.from(2),
            tokens: [],
            amounts: [],
            receiver: receiver,
            executor: ethers.constants.AddressZero,
            data: ethers.constants.HashZero,
          },
        },
        {
          sourceChainId: BigNumber.from(1),
          sequenceNumber: BigNumber.from(3),
          sender: receiver,
          payload: {
            destinationChainId: BigNumber.from(2),
            tokens: [],
            amounts: [],
            receiver: receiver,
            executor: ethers.constants.AddressZero,
            data: ethers.constants.HashZero,
          },
        },
        {
          sourceChainId: BigNumber.from(1),
          sequenceNumber: BigNumber.from(4),
          sender: receiver,
          payload: {
            destinationChainId: BigNumber.from(2),
            tokens: [],
            amounts: [],
            receiver: receiver,
            executor: ethers.constants.AddressZero,
            data: ethers.constants.HashZero,
          },
        },
      ]
      tree = new MerkleMultiTree(messages)
    })

    describe('contract root verification', async () => {
      it('leaf 1', async () => {
        const execReport: ExecutionReport = tree.generateExecutionReport([0])
        const response = await ramp.merkleRoot(execReport)
        expect(response).to.equal(tree.getRoot())
      })

      it('2 leaves', async () => {
        const indices = [0, 3]
        const execReport: ExecutionReport =
          tree.generateExecutionReport(indices)
        const response = await ramp.merkleRoot(execReport)
        expect(response).to.equal(tree.getRoot())
      })

      it('3 leaves', async () => {
        const indices = [0, 1, 3]
        const execReport: ExecutionReport =
          tree.generateExecutionReport(indices)
        const response = await ramp.merkleRoot(execReport)
        expect(response).to.equal(tree.getRoot())
      })

      it('4 leaves', async () => {
        const indices = [0, 1, 2, 3]
        const execReport: ExecutionReport =
          tree.generateExecutionReport(indices)
        const response = await ramp.merkleRoot(execReport)
        expect(response).to.equal(tree.getRoot())
      })
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
          ramp.connect(roles.defaultAccount).report(encodeRelayReport(report)),
          'RelayReportError()',
        )
      })

      it('reverts when the minSequenceNumber is not 1 greater than the previous report maxSequenceNumber', async () => {
        await ramp
          .connect(roles.defaultAccount)
          .report(encodeRelayReport(report))
        report = {
          merkleRoot: numToBytes32(2),
          minSequenceNumber: BigNumber.from(3),
          maxSequenceNumber: BigNumber.from(4),
        }
        await evmRevert(
          ramp.connect(roles.defaultAccount).report(encodeRelayReport(report)),
          `SequenceError(3, 3)`,
        )
      })
    })

    describe('success', () => {
      let report: RelayReport
      let root: string
      let response: ContractTransaction
      let gasUsed: BigNumber
      beforeEach(async () => {
        gasUsed = BigNumber.from(0)
        root = numToBytes32(1)
        report = {
          merkleRoot: root,
          minSequenceNumber: BigNumber.from(1),
          maxSequenceNumber: BigNumber.from(2),
        }
        response = await ramp
          .connect(roles.defaultAccount)
          .report(encodeRelayReport(report))
        gasUsed = gasUsed.add((await response.wait()).gasUsed)
      })
      it('GASTEST [ @skip-coverage ]', async () => {
        expectGasWithinDeviation(gasUsed, GAS.OffRamp.report)
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
    let message: CCIPMessage
    let payload: CCIPMessagePayload
    beforeEach(async () => {
      sequenceNumber = BigNumber.from(1)
      sourceId = BigNumber.from(sourceChainId)
      destinationId = BigNumber.from(destinationChainId)
      sender = await roles.oracleNode.getAddress()
      messagedata = stringToBytes('Message')
      amount = BigNumber.from('10000000000')
      payload = {
        receiver: receiver.address,
        data: messagedata,
        tokens: [sourceToken1.address, sourceToken2.address],
        amounts: [amount, amount],
        executor: hre.ethers.constants.AddressZero,
        destinationChainId: destinationId,
      }
      message = {
        sourceChainId: sourceId,
        sequenceNumber: sequenceNumber,
        sender: sender,
        payload: payload,
      }
    })

    describe('failure', () => {
      describe('verifyMerkleProof failures', () => {
        let tree: MerkleMultiTree
        let relayReport: RelayReport
        let executionReport: ExecutionReport

        beforeEach(async () => {
          const sequenceNumber2 = BigNumber.from(2)
          const payload2 = {
            receiver: receiver.address,
            data: messagedata,
            tokens: [sourceToken1.address],
            amounts: [BigNumber.from('9999999')],
            executor: hre.ethers.constants.AddressZero,
            destinationChainId: destinationId,
          }
          const message2 = {
            sourceChainId: sourceId,
            sequenceNumber: sequenceNumber2,
            sender: sender,
            payload: payload2,
          }
          tree = new MerkleMultiTree([message, message2])
          relayReport = tree.generateRelayReport()
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(relayReport))
          executionReport = tree.generateExecutionReport([0])
        })

        it('fails when the payload is wrong', async () => {
          executionReport.messages[0].payload.data = stringToBytes('loremipsum')

          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(executionReport, false),
          )
        })

        it('fails when the proofs is wrong', async () => {
          executionReport.proofs = []
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(executionReport, false),
          )
        })

        it('fails when the execution delay has not yet passed', async () => {
          let newConfig = initialConfig
          newConfig.executionDelaySeconds = 60 * 60
          await ramp.connect(roles.defaultAccount).setOffRampConfig(newConfig)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(executionReport, false),
            `ExecutionDelayError()`,
          )
        })
      })
      describe('validation fails', () => {
        it('fails if the receiver is the ramp', async () => {
          message.payload.receiver = ramp.address
          await executionValidationFail(
            ramp,
            [message],
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails if the receiver is the pool1', async () => {
          message.payload.receiver = pool1.address
          await executionValidationFail(
            ramp,
            [message],
            `InvalidReceiver("${message.payload.receiver}")`,
          )
        })
        it('fails when the message executor is invalid', async () => {
          // Set the executor to a specific address, then executing with a different
          // one should revert.
          message.payload.executor = await roles.oracleNode1.getAddress()
          await executionValidationFail(
            ramp,
            [message],
            `InvalidExecutor(${message.sequenceNumber})`,
          )
        })
        it('fails when the message is already executed', async () => {
          const tree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(tree.generateRelayReport()))
          const execReport = tree.generateExecutionReport([0])
          await ramp
            .connect(roles.defaultAccount)
            .executeTransaction(execReport, false)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(execReport, false),
            `AlreadyExecuted(${message.sequenceNumber})`,
          )
        })
        it('should fail if sent from an unsupported source chain', async () => {
          message.sourceChainId = BigNumber.from(999)
          await executionValidationFail(
            ramp,
            [message],
            `InvalidSourceChain(${message.sourceChainId})`,
          )
        })
        it('should fail if the number of tokens sent is not 1', async () => {
          message.payload.tokens.push(await roles.oracleNode.getAddress())
          await executionValidationFail(
            ramp,
            [message],
            `UnsupportedNumberOfTokens()`,
          )
        })
        it('should fail if the number of amounts of tokens to send is not 1', async () => {
          message.payload.amounts.push(BigNumber.from(50000))
          await executionValidationFail(
            ramp,
            [message],
            `UnsupportedNumberOfTokens()`,
          )
        })
        it('should fail if sent using an unsupported source token', async () => {
          message.payload.tokens[0] = await roles.oracleNode2.getAddress()
          await executionValidationFail(
            ramp,
            [message],
            `UnsupportedToken("${message.payload.tokens[0]}")`,
          )
        })
        it('should fail if sending more tokens than the tokenBucket allows', async () => {
          message.payload.amounts[0] = bucketCapactiy.add(1)
          await executionValidationFail(
            ramp,
            [message],
            `ExceedsTokenLimit(${bucketCapactiy}, ${message.payload.amounts[0]})`,
          )
        })
        it('should fail if the contract is paused', async () => {
          const tree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(tree.generateRelayReport()))
          await ramp.connect(roles.defaultAccount).pause()
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(tree.generateExecutionReport([0]), false),
            `Pausable: paused`,
          )
        })
        it('fails when the AFN signal is bad', async () => {
          const tree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(tree.generateRelayReport()))
          await afn.voteBad()
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(tree.generateExecutionReport([0]), false),
            `BadAFNSignal()`,
          )
        })
        it('fails when the AFN signal is stale', async () => {
          const tree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(tree.generateRelayReport()))
          await afn.setTimestamp(BigNumber.from(1))
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(tree.generateExecutionReport([0]), false),
            `StaleAFNHeartbeat()`,
          )
        })
      })
      describe('fee taking fails', () => {
        it('fails if the price feed does not exist', async () => {
          const tree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(tree.generateRelayReport()))
          await ramp
            .connect(roles.defaultAccount)
            .removeFeed(sourceToken1.address, priceFeed1.address)
          await evmRevert(
            ramp
              .connect(roles.defaultAccount)
              .executeTransaction(tree.generateExecutionReport([0]), true),
            `FeeError()`,
          )
        })
        it('fails if the fee exceeds the amount sent', async () => {
          message.payload.amounts[0] = 1
          await executionValidationFail(
            ramp,
            [message],
            `panic code 0x11 (Arithmetic operation underflowed or overflowed outside of an unchecked block)`,
            true,
          )
        })
      })
    })

    describe('success - 2 tokens', () => {
      let tx: ContractTransaction
      let tree: MerkleMultiTree
      beforeEach(async () => {
        tree = new MerkleMultiTree([message])
        await ramp
          .connect(roles.defaultAccount)
          .report(encodeRelayReport(tree.generateRelayReport()))
      })
      describe('without fees', () => {
        beforeEach(async () => {
          tx = await ramp
            .connect(roles.defaultAccount)
            .executeTransaction(tree.generateExecutionReport([0]), false)
        })

        describe('GASTEST', () => {
          it('GASTEST - contract receiver execution [ @skip-coverage ]', async () => {
            expectGasWithinDeviation(
              (await tx.wait()).gasUsed,
              GAS.OffRamp.executeTransaction.ONE_MESSAGE.TWO_TOKENS.NO_FEES
                .CONTRACT_RECEIVER,
            )
          })

          it('GASTEST - EOA receiver [ @skip-coverage ]', async () => {
            const nextSequenceNumber = sequenceNumber.add(1)
            message.payload.receiver = await roles.consumer.getAddress()
            message.sequenceNumber = nextSequenceNumber
            message.payload.data = []
            const newTree: MerkleMultiTree = new MerkleMultiTree([message])
            await ramp
              .connect(roles.defaultAccount)
              .report(encodeRelayReport(newTree.generateRelayReport()))
            tx = await ramp
              .connect(roles.oracleNode)
              .executeTransaction(newTree.generateExecutionReport([0]), true)
            expectGasWithinDeviation(
              (await tx.wait()).gasUsed,
              GAS.OffRamp.executeTransaction.ONE_MESSAGE.TWO_TOKENS.NO_FEES
                .EOA_RECEIVER,
            )
          })
        })

        it('should set s_executed to true', async () => {
          expect(await ramp.getExecuted(message.sequenceNumber)).to.be.true
        })
        it('should deliver the message to the receiver', async () => {
          messageDeepEqual(await receiver.s_message(), message)
        })
        it('should send the funds to the receiver contract', async () => {
          expect(await destinationToken1.balanceOf(receiver.address)).to.equal(
            message.payload.amounts[0],
          )
          expect(await destinationToken2.balanceOf(receiver.address)).to.equal(
            message.payload.amounts[1],
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
          const newTree: MerkleMultiTree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(newTree.generateRelayReport()))
          // Should not revert
          await expect(
            ramp
              .connect(roles.oracleNode1)
              .executeTransaction(newTree.generateExecutionReport([0]), false),
          )
            .to.emit(ramp, 'CrossChainMessageExecuted')
            .withArgs(message.sequenceNumber)
        })
      })
      describe('with fees', () => {
        beforeEach(async () => {
          tx = await ramp
            .connect(roles.oracleNode)
            .executeTransaction(tree.generateExecutionReport([0]), true)
        })

        describe('GASTEST', () => {
          it('GASTEST - contract receiver execution [ @skip-coverage ]', async () => {
            expectGasWithinDeviation(
              (await tx.wait()).gasUsed,
              GAS.OffRamp.executeTransaction.ONE_MESSAGE.TWO_TOKENS.FEES
                .CONTRACT_RECEIVER,
            )
          })

          it('GASTEST - EOA receiver [ @skip-coverage ]', async () => {
            const nextSequenceNumber = sequenceNumber.add(1)
            message.payload.receiver = await roles.consumer.getAddress()
            message.sequenceNumber = nextSequenceNumber
            message.payload.data = []
            const newTree: MerkleMultiTree = new MerkleMultiTree([message])
            await ramp
              .connect(roles.defaultAccount)
              .report(encodeRelayReport(newTree.generateRelayReport()))
            tx = await ramp
              .connect(roles.oracleNode)
              .executeTransaction(newTree.generateExecutionReport([0]), true)
            expectGasWithinDeviation(
              (await tx.wait()).gasUsed,
              GAS.OffRamp.executeTransaction.ONE_MESSAGE.TWO_TOKENS.FEES
                .EOA_RECEIVER,
            )

            expect(tx)
              .to.emit(ramp, 'CrossChainMessageExecuted')
              .withArgs(message.sequenceNumber)
          })
        })

        it('should set s_executed to true', async () => {
          expect(await ramp.getExecuted(message.sequenceNumber)).to.be.true
        })
        it('should deliver the message to the receiver', async () => {
          messageDeepEqual(await receiver.s_message(), message)
        })
        it('should mint fee funds to the executor', async () => {
          expect(
            await destinationToken1.balanceOf(
              await roles.oracleNode.getAddress(),
            ),
          ).to.equal(priceFeed1LatestAnswer)
          await expect(tx)
            .to.emit(pool1, 'Released')
            .withArgs(
              ramp.address,
              await roles.oracleNode.getAddress(),
              priceFeed1LatestAnswer,
            )
        })
        it('should not extract a fee if fee is zero', async () => {
          let newConfig = initialConfig
          newConfig.executionFeeJuels = 0
          newConfig.executionDelaySeconds = 0
          await ramp.connect(roles.defaultAccount).setOffRampConfig(newConfig)
          const newSequenceNumber = message.sequenceNumber.add(1)
          message.sequenceNumber = newSequenceNumber
          const newTree: MerkleMultiTree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(newTree.generateRelayReport()))
          tx = await ramp
            .connect(roles.oracleNode)
            .executeTransaction(newTree.generateExecutionReport([0]), true)

          const receipt: ContractReceipt = await tx.wait()
          for (let i = 0; i < receipt.logs.length; i++) {
            const log = receipt.logs[i]
            if (log.address == pool1.address) {
              const parsedLog = pool1.interface.parseLog(log)
              // ensure that no Released events are emitted with the executor as recipient
              expect(parsedLog.args.recipient).to.not.equal(
                await roles.oracleNode.getAddress(),
              )
            }
          }
        })
        it('should send the funds to the receiver contract', async () => {
          const amountAfterFee = BigNumber.from(message.payload.amounts[0]).sub(
            priceFeed1LatestAnswer,
          )
          expect(await destinationToken1.balanceOf(receiver.address)).to.equal(
            amountAfterFee,
          )
          expect(await destinationToken2.balanceOf(receiver.address)).to.equal(
            message.payload.amounts[1],
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
          const newTree: MerkleMultiTree = new MerkleMultiTree([message])
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(newTree.generateRelayReport()))
          // Should not revert
          await expect(
            ramp
              .connect(roles.oracleNode1)
              .executeTransaction(newTree.generateExecutionReport([0]), true),
          )
            .to.emit(ramp, 'CrossChainMessageExecuted')
            .withArgs(message.sequenceNumber)
        })
      })
    })

    describe('GASTEST - Tree of 20 [ @skip-coverage ]', () => {
      describe('with no tokens', () => {
        let messages: CCIPMessage[] = []
        let tree: MerkleMultiTree
        let relayReport: RelayReport
        let executionReport: ExecutionReport
        beforeEach(async () => {
          const lastReport: RelayReport = await ramp.getLastReport()
          for (let i = 0; i < 20; i++) {
            const tempReceiver = <SimpleMessageReceiver>(
              await deployContract(
                roles.defaultAccount,
                SimpleMessageReceiverArtifact,
              )
            )
            messages.push({
              sourceChainId: BigNumber.from(sourceChainId),
              sequenceNumber: lastReport.maxSequenceNumber.add(
                BigNumber.from(i + 1),
              ),
              sender: await roles.defaultAccount.getAddress(),
              payload: {
                tokens: [],
                amounts: [],
                destinationChainId: BigNumber.from(destinationChainId),
                receiver: tempReceiver.address,
                executor: ethers.constants.AddressZero,
                data: ethers.utils.defaultAbiCoder.encode(
                  ['string'],
                  [`no tokens message ${i + 1}`],
                ),
              },
            })
          }

          tree = new MerkleMultiTree(messages)
          relayReport = tree.generateRelayReport()
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(relayReport))
        })

        it('GASTEST - executing 10 messages', async () => {
          const messageIndices = [...Array(10).keys()]
          executionReport = tree.generateExecutionReport(messageIndices)

          const tx = await ramp
            .connect(roles.defaultAccount)
            .executeTransaction(executionReport, false)
          const receipt: ContractReceipt = await tx.wait()
          expectGasWithinDeviation(
            receipt.gasUsed,
            GAS.OffRamp.executeTransaction.TEN_MESSAGES.NO_FEES.NO_TOKENS,
          )
        })
      })

      describe('with 1 token', () => {
        let messages: CCIPMessage[] = []
        let tree: MerkleMultiTree
        let relayReport: RelayReport
        let executionReport: ExecutionReport
        beforeEach(async () => {
          const lastReport: RelayReport = await ramp.getLastReport()
          for (let i = 0; i < 20; i++) {
            const tempReceiver = <SimpleMessageReceiver>(
              await deployContract(
                roles.defaultAccount,
                SimpleMessageReceiverArtifact,
              )
            )
            messages.push({
              sourceChainId: BigNumber.from(sourceChainId),
              sequenceNumber: lastReport.maxSequenceNumber.add(
                BigNumber.from(i + 1),
              ),
              sender: await roles.defaultAccount.getAddress(),
              payload: {
                tokens: [sourceToken1.address],
                amounts: [1],
                destinationChainId: BigNumber.from(destinationChainId),
                receiver: tempReceiver.address,
                executor: ethers.constants.AddressZero,
                data: ethers.utils.defaultAbiCoder.encode(
                  ['string'],
                  [`with 1 token message ${i + 1}`],
                ),
              },
            })
          }

          tree = new MerkleMultiTree(messages)
          relayReport = tree.generateRelayReport()
          await ramp
            .connect(roles.defaultAccount)
            .report(encodeRelayReport(relayReport))
        })

        it('GASTEST - executing 10 messages', async () => {
          const messageIndices = [...Array(10).keys()]
          executionReport = tree.generateExecutionReport(messageIndices)

          const tx = await ramp
            .connect(roles.defaultAccount)
            .executeTransaction(executionReport, false)
          const receipt: ContractReceipt = await tx.wait()
          expectGasWithinDeviation(
            receipt.gasUsed,
            GAS.OffRamp.executeTransaction.TEN_MESSAGES.NO_FEES.ONE_TOKEN,
          )
        })
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

    it('unknown account cannot pause pool1', async function () {
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

    it('unknown account cannot unpause pool1', async function () {
      const account = roles.stranger
      await expect(ramp.connect(account).unpause()).to.be.revertedWith(
        'Only callable by owner',
      )
    })
  })

  describe('#setOffRampConfig', () => {
    it('can only be called by the owner', async () => {
      await evmRevert(
        ramp.connect(roles.stranger).setOffRampConfig(initialConfig),
        'Only callable by owner',
      )
    })

    it('sets the config', async () => {
      let newConfig = initialConfig
      newConfig.executionDelaySeconds = newConfig.executionDelaySeconds * 2
      newConfig.executionFeeJuels = newConfig.executionFeeJuels * 2
      newConfig.maxDataSize = newConfig.maxDataSize * 2
      newConfig.maxTokensLength = newConfig.maxTokensLength * 2
      const tx = await ramp
        .connect(roles.defaultAccount)
        .setOffRampConfig(newConfig)
      await expect(tx)
        .to.emit(ramp, 'OffRampConfigSet')
        .withArgs([
          newConfig.executionFeeJuels,
          newConfig.executionDelaySeconds,
          newConfig.maxDataSize,
          newConfig.maxTokensLength,
        ])
      const actualConfig = await ramp.getOffRampConfig()
      expect(actualConfig).to.deep.equal([
        BigNumber.from(newConfig.executionFeeJuels),
        BigNumber.from(newConfig.executionDelaySeconds),
        BigNumber.from(newConfig.maxDataSize),
        BigNumber.from(newConfig.maxTokensLength),
      ])
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
      await expect(response).to.equal('OffRamp 0.0.1')
    })
  })
})
