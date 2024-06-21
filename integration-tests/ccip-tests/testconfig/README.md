# CCIP Configuration

This README provides details on the configuration file used for setting up and running integration tests for the Cross-Chain Interoperability Protocol (CCIP).

## CCIP.ContractVersions
Specifies contract versions of different contracts to be referred by test.
Supported versions are:
- **PriceRegistry**: '1.2.0', 'Latest'
- **OffRamp**: '1.2.0', 'Latest'
- **OnRamp**: '1.2.0', 'Latest'
- **TokenPool**: '1.4.0', 'Latest'
- **CommitStore**: '1.2.0', 'Latest'

Example Usage:
```toml
[CCIP.ContractVersions]
PriceRegistry = "1.2.0"
OffRamp = "1.2.0"
OnRamp = "1.2.0"
TokenPool = "1.4.0"
CommitStore = "1.2.0"
```

## CCIP.Deployments
CCIP Deployment contains all necessary contract addresses for various networks. This is mandatory if the test are to be run for existing deployments. 
The deployment data can be specified -
 - Under `Data` field with value as stringify format of json. 
 - Under `DataFile` field with value as the path of the file containing the deployment data in json format.

The json schema is specified in https://github.com/smartcontractkit/ccip/blob/ccip-develop/integration-tests/ccip-tests/contracts/laneconfig/parse_contracts.go#L96

Example Usage:
```toml
[CCIP.Deployments]
Data = """
{
    "lane_configs": {
        "Arbitrum Mainnet": {
            "is_native_fee_token": true,
            "fee_token": "0xf97f4df75117a78c1A5a0DBb814Af92458539FB4",
            "bridge_tokens": ["0x82aF49447D8a07e3bd95BD0d56f35241523fBab1"],
            "bridge_tokens_pools": ["0x82aF49947D8a07e3bd95BD0d56f35241523fBab1"],
            "arm": "0xe06b0e8c4bd455153e8794ad7Ea8Ff5A14B64E4b",
            "router": "0x141fa059441E0ca23ce184B6A78bafD2A517DdE8",
            "price_registry": "0x13015e4E6f839E1Aa1016DF521ea458ecA20438c",
            "wrapped_native": "0x82aF49447D8a07e3bd95BD0d56f35241523fBab1",
            "src_contracts": {
                "Ethereum Mainnet": {
                    "on_ramp": "0xCe11020D56e5FDbfE46D9FC3021641FfbBB5AdEE",
                    "deployed_at": 11111111
                }
            },
            "dest_contracts": {
                "Ethereum Mainnet": {
                    "off_ramp": "0x542ba1902044069330e8c5b36A84EC503863722f",
                    "commit_store": "0x060331fEdA35691e54876D957B4F9e3b8Cb47d20",
                    "receiver_dapp": "0x1A2A69e3eB1382FE34Bc579AdD5Bae39e31d4A2c"
                }
            }
        },
        "Ethereum Mainnet": {
            "is_native_fee_token": true,
            "fee_token": "0x514910771AF9Ca656af840dff83E8264EcF986CA",
            "bridge_tokens": ["0x8B63b3DE93431C0f756A493644d128134291fA1b"],
            "bridge_tokens_pools": ["0x8B63b3DE93431C0f756A493644d128134291fA1b"],
            "arm": "0x8B63b3DE93431C0f756A493644d128134291fA1b",
            "router": "0x80226fc0Ee2b096224EeAc085Bb9a8cba1146f7D",
            "price_registry": "0x8c9b2Efb7c64C394119270bfecE7f54763b958Ad",
            "wrapped_native": "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
            "src_contracts": {
                "Arbitrum Mainnet": {
                    "on_ramp": "0x925228D7B82d883Dde340A55Fe8e6dA56244A22C",
                    "deployed_at": 11111111
                }
            },
            "dest_contracts": {
                "Arbitrum Mainnet": {
                    "off_ramp": "0xeFC4a18af59398FF23bfe7325F2401aD44286F4d",
                    "commit_store": "0x9B2EEd6A1e16cB50Ed4c876D2dD69468B21b7749",
                    "receiver_dapp": "0x1A2A69e3eB1382FE34Bc579AdD5Bae39e31d4A2c"
                }
            }
        }
    }
}
"""
```
Or 
```toml
[CCIP.Deployments]
DataFile = '<path/to/deployment.json>'
```

## CCIP.Env 
Specifies the environment details for the test to be run on.

