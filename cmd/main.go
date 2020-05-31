package main

import (
	"fmt"

	base "memoriesbot/pkg/controllers"
	"memoriesbot/pkg/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func routeRequests(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	route := fmt.Sprintf("%s %s", req.HTTPMethod, req.Resource)

	switch route {
	case "GET /":
		return logger.AccessLog(base.ServeWhoami)(req)
	case "GET /v1/status":
		return logger.AccessLog(base.ServeHealthCheck)(req)
	// case "GET /v1/memories":
	// 	return logger.AccessLog(serveMemories)(req)
	default:
		return logger.AccessLog(base.HandleUnknownEndpoint)(req)
	}
}

func main() {
	lambda.Start(routeRequests)
}
