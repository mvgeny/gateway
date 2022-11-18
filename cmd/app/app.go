package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type listener interface {
	Start(ctx context.Context) error
	Compensate()
}

func Run(ctx context.Context, listeners ...listener) error {
	//run all listeners with graceful shutdown
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	g, gctx := errgroup.WithContext(ctx)
	for i := range listeners {
		id := i
		g.Go(func() error {
			doneErr := make(chan error, 1)
			go func() {
				doneErr <- listeners[id].Start(gctx)
			}()
			select {
			case <-gctx.Done():
				listeners[id].Compensate()
				return gctx.Err()
			case result := <-doneErr:
				cancel()
				return result
			}
		})
	}
	return g.Wait()
}
