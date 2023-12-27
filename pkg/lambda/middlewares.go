package lambda

import (
	"context"
	"go.uber.org/zap"
	"sebsegura/sample-lambda/pkg/logger"
)

func WithLogger(ctx context.Context) (context.Context, *zap.Logger) {
	log := logger.NewLogger()
	return logger.NewContextWithLogger(ctx, log), log
}

type SyncMiddleware[I, O any] func(next SyncHandlerFn[I, O]) SyncHandlerFn[I, O]

func SyncLoggerMiddleware[I, O any](next SyncHandlerFn[I, O]) SyncHandlerFn[I, O] {
	return func(ctx context.Context, in *I) (*O, error) {
		ctx, log := WithLogger(ctx)
		log.With(zap.Any("event", in)).Info("starting execution...")

		out, err := next(ctx, in)
		if err != nil {
			log.With(zap.Any("error", err)).Error("execution has an error")
		} else {
			log.With(zap.Any("response", out)).Info("success")
		}

		return out, err
	}
}

type AsyncMiddleware[I any] func(next AsyncHandlerFn[I]) AsyncHandlerFn[I]

func AsyncLoggerMiddleware[I any](next AsyncHandlerFn[I]) AsyncHandlerFn[I] {
	return func(ctx context.Context, in *I) error {
		ctx, log := WithLogger(ctx)
		log.With(zap.Any("event", in)).Info("starting execution...")

		err := next(ctx, in)
		if err != nil {
			log.With(zap.Any("error", err)).Error("execution has an error")
		} else {
			log.Info("success")
		}

		return err
	}
}
