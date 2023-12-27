package lambda

import "github.com/aws/aws-lambda-go/lambda"

func StartSync[I, O any](svc SyncHandlerFn[I, O]) {
	h := NewSyncHandler[I, O](svc, SyncLoggerMiddleware[I, O])
	lambda.Start(h)
}

func StartAsync[I any](svc AsyncHandlerFn[I]) {
	h := NewAsyncHandler[I](svc, AsyncLoggerMiddleware[I])
	lambda.Start(h)
}
