package lambda

import (
	"context"
	"encoding/json"
)

type AdapterFunc[I any] func(raw json.RawMessage) (*I, error)

type HandlerFunc[I, O any] func(ctx context.Context, evt *I) (*O, error)

type Handler[I, O any] struct {
	fn          HandlerFunc[I, O]
	middlewares []Middleware[I, O]
}

func NewHandler[I, O any](fn HandlerFunc[I, O], middlewares ...Middleware[I, O]) *Handler[I, O] {
	return &Handler[I, O]{
		middlewares: middlewares,
	}
}

func (h *Handler[I, O]) EventHandler(ctx context.Context, in *I) (*O, error) {
	next := h.fn
	// Apply middlewares in reverse order
	for i := len(h.middlewares) - 1; i >= 0; i-- {
		next = h.middlewares[i](next)
	}
	return next(ctx, in)
}

func Interceptor[I, O any](h *Handler[I, O], adapt AdapterFunc[I]) func(ctx context.Context, raw json.RawMessage) (*O, error) {
	return func(ctx context.Context, raw json.RawMessage) (*O, error) {
		in, err := adapt(raw)
		if err != nil {
			return nil, err
		}
		return h.EventHandler(ctx, in)
	}
}
