package smoke

import (
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-testing-framework/logging"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/actions"
	"github.com/smartcontractkit/chainlink/integration-tests/ccip-tests/testsetups"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/evm_2_evm_onramp"
	"github.com/smartcontractkit/chainlink/v2/core/gethwrappers/ccip/generated/lock_release_token_pool"
)

func TestSmokeCCIPForBidirectionalLane(t *testing.T) {
	t.Parallel()
	type subtestInput struct {
		testName string
		lane     *actions.CCIPLane
	}
	l := logging.GetTestLogger(t)
	TestCfg := testsetups.NewCCIPTestConfig(t, l, testsetups.Smoke)
	transferAmounts := []*big.Int{big.NewInt(1e14), big.NewInt(1e14)}
	setUpOutput := testsetups.CCIPDefaultTestSetUp(t, l, "smoke-ccip", 6, transferAmounts, nil, 5, true, true, TestCfg)
	var tcs []subtestInput
	if len(setUpOutput.Lanes) == 0 {
		return
	}

	t.Cleanup(func() {
		if TestCfg.MsgType == actions.TokenTransfer {
			setUpOutput.Balance.Verify(t)
		}
		require.NoError(t, setUpOutput.TearDown())
	})
	for i := range setUpOutput.Lanes {
		tcs = append(tcs, subtestInput{
			testName: fmt.Sprintf("CCIP message transfer from network %s to network %s",
				setUpOutput.Lanes[i].ForwardLane.SourceNetworkName, setUpOutput.Lanes[i].ForwardLane.DestNetworkName),
			lane: setUpOutput.Lanes[i].ForwardLane,
		})
		if setUpOutput.Lanes[i].ReverseLane != nil {
			tcs = append(tcs, subtestInput{
				testName: fmt.Sprintf("CCIP message transfer from network %s to network %s",
					setUpOutput.Lanes[i].ReverseLane.SourceNetworkName, setUpOutput.Lanes[i].ReverseLane.DestNetworkName),
				lane: setUpOutput.Lanes[i].ReverseLane,
			})
		}
	}
	l.Info().Int("Total Lanes", len(tcs)).Msg("Starting CCIP test")
	for _, testcase := range tcs {
		tc := testcase
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()
			tc.lane.Test = t
			l.Info().
				Str("Source", tc.lane.SourceNetworkName).
				Str("Destination", tc.lane.DestNetworkName).
				Msgf("Starting lane %s -> %s", tc.lane.SourceNetworkName, tc.lane.DestNetworkName)

			tc.lane.RecordStateBeforeTransfer()
			err := tc.lane.SendRequests(1, TestCfg.MsgType)
			require.NoError(t, err)
			tc.lane.ValidateRequests()
		})
	}
}

