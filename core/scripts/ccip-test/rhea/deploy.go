package rhea

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"

	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/ccip"

	"github.com/smartcontractkit/chainlink/core/scripts/ccip-test/shared"
	helpers "github.com/smartcontractkit/chainlink/core/scripts/common"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/commit_store"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_offramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/generated/ping_pong_demo"
)

const (
	RATE_LIMIT_CAPACITY_DOLLAR        = 1e9
	RATE_LIMIT_RATE_DOLLAR            = 1e6
	MAX_DATA_SIZE                     = 200_000
	MAX_TOKEN_LENGTH                  = 5
	MAX_TX_GAS_LIMIT                  = 4e6
	PERMISSIONLESS_EXEC_THRESHOLD_SEC = 60
	DEST_GAS_OVERHEAD                 = 5000
	DEFAULT_GAS_LIMIT                 = 200_000
)

// DeployLanes will deploy all source and Destination chain contracts using the
// owner key. Only run this of the currently deployed contracts are outdated or
// when initializing a new chain.
func DeployLanes(t *testing.T, source *EvmDeploymentConfig, destination *EvmDeploymentConfig) {
	sourceChainSelector, destChainSelector := GetCCIPChainId(source.ChainConfig.EvmChainId), GetCCIPChainId(destination.ChainConfig.EvmChainId)

	// After running this code please update the configuration to reflect the newly
	// deployed contract addresses.
	// Deploy onRamps on both chains
	deploySourceContracts(t, source, destChainSelector, destination.ChainConfig.SupportedTokens)
	deploySourceContracts(t, destination, sourceChainSelector, source.ChainConfig.SupportedTokens)

	// Deploy commitStores and offRamps on both chains
	deployDestinationContracts(t, destination, sourceChainSelector, source.LaneConfig.OnRamp, source.ChainConfig.SupportedTokens)
	deployDestinationContracts(t, source, destChainSelector, destination.LaneConfig.OnRamp, destination.ChainConfig.SupportedTokens)

	SetPriceRegistryPrices(t, source, destChainSelector)
	SetPriceRegistryPrices(t, destination, sourceChainSelector)

	DeployPingPongDapps(t, source, destination)

	UpdateDeployedAt(t, source, destination)
}

func deploySourceContracts(t *testing.T, source *EvmDeploymentConfig, destChainSelector uint64, destSupportedTokens map[Token]EVMBridgedToken) {
	if source.LaneConfig.DeploySettings.DeployRamp {
		// Updates source.OnRamp if any new contracts are deployed
		deployOnRamp(t, source, destChainSelector, destSupportedTokens)
		setOnRampOnRouter(t, source, destChainSelector)
	}

	// Skip if we reuse both the onRamp and the token pools
	if source.LaneConfig.DeploySettings.DeployRamp || source.ChainConfig.DeploySettings.DeployTokenPools {
		setOnRampOnTokenPools(t, source)
	}
	source.Logger.Infof("%s contracts deployed as source chain", ccip.ChainName(int64(source.ChainConfig.EvmChainId)))
}

func deployDestinationContracts(t *testing.T, client *EvmDeploymentConfig, sourceChainSelector uint64, onRamp common.Address, supportedTokens map[Token]EVMBridgedToken) {
	// Updates destClient.LaneConfig.CommitStore if any new contracts are deployed
	deployCommitStore(t, client, sourceChainSelector, onRamp)

	if client.LaneConfig.DeploySettings.DeployCommitStore || client.ChainConfig.DeploySettings.DeployPriceRegistry {
		setPriceRegistryUpdater(t, client)
	}

	if client.LaneConfig.DeploySettings.DeployRamp {
		// Updates destClient.LaneConfig.OffRamp if any new contracts are deployed
		deployOffRamp(t, client, sourceChainSelector, supportedTokens, onRamp)
	}

	if client.LaneConfig.DeploySettings.DeployRamp || client.ChainConfig.DeploySettings.DeployTokenPools {
		setOffRampOnTokenPools(t, client)
	}

	if client.LaneConfig.DeploySettings.DeployRamp || client.ChainConfig.DeploySettings.DeployRouter {
		setOffRampOnRouter(t, sourceChainSelector, client)
	}

	client.Logger.Infof("%s contracts fully deployed as destination chain", ccip.ChainName(int64(client.ChainConfig.EvmChainId)))
}

