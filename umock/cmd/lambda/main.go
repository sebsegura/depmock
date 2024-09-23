package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"log"
	"sebsegura/umock/internal/api"
	"sebsegura/umock/internal/config"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	config.LoadConfig()

	router, err := api.SetupRouter(context.Background(), "spec.yaml")
	if err != nil {
		log.Fatalf("cannot setup router: %v", err)
	}

	ginLambda := ginadapter.New(router)

	return ginLambda.ProxyWithContext(ctx, req)
}
