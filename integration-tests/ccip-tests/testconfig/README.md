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
CCIP Deployment data containing all necessary contract addresses for various networks. This is mandatory if the test are to be run for existing deployments. 
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
            "version" : "1.4.0",
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
Specifies the environment details for the test to be run.
### CCIP.Env.Networks
Specifies the network details for the test to be run.