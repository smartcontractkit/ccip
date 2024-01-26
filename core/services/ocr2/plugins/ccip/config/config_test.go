package config

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip/internal/pricegetter"
)

func TestCommitConfig(t *testing.T) {

	tests := []struct {
		name                    string
		cfg                     CommitPluginJobSpecConfig
		expectedValidationError error
	}{
		{
			name: "valid config",
			cfg: CommitPluginJobSpecConfig{
				SourceStartBlock: 222,
				DestStartBlock:   333,
				OffRamp:          common.HexToAddress("0x123"),
				PriceGetterConfig: `
				{
					"aggregatorPrices": {
					 "0x0820c05e1fba1244763a494a52272170c321cad3": {
					  "chainID": "1000",
					  "contractAddress": "0xb8dabd288955d302d05ca6b011bb46dfa3ea7acf"
					 },
					 "0x4a98bb4d65347016a7ab6f85bea24b129c9a1272": {
					  "chainID": "1337",
					  "contractAddress": "0xb80244cc8b0bb18db071c150b36e9bcb8310b236"
					 }
					},
					"staticPrices": {
						"0xec8c353470ccaa4f43067fcde40558e084a12927": {
							"chainID": "1057",
							"price": "1000000000000000000"
						}
					}
				}
				`,
			},
			expectedValidationError: nil,
		},
		{
			name: "invalid chainID format",
			cfg: CommitPluginJobSpecConfig{
				SourceStartBlock: 222,
				DestStartBlock:   333,
				OffRamp:          common.HexToAddress("0x123"),
				PriceGetterConfig: `
				{
					"aggregatorPrices": {
					 "0x0820c05e1fba1244763a494a52272170c321cad3": {
					  "chainID": "1000",
					  "contractAddress": "0xb8dabd288955d302d05ca6b011bb46dfa3ea7acf"
					 },
					 "0x4a98bb4d65347016a7ab6f85bea24b129c9a1272": {
					  "chainID": "1337",
					  "contractAddress": "0xb80244cc8b0bb18db071c150b36e9bcb8310b236"
					 }
					},
					"staticPrices": {
						"0xec8c353470ccaa4f43067fcde40558e084a12927": {
							"chainID": 1057,
							"price": "1000000000000000000"
						}
					}
				}
				`,
			},
			expectedValidationError: fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal unquoted value into uint64"),
		},
		{
			name: "missing dynamic aggregator contract address",
			cfg: CommitPluginJobSpecConfig{
				SourceStartBlock: 222,
				DestStartBlock:   333,
				OffRamp:          common.HexToAddress("0x123"),
				PriceGetterConfig: `
				{
					"aggregatorPrices": {
					 "0x0820c05e1fba1244763a494a52272170c321cad3": {
					  "chainID": "1000",
					  "contractAddress": "0xb8dabd288955d302d05ca6b011bb46dfa3ea7acf"
					 },
					 "0x4a98bb4d65347016a7ab6f85bea24b129c9a1272": {
					  "chainID": "1337",
					  "contractAddress": ""
					 }
					},
					"staticPrices": {
						"0xec8c353470ccaa4f43067fcde40558e084a12927": {
							"chainID": "1057",
							"price": "1000000000000000000"
						}
					}
				}
				`,
			},
			expectedValidationError: fmt.Errorf("hex string has length 0, want 40 for common.Address"),
		},
		{
			name: "missing static price",
			cfg: CommitPluginJobSpecConfig{
				SourceStartBlock: 222,
				DestStartBlock:   333,
				OffRamp:          common.HexToAddress("0x123"),
				PriceGetterConfig: `
				{
					"aggregatorPrices": {
					 "0x0820c05e1fba1244763a494a52272170c321cad3": {
					  "chainID": "1000",
					  "contractAddress": "0xb8dabd288955d302d05ca6b011bb46dfa3ea7acf"
					 },
					 "0x4a98bb4d65347016a7ab6f85bea24b129c9a1272": {
					  "chainID": "1337",
					  "contractAddress": "0xb80244cc8b0bb18db071c150b36e9bcb8310b236"
					 }
					},
					"staticPrices": {
						"0xec8c353470ccaa4f43067fcde40558e084a12927": {
							"chainID": "1057",
							"price": ""
						}
					}
				}
				`,
			},
			expectedValidationError: fmt.Errorf("json: invalid use of ,string struct tag, trying to unmarshal \"\" into uint64"),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bts, err := json.Marshal(test.cfg)
			require.NoError(t, err)
			parsedConfig := CommitPluginJobSpecConfig{}
			require.NoError(t, json.Unmarshal(bts, &parsedConfig))
			require.Equal(t, test.cfg, parsedConfig)

			// Ensure correctness of price getter configuration.
			pgc, err := pricegetter.NewDynamicPriceGetterConfig(test.cfg.PriceGetterConfig)
			if test.expectedValidationError != nil {
				require.ErrorContains(t, err, test.expectedValidationError.Error())
			} else {
				require.NoError(t, err)
				require.Equal(t, uint64(1000), pgc.AggregatorPrices[common.HexToAddress("0x0820c05e1fba1244763a494a52272170c321cad3")].ChainID)
				require.Equal(t, uint64(1337), pgc.AggregatorPrices[common.HexToAddress("0x4a98bb4d65347016a7ab6f85bea24b129c9a1272")].ChainID)
				require.Equal(t, uint64(1057), pgc.StaticPrices[common.HexToAddress("0xec8c353470ccaa4f43067fcde40558e084a12927")].ChainID)
			}
		})
	}
}

