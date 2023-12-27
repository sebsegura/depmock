package lambda

import (
	"context"
	"encoding/json"
)

type AsyncHandlerFn[I any] func(ctx context.Context, evt *I) error
type SyncHandlerFn[I, O any] func(ctx context.Context, evt *I) (*O, error)

type SyncHandler[I, O any] struct {
	fn SyncHandlerFn[I, O]
}

type AsyncHandler[I any] struct {
	fn AsyncHandlerFn[I]
}

func NewSyncHandler[I, O any](fn SyncHandlerFn[I, O], middlewares ...SyncMiddleware[I, O]) *SyncHandler[I, O] {
	fn = applyMiddlewares[I, O](fn, middlewares).(SyncHandlerFn[I, O])
	return &SyncHandler[I, O]{
		fn: fn,
	}
}

func NewAsyncHandler[I any](fn AsyncHandlerFn[I], middlewares ...AsyncMiddleware[I]) *AsyncHandler[I] {
	fn = applyMiddlewares[I, any](fn, middlewares).(AsyncHandlerFn[I])
	return &AsyncHandler[I]{
		fn: fn,
	}
}

func applyMiddlewares[I, O any](fn any, middlewares any) any {
	switch f := fn.(type) {
	case SyncHandlerFn[I, O]:
		mw := middlewares.([]SyncMiddleware[I, O])
		for i := len(mw) - 1; i >= 0; i-- {
			f = mw[i](f)
		}
		return f
	case AsyncHandlerFn[I]:
		mw := middlewares.([]AsyncMiddleware[I])
		for i := len(mw) - 1; i >= 0; i-- {
			f = mw[i](f)
		}
		return f
	}
	return nil
}

func (h *SyncHandler[I, O]) EventHandler(ctx context.Context, raw json.RawMessage) (*O, error) {
	in, err := AdaptEvent[I](raw)
	if err != nil {
		return nil, err
	}
	return h.fn(ctx, in)
}

func (h *AsyncHandler[I]) EventHandler(ctx context.Context, raw json.RawMessage) error {
	in, err := AdaptEvent[I](raw)
	if err != nil {
		return err
	}
	return h.fn(ctx, in)
}
