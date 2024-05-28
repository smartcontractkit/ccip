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
			token: [8]byte{},
		},
		{
			name:  "eth",
			sym:   "ETH",
			token: [8]byte{'E', 'T', 'H', 0, 0, 0, 0, 0},
		},
		{
			name:  "usdc",
			sym:   "USDC",
			token: [8]byte{'U', 'S', 'D', 'C', 0, 0, 0, 0},
		},
		{
			name:  "usdt",
			sym:   "USDT",
			token: [8]byte{'U', 'S', 'D', 'T', 0, 0, 0, 0},
		},
		{
			name:  "btc",
			sym:   "BTC",
			token: [8]byte{'B', 'T', 'C', 0, 0, 0, 0, 0},
		},
		{
			name:  "long symbol",
			sym:   "BTCCCXXXY",
			token: [8]byte{'B', 'T', 'C', 'C', 'C', 'X', 'X', 'X'},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tid := models.NewTokenID(tc.sym)
			require.Equal(t, tc.token, tid)
			s := tc.sym
			if len(s) > 8 {
				s = s[:8]
			}
			require.Equal(t, s, tid.String())
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
