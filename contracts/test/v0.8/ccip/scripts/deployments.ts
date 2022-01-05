import { Contract, ContractFactory } from '@ethersproject/contracts'
import { NonceManager } from '@ethersproject/experimental'
import { HardhatRuntimeEnvironment } from 'hardhat/types'

interface PoolInformation {
  chainId: number
  link: string
  pool: Contract
  wallet: NonceManager
}

interface Lane {
  source: {
    eoaSender: Contract
    onRamp: Contract
    poolInfo: PoolInformation
  }
  destination: {
    eoaReceiver: Contract
    offRamp: Contract
    poolInfo: PoolInformation
  }
}

export async function deployPools(
  _: any,
  hre: HardhatRuntimeEnvironment,
  wallets: any,
  envVars: any,
) {
  const poolFactory: ContractFactory = await hre.ethers.getContractFactory(
    'LockUnlockPool',
  )
  console.log('Deploying to Kovan...')
  const kovanPool = await poolFactory
    .connect(wallets.kovan)
    .deploy(envVars.KOVAN_LINK)
  const kovanPoolVerifyCmd = `yarn hardhat verify --network kovan ${kovanPool.address} ${envVars.KOVAN_LINK}`
  console.log('Deploying to Rinkeby...')
  const rinkebyPool = await poolFactory
    .connect(wallets.rinkeby)
    .deploy(envVars.RINKEBY_LINK)
  const rinkebyPoolVerifyCmd = `yarn hardhat verify --network rinkeby ${rinkebyPool.address} ${envVars.RINKEBY_LINK}`
  console.log('Deploying to Georli...')
  const goerliPool = await poolFactory
    .connect(wallets.goerli)
    .deploy(envVars.GOERLI_LINK)
  const goerliPoolVerifyCmd = `yarn hardhat verify --network goerli ${goerliPool.address} ${envVars.GOERLI_LINK}`

  const output = `
==== DEPLOYED POOLS ====
Kovan:    ${kovanPool.address}
Rinkeby:  ${rinkebyPool.address}
Goerli:   ${goerliPool.address}
==== -------------- ====
To verify the pools, run the following commands:

${kovanPoolVerifyCmd}

${rinkebyPoolVerifyCmd}

${goerliPoolVerifyCmd}`

  console.log(output)
}

/**
 * Create a lane between two networks
 * @param hre
 * @param wallets
 * @param envVars
 */
async function createLane(
  hre: HardhatRuntimeEnvironment,
  network1: PoolInformation,
  network2: PoolInformation,
): Promise<Lane> {
  const offRampFactory: ContractFactory = await hre.ethers.getContractFactory(
    'SingleTokenOffRamp',
  )
  const eoaReceiverFactory: ContractFactory =
    await hre.ethers.getContractFactory('EOASingleTokenReceiver')
  const onRampFactory: ContractFactory = await hre.ethers.getContractFactory(
    'SingleTokenOnRamp',
  )
  const eoaSenderFactory: ContractFactory = await hre.ethers.getContractFactory(
    'EOASingleTokenSender',
  )

  // Set up network1 -> network2 lane end to end
  // network2 side receiving contracts first
  console.log('deploy offramp')
  const oneToTwoOffRamp = await offRampFactory
    .connect(network2.wallet)
    .deploy(
      network1.chainId!,
      network2.chainId!,
      network2.link,
      network2.pool.address,
    )
  const setOffRampGas = await network2.pool
    .connect(network2.wallet)
    .estimateGas.setOnRamp(oneToTwoOffRamp.address, true)
  console.log('set off ramp with gas', setOffRampGas.toString())
  await (
    await network2.pool
      .connect(network2.wallet)
      .setOffRamp(oneToTwoOffRamp.address, true, {
        gasLimit: setOffRampGas.add(10000),
      })
  ).wait()
  console.log('deploy receiver')
  const oneToTwoEOAReceiver = await eoaReceiverFactory
    .connect(network2.wallet)
    .deploy(oneToTwoOffRamp.address)
  // network 1 side sending contracts next
  console.log('deploy onramp')
  const oneToTwoOnRamp = await onRampFactory
    .connect(network1.wallet)
    .deploy(
      network1.chainId,
      network1.link,
      network1.pool.address,
      network2.chainId,
      network2.link,
    )
  const setOnRampGas = await network1.pool
    .connect(network1.wallet)
    .estimateGas.setOnRamp(oneToTwoOnRamp.address, true)
  console.log('set on ramp with gas', setOnRampGas.toString())
  await (
    await network1.pool
      .connect(network1.wallet)
      .setOnRamp(oneToTwoOnRamp.address, true, {
        gasLimit: setOnRampGas.add(10000),
      })
  ).wait()
  console.log('deploy sender')
  const oneToTwoEOASender = await eoaSenderFactory
    .connect(network1.wallet)
    .deploy(oneToTwoOnRamp.address, oneToTwoEOAReceiver.address)

  const lane: Lane = {
    source: {
      eoaSender: oneToTwoEOASender,
      onRamp: oneToTwoOnRamp,
      poolInfo: network1,
    },
    destination: {
      eoaReceiver: oneToTwoEOAReceiver,
      offRamp: oneToTwoOffRamp,
      poolInfo: network2,
    },
  }
  return lane
}

