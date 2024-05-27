package models_test

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/require"

	ubig "github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils/big"
	"github.com/smartcontractkit/chainlink/v2/core/services/ocr2/plugins/liquiditymanager/models"
)

func TestTokenID(t *testing.T) {
	tests := []struct {
		name  string
		sym   string
		token models.TokenID
	}{
		{
			name:  "empty",
			sym:   "",
			token: models.TokenID(0),
		},
		{
			name:  "eth",
			sym:   "ETH",
			token: models.TokenID(17628063135754642611), // 0xf4a3760644d064b3
		},
		{
			name:  "usdc",
			sym:   "USDC",
			token: models.TokenID(9624736183338713132), // 0x8591ee9090c0c02c
		},
		{
			name:  "usdt",
			sym:   "USDT",
			token: models.TokenID(14596173802194877749), // 0xca9006bd3fb03d35
		},
		{
			name:  "btc",
			sym:   "BTC",
			token: models.TokenID(15746100419329268335), // 0xda8562e7abc01a6f
		},
		{
			name:  "long symbol",
			sym:   "this_is_a_very_long_symbol_with_multiple_characters",
			token: models.TokenID(3024370980551524913),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(t, tc.token, models.NewTokenID(tc.sym))
		})
	}
}

func TestJSONEncoding(t *testing.T) {
	t.Parallel()

	t.Run("marshal proposed transfer to json", func(t *testing.T) {
		proposedTransfer := models.ProposedTransfer{
			From:   models.NetworkSelector(1),
			To:     models.NetworkSelector(2),
			Amount: ubig.NewI(3),
		}
		jsonBytes, err := json.Marshal(proposedTransfer)
		require.NoError(t, err, "failed to marshal proposed transfer to json")
		var unmarshaled models.ProposedTransfer
		err = json.Unmarshal(jsonBytes, &unmarshaled)
		require.NoError(t, err, "failed to unmarshal proposed transfer from json")
		require.Equal(t, proposedTransfer, unmarshaled, "marshalled and unmarshalled proposed transfer should be equal")
	})

	t.Run("marshal transfer to json", func(t *testing.T) {
		transfer := models.Transfer{
			From:               models.NetworkSelector(1),
			To:                 models.NetworkSelector(2),
			Amount:             ubig.NewI(3),
			Sender:             models.Address(common.HexToAddress("0x1")),
			Receiver:           models.Address(common.HexToAddress("0x2")),
			LocalTokenAddress:  models.Address(common.HexToAddress("0x3")),
			RemoteTokenAddress: models.Address(common.HexToAddress("0x4")),
			BridgeData:         hexutil.Bytes{0x1, 0x2, 0x3},
			NativeBridgeFee:    ubig.NewI(4),
		}
		jsonBytes, err := json.Marshal(transfer)
		require.NoError(t, err, "failed to marshal transfer to json")
		var unmarshaled models.Transfer
		err = json.Unmarshal(jsonBytes, &unmarshaled)
		require.NoError(t, err, "failed to unmarshal transfer from json")
		require.Equal(t, transfer, unmarshaled, "marshalled and unmarshalled transfer should be equal")
	})

	t.Run("marshal pending transfer to json", func(t *testing.T) {
		pendingTransfer := models.PendingTransfer{
			Transfer: models.Transfer{
				From:               models.NetworkSelector(1),
				To:                 models.NetworkSelector(2),
				Amount:             ubig.NewI(3),
				Sender:             models.Address(common.HexToAddress("0x1")),
				Receiver:           models.Address(common.HexToAddress("0x2")),
				LocalTokenAddress:  models.Address(common.HexToAddress("0x3")),
				RemoteTokenAddress: models.Address(common.HexToAddress("0x4")),
				BridgeData:         hexutil.Bytes{0x1, 0x2, 0x3},
				NativeBridgeFee:    ubig.NewI(4),
			},
			Status: models.TransferStatusReady,
		}
		jsonBytes, err := json.Marshal(pendingTransfer)
		require.NoError(t, err, "failed to marshal pending transfer to json")
		var unmarshaled models.PendingTransfer
		err = json.Unmarshal(jsonBytes, &unmarshaled)
		require.NoError(t, err, "failed to unmarshal pending transfer from json")
		require.Equal(t, pendingTransfer, unmarshaled, "marshalled and unmarshalled pending transfer should be equal")
	})
}
