package ccipdata

import (
	"encoding/hex"
	"strconv"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/client"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/logpoller"
	"github.com/smartcontractkit/chainlink/v2/core/chains/evm/utils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils/pgtest"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestParse(t *testing.T) {
	expectedBody, err := hexutil.Decode("0x000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000f80000000000000001000000020000000000048d71000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d000000000000000000000000023a04d5935ed8bc8e3eb78db3541f0abfb001c6e0000000000000000000000006cb3ed9b441eb674b58495c8b3324b59faff5243000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc65000000000000000000000000ab4f961939bfe6a93567cc57c59eed7084ce2131000000000000000000000000000000000000000000000000000000000000271000000000000000000000000035e08285cfed1ef159236728f843286c55fc08610000000000000000")
	require.NoError(t, err)

	parsedBody, err := parseUSDCMessageSent(expectedBody)
	require.NoError(t, err)

	expectedPostParse := "0x0000000000000001000000020000000000048d71000000000000000000000000eb08f243e5d3fcff26a9e38ae5520a669f4019d000000000000000000000000023a04d5935ed8bc8e3eb78db3541f0abfb001c6e0000000000000000000000006cb3ed9b441eb674b58495c8b3324b59faff5243000000000000000000000000000000005425890298aed601595a70ab815c96711a31bc65000000000000000000000000ab4f961939bfe6a93567cc57c59eed7084ce2131000000000000000000000000000000000000000000000000000000000000271000000000000000000000000035e08285cfed1ef159236728f843286c55fc0861"

	require.Equal(t, expectedPostParse, hexutil.Encode(parsedBody))
}

func TestFilters(t *testing.T) {
	t.Run("filters of different jobs should be distinct", func(t *testing.T) {
		lggr := logger.TestLogger(t)
		chainID := testutils.NewRandomEVMChainID()
		db := pgtest.NewSqlxDB(t)
		o := logpoller.NewORM(chainID, db, lggr)
		ec := backends.NewSimulatedBackend(map[common.Address]core.GenesisAccount{}, 10e6)
		esc := client.NewSimulatedBackendClient(t, ec, chainID)
		lpOpts := logpoller.Opts{
			PollPeriod:               1 * time.Hour,
			FinalityDepth:            1,
			BackfillBatchSize:        1,
			RpcBatchSize:             1,
			KeepFinalizedBlocksDepth: 100,
		}
		lp := logpoller.NewLogPoller(o, esc, lggr, lpOpts)

		jobID1 := "job-1"
		jobID2 := "job-2"
		transmitter := utils.RandomAddress()

		f1 := logpoller.FilterName("USDC message sent", jobID1, transmitter.Hex())
		f2 := logpoller.FilterName("USDC message sent", jobID2, transmitter.Hex())

		_, err := NewUSDCReader(lggr, jobID1, transmitter, lp, true)
		assert.NoError(t, err)
		assert.True(t, lp.HasFilter(f1))

		_, err = NewUSDCReader(lggr, jobID2, transmitter, lp, true)
		assert.NoError(t, err)
		assert.True(t, lp.HasFilter(f2))

		err = CloseUSDCReader(lggr, jobID2, transmitter, lp)
		assert.NoError(t, err)
		assert.True(t, lp.HasFilter(f1))
		assert.False(t, lp.HasFilter(f2))
	})
}

func TestEncoding(t *testing.T) {
	t.Run("encode expectedSlot", func(t *testing.T) {
		sourceDomainHex := "e0f516f1"
		destDomainHex := "f516f1ee"
		nonceHex := "f56f101748232938"
		version := "00000000"
		sender := "000000000000000000000000"

		sourceDomain, err := strconv.ParseUint(sourceDomainHex, 16, 32)
		require.NoError(t, err)
		destDomain, err := strconv.ParseUint(destDomainHex, 16, 32)
		require.NoError(t, err)
		nonce, err := strconv.ParseUint(nonceHex, 16, 64)
		require.NoError(t, err)

		expected := version + sourceDomainHex + destDomainHex + nonceHex + sender

		actualSlot := GetExpectedNonceSlotData(uint32(sourceDomain), uint32(destDomain), nonce)

		assert.Equal(t, expected, hex.EncodeToString(actualSlot[:]))
	})
}