func TestSmokeCCIPRateLimit(t *testing.T) {
	t.Parallel()
	type subtestInput struct {
		testName string
		lane     *actions.CCIPLane
	}
	l := logging.GetTestLogger(t)
	TestCfg := testsetups.NewCCIPTestConfig(t, l, testsetups.Smoke)
	require.Equal(t, actions.TokenTransfer, TestCfg.MsgType, "Test config should have token transfer message type")
	transferAmounts := []*big.Int{big.NewInt(1e14)}
	setUpOutput := testsetups.CCIPDefaultTestSetUp(
		t, l, "smoke-ccip", 6, transferAmounts, nil,
		5, true, true, TestCfg)
	var tcs []subtestInput
	if len(setUpOutput.Lanes) == 0 {
		return
	}

	t.Cleanup(func() {
		require.NoError(t, setUpOutput.TearDown())
	})

	for i := range setUpOutput.Lanes {
		tcs = append(tcs, subtestInput{
			testName: fmt.Sprintf("Network %s to network %s",
				setUpOutput.Lanes[i].ForwardLane.SourceNetworkName, setUpOutput.Lanes[i].ForwardLane.DestNetworkName),
			lane: setUpOutput.Lanes[i].ForwardLane,
		})
	}

	// if we are running in simulated or in testnet mode, we can set the rate limit to test friendly values
	// For mainnet, we need to set this as false to avoid changing the deployed contract config
	SetRateLimit := true
	AggregatedRateLimitCapacity := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(100))
	AggregatedRateLimitRate := big.NewInt(1e18)

	TokenPoolRateLimitCapacity := new(big.Int).Mul(big.NewInt(1e18), big.NewInt(10))
	TokenPoolRateLimitRate := big.NewInt(1e17)

	for _, testcase := range tcs {
		tc := testcase
		t.Run(fmt.Sprintf("%s - Rate Limit", tc.testName), func(t *testing.T) {
			tc.lane.Test = t
			src := tc.lane.Source
			l.Info().
				Str("Source", tc.lane.SourceNetworkName).
				Str("Destination", tc.lane.DestNetworkName).
				Msgf("Starting lane %s -> %s", tc.lane.SourceNetworkName, tc.lane.DestNetworkName)

			// capture the rate limit config before we change it
			prevRLOnRamp, err := src.OnRamp.Instance.CurrentRateLimiterState(nil)
			require.NoError(t, err)
			tc.lane.Logger.Info().Interface("rate limit", prevRLOnRamp).Msg("Initial OnRamp rate limiter state")

			prevOnRampRLTokenPool, err := src.Common.BridgeTokenPools[0].Instance.CurrentOnRampRateLimiterState(nil,
				src.OnRamp.EthAddress)
			require.NoError(t, err)
			tc.lane.Logger.Info().
				Interface("rate limit", prevOnRampRLTokenPool).
				Str("pool", src.Common.BridgeTokenPools[0].Address()).
				Str("onRamp", src.OnRamp.Address()).
				Msg("Initial tokenpool rate limiter state")

			// some sanity checks
			rlOffRamp, err := tc.lane.Dest.OffRamp.Instance.CurrentRateLimiterState(nil)
			require.NoError(t, err)
			tc.lane.Logger.Info().Interface("rate limit", rlOffRamp).Msg("Initial OffRamp rate limiter state")
			if rlOffRamp.IsEnabled {
				require.GreaterOrEqual(t, rlOffRamp.Capacity.Cmp(prevRLOnRamp.Capacity), 0, "OffRamp Aggregated capacity should be greater than or equal to OnRamp Aggregated capacity")
			}

			prevOffRampRLTokenPool, err := tc.lane.Dest.Common.BridgeTokenPools[0].Instance.CurrentOffRampRateLimiterState(nil, tc.lane.Dest.OffRamp.EthAddress)
			require.NoError(t, err)
			tc.lane.Logger.Info().
				Interface("rate limit", prevOffRampRLTokenPool).
				Str("pool", tc.lane.Dest.Common.BridgeTokenPools[0].Address()).
				Str("offRamp", tc.lane.Dest.OffRamp.Address()).
				Msg("Initial tokenpool rate limiter state")
			if prevOffRampRLTokenPool.IsEnabled {
				require.GreaterOrEqual(t, prevOffRampRLTokenPool.Capacity.Cmp(prevOnRampRLTokenPool.Capacity), 0, "OffRamp Token Pool capacity should be greater than or equal to OnRamp Token Pool capacity")
			}

			AggregatedRateLimitChanged := false
			TokenPoolRateLimitChanged := false

			// reset the rate limit config to what it was before the test
			t.Cleanup(func() {
				if AggregatedRateLimitChanged {
					require.NoError(t, src.OnRamp.SetRateLimit(evm_2_evm_onramp.RateLimiterConfig{
						IsEnabled: prevRLOnRamp.IsEnabled,
						Capacity:  prevRLOnRamp.Capacity,
						Rate:      prevRLOnRamp.Rate,
					}), "setting rate limit")
					require.NoError(t, src.Common.ChainClient.WaitForEvents(), "waiting for events")
				}
				if TokenPoolRateLimitChanged {
					require.NoError(t, src.Common.BridgeTokenPools[0].SetOnRampRateLimit(src.OnRamp.EthAddress,
						lock_release_token_pool.RateLimiterConfig{
							Capacity:  prevOnRampRLTokenPool.Capacity,
							IsEnabled: prevOnRampRLTokenPool.IsEnabled,
							Rate:      prevOnRampRLTokenPool.Rate,
						}))
					require.NoError(t, src.Common.ChainClient.WaitForEvents(), "waiting for events")
				}
			})

			if SetRateLimit {
				if prevRLOnRamp.Capacity.Cmp(AggregatedRateLimitCapacity) != 0 ||
					prevRLOnRamp.Rate.Cmp(AggregatedRateLimitRate) != 0 ||
					!prevRLOnRamp.IsEnabled {
					require.NoError(t, src.OnRamp.SetRateLimit(evm_2_evm_onramp.RateLimiterConfig{
						IsEnabled: true,
						Capacity:  AggregatedRateLimitCapacity,
						Rate:      AggregatedRateLimitRate,
					}),
						"setting rate limit on onramp")
					require.NoError(t, src.Common.ChainClient.WaitForEvents(), "waiting for events")
					AggregatedRateLimitChanged = true
				}
			} else {
				AggregatedRateLimitCapacity = prevRLOnRamp.Capacity
				AggregatedRateLimitRate = prevRLOnRamp.Rate
			}

			rlOnRamp, err := src.OnRamp.Instance.CurrentRateLimiterState(nil)
			require.NoError(t, err)
			tc.lane.Logger.Info().Interface("rate limit", rlOnRamp).Msg("OnRamp rate limiter state")
			require.True(t, rlOnRamp.IsEnabled, "OnRamp rate limiter should be enabled")

			tokenPrice, err := src.Common.PriceRegistry.Instance.GetTokenPrice(nil, src.Common.BridgeTokens[0].ContractAddress)
			require.NoError(t, err)
			tc.lane.Logger.Info().Str("tokenPrice.Value", tokenPrice.Value.String()).Msg("Price Registry Token Price")

			totalTokensForOnRampCapacity := new(big.Int).Mul(
				big.NewInt(1e18),
				new(big.Int).Div(rlOnRamp.Capacity, tokenPrice.Value))

			tc.lane.Source.Common.ChainClient.ParallelTransactions(true)

			// current tokens are equal to the full capacity  - should fail
			src.TransferAmount[0] = rlOnRamp.Tokens
			tc.lane.Logger.Info().Str("tokensTobeSent", rlOnRamp.Tokens.String()).Msg("Aggregated Capacity")
			failedTx, _, _, err := tc.lane.Source.SendRequest(
				tc.lane.Dest.ReceiverDapp.EthAddress,
				actions.TokenTransfer, "msg with token more than aggregated capacity",
				common.HexToAddress(tc.lane.Source.Common.FeeToken.Address()))
			require.NoError(t, err)
			require.Error(t, tc.lane.Source.Common.ChainClient.WaitForEvents())
			errReason, v, err := tc.lane.Source.Common.ChainClient.RevertReasonFromTx(failedTx, evm_2_evm_onramp.EVM2EVMOnRampABI)
			require.NoError(t, err)
			tc.lane.Logger.Info().
				Str("Revert Reason", errReason).
				Interface("Args", v).
				Str("TokensSent", src.TransferAmount[0].String()).
				Str("Token", tc.lane.Source.Common.BridgeTokens[0].Address()).
				Str("FailedTx", failedTx.Hex()).
				Msg("Msg sent with tokens more than AggregateValueMaxCapacity")
			require.Equal(t, "AggregateValueMaxCapacityExceeded", errReason)

			// 99% of the aggregated capacity - should succeed
			tokensTobeSent := new(big.Int).Div(new(big.Int).Mul(totalTokensForOnRampCapacity, big.NewInt(99)), big.NewInt(100))
			tc.lane.Logger.Info().Str("tokensTobeSent", tokensTobeSent.String()).Msg("99% of Aggregated Capacity")
			tc.lane.RecordStateBeforeTransfer()
			src.TransferAmount[0] = tokensTobeSent
			err = tc.lane.SendRequests(1, TestCfg.MsgType)
			require.NoError(t, err)

			// try to send again with amount more than the amount refilled by rate and
			// this should fail, as the refill rate is not enough to refill the capacity
			tokensTobeSent = new(big.Int).Mul(AggregatedRateLimitRate, big.NewInt(5))
			src.TransferAmount[0] = tokensTobeSent
			tc.lane.Logger.Info().Str("tokensTobeSent", tokensTobeSent.String()).Msg("More than Aggregated Rate")
			failedTx, _, _, err = tc.lane.Source.SendRequest(
				tc.lane.Dest.ReceiverDapp.EthAddress,
				actions.TokenTransfer, "msg with token more than aggregated rate",
				common.HexToAddress(tc.lane.Source.Common.FeeToken.Address()))
			require.NoError(t, err)
			require.Error(t, tc.lane.Source.Common.ChainClient.WaitForEvents())
			errReason, v, err = tc.lane.Source.Common.ChainClient.RevertReasonFromTx(failedTx, evm_2_evm_onramp.EVM2EVMOnRampABI)
			require.NoError(t, err)
			tc.lane.Logger.Info().
				Str("Revert Reason", errReason).
				Interface("Args", v).
				Str("TokensSent", src.TransferAmount[0].String()).
				Str("Token", tc.lane.Source.Common.BridgeTokens[0].Address()).
				Str("FailedTx", failedTx.Hex()).
				Msg("Msg sent with tokens more than AggregateValueRate")
			require.Equal(t, "AggregateValueRateLimitReached", errReason)

			// validate the  successful request was delivered to the destination
			tc.lane.ValidateRequests()

			// now set the token pool rate limit
			if SetRateLimit {
				if prevOnRampRLTokenPool.Capacity.Cmp(TokenPoolRateLimitCapacity) != 0 ||
					prevOnRampRLTokenPool.Rate.Cmp(TokenPoolRateLimitRate) != 0 ||
					!prevOnRampRLTokenPool.IsEnabled {
					require.NoError(t, src.Common.BridgeTokenPools[0].SetOnRampRateLimit(
						src.OnRamp.EthAddress,
						lock_release_token_pool.RateLimiterConfig{
							IsEnabled: true,
							Capacity:  TokenPoolRateLimitCapacity,
							Rate:      TokenPoolRateLimitRate,
						}),
						"setting rate limit on token pool")
					require.NoError(t, src.Common.ChainClient.WaitForEvents(), "waiting for events")
					TokenPoolRateLimitChanged = true
				}
			} else {
				TokenPoolRateLimitCapacity = prevOnRampRLTokenPool.Capacity
				TokenPoolRateLimitRate = prevOnRampRLTokenPool.Rate
			}

			rlOnPool, err := src.Common.BridgeTokenPools[0].Instance.CurrentOnRampRateLimiterState(nil,
				src.OnRamp.EthAddress)
			require.NoError(t, err)
			require.True(t, rlOnPool.IsEnabled, "Token Pool rate limiter should be enabled")

			// wait for the AggregateCapacity to be refilled
			time.Sleep(1 * time.Minute)

			// try to send more than token pool capacity - should fail
			tokensTobeSent = new(big.Int).Add(TokenPoolRateLimitCapacity, big.NewInt(2))
			src.TransferAmount[0] = tokensTobeSent
			tc.lane.Logger.Info().Str("tokensTobeSent", tokensTobeSent.String()).Msg("More than Token Pool Capacity")
			failedTx, _, _, err = tc.lane.Source.SendRequest(
				tc.lane.Dest.ReceiverDapp.EthAddress,
				actions.TokenTransfer, "msg with token more than token pool capacity",
				common.HexToAddress(tc.lane.Source.Common.FeeToken.Address()))
			require.NoError(t, err)
			require.Error(t, tc.lane.Source.Common.ChainClient.WaitForEvents())
			errReason, v, err = tc.lane.Source.Common.ChainClient.RevertReasonFromTx(failedTx, lock_release_token_pool.LockReleaseTokenPoolABI)
			require.NoError(t, err)
			tc.lane.Logger.Info().
				Str("Revert Reason", errReason).
				Interface("Args", v).
				Str("TokensSent", src.TransferAmount[0].String()).
				Str("Token", tc.lane.Source.Common.BridgeTokens[0].Address()).
				Str("FailedTx", failedTx.Hex()).
				Msg("Msg sent with tokens more than token pool capacity")
			require.Equal(t, "TokenMaxCapacityExceeded", errReason)

			// try to send 99% of token pool capacity - should succeed
			tokensTobeSent = new(big.Int).Div(new(big.Int).Mul(TokenPoolRateLimitCapacity, big.NewInt(99)), big.NewInt(100))
			src.TransferAmount[0] = tokensTobeSent
			tc.lane.Logger.Info().Str("tokensTobeSent", tokensTobeSent.String()).Msg("99% of Token Pool Capacity")
			tc.lane.RecordStateBeforeTransfer()
			err = tc.lane.SendRequests(1, TestCfg.MsgType)
			require.NoError(t, err)

			// try to send again with amount more than the amount refilled by token pool rate and
			// this should fail, as the refill rate is not enough to refill the capacity
			tokensTobeSent = new(big.Int).Mul(TokenPoolRateLimitRate, big.NewInt(50))
			tc.lane.Logger.Info().Str("tokensTobeSent", tokensTobeSent.String()).Msg("More than TokenPool Rate")
			src.TransferAmount[0] = tokensTobeSent
			failedTx, _, _, err = tc.lane.Source.SendRequest(
				tc.lane.Dest.ReceiverDapp.EthAddress,
				actions.TokenTransfer, "msg with token more than token pool rate",
				common.HexToAddress(tc.lane.Source.Common.FeeToken.Address()))
			require.NoError(t, err)
			require.Error(t, tc.lane.Source.Common.ChainClient.WaitForEvents())
			errReason, v, err = tc.lane.Source.Common.ChainClient.RevertReasonFromTx(failedTx, lock_release_token_pool.LockReleaseTokenPoolABI)
			require.NoError(t, err)
			tc.lane.Logger.Info().
				Str("Revert Reason", errReason).
				Interface("Args", v).
				Str("TokensSent", src.TransferAmount[0].String()).
				Str("Token", tc.lane.Source.Common.BridgeTokens[0].Address()).
				Str("FailedTx", failedTx.Hex()).
				Msg("Msg sent with tokens more than TokenPool Rate")
			require.Equal(t, "TokenRateLimitReached", errReason)

			// validate that the successful transfers are reflected in destination
			tc.lane.ValidateRequests()
		})
	}
}
