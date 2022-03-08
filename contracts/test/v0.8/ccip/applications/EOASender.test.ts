import hre, { ethers } from 'hardhat'
import { expect } from 'chai'
import { Roles, getUsers } from '../../../test-helpers/setup'
import { MockOnRamp, SenderDapp, MockERC20 } from '../../../../typechain'
import { Artifact } from 'hardhat/types'
import { BigNumber } from '@ethersproject/bignumber'
import { evmRevert } from '../../../test-helpers/matchers'

const { deployContract } = hre.waffle

let roles: Roles

let SenderArtifact: Artifact
let RampArtifact: Artifact
let TokenArtifact: Artifact

let sourceChainId: BigNumber
let token: MockERC20
let destinationToken: string
let destinationChainId: BigNumber

let ramp: MockOnRamp
let senderContract: SenderDapp
let destinationContract: string

beforeEach(async () => {
  const users = await getUsers()
  roles = users.roles
  destinationContract = await users.contracts.contract8.getAddress()
})

describe('SenderDapp', () => {
  beforeEach(async () => {
    sourceChainId = BigNumber.from(1)
    destinationToken = await roles.oracleNode2.getAddress()
    destinationChainId = BigNumber.from(2)

    SenderArtifact = await hre.artifacts.readArtifact('SenderDapp')
    RampArtifact = await hre.artifacts.readArtifact('MockOnRamp')
    TokenArtifact = await hre.artifacts.readArtifact('MockERC20')

    token = <MockERC20>(
      await deployContract(roles.defaultAccount, TokenArtifact, [
        'LINK Token',
        'LINK',
        await roles.defaultAccount.getAddress(),
        BigNumber.from('10000000000000000000'),
      ])
    )

    ramp = <MockOnRamp>(
      await deployContract(roles.defaultAccount, RampArtifact, [
        sourceChainId,
        token.address,
        destinationToken,
        await roles.oracleNode3.getAddress(),
        destinationChainId,
      ])
    )

    senderContract = <SenderDapp>(
      await deployContract(roles.defaultAccount, SenderArtifact, [
        ramp.address,
        destinationChainId,
        destinationContract,
      ])
    )
  })

  describe('#constructor', () => {
    it('should set the onRamp', async () => {
      const onRamp = await senderContract.ON_RAMP()
      expect(onRamp).to.equal(ramp.address)
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
    let options: string

    beforeEach(async () => {
      senderAddress = await roles.defaultAccount.getAddress()
      destinationAddress = senderAddress
      data = ethers.utils.defaultAbiCoder.encode(
        ['address', 'address'],
        [senderAddress, destinationAddress],
      )
      amount = BigNumber.from('1000000000000000000')
      options = '0x'
    })

    it('should send a request to the onRamp', async () => {
      const expectedResponse = [
        destinationContract,
        data,
        [token.address],
        [amount],
        options,
      ]

      await token.approve(senderContract.address, amount)
      await senderContract.sendTokens(
        destinationAddress,
        [token.address],
        [amount],
        ethers.constants.AddressZero,
      )
      const response = await ramp.getMessagePayload()
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
