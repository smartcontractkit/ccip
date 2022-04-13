import { expect } from 'chai'
import {
  BigNumber,
  Contract,
  ContractFactory,
  ContractTransaction,
} from 'ethers'
import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { SimpleMessageReceiver } from '../../../../typechain'
import { CCIPMessage, messageDeepEqual } from '../../../test-helpers/ccip'
import { evmRevert } from '../../../test-helpers/matchers'

import { getUsers, Roles } from '../../../test-helpers/setup'

const { deployContract } = hre.waffle

let roles: Roles

let OffRampRouterFactory: ContractFactory
let SimpleMessageReceiverArtifact: Artifact

let router: Contract
let receiver: SimpleMessageReceiver

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('OffRampRouter', () => {
  beforeEach(async () => {
    OffRampRouterFactory = await hre.ethers.getContractFactory('OffRampRouter')
    SimpleMessageReceiverArtifact = await hre.artifacts.readArtifact(
      'SimpleMessageReceiver',
    )

    router = await OffRampRouterFactory.connect(roles.defaultAccount).deploy([
      await roles.defaultAccount.getAddress(),
    ])
    receiver = <SimpleMessageReceiver>(
      await deployContract(roles.defaultAccount, SimpleMessageReceiverArtifact)
    )
  })

  describe('#addOffRamp', () => {
    describe('failure', () => {
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          router
            .connect(roles.oracleNode)
            .addOffRamp(await roles.oracleNode.getAddress()),
          'Only callable by owner',
        )
      })
      it('fails when the offRamp is already configured', async () => {
        const addr = await roles.defaultAccount.getAddress()
        await evmRevert(
          router.connect(roles.defaultAccount).addOffRamp(addr),
          `AlreadyConfigured("${addr}")`,
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      let addr: string
      beforeEach(async () => {
        addr = await roles.oracleNode.getAddress()
        tx = await router.connect(roles.defaultAccount).addOffRamp(addr)
      })
      it('adds then offRamp', async () => {
        const ramps = await router.getOffRamps()
        expect(ramps).to.contain(addr)
        const response = await router.isOffRamp(addr)
        expect(response).to.be.true
      })
      it('emits an event', async () => {
        await expect(tx).to.emit(router, 'OffRampAdded').withArgs(addr)
      })
    })
  })

  describe('#removeOffRamp', () => {
    describe('failure', () => {
      it('fails when called by a non-owner', async () => {
        await evmRevert(
          router
            .connect(roles.oracleNode)
            .removeOffRamp(await roles.oracleNode.getAddress()),
          'Only callable by owner',
        )
      })
      it('fails when the offramp is not configured', async () => {
        const addr = await roles.oracleNode.getAddress()
        await evmRevert(
          router.connect(roles.defaultAccount).removeOffRamp(addr),
          `OffRampNotConfigured("${addr}")`,
        )
      })
      it('fails if there are no offramps configured', async () => {
        const addr = await roles.defaultAccount.getAddress()
        await router.connect(roles.defaultAccount).removeOffRamp(addr)
        await evmRevert(
          router.connect(roles.defaultAccount).removeOffRamp(addr),
          `NoOffRampsConfigured()`,
        )
      })
    })
    describe('success', () => {
      let tx: ContractTransaction
      let addr: string
      beforeEach(async () => {
        addr = await roles.defaultAccount.getAddress()
        tx = await router.connect(roles.defaultAccount).removeOffRamp(addr)
      })
      it('removes the offramp', async () => {
        const ramps = await router.getOffRamps()
        expect(ramps).to.not.contain(addr)
        const response = await router.isOffRamp(addr)
        expect(response).to.be.false
      })
      it('emits an event', async () => {
        it('emits an event', async () => {
          await expect(tx).to.emit(router, 'OffRampRemoved').withArgs(addr)
        })
      })
    })
  })

  describe('#routeMessage', () => {
    let message: CCIPMessage
    beforeEach(async () => {
      message = {
        sequenceNumber: BigNumber.from(1),
        sourceChainId: BigNumber.from(1),
        sender: await roles.consumer.getAddress(),
        payload: {
          tokens: [],
          amounts: [],
          destinationChainId: BigNumber.from(2),
          receiver: receiver.address,
          executor: hre.ethers.constants.AddressZero,
          data: hre.ethers.constants.HashZero,
          options: hre.ethers.constants.HashZero,
        },
      }
    })
    describe('failure', () => {
      it('fails if called by a non-offRamp', async () => {
        await evmRevert(
          router
            .connect(roles.oracleNode)
            .routeMessage(receiver.address, message),
          `OffRampNotConfigured("${await roles.oracleNode.getAddress()}")`,
        )
      })
      it('emits a message failure if the receiver is not a contract', async () => {
        await evmRevert(
          router
            .connect(roles.defaultAccount)
            .routeMessage(hre.ethers.constants.AddressZero, message),
          `function call to a non-contract account`,
        )
      })
    })
    describe('success', () => {
      it('routes message to the receiver', async () => {
        const tx: ContractTransaction = await router
          .connect(roles.defaultAccount)
          .routeMessage(receiver.address, message)
        await expect(tx).to.emit(receiver, 'MessageReceived')
        const response = await receiver.s_message()
        messageDeepEqual(response, message)
      })
    })
  })
})
