package ccipexec

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	cciptypes "github.com/smartcontractkit/chainlink-common/pkg/types/ccip"

	"github.com/smartcontractkit/chainlink/v2/core/logger"
)

func TestInflightReportsContainer_add(t *testing.T) {
	lggr := logger.TestLogger(t)
	container := newInflightExecReportsContainer(time.Second)

	container.add(lggr, []cciptypes.EVM2EVMMessage{
		{SequenceNumber: 1}, {SequenceNumber: 2}, {SequenceNumber: 3},
	})
	require.Equal(t, 1, len(container.getAll()))
	insertTime := container.reports[0].createdAt

	time.Sleep(100 * time.Millisecond)
	container.add(lggr, []cciptypes.EVM2EVMMessage{
		{SequenceNumber: 1},
	})
	require.Equal(t, 1, len(container.getAll()))
	updateTime := container.reports[0].createdAt

	require.NotEqual(t, insertTime, updateTime)
}

func TestInflightReportsContainer_expire(t *testing.T) {
	lggr := logger.TestLogger(t)
	container := newInflightExecReportsContainer(time.Second)

	container.add(lggr, []cciptypes.EVM2EVMMessage{
		{SequenceNumber: 1}, {SequenceNumber: 2}, {SequenceNumber: 3},
	})
	container.reports[0].createdAt = time.Now().Add(-time.Second * 5)
	require.Equal(t, 1, len(container.getAll()))

	container.expire(lggr)
	require.Equal(t, 0, len(container.getAll()))
}
