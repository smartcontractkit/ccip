import { expect } from 'chai'
import { BigNumber } from 'ethers'
import { ContractTransaction } from 'ethers'
import { constants } from 'ethers'
import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import { MockERC20, MockOnRamp, OnRampRouter } from '../../../../typechain'
import { CCIPMessagePayload } from '../../../test-helpers/ccip/ccip'
import { evmRevert } from '../../../test-helpers/matchers'
import { getUsers, Roles } from '../../../test-helpers/setup'

const { deployContract } = hre.waffle

let roles: Roles

let TokenArtifact: Artifact
let OnRampArtifact: Artifact
let RouterArtifact: Artifact

const chainId: number = 1
const destinationChainIds: Array<number> = [2, 3]
const pool: string = constants.AddressZero

let tokens: Array<MockERC20>
let onRamps: Array<MockOnRamp>
let router: OnRampRouter
let mintAmount: BigNumber

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('OnRampRouter', () => {
  beforeEach(async () => {
    mintAmount = BigNumber.from('1000000000000000000')

    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    OnRampArtifact = await hre.artifacts.readArtifact('MockOnRamp')
    RouterArtifact = await hre.artifacts.readArtifact('OnRampRouter')

    tokens = new Array<MockERC20>()
    onRamps = new Array<MockOnRamp>()
    router = <OnRampRouter>(
      await deployContract(roles.defaultAccount, RouterArtifact)
    )
    for (let i = 0; i < destinationChainIds.length; i++) {
      const destinationChainId = destinationChainIds[i]

      tokens.push(
        <MockERC20>(
          await deployContract(roles.defaultAccount, TokenArtifact, [
            'TOKEN',
            'TOKEN',
            await roles.defaultAccount.getAddress(),
            mintAmount,
          ])
        ),
      )

      onRamps.push(
        <MockOnRamp>(
          await deployContract(roles.defaultAccount, OnRampArtifact, [
            chainId,
            tokens[i].address,
            tokens[i].address,
            pool,
            destinationChainId,
          ])
        ),
      )

      await router
        .connect(roles.defaultAccount)
        .setOnRamp(destinationChainId, onRamps[i].address)
    }
  })

  describe('#requestCrossChainSend', () => {
    let receiver: string
    let messagedata: string
    let options: string
    let amounts: Array<BigNumber>
    let payload: CCIPMessagePayload

    beforeEach(async () => {
      const amount = BigNumber.from('10000000000000000')
      receiver = await roles.stranger.getAddress()
      messagedata = hre.ethers.constants.HashZero
      options = hre.ethers.constants.HashZero
      amounts = [amount, amount]
      payload = {
        receiver: receiver,
        data: messagedata,
        tokens: tokens.map((t) => t.address),
        amounts: amounts,
        destinationChainId: BigNumber.from(destinationChainIds[0]),
        executor: hre.ethers.constants.AddressZero,
        options: options,
      }
    })

    describe('failure', () => {
      it('fails if the onRamp is not supported', async () => {
        payload.destinationChainId = BigNumber.from(55)
        await evmRevert(
          router.connect(roles.defaultAccount).requestCrossChainSend(payload),
          `UnsupportedDestinationChain(${payload.destinationChainId})`,
        )
      })
      it('fails if the number of tokens does not equal the amounts', async () => {
        payload.tokens = [tokens[0].address]
        await evmRevert(
          router.connect(roles.defaultAccount).requestCrossChainSend(payload),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if the onRamp does not have approval on the token', async () => {
        await evmRevert(
          router.connect(roles.defaultAccount).requestCrossChainSend(payload),
          `ERC20: transfer amount exceeds allowance`,
        )
      })
    })
    describe('success', () => {
      let senderAddress: string
      beforeEach(async () => {
        senderAddress = await roles.defaultAccount.getAddress()
        for (let i = 0; i < tokens.length; i++) {
          const token = tokens[i]
          await token
            .connect(roles.defaultAccount)
            .approve(router.address, amounts[i])
        }
      })
      it('transfers the tokens from the sender to the router', async () => {
        const token1SenderBalanceBefore = await tokens[0].balanceOf(
          senderAddress,
        )
        const token2SenderBalanceBefore = await tokens[1].balanceOf(
          senderAddress,
        )
        await router
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload)
        const token1SenderBalanceAfter = await tokens[0].balanceOf(
          senderAddress,
        )
        const token2SenderBalanceAfter = await tokens[1].balanceOf(
          senderAddress,
        )
        expect(token1SenderBalanceAfter).to.equal(
          token1SenderBalanceBefore.sub(amounts[0]),
        )
        expect(token2SenderBalanceAfter).to.equal(
          token2SenderBalanceBefore.sub(amounts[1]),
        )
      })
      it('approves the onRamp to spend the tokens', async () => {
        await router
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload)
        const allowance1 = await tokens[0].allowance(
          router.address,
          onRamps[0].address,
        )
        const allowance2 = await tokens[1].allowance(
          router.address,
          onRamps[0].address,
        )
        expect(allowance1).to.equal(payload.amounts[0])
        expect(allowance2).to.equal(payload.amounts[1])
      })
      it('calls requestCrossChainSend on the onRamp with the payload', async () => {
        await router
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload)
        const rampPayload = await onRamps[0].getMessagePayload()
        expect(rampPayload.receiver).to.equal(payload.receiver)
        expect(rampPayload.tokens).to.deep.equal(payload.tokens)
        expect(rampPayload.amounts.map((a) => a.toString())).to.deep.equal(
          payload.amounts.map((a) => a.toString()),
        )
      })
    })
  })

  describe('#setOnRamp', () => {
    let newOnRamp: MockOnRamp

    beforeEach(async () => {
      newOnRamp = <MockOnRamp>(
        await deployContract(roles.defaultAccount, OnRampArtifact, [
          chainId,
          tokens[0].address,
          tokens[0].address,
          pool,
          destinationChainIds[0],
        ])
      )
    })

    describe('failure', () => {
      it('should only be called by the owner', async () => {
        await evmRevert(
          router
            .connect(roles.stranger)
            .setOnRamp(destinationChainIds[0], newOnRamp.address),
          'Only callable by owner',
        )
      })

      it('should revert if the proposal is the same as the current', async () => {
        await evmRevert(
          router
            .connect(roles.defaultAccount)
            .setOnRamp(destinationChainIds[0], onRamps[0].address),
          `OnRampAlreadySet(${destinationChainIds[0]}, "${onRamps[0].address}")`,
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction
      let destChainId: number

      beforeEach(async () => {
        destChainId = destinationChainIds[0]
        tx = await router
          .connect(roles.defaultAccount)
          .setOnRamp(destChainId, newOnRamp.address)
      })

      it('should emit an event', async () => {
        await expect(tx)
          .to.emit(router, 'OnRampSet')
          .withArgs(destChainId, newOnRamp.address)
      })

      it('should set the correct value', async () => {
        const response = await router.getOnRamp(destChainId)
        expect(response).to.equal(newOnRamp.address)
      })
    })
  })

  describe('#isChainSupported', () => {
    it('returns true when the chain is supported', async () => {
      expect(await router.isChainSupported(destinationChainIds[0])).to.be.true
    })
    it('returns false when the chain is not supported', async () => {
      expect(await router.isChainSupported(55)).to.be.false
    })
  })
})
