package lambda

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/lambda"
)

func StartSync[I, O any](svc SyncHandlerFn[I, O]) {
	h := NewSyncHandler[I, O](svc, SyncLoggerMiddleware[I, O])
	lambda.Start(func(ctx context.Context, raw json.RawMessage) (*O, error) {
		return h.EventHandler(ctx, raw)
	})
}

func StartAsync[I any](svc AsyncHandlerFn[I]) {
	h := NewAsyncHandler[I](svc, AsyncLoggerMiddleware[I])
	lambda.Start(func(ctx context.Context, raw json.RawMessage) error {
		return h.EventHandler(ctx, raw)
	})
}