/**
 * Deploy ramps to create lanes between networks
 * @param args
 * @param hre
 * @param wallets
 * @param envVars
 */
export async function deployRamps(
  args: any,
  hre: HardhatRuntimeEnvironment,
  wallets: any,
  envVars: any,
) {
  const poolFactory: ContractFactory = await hre.ethers.getContractFactory(
    'LockUnlockPool',
  )

  const kovan: PoolInformation = {
    chainId: envVars.KOVAN_ID,
    link: envVars.KOVAN_LINK,
    pool: await poolFactory.connect(wallets.kovan).attach(args.kp),
    wallet: wallets.kovan,
  }
  const rinkeby: PoolInformation = {
    chainId: envVars.RINKEBY_ID,
    link: envVars.RINKEBY_LINK,
    pool: await poolFactory.connect(wallets.rinkeby).attach(args.rp),
    wallet: wallets.rinkeby,
  }
  const goerli: PoolInformation = {
    chainId: envVars.GOERLI_ID,
    link: envVars.GOERLI_LINK,
    pool: await poolFactory.connect(wallets.goerli).attach(args.gp),
    wallet: wallets.goerli,
  }

  // kovan->rinkeby
  console.log('deploying kovan to rinkeby lane')
  const kovanToRinkeby: Lane = await createLane(hre, kovan, rinkeby)
  // rinkeby->kovan
  console.log('deploying rinkeby to kovan lane')
  const rinkebyToKovan: Lane = await createLane(hre, rinkeby, kovan)

  // kovan->goerli
  console.log('deploying kovan to goerli lane')
  const kovanToGoerli: Lane = await createLane(hre, kovan, goerli)
  // goerli->kovan
  console.log('deploying goerli to kovan lane')
  const goerliToKovan: Lane = await createLane(hre, goerli, kovan)

  // rinkeby->goerli
  console.log('deploying rinkeby to goerli lane')
  const rinkebyToGoerli: Lane = await createLane(hre, rinkeby, goerli)
  // goerli->rinkeby
  console.log('deploying goerli to rinkeby lane')
  const goerliToRinkeby: Lane = await createLane(hre, goerli, rinkeby)

  console.log(
    `
==== DEPLOYED LANES ====

--- KOVAN <-> RINKEBY ---
- KOVAN -
EOASender:      ${kovanToRinkeby.source.eoaSender.address}
OnRamp:         ${kovanToRinkeby.source.onRamp.address}
OffRamp:        ${rinkebyToKovan.destination.offRamp.address}
EOAReceiver:    ${rinkebyToKovan.destination.eoaReceiver.address}

- Rinkeby -
EOASender:      ${rinkebyToKovan.source.eoaSender.address}
OnRamp:         ${rinkebyToKovan.source.onRamp.address}
OffRamp:        ${kovanToRinkeby.destination.offRamp.address}
EOAReceiver:    ${kovanToRinkeby.destination.eoaReceiver.address}

--- KOVAN <-> GOERLI ---
- KOVAN -
EOASender:      ${kovanToGoerli.source.eoaSender.address}
OnRamp:         ${kovanToGoerli.source.onRamp.address}
OffRamp:        ${goerliToKovan.destination.offRamp.address}
EOAReceiver:    ${goerliToKovan.destination.eoaReceiver.address}

- Goerli -
EOASender:      ${goerliToKovan.source.eoaSender.address}
OnRamp:         ${goerliToKovan.source.onRamp.address}
OffRamp:        ${kovanToGoerli.destination.offRamp.address}
EOAReceiver:    ${kovanToGoerli.destination.eoaReceiver.address}

--- RINKEBY <-> GOERLI ---
- Rinkeby -
EOASender:      ${rinkebyToGoerli.source.eoaSender.address}
OnRamp:         ${rinkebyToGoerli.source.onRamp.address}
OffRamp:        ${goerliToRinkeby.destination.offRamp.address}
EOAReceiver:    ${goerliToRinkeby.destination.eoaReceiver.address}

- Goerli -
EOASender:      ${goerliToRinkeby.source.eoaSender.address}
OnRamp:         ${goerliToRinkeby.source.onRamp.address}
OffRamp:        ${rinkebyToGoerli.destination.offRamp.address}
EOAReceiver:    ${rinkebyToGoerli.destination.eoaReceiver.address}
`,
  )
}

export async function transferOwnership(
  args: any,
  hre: HardhatRuntimeEnvironment,
  wallets: any,
) {
  const contractFactory = await hre.ethers.getContractFactory(
    'src/v0.8/ConfirmedOwner.sol:ConfirmedOwner',
  )
  let wallet
  if (args.chain == 'kovan') {
    wallet = wallets.kovan
  } else if (args.chain == 'rinkeby') {
    wallet = wallets.rinkeby
  } else if (args.chain == 'goerli') {
    wallet = wallets.goerli
  } else {
    throw new Error("Chain config doesn't exist")
  }

  const contract = await contractFactory.connect(wallet).attach(args.contract)
  const gasLimit = await contract
    .connect(wallet)
    .estimateGas.transferOwnership(args.to)
  console.log('Transfer ownership with gas', gasLimit.toString())
  await (
    await contract
      .connect(wallet)
      .transferOwnership(args.to, { gasLimit: gasLimit.add(10000) })
  ).wait()
}
