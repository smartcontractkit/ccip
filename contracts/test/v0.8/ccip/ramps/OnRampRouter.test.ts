import { expect } from 'chai'
import { BigNumber, Signer } from 'ethers'
import { ContractTransaction } from 'ethers'
import hre from 'hardhat'
import { Artifact } from 'hardhat/types'
import {
  EVM2AnyTollOnRampRouter,
  MockERC20,
  MockOnRamp,
  MockPool,
} from '../../../../typechain'
import { evmRevert } from '../../../test-helpers/matchers'
import { getUsers, Roles } from '../../../test-helpers/setup'
import { EVM2AnyTollMessage } from '../../../test-helpers/ccip/ccip'

const { deployContract } = hre.waffle

let roles: Roles

let TokenArtifact: Artifact
let OnRampArtifact: Artifact
let RouterArtifact: Artifact
let MockPoolArtifact: Artifact

const chainId: number = 1
const destinationChainId = 2

let tokens: Array<MockERC20>
let onRamp: MockOnRamp
let pool: MockPool
let router: EVM2AnyTollOnRampRouter
let mintAmount: BigNumber
let fee = BigNumber.from(1)

before(async () => {
  const users = await getUsers()
  roles = users.roles
})

describe('EVM2AnyTollOnRampRouter', () => {
  beforeEach(async () => {
    mintAmount = BigNumber.from('1000000000000000000')

    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')
    OnRampArtifact = await hre.artifacts.readArtifact('MockOnRamp')
    RouterArtifact = await hre.artifacts.readArtifact('EVM2AnyTollOnRampRouter')
    MockPoolArtifact = await hre.artifacts.readArtifact('MockPool')

    pool = <MockPool>(
      await deployContract(roles.defaultAccount, MockPoolArtifact, [1])
    )
    tokens = new Array<MockERC20>()
    router = <EVM2AnyTollOnRampRouter>(
      await deployContract(roles.defaultAccount, RouterArtifact)
    )

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

    onRamp = <MockOnRamp>(
      await deployContract(roles.defaultAccount, OnRampArtifact, [
        chainId,
        pool.address,
        destinationChainId,
        fee,
      ])
    )

    await router
      .connect(roles.defaultAccount)
      .setOnRamp(destinationChainId, onRamp.address)
  })

  describe('#ccipSend', () => {
    let receiver: string
    let messagedata: string
    let amounts: Array<BigNumber>
    let payload: EVM2AnyTollMessage

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
        feeToken: tokens.map((t) => t.address)[0],
        feeTokenAmount: fee,
        gasLimit: 0,
      }
    })

    describe('failure', () => {
      it('fails if the number of tokens does not equal the amounts', async () => {
        payload.tokens = [tokens[0].address]
        await evmRevert(
          router
            .connect(roles.defaultAccount)
            .ccipSend(destinationChainId, payload),
          `UnsupportedNumberOfTokens()`,
        )
      })
      it('fails if the onRamp does not have approval on the token', async () => {
        await evmRevert(
          router
            .connect(roles.defaultAccount)
            .ccipSend(destinationChainId, payload),
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
            .approve(router.address, amounts[i].add(fee))
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
          .ccipSend(destinationChainId, payload)
        const token1SenderBalanceAfter = await tokens[0].balanceOf(
          senderAddress,
        )
        const token2SenderBalanceAfter = await tokens[1].balanceOf(
          senderAddress,
        )
        expect(token1SenderBalanceAfter).to.equal(
          token1SenderBalanceBefore.sub(amounts[0]).sub(fee),
        )
        expect(token2SenderBalanceAfter).to.equal(
          token2SenderBalanceBefore.sub(amounts[1]),
        )
      })
      it('sends the tokens to the pool', async () => {
        await router
          .connect(roles.defaultAccount)
          .ccipSend(destinationChainId, payload)
        const balance1 = await tokens[0].balanceOf(pool.address)
        const balance2 = await tokens[1].balanceOf(pool.address)
        expect(balance1).to.equal(BigNumber.from(payload.amounts[0]))
        expect(balance2).to.equal(payload.amounts[1])
      })
      it('calls ccipSend on the onRamp with the payload', async () => {
        await router
          .connect(roles.defaultAccount)
          .ccipSend(destinationChainId, payload)
        const rampPayload = await onRamp.getMessagePayload()
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
          destinationChainId,
          fee,
        ])
      )
    })

    describe('failure', () => {
      it('should only be called by the owner', async () => {
        await evmRevert(
          router
            .connect(roles.stranger)
            .setOnRamp(destinationChainId, newOnRamp.address),
          'Only callable by owner',
        )
      })

      it('should revert if the proposal is the same as the current', async () => {
        await evmRevert(
          router
            .connect(roles.defaultAccount)
            .setOnRamp(destinationChainId, onRamp.address),
          `OnRampAlreadySet(${destinationChainId}, "${onRamp.address}")`,
        )
      })
    })

    describe('success', () => {
      let tx: ContractTransaction

      beforeEach(async () => {
        tx = await router
          .connect(roles.defaultAccount)
          .setOnRamp(destinationChainId, newOnRamp.address)
      })

      it('should emit an event', async () => {
        await expect(tx)
          .to.emit(router, 'OnRampSet')
          .withArgs(destinationChainId, newOnRamp.address)
      })

      it('should set the correct value', async () => {
        const response = await router.getOnRamp(destinationChainId)
        expect(response).to.equal(newOnRamp.address)
      })
    })
  })

  describe('#isChainSupported', () => {
    it('returns true when the chain is supported', async () => {
      expect(await router.isChainSupported(destinationChainId)).to.be.true
    })
    it('returns false when the chain is not supported', async () => {
      expect(await router.isChainSupported(55)).to.be.false
    })
  })
})
