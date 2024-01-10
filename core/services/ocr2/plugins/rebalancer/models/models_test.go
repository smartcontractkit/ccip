package models_test

import (
	"encoding/json"
	"math/big"
	"testing"

	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/rebalancer/models"
	"github.com/test-go/testify/require"
)

func TestMarshalReportMetadata(t *testing.T) {
	t.Parallel()

	rm := models.ReportMetadata{}
	b, err := json.Marshal(rm)
	require.NoError(t, err, "failed to marshal empty ReportMetadata")

	var unmarshalled models.ReportMetadata
	err = json.Unmarshal(b, &unmarshalled)
	require.NoError(t, err, "failed to unmarshal empty ReportMetadata")
	require.Equal(t, rm, unmarshalled, "marshalled and unmarshalled ReportMetadata should be equal")

	rm = models.ReportMetadata{
		Transfers: []models.Transfer{
			models.NewTransfer(1, 2, big.NewInt(3)),
		},
		LiquidityManagerAddress: models.Address(testutils.NewAddress()),
		NetworkID:               1,
	}
	b, err = json.Marshal(rm)
	require.NoError(t, err, "failed to marshal ReportMetadata")

	err = json.Unmarshal(b, &unmarshalled)
	require.NoError(t, err, "failed to unmarshal ReportMetadata")
	require.Equal(t, rm, unmarshalled, "marshalled and unmarshalled ReportMetadata should be equal")
}
