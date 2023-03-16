package ccip

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/smartcontractkit/chainlink/core/logger"
)

func TestInflightReportsContainer_add(t *testing.T) {
	lggr := logger.TestLogger(t)
	container := newInflightReportsContainer(time.Second)

	err := container.add(lggr, []uint64{1}, [][]byte{
		{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3},
	})
	require.NoError(t, err)
	err = container.add(lggr, []uint64{1}, nil)
	require.Error(t, err)
	require.Equal(t, "report is already in flight", err.Error())
	require.Equal(t, 1, len(container.getAll()))
}

func TestInflightReportsContainer_expire(t *testing.T) {
	lggr := logger.TestLogger(t)
	container := newInflightReportsContainer(time.Second)

	err := container.add(lggr, []uint64{1}, [][]byte{
		{1, 1, 1, 1}, {2, 2, 2, 2}, {3, 3, 3, 3},
	})
	require.NoError(t, err)
	container.reports[0].createdAt = time.Now().Add(-time.Second * 5)
	require.Equal(t, 1, len(container.getAll()))

	container.expire(lggr)
	require.Equal(t, 0, len(container.getAll()))
}
