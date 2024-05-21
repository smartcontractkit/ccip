package evm

import (
	commontypes "github.com/smartcontractkit/chainlink-common/pkg/types"
	"github.com/smartcontractkit/chainlink/v2/core/internal/testutils"
	"github.com/smartcontractkit/chainlink/v2/core/logger"
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_CCIPRelayerSet(t *testing.T) {
	ctx := testutils.Context(t)

	relayerMap := make(map[commontypes.RelayID]Relayer, 2)
	sourceRelayer := Relayer{}
	destRelayer := Relayer{}
	relayerMap[commontypes.RelayID{Network: "N1", ChainID: "C1"}] = sourceRelayer
	relayerMap[commontypes.RelayID{Network: "N2", ChainID: "C2"}] = destRelayer

	rs := NewCCIPRelayerSet(relayerMap, sourceRelayer, destRelayer, logger.TestLogger(t))

	t.Run("Get", func(t *testing.T) {
		r1, err := rs.Get(ctx, commontypes.RelayID{Network: "N1", ChainID: "C1"})
		require.NoError(t, err)
		require.Equal(t, sourceRelayer, r1)

		r2, err := rs.Get(ctx, commontypes.RelayID{Network: "N2", ChainID: "C2"})
		require.NoError(t, err)
		require.Equal(t, destRelayer, r2)
	})

	t.Run("List", func(t *testing.T) {
		relays, err := rs.List(ctx, []commontypes.RelayID{}...)
		require.NoError(t, err)
		require.Equal(t, 2, len(relays))

		relays1, err1 := rs.List(ctx, commontypes.RelayID{Network: "N1", ChainID: "C1"})
		require.NoError(t, err1)
		require.Equal(t, 1, len(relays1))
		require.Equal(t, sourceRelayer, relays1)

		relays2, err2 := rs.List(ctx, commontypes.RelayID{Network: "N2", ChainID: "C2"})
		require.NoError(t, err2)
		require.Equal(t, 1, len(relays2))
		require.Equal(t, destRelayer, relays2)
	})

}
