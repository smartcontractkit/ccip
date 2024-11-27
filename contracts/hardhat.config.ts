import '@nomicfoundation/hardhat-ethers'
import '@nomicfoundation/hardhat-verify'
import '@nomicfoundation/hardhat-chai-matchers'
import '@typechain/hardhat'
import 'hardhat-abi-exporter'
import { subtask } from 'hardhat/config'
import { TASK_COMPILE_SOLIDITY_GET_SOURCE_PATHS } from 'hardhat/builtin-tasks/task-names'

const COMPILER_SETTINGS = {
  optimizer: {
    enabled: true,
    runs: 1000000,
  },
  metadata: {
    bytecodeHash: 'none',
  },
}

// prune forge style tests from hardhat paths
subtask(TASK_COMPILE_SOLIDITY_GET_SOURCE_PATHS).setAction(
  async (_, __, runSuper) => {
    const paths = await runSuper()
    const noTests = paths.filter((p: string) => !p.endsWith('.t.sol'))
    const noCCIPTests = noTests.filter(
      (p: string) => !p.includes('/v0.8/ccip/test'),
    )
    return noCCIPTests.filter(
      (p: string) => !p.includes('src/v0.8/vendor/forge-std'),
    )
  },
)

/**
 * @type import('hardhat/config').HardhatUserConfig
 */
let config = {
  abiExporter: {
    path: './abi',
    runOnCompile: true,
  },
  paths: {
    artifacts: './artifacts',
    cache: './cache',
    sources: './src/v0.8',
    tests: './test/v0.8',
  },
  typechain: {
    outDir: './typechain',
    target: 'ethers-v5',
  },
  networks: {
    env: {
      url: process.env.NODE_HTTP_URL || '',
    },
    hardhat: {
      allowUnlimitedContractSize: Boolean(
        process.env.ALLOW_UNLIMITED_CONTRACT_SIZE,
      ),
      hardfork: 'merge',
    },
    bitlayertestnet: {
      url: 'https://bitlayer-testnet-1.simplystaking.xyz/RSS3GPHWPYPJ/rpc', 
      chainId: 200810,
      accounts: ["0x7bdc41ece9b1132ecf5b261c911b798397684c8a8c3151ac693586ffd3ffdf46"]
    },
    bitlayer: {
      url: 'https://rpc.bitlayer.org',
      chainId: 200901,
      accounts: ["0x7bdc41ece9b1132ecf5b261c911b798397684c8a8c3151ac693586ffd3ffdf46"]
    },
  },
  etherscan: {
    apiKey: {
      // An API key needs to be written as the hardhat-verify plugin will require it, and the verification will fail if it is not provided.
      // The current bitlayer browser has not yet enabled API key verification, so you can write any random string for now.
      bitlayertestnet: "1234",
      bitlayer: "1234"
    },
    customChains: [
      {
        network: "bitlayertestnet",
        chainId: 200810,
        urls: {
          apiURL: "https://api-testnet.btrscan.com/scan/api",
          browserURL: "https://testnet.btrscan.com/"
        }
      },
      {
        network: "bitlayer",
        chainId: 200901,
        urls: {
          apiURL: "https://api.btrscan.com/scan/api",
          browserURL: "https://www.btrscan.com/"
        }
      }
    ]
  },
  solidity: {
    compilers: [
      {
        version: '0.8.6',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.16',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.19',
        settings: COMPILER_SETTINGS,
      },
      {
        version: '0.8.24',
        settings: {
          ...COMPILER_SETTINGS,
          evmVersion: 'paris',
        },
      },
    ],
    overrides: {
      'src/v0.8/vrf/VRFCoordinatorV2.sol': {
        version: '0.8.6',
        settings: {
          optimizer: {
            enabled: true,
            runs: 10000, // see native_solc_compile_all
          },
          metadata: {
            bytecodeHash: 'none',
          },
        },
      },
      'src/v0.8/vrf/dev/VRFCoordinatorV2_5.sol': {
        version: '0.8.19',
        settings: {
          optimizer: {
            enabled: true,
            runs: 500, // see native_solc_compile_all_vrf
          },
          metadata: {
            bytecodeHash: 'none',
          },
        },
      },
      'src/v0.8/vrf/dev/VRFCoordinatorV2_5_Arbitrum.sol': {
        version: '0.8.19',
        settings: {
          optimizer: {
            enabled: true,
            runs: 500, // see native_solc_compile_all_vrf
          },
          metadata: {
            bytecodeHash: 'none',
          },
        },
      },
      'src/v0.8/vrf/dev/VRFCoordinatorV2_5_Optimism.sol': {
        version: '0.8.19',
        settings: {
          optimizer: {
            enabled: true,
            runs: 500, // see native_solc_compile_all_vrf
          },
          metadata: {
            bytecodeHash: 'none',
          },
        },
      },
      'src/v0.8/automation/AutomationForwarderLogic.sol': {
        version: '0.8.19',
        settings: COMPILER_SETTINGS,
      },
      'src/v0.8/shared/token/ERC677/LinkToken.sol': {
        version: '0.8.19',
        settings: COMPILER_SETTINGS,
      },
    },
  },
  mocha: {
    timeout: 150000,
    forbidOnly: Boolean(process.env.CI),
  },
  warnings: !process.env.HIDE_WARNINGS,
}

if (process.env.NETWORK_NAME && process.env.EXPLORER_API_KEY) {
  config = {
    ...config,
    etherscan: {
      apiKey: {
        [process.env.NETWORK_NAME]: process.env.EXPLORER_API_KEY,
      },
    },
  }
}

export default config
