import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockERC20, MockOffRamp, ReceiverDapp } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { evmRevert } from '../../../test-helpers/matchers'
import { CCIPMessage } from '../../../test-helpers/ccip/ccip'
import { BigNumber } from '@ethersproject/bignumber'

const { deployContract } = hre.waffle

let roles: Roles

let ReceiverArtifact: Artifact
let RampArtifact: Artifact
let TokenArtifact: Artifact

let ramp: MockOffRamp
let receiverContract: ReceiverDapp
let token: MockERC20

let balance: BigNumber

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('ReceiverDapp', () => {
  beforeEach(async () => {
    balance = BigNumber.from('12000000000000000000')
    ReceiverArtifact = await hre.artifacts.readArtifact('ReceiverDapp')
    RampArtifact = await hre.artifacts.readArtifact('MockOffRamp')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')

    ramp = <MockOffRamp>await deployContract(roles.defaultAccount, RampArtifact)

    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        balance,
      ])
    )

    receiverContract = <ReceiverDapp>(
      await deployContract(roles.defaultAccount, ReceiverArtifact, [
        ramp.address,
        token.address,
      ])
    )

    await token
      .connect(roles.defaultAccount)
      .transfer(receiverContract.address, balance)

    await ramp.setToken(token.address)
  })

  describe('#constructor', () => {
    it('sets the off ramp', async () => {
      const response = await receiverContract.ROUTER()
      expect(response).to.equal(ramp.address)
    })
  })

  describe('#receiveMessage', () => {
    let accountAddr: string

    it('fails if the sender is not the off ramp', async () => {
      const message: CCIPMessage = {
        sequenceNumber: BigNumber.from(1),
        sourceChainId: BigNumber.from(1),
        sender: ethers.constants.AddressZero,
        payload: {
          destinationChainId: BigNumber.from(2),
          receiver: ethers.constants.AddressZero,
          data: ethers.constants.HashZero,
          tokens: [],
          amounts: [],
          executor: ethers.constants.AddressZero,
        },
      }
      accountAddr = await roles.defaultAccount.getAddress()
      await evmRevert(
        receiverContract.connect(roles.defaultAccount).receiveMessage(message),
        `InvalidDeliverer("${accountAddr}")`,
      )
    })
    describe('success', () => {
      let data: string
      let sequenceNumber: BigNumber
      let amount: BigNumber

      beforeEach(async () => {
        accountAddr = await roles.defaultAccount.getAddress()
        data = ethers.utils.defaultAbiCoder.encode(
          ['address', 'address'],
          [accountAddr, accountAddr],
        )
        sequenceNumber = BigNumber.from(1)
        amount = balance
        const message: CCIPMessage = {
          sequenceNumber,
          sourceChainId: BigNumber.from(5),
          sender: receiverContract.address,
          payload: {
            destinationChainId: BigNumber.from(2),
            receiver: receiverContract.address,
            data,
            tokens: [token.address],
            amounts: [amount],
            executor: ethers.constants.AddressZero,
          },
        }
        await ramp.deliverMessageTo(receiverContract.address, message)
      })

      it('forwards the tokens', async () => {
        expect(await token.balanceOf(accountAddr)).to.equal(amount)
      })
    })
  })

  describe('#typeAndVersion', () => {
    it('returns the type and version', async () => {
      const response = await receiverContract.typeAndVersion()
      expect(response).to.equal('ReceiverDapp 0.0.1')
    })
  })
})
