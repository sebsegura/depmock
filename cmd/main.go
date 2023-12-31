package main

import (
	"sebsegura/sample-lambda/internal/service"
	"sebsegura/sample-lambda/pkg/client"
	"sebsegura/sample-lambda/pkg/lambda"
)

func main() {
	lambda.StartAsync(service.New(client.New()).Credit)
}