func TestExecutionConfig(t *testing.T) {
	exampleConfig := ExecutionPluginJobSpecConfig{
		SourceStartBlock: 222,
		DestStartBlock:   333,
	}

	bts, err := json.Marshal(exampleConfig)
	require.NoError(t, err)

	parsedConfig := ExecutionPluginJobSpecConfig{}
	require.NoError(t, json.Unmarshal(bts, &parsedConfig))

	require.Equal(t, exampleConfig, parsedConfig)
}

func TestUSDCValidate(t *testing.T) {
	testcases := []struct {
		config USDCConfig
		err    string
	}{
		{
			config: USDCConfig{},
			err:    "AttestationAPI is required",
		},
		{
			config: USDCConfig{
				AttestationAPI: "api",
			},
			err: "SourceTokenAddress is required",
		},
		{
			config: USDCConfig{
				AttestationAPI:     "api",
				SourceTokenAddress: utils.ZeroAddress,
			},
			err: "SourceTokenAddress is required",
		},
		{
			config: USDCConfig{
				AttestationAPI:     "api",
				SourceTokenAddress: utils.RandomAddress(),
			},
			err: "SourceMessageTransmitterAddress is required",
		},
		{
			config: USDCConfig{
				AttestationAPI:                  "api",
				SourceTokenAddress:              utils.RandomAddress(),
				SourceMessageTransmitterAddress: utils.ZeroAddress,
			},
			err: "SourceMessageTransmitterAddress is required",
		},
		{
			config: USDCConfig{
				AttestationAPI:                  "api",
				SourceTokenAddress:              utils.RandomAddress(),
				SourceMessageTransmitterAddress: utils.RandomAddress(),
			},
			err: "",
		},
		{
			config: USDCConfig{
				AttestationAPI:                  "api",
				SourceTokenAddress:              utils.RandomAddress(),
				SourceMessageTransmitterAddress: utils.RandomAddress(),
				AttestationAPITimeoutSeconds:    -1,
			},
			err: "AttestationAPITimeoutSeconds must be non-negative",
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(fmt.Sprintf("error = %s", tc.err), func(t *testing.T) {
			t.Parallel()
			err := tc.config.ValidateUSDCConfig()
			if tc.err != "" {
				require.ErrorContains(t, err, tc.err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
