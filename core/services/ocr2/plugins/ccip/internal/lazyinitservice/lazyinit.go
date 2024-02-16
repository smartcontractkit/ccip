// Package lazyinitservice provides an implementation of the job.ServiceCtx interface, LazyInitService.
//
// This implementation executes the service initialization lazily on the first Start method invocation.
// If the initialization fails, the service keeps trying to initialize the underlying service periodically until the first success.
// The initialization function can indicate that there is no point in retrying using the Unrecoverable error wrapper.
package lazyinitservice

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/avast/retry-go/v4"

	"github.com/smartcontractkit/chainlink/v2/core/services/job"
)

var ErrNoService = errors.New("LazyInitService: the init function did not return a service")

// An InitFunc represents an expensive blocking computation producing a service.
// Init functions must respect the context passed as the argument and quit promptly if the context is canceled.
type InitFunc = func(context.Context) (job.ServiceCtx, error)

// A LogErrorFunc is a callback for reporting background initialization and startup errors.
type LogErrorFunc = func(error)

type Option = func(*LazyInitService)

type LazyInitService struct {
	// initFunc is the function creating the service.
	initFunc InitFunc
	// initComplete guards the initialization process allowing for a graceful shutdown.
	initComplete sync.WaitGroup
	// logErrorFunc is the function for logging errors occurring in background.
	logErrorFunc LogErrorFunc
	// cancelFunc is the function canceling the initialization process.
	cancelFunc context.CancelFunc
	// initializedService contains the service the initFunc returns.
	initializedService job.ServiceCtx
}

// WithLogErrorFunc instructs the service constructor to use the given function for error reporting.
func WithLogErrorFunc(f LogErrorFunc) Option {
	return func(s *LazyInitService) {
		s.logErrorFunc = f
	}
}

// New creates a new service with the given initialization function.
func New(f InitFunc, opts ...Option) *LazyInitService {
	s := &LazyInitService{
		initFunc: f,
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Start initiates the underlying service initialization and starts it.
//
// Start ignores the given ctx cancellation if the service is not initialized yet.
// Use Close to stop the initialization process and the service.
func (s *LazyInitService) Start(ctx context.Context) error {
	s.initComplete.Wait()

	if s.initializedService != nil {
		return s.initializedService.Start(ctx)
	}

	s.initComplete.Add(1)

	// We create a new context because the original context will be cancelled once `Start` returns.
	ctx, s.cancelFunc = context.WithCancel(context.Background())
	go s.initAndRun(ctx)
	return nil
}

// initAndRun implements the lazy initialization logic.
func (s *LazyInitService) initAndRun(ctx context.Context) {
	defer s.initComplete.Done()

	service, err := retry.DoWithData[job.ServiceCtx](
		func() (job.ServiceCtx, error) { return s.initFunc(ctx) },
		retry.Context(ctx),
		retry.OnRetry(func(n uint, err error) {
			s.reportError(fmt.Errorf("initialization attempt %d failed: %w", n, err))
		}),
	)
	if err != nil {
		s.reportError(err)
		return
	}
	if service == nil {
		s.reportError(ErrNoService)
		return
	}
	s.initializedService = service
	if err = s.initializedService.Start(ctx); err != nil {
		s.reportError(fmt.Errorf("service failed to start: %w", err))
	}
}

// reportError records the given error using the service log error function.
func (s *LazyInitService) reportError(err error) {
	if s.logErrorFunc != nil {
		s.logErrorFunc(err)
	}
}

// Close implements graceful service shutdown logic.
func (s *LazyInitService) Close() error {
	// First, cancel the context to break the initialization retry loop.
	if s.cancelFunc != nil {
		s.cancelFunc()
		s.cancelFunc = nil
	}
	// Second, wait for the initialization to complete.
	s.initComplete.Wait()
	// Now, we can close the internal service if it was initialized.
	if s.initializedService != nil {
		return s.initializedService.Close()
	}
	return nil
}

// Unrecoverable wraps the given error into an error that signals to the retry mechanism to stop trying.
func Unrecoverable(err error) error {
	return retry.Unrecoverable(err)
}
