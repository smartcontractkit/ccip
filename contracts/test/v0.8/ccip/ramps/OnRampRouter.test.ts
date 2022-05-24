import { expect } from 'chai'
import { BigNumber, Signer } from 'ethers'
import { ContractTransaction } from 'ethers'
import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import {
  MockERC20,
  MockOnRamp,
  MockPool,
  OnRampRouter,
} from '../../../../typechain'
import { CCIPMessagePayload } from '../../../test-helpers/ccip/ccip'
import { evmRevert } from '../../../test-helpers/matchers'
import { getUsers, Roles } from '../../../test-helpers/setup'

const { deployContract } = hre.waffle

let roles: Roles

let TokenArtifact: Artifact
let OnRampArtifact: Artifact
let RouterArtifact: Artifact
let MockPoolArtifact: Artifact

const chainId: number = 1
const destinationChainIds: Array<number> = [2, 3]

let tokens: Array<MockERC20>
let onRamps: Array<MockOnRamp>
let pool: MockPool
let router: OnRampRouter
let mintAmount: BigNumber
let fee = BigNumber.from(1)

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
    MockPoolArtifact = await hre.artifacts.readArtifact('MockPool')

    pool = <MockPool>(
      await deployContract(roles.defaultAccount, MockPoolArtifact, [1])
    )
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
            pool.address,
            destinationChainId,
            fee,
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
    let amounts: Array<BigNumber>
    let payload: CCIPMessagePayload

    beforeEach(async () => {
      const amount = BigNumber.from('10000000000000000')
      receiver = await roles.stranger.getAddress()
      messagedata = hre.ethers.constants.HashZero
      amounts = [amount, amount]
      payload = {
        receiver: receiver,
        data: messagedata,
        tokens: tokens.map((t) => t.address),
        amounts: amounts,
        destinationChainId: BigNumber.from(destinationChainIds[0]),
        executor: hre.ethers.constants.AddressZero,
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
      it('sends the tokens to the pool', async () => {
        await router
          .connect(roles.defaultAccount)
          .requestCrossChainSend(payload)
        const balance1 = await tokens[0].balanceOf(pool.address)
        const balance2 = await tokens[1].balanceOf(pool.address)
        expect(balance1).to.equal(BigNumber.from(payload.amounts[0]).sub(fee))
        expect(balance2).to.equal(payload.amounts[1])
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

  describe('#withdrawAccumulatedFees', () => {
    let recipient: Signer
    let recipientAddress: string
    let feeToken: MockERC20
    let withdrawAmount: BigNumber

    beforeEach(async () => {
      withdrawAmount = BigNumber.from(123)
      recipient = roles.defaultAccount
      recipientAddress = await recipient.getAddress()
      feeToken = tokens[0]
      await feeToken
        .connect(roles.defaultAccount)
        .transfer(router.address, withdrawAmount)
    })

    it('success', async () => {
      const recipientBalanceBefore = await feeToken.balanceOf(recipientAddress)
      const routerBalanceBefore = await feeToken.balanceOf(router.address)

      const tx = await router
        .connect(roles.defaultAccount)
        .withdrawAccumulatedFees(
          feeToken.address,
          recipientAddress,
          withdrawAmount,
        )

      const recipientBalanceAfter = await feeToken.balanceOf(recipientAddress)
      const routerBalanceAfter = await feeToken.balanceOf(router.address)

      expect(recipientBalanceAfter).to.equal(
        recipientBalanceBefore.add(withdrawAmount),
      )
      expect(routerBalanceAfter).to.equal(
        routerBalanceBefore.sub(withdrawAmount),
      )

      await expect(tx)
        .to.emit(router, 'FeesWithdrawn')
        .withArgs(feeToken.address, recipientAddress, withdrawAmount)
    })

    describe('failure', () => {
      it('fails if called by a non-owner', async () => {
        await evmRevert(
          router
            .connect(roles.stranger)
            .withdrawAccumulatedFees(
              feeToken.address,
              recipientAddress,
              withdrawAmount,
            ),
          'Only callable by owner',
        )
      })
      it('fails if amount is greater than OnRamp balance', async () => {
        await evmRevert(
          router
            .connect(roles.defaultAccount)
            .withdrawAccumulatedFees(
              feeToken.address,
              recipientAddress,
              withdrawAmount.mul(2),
            ),
          'ERC20: transfer amount exceeds balance',
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
          pool.address,
          destinationChainIds[0],
          fee,
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
