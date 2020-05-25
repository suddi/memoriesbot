package main

import (
	"fmt"

	base "memoriesbot/pkg/controllers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func routeRequests(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	route := fmt.Sprintf("%s %s", req.HTTPMethod, req.Resource)

	switch route {
	case "GET /":
		return base.ServeWhoami(req)
	case "GET /v1/status":
		return base.ServeHealthCheck(req)
	// case "GET /v1/memories":
	// 	return serveMemories()
	default:
		return base.HandleUnknownEndpoint(req)
	}
}

func main() {
	lambda.Start(routeRequests)
}
