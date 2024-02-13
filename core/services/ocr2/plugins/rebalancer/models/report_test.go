package models_test

import (
	"encoding/json"
	"math/big"
	"testing"
	"time"

	"github.com/test-go/testify/require"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
)

func TestMarshalReportMetadata(t *testing.T) {
	t.Parallel()

	t.Run("marshal json", func(t *testing.T) {
		rm := models.Report{}
		b, err := json.Marshal(rm)
		require.NoError(t, err, "failed to marshal empty ReportMetadata")

		var unmarshalled models.Report
		err = json.Unmarshal(b, &unmarshalled)
		require.NoError(t, err, "failed to unmarshal empty ReportMetadata")
		require.Equal(t, rm, unmarshalled, "marshalled and unmarshalled ReportMetadata should be equal")

		rm = models.Report{
			Transfers: []models.Transfer{
				models.NewTransfer(1, 2, big.NewInt(3), time.Now().UTC(), []byte{}),
			},
			LiquidityManagerAddress: models.Address(testutils.NewAddress()),
			NetworkID:               1,
			ConfigDigest: models.ConfigDigest{
				ConfigDigest: testutils.Random32Byte(),
			},
		}
		b, err = json.Marshal(rm)
		require.NoError(t, err, "failed to marshal ReportMetadata")

		err = json.Unmarshal(b, &unmarshalled)
		require.NoError(t, err, "failed to unmarshal ReportMetadata")
		require.Equal(t, rm, unmarshalled, "marshalled and unmarshalled ReportMetadata should be equal")
	})

	t.Run("marshal onchain", func(t *testing.T) {
		bridgeData1 := testutils.Random32Byte()
		bridgeData2 := testutils.Random32Byte()
		rm := models.Report{
			NetworkID:               1,
			LiquidityManagerAddress: models.Address(testutils.NewAddress()),
			Transfers: []models.Transfer{
				// send instruction
				{
					From:               1,
					To:                 2,
					Amount:             ubig.NewI(3),
					Sender:             models.Address(testutils.NewAddress()),
					Receiver:           models.Address(testutils.NewAddress()),
					LocalTokenAddress:  models.Address(testutils.NewAddress()),
					RemoteTokenAddress: models.Address(testutils.NewAddress()),
					Date:               time.Now().UTC(),
					BridgeData:         bridgeData1[:],
					NativeBridgeFee:    ubig.NewI(4),
				},
				// receive instruction
				{
					From:               3,
					To:                 1,
					Amount:             ubig.NewI(5),
					Sender:             models.Address(testutils.NewAddress()),
					Receiver:           models.Address(testutils.NewAddress()),
					LocalTokenAddress:  models.Address(testutils.NewAddress()),
					RemoteTokenAddress: models.Address(testutils.NewAddress()),
					Date:               time.Now().UTC(),
					BridgeData:         bridgeData2[:],
					NativeBridgeFee:    ubig.NewI(6),
				},
			},
		}
		instructions, err := rm.ToLiquidityInstructions()
		require.NoError(t, err, "failed to convert ReportMetadata to LiquidityInstructions")

		encoded, err := rm.OnchainEncode()
		require.NoError(t, err, "failed to encode ReportMetadata")

		r, decodedInstructions, err := models.DecodeReport(rm.NetworkID, rm.LiquidityManagerAddress, encoded)
		require.NoError(t, err, "failed to unmarshal ReportMetadata")
		require.Equal(t, instructions, decodedInstructions, "marshalled and unmarshalled instructions should be equal")
		require.Equal(t, rm.NetworkID, r.NetworkID, "marshalled and unmarshalled NetworkID should be equal")
		require.Equal(t, rm.LiquidityManagerAddress, r.LiquidityManagerAddress, "marshalled and unmarshalled LiquidityManagerAddress should be equal")
		require.Equal(t, rm.Transfers[0].Amount, r.Transfers[0].Amount, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[0].From, r.Transfers[0].From, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[0].To, r.Transfers[0].To, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[0].BridgeData, r.Transfers[0].BridgeData, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[1].Amount, r.Transfers[1].Amount, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[1].From, r.Transfers[1].From, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[1].To, r.Transfers[1].To, "marshalled and unmarshalled Transfers should be equal")
		require.Equal(t, rm.Transfers[1].BridgeData, r.Transfers[1].BridgeData, "marshalled and unmarshalled Transfers should be equal")
	})
}
