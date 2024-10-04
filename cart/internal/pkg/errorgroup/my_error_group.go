package errorgroup

import (
	"context"
	"sync"
)

type ErrorGroup struct {
	cancel  func(error)
	wg      sync.WaitGroup
	errOnce sync.Once
	err     error
}

func New(ctx context.Context) (*ErrorGroup, context.Context) {
	ctx, cancel := context.WithCancelCause(ctx)
	return &ErrorGroup{cancel: cancel}, ctx
}

func (g *ErrorGroup) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.errOnce.Do(func() {
				g.err = err
				if g.cancel != nil {
					g.cancel(g.err)
				}
			})
		}
	}()
}

func (g *ErrorGroup) Wait() error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel(g.err)
	}
	return g.err
}