Test needs network/chain details to be set through configuration. This configuration is mandatory for running the tests.
you have option to set the network details in two ways:
1. Using [CCIP.Env.Networks](#ccipenvnetworks) 
2. Using a separate network config file -
   * refer to the example - [network_config.toml.example](/integration-tests/ccip-tests/testconfig/examples/network_config.toml.example)
   * once all necessary values are set, encode the toml file content in base64 format,
   * set the base64'ed string content in `BASE64_NETWORK_CONFIG` environment variable.

### CCIP.Env.Networks
Specifies the network details for the test to be run.
The NetworkConfig is imported from https://github.com/smartcontractkit/chainlink-testing-framework/blob/main/config/network.go#L39

#### selected_networks
It denotes the network names in which tests will be run. These networks are used to deploy ccip contracts and set up lanes between them.
If more than 2 networks are specified, then lanes will be set up between all possible pairs of networks.

For example , if `selected_networks = ['SIMULATED_1', 'SIMULATED_2', 'SIMULATED_3']`, it denotes that lanes will be set up between SIMULATED_1 and SIMULATED_2, SIMULATED_1 and SIMULATED_3, SIMULATED_2 and SIMULATED_3
This behaviour can be varied based on `NoOfNetworks`, `NetworkPairs`,`MaxNoOfLanes` values in the TestGroupConfig.

The name of the networks are taken from [known_networks](https://github.com/smartcontractkit/chainlink-testing-framework/blob/main/networks/known_networks.go#L884) in chainlink-testing-framework
If the network is not present in known_networks, then the network details can be specified in the config file itself under the following `EVMNetworks` key.

#### CCIP.Env.Network.EVMNetworks
Specifies the network config to be used while creating blockchain EVMClient for test. 
It is a map of network name to EVMNetworks where key is network name specified under `CCIP.Env.Networks.selected_networks` and value is `EVMNetwork`. 
The EVMNetwork is imported from [EVMNetwork](https://github.com/smartcontractkit/chainlink-testing-framework/blob/main/blockchain/config.go#L43) in chainlink-testing-framework.

If `CCIP.Env.Network.EVMNetworks` config is not set for a network name specified under `CCIP.Env.Networks.selected_networks`, test will try to find the corresponding network config from defined networks in `MappedNetworks` under [known_networks.go](https://github.com/smartcontractkit/chainlink-testing-framework/blob/main/networks/known_networks.go)

#### CCIP.Env.Network.AnvilConfigs
If the test needs to run on chains created using Anvil, then the AnvilConfigs can be specified. 
It is a map of network name to `AnvilConfig` where key is network name specified under `CCIP.Env.Networks.selected_networks` and value is `AnvilConfig`. 
The AnvilConfig is imported from [AnvilConfig](https://github.com/smartcontractkit/chainlink-testing-framework/blob/main/config/network.go#L20) in chainlink-testing-framework.

#### CCIP.ENV.Network.RpcHttpUrls
RpcHttpUrls is the RPC HTTP endpoints for each network,
// key is the network name as declared in selected_networks slice

Example Usage of Network Config:
```toml
[CCIP.Env.Network]
selected_networks= ['PRIVATE-CHAIN-1', 'PRIVATE-CHAIN-2']

[CCIP.Env.Network.EVMNetworks.PRIVATE-CHAIN-1]
evm_name = 'private-chain-1'
evm_chain_id = 2337
evm_urls = ['wss://ignore-this-url.com']
evm_http_urls = ['https://ignore-this-url.com']
evm_keys = ['59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d']
evm_simulated = true
client_implementation = 'Ethereum'
evm_chainlink_transaction_limit = 5000
evm_transaction_timeout = '3m'
evm_minimum_confirmations = 1
evm_gas_estimation_buffer = 1000
evm_supports_eip1559 = true
evm_default_gas_limit = 6000000
evm_finality_depth = 400

[CCIP.Env.Network.EVMNetworks.PRIVATE-CHAIN-2]
evm_name = 'private-chain-2'
evm_chain_id = 1337
evm_urls = ['wss://ignore-this-url.com']
evm_http_urls = ['https://ignore-this-url.com']
evm_keys = ['ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80']
evm_simulated = true
client_implementation = 'Ethereum'
evm_chainlink_transaction_limit = 5000
evm_transaction_timeout = '3m'
evm_minimum_confirmations = 1
evm_gas_estimation_buffer = 1000
evm_supports_eip1559 = true
evm_default_gas_limit = 6000000
evm_finality_depth = 400

[CCIP.Env.Network.AnvilConfigs.PRIVATE-CHAIN-1]
block_time = 1
#
[CCIP.Env.Network.AnvilConfigs.PRIVATE-CHAIN-2]
block_time = 1
```