package main

import (
	"context"

	"github.com/Bancar/goala/ulog"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(Handle)
}

func Handle(ctx context.Context, evt *events.APIGatewayV2HTTPRequest) (*events.APIGatewayV2HTTPResponse, error) {
	log := ulog.With(
		ulog.Str("request.path", evt.RawPath),
		ulog.Str("request.body", evt.Body),
	)

	log.Info("notification received")

	return &events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       `{"message": "ok"}`,
	}, nil
}
