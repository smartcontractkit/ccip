package lazyinitservice

import (
	"context"
	"errors"
	"testing"

	"github.com/test-go/testify/require"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

type dummyService struct {
	startCallCount int
	completeStart  chan struct{}
	closeCallCount int
}

func newDummyService() *dummyService {
	return &dummyService{
		completeStart: make(chan struct{}, 1),
	}
}

func (s *dummyService) AwaitCompleteStart() {
	<-s.completeStart
}

func (s *dummyService) Start(context.Context) error {
	s.startCallCount++
	s.completeStart <- struct{}{}
	return nil
}

func (s *dummyService) Close() error {
	s.closeCallCount++
	return nil
}

func TestLazyInitService_AsyncInit(t *testing.T) {
	dummy := newDummyService()

	s := New(func(context.Context) (job.ServiceCtx, error) {
		return dummy, nil
	})
	require.NoError(t, s.Start(context.Background()))
	dummy.AwaitCompleteStart()
	require.NoError(t, s.Close())
	require.Equal(t, 1, dummy.startCallCount)
	require.Equal(t, 1, dummy.closeCallCount)
}

func TestLazyInitService_NoStartOnUnrecoverableFailure(t *testing.T) {
	tries := 0
	ch := make(chan struct{})
	s := New(func(context.Context) (job.ServiceCtx, error) {
		tries++
		close(ch)
		return nil, Unrecoverable(errors.New("boom"))
	})
	require.NoError(t, s.Start(context.Background()))
	<-ch
	require.NoError(t, s.Close())
	require.Equal(t, 1, tries)
}

func TestLazyInitService_RetryOnRecoverableFailure(t *testing.T) {
	tries := 0
	var msgs []string
	dummy := newDummyService()
	s := New(func(context.Context) (job.ServiceCtx, error) {
		tries++
		if tries <= 3 {
			return nil, errors.New("boom")
		}
		return dummy, nil
	}, WithLogErrorFunc(func(msg string) { msgs = append(msgs, msg) }))
	require.NoError(t, s.Start(context.Background()))
	dummy.AwaitCompleteStart()
	require.NoError(t, s.Close())
	require.Equal(t, 1, dummy.startCallCount)
	require.Equal(t, 1, dummy.closeCallCount)
	require.Equal(t, 4, tries)
	require.Equal(t, 3, len(msgs))
}

func TestLazyInitService_ParentContextCancel(t *testing.T) {
	s := New(func(context.Context) (job.ServiceCtx, error) {
		return nil, errors.New("boom")
	})
	ctx, cancelFunc := context.WithCancel(context.Background())
	require.NoError(t, s.Start(ctx))
	cancelFunc()

	require.NoError(t, s.Close())
}

func TestLazyInitService_Restart(t *testing.T) {
	initCount := 0
	dummy := newDummyService()
	s := New(func(context.Context) (job.ServiceCtx, error) {
		initCount++
		return dummy, nil
	})
	require.NoError(t, s.Start(context.Background()))
	dummy.AwaitCompleteStart()
	require.NoError(t, s.Close())

	require.Equal(t, 1, initCount)
	require.Equal(t, 1, dummy.startCallCount)
	require.Equal(t, 1, dummy.closeCallCount)

	require.NoError(t, s.Start(context.Background()))
	dummy.AwaitCompleteStart()
	require.NoError(t, s.Close())

	require.Equal(t, 1, initCount)
	require.Equal(t, 2, dummy.startCallCount)
	require.Equal(t, 2, dummy.closeCallCount)
}
