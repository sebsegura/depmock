package lambda

import "github.com/aws/aws-lambda-go/lambda"

type Lambda[I, O any] struct {
	svc         HandlerFunc[I, O]
	middlewares []Middleware[I, O]
	options     []lambda.Option
	adapter     AdapterFunc[I]
}

func NewLambda[I, O any](svc HandlerFunc[I, O]) *Lambda[I, O] {
	return &Lambda[I, O]{}
}

func (l *Lambda[I, O]) Use(middlewares ...Middleware[I, O]) *Lambda[I, O] {
	if len(middlewares) > 0 {
		l.middlewares = append(l.middlewares, middlewares...)
	}
	return l
}

func (l *Lambda[I, O]) Options(opts ...lambda.Option) *Lambda[I, O] {
	if len(opts) > 0 {
		l.options = append(l.options, opts...)
	}
	return l
}

func (l *Lambda[I, O]) Adapt(fn AdapterFunc[I]) *Lambda[I, O] {
	l.adapter = fn
	return l
}