func deployOnRamp(t *testing.T, client *EvmDeploymentConfig, destChainSelector uint64, destSupportedTokens map[Token]EVMBridgedToken) {
	if !client.LaneConfig.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OnRamp deployment, using onRamp on %s", client.LaneConfig.OnRamp)
		return
	}

	var tokensAndPools []evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool
	var tokenTransferFeeConfig []evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs
	for token, tokenConfig := range client.ChainConfig.SupportedTokens {
		if _, ok := destSupportedTokens[token]; !ok {
			// If the token is not supported on the destination chain we
			// should not enable it for this ramp. If we enable the token,
			// txs could be sent but not executed, keeping the tokens in limbo.
			continue
		}

		tokensAndPools = append(tokensAndPools, evm_2_evm_onramp.EVM2EVMOnRampTokenAndPool{
			Token: tokenConfig.Token,
			Pool:  tokenConfig.Pool,
		})
		tokenTransferFeeConfig = append(tokenTransferFeeConfig, evm_2_evm_onramp.EVM2EVMOnRampTokenTransferFeeConfigArgs{
			Token:  tokenConfig.Token,
			MinFee: 1_00,    // $1,
			MaxFee: 5000_00, // $5,000
			Ratio:  5_0,     // 5 bps
		})
	}

	var feeTokenConfig []evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs

	for _, feeToken := range client.ChainConfig.FeeTokens {
		tokenConfig := client.ChainConfig.SupportedTokens[feeToken]
		multiplier := uint64(1e18)
		// Let link cost 10% of the non-link fee. This helps with our ping pong running out of funds.
		if feeToken == LINK {
			multiplier = 1e17
		}

		feeTokenConfig = append(feeTokenConfig, evm_2_evm_onramp.EVM2EVMOnRampFeeTokenConfigArgs{
			Token:               tokenConfig.Token,
			Multiplier:          multiplier,
			NetworkFeeAmountUSD: big.NewInt(1e16),
			DestGasOverhead:     DEST_GAS_OVERHEAD,
		})
	}

	client.Logger.Infof("Deploying OnRamp: destinationChains %+v, tokensAndPools %+v", destChainSelector, tokensAndPools)
	onRampAddress, tx, _, err := evm_2_evm_onramp.DeployEVM2EVMOnRamp(
		client.Owner,  // user
		client.Client, // client
		evm_2_evm_onramp.EVM2EVMOnRampStaticConfig{
			LinkToken:         client.ChainConfig.SupportedTokens[LINK].Token,
			ChainSelector:     GetCCIPChainId(client.ChainConfig.EvmChainId),
			DestChainSelector: destChainSelector,
			DefaultTxGasLimit: DEFAULT_GAS_LIMIT,
			PrevOnRamp:        common.HexToAddress(""),
		},
		evm_2_evm_onramp.EVM2EVMOnRampDynamicConfig{
			Router:          client.ChainConfig.Router,
			PriceRegistry:   client.ChainConfig.PriceRegistry,
			MaxDataSize:     MAX_DATA_SIZE,
			MaxTokensLength: MAX_TOKEN_LENGTH,
			MaxGasLimit:     MAX_TX_GAS_LIMIT,
			Afn:             client.ChainConfig.Afn,
		},
		tokensAndPools,
		[]common.Address{}, // allow list
		evm_2_evm_onramp.RateLimiterConfig{
			Capacity:  UsdToRateLimitValue(RATE_LIMIT_CAPACITY_DOLLAR),
			Rate:      UsdToRateLimitValue(RATE_LIMIT_RATE_DOLLAR),
			IsEnabled: true,
		},
		feeTokenConfig,
		tokenTransferFeeConfig,
		[]evm_2_evm_onramp.EVM2EVMOnRampNopAndWeight{},
	)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	client.Logger.Infof(fmt.Sprintf("Onramp deployed on %s in tx %s", onRampAddress.String(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash())))
	client.LaneConfig.OnRamp = onRampAddress
}

func deployOffRamp(t *testing.T, client *EvmDeploymentConfig, sourceChainSelector uint64, sourceTokens map[Token]EVMBridgedToken, onRamp common.Address) {
	if !client.LaneConfig.DeploySettings.DeployRamp {
		client.Logger.Infof("Skipping OffRamp deployment, using offRamp on %s", client.LaneConfig.OnRamp)
		return
	}

	var syncedSourceTokens []common.Address
	var syncedDestPools []common.Address

	for tokenName, tokenConfig := range sourceTokens {
		if _, ok := client.ChainConfig.SupportedTokens[tokenName]; ok {
			syncedSourceTokens = append(syncedSourceTokens, tokenConfig.Token)
			syncedDestPools = append(syncedDestPools, client.ChainConfig.SupportedTokens[tokenName].Pool)
		} else {
			client.Logger.Warnf("Token %s not supported by destination chain", tokenName)
		}
	}

	client.Logger.Infof("Deploying OffRamp")
	offRampAddress, tx, _, err := evm_2_evm_offramp.DeployEVM2EVMOffRamp(
		client.Owner,
		client.Client,
		evm_2_evm_offramp.EVM2EVMOffRampStaticConfig{
			CommitStore:         client.LaneConfig.CommitStore,
			ChainSelector:       GetCCIPChainId(client.ChainConfig.EvmChainId),
			SourceChainSelector: sourceChainSelector,
			OnRamp:              onRamp,
			PrevOffRamp:         common.HexToAddress(""),
		},
		syncedSourceTokens,
		syncedDestPools,
		evm_2_evm_offramp.RateLimiterConfig{
			Capacity:  UsdToRateLimitValue(RATE_LIMIT_CAPACITY_DOLLAR),
			Rate:      UsdToRateLimitValue(RATE_LIMIT_RATE_DOLLAR),
			IsEnabled: true,
		},
	)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)

	client.Logger.Infof("OffRamp contract deployed on %s in tx: %s", offRampAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	client.LaneConfig.OffRamp = offRampAddress

	client.Logger.Infof(fmt.Sprintf("Offramp configured for already deployed router in tx %s", helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash())))
}

