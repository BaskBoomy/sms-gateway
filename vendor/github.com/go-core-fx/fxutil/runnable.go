package fxutil

import (
	"context"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Runnable interface {
	Run(ctx context.Context) error
}

func RegisterRunnable[T Runnable]() func(logger *zap.Logger, lc fx.Lifecycle, sh fx.Shutdowner, r T) {
	return func(logger *zap.Logger, lc fx.Lifecycle, sh fx.Shutdowner, r T) {
		ctx, cancel := context.WithCancel(context.Background())
		waitCh := make(chan struct{})

		lc.Append(fx.Hook{
			OnStart: func(_ context.Context) error {
				go func() {
					if err := r.Run(ctx); err != nil {
						logger.Error("runnable failed", zap.Error(err))
						_ = sh.Shutdown()
					}
					close(waitCh)
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				cancel()
				select {
				case <-waitCh:
				case <-ctx.Done():
				}
				return nil
			},
		})
	}
}
