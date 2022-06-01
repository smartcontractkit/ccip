import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockOnRampRouter, SenderDapp, MockERC20 } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle

let roles: Roles

let SenderArtifact: Artifact
let RampRouterArtifact: Artifact
let TokenArtifact: Artifact

let token: MockERC20
let destinationChainId: BigNumber

let router: MockOnRampRouter
let senderContract: SenderDapp
let destinationContract: string

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
  destinationContract = await users.contracts.contract8.getAddress()
})

describe('SenderDapp', () => {
  beforeEach(async () => {
    destinationChainId = BigNumber.from(2)

    SenderArtifact = await hre.artifacts.readArtifact('SenderDapp')
    RampRouterArtifact = await hre.artifacts.readArtifact('MockOnRampRouter')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')

    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        BigNumber.from('10000000000000000000'),
      ])
    )

    router = <MockOnRampRouter>(
      await deployContract(roles.defaultAccount, RampRouterArtifact)
    )

    senderContract = <SenderDapp>(
      await deployContract(roles.defaultAccount, SenderArtifact, [
        router.address,
        destinationChainId,
        destinationContract,
      ])
    )
  })

  describe('#constructor', () => {
    it('should set the onRamp', async () => {
      const onRamp = await senderContract.ON_RAMP_ROUTER()
      expect(onRamp).to.equal(router.address)
    })

    it('#should set the destination contract', async () => {
      const destContract = await senderContract.DESTINATION_CONTRACT()
      expect(destContract).to.equal(destinationContract)
    })
  })

  describe('#sendMessage', () => {
    let senderAddress: string
    let destinationAddress: string
    let data: string
    let amount: BigNumber

    beforeEach(async () => {
      senderAddress = await roles.defaultAccount.getAddress()
      destinationAddress = senderAddress
      data = ethers.utils.defaultAbiCoder.encode(
        ['address', 'address'],
        [senderAddress, destinationAddress],
      )
      amount = BigNumber.from('1000000000000000000')
    })

    it('should send a request to the onRamp', async () => {
      const expectedResponse = [
        destinationAddress,
        data,
        [token.address],
        [amount],
      ]

      await token.approve(senderContract.address, amount)
      await senderContract.sendTokens(
        destinationAddress,
        [token.address],
        [amount],
        ethers.constants.AddressZero,
      )
      const response = await router.getMessagePayload()
      console.log(expectedResponse)
      console.log(response)
      for (let i = 0; i < response.length; i++) {
        const actual = response[i].toString()
        const expected = expectedResponse[i].toString()
        expect(actual).to.deep.equal(expected)
      }
    })

    it('should fail if the destination address is zero address', async () => {
      await evmRevert(
        senderContract.sendTokens(
          ethers.constants.AddressZero,
          [token.address],
          [amount],
          ethers.constants.AddressZero,
        ),
        `InvalidDestinationAddress("${ethers.constants.AddressZero}")`,
      )
    })
  })

  describe('#typeAndVersion', () => {
    it('should return the correct type and version', async () => {
      expect(await senderContract.typeAndVersion()).to.equal('SenderDapp 1.0.0')
    })
  })
})
