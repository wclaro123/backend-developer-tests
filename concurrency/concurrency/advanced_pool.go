package concurrency

import (
	"context"
	"errors"
	"fmt"
	"sync"
)

// ErrPoolClosed is returned from AdvancedPool.Submit when the pool is closed
// before submission can be sent.
var ErrPoolClosed = errors.New("pool closed")

// AdvancedPool is a more advanced worker pool that supports cancelling the
// submission and closing the pool. All functions are safe to call from multiple
// goroutines.
type AdvancedPool interface {
	// Submit submits the given task to the pool, blocking until a slot becomes
	// available or the context is closed. The given context and its lifetime only
	// affects this function and is not the context passed to the callback. If the
	// context is closed before a slot becomes available, the context error is
	// returned. If the pool is closed before a slot becomes available,
	// ErrPoolClosed is returned. Otherwise the task is submitted to the pool and
	// no error is returned. The context passed to the callback will be closed
	// when the pool is closed.
	Submit(context.Context, func(context.Context)) error

	// Close closes the pool and waits until all submitted tasks have completed
	// before returning. If the pool is already closed, ErrPoolClosed is returned.
	// If the given context is closed before all tasks have finished, the context
	// error is returned. Otherwise, no error is returned.
	Close(context.Context) error
}

// NewAdvancedPool creates a new AdvancedPool. maxSlots is the maximum total
// submitted tasks, running or waiting, that can be submitted before Submit
// blocks waiting for more room. maxConcurrent is the maximum tasks that can be
// running at any one time. An error is returned if maxSlots is less than
// maxConcurrent or if either value is not greater than zero.
func NewAdvancedPool(maxSlots, maxConcurrent int) (AdvancedPool, error) {
	if maxSlots < 1 || maxConcurrent < 1 || maxSlots < maxConcurrent {
		return advancedPool{}, errors.New("could not create a new advanced pool")
	}

	jobsCh := make(chan func(context.Context), maxSlots)
	closeCh := make(chan struct{})
	doneCh := make(chan struct{})
	ctx := context.Background()
	once := sync.Once{}
	var wg sync.WaitGroup

	func() {
		defer close(doneCh)
		wg.Wait()
	}()

	go func() {
		defer ctx.Done()
		<-closeCh
	}()

	for i := 0; i < maxConcurrent; i++ {
		go workerAdvanced(ctx, jobsCh, &wg)
	}

	return &advancedPool{jobsCh: jobsCh, closeCh: closeCh, doneCh: doneCh, once: &once}, nil
}

type advancedPool struct {
	jobsCh  chan func(context.Context)
	closeCh chan struct{}
	doneCh  chan struct{}
	once    *sync.Once
}

func (a advancedPool) Submit(ctx context.Context, f func(context.Context)) error {
	select {
	case <-a.closeCh:
		return ErrPoolClosed

	case a.jobsCh <- f:
		return nil

	case <-ctx.Done():
		return ctx.Err()
	}
}

func (a advancedPool) Close(ctx context.Context) error {
	select {
	case <-a.closeCh:
		return ErrPoolClosed

	case <-ctx.Done():
		return ctx.Err()

	case <-a.doneCh:
		a.once.Do(func() {
			close(a.closeCh)
		})
		return nil
	}
}

func workerAdvanced(ctx context.Context, jobsCh chan func(context.Context), wg *sync.WaitGroup) {
	for f := range jobsCh {
		wg.Add(1)
		f(ctx)
		wg.Done()
	}
	fmt.Println("closing worker")
}