func deployCommitStore(t *testing.T, client *EvmDeploymentConfig, sourceChainSelector uint64, onRamp common.Address) {
	if !client.LaneConfig.DeploySettings.DeployCommitStore {
		client.Logger.Infof("Skipping CommitStore deployment, using CommitStore on %s", client.LaneConfig.CommitStore)
		return
	}

	client.Logger.Infof("Deploying commitStore")
	commitStoreAddress, tx, _, err := commit_store.DeployCommitStore(
		client.Owner,  // user
		client.Client, // client
		commit_store.CommitStoreStaticConfig{
			ChainSelector:       GetCCIPChainId(client.ChainConfig.EvmChainId),
			SourceChainSelector: sourceChainSelector,
			OnRamp:              onRamp,
		},
	)
	shared.RequireNoError(t, err)
	err = shared.WaitForMined(client.Logger, client.Client, tx.Hash(), true)
	shared.RequireNoError(t, err)
	client.Logger.Infof("CommitStore deployed on %s in tx: %s", commitStoreAddress.Hex(), helpers.ExplorerLink(int64(client.ChainConfig.EvmChainId), tx.Hash()))
	client.LaneConfig.CommitStore = commitStoreAddress
}

func DeployPingPongDapps(t *testing.T, sourceClient *EvmDeploymentConfig, destClient *EvmDeploymentConfig) {
	fundingAmount := big.NewInt(1e18)

	if sourceClient.LaneConfig.DeploySettings.DeployPingPongDapp {
		feeToken := sourceClient.ChainConfig.SupportedTokens[LINK].Token
		sourceClient.Logger.Infof("Deploying source chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(
			sourceClient.Owner,
			sourceClient.Client,
			sourceClient.ChainConfig.Router,
			feeToken,
		)
		shared.RequireNoError(t, err)

		err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		sourceClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(int64(sourceClient.ChainConfig.EvmChainId), tx.Hash()))

		sourceClient.LaneConfig.PingPongDapp = pingPongDappAddress
		err = FundPingPong(sourceClient, fundingAmount, feeToken)
		shared.RequireNoError(t, err)
	}

	if destClient.LaneConfig.DeploySettings.DeployPingPongDapp {
		feeToken := destClient.ChainConfig.SupportedTokens[LINK].Token
		destClient.Logger.Infof("Deploying destination chain ping pong dapp")

		pingPongDappAddress, tx, _, err := ping_pong_demo.DeployPingPongDemo(
			destClient.Owner,
			destClient.Client,
			destClient.ChainConfig.Router,
			feeToken,
		)
		shared.RequireNoError(t, err)

		err = shared.WaitForMined(destClient.Logger, destClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		destClient.Logger.Infof("Ping pong deployed on %s in tx: %s", pingPongDappAddress.Hex(), helpers.ExplorerLink(int64(destClient.ChainConfig.EvmChainId), tx.Hash()))

		destClient.LaneConfig.PingPongDapp = pingPongDappAddress
		err = FundPingPong(destClient, fundingAmount, feeToken)
		shared.RequireNoError(t, err)
	}

	if sourceClient.LaneConfig.DeploySettings.DeployPingPongDapp || destClient.LaneConfig.DeploySettings.DeployPingPongDapp {
		pingDapp, err := ping_pong_demo.NewPingPongDemo(sourceClient.LaneConfig.PingPongDapp, sourceClient.Client)
		shared.RequireNoError(t, err)

		tx, err := pingDapp.SetCounterpart(sourceClient.Owner, GetCCIPChainId(destClient.ChainConfig.EvmChainId), destClient.LaneConfig.PingPongDapp)
		shared.RequireNoError(t, err)
		err = shared.WaitForMined(sourceClient.Logger, sourceClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		sourceClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(int64(sourceClient.ChainConfig.EvmChainId), tx.Hash()))

		pongDapp, err := ping_pong_demo.NewPingPongDemo(destClient.LaneConfig.PingPongDapp, destClient.Client)
		shared.RequireNoError(t, err)

		tx, err = pongDapp.SetCounterpart(destClient.Owner, GetCCIPChainId(sourceClient.ChainConfig.EvmChainId), sourceClient.LaneConfig.PingPongDapp)
		shared.RequireNoError(t, err)
		err = shared.WaitForMined(destClient.Logger, destClient.Client, tx.Hash(), true)
		shared.RequireNoError(t, err)
		destClient.Logger.Infof("Ping pong dapp configured in tx: %s", helpers.ExplorerLink(int64(destClient.ChainConfig.EvmChainId), tx.Hash()))
	} else {
		sourceClient.Logger.Infof("Skipping ping pong deployment")
	}
}

func UsdToRateLimitValue(usd int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(1e18), big.NewInt(usd))
}
