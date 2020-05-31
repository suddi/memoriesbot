package main

import (
	"fmt"

	"memoriesbot/pkg/controllers"
	"memoriesbot/pkg/logger"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func routeRequests(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	route := fmt.Sprintf("%s %s", req.HTTPMethod, req.Resource)

	switch route {
	case "GET /":
		return logger.AccessLog(controllers.ServeWhoami)(req)
	case "GET /v1/status":
		return logger.AccessLog(controllers.ServeHealthCheck)(req)
	case "GET /v1/auth/google":
		return logger.AccessLog(controllers.ServeRequestToken)(req)
	case "GET /v1/auth/google/code":
		return logger.AccessLog(controllers.ServeExchangeAuthCode)(req)
	// case "GET /v1/memories":
	// 	return log.AccessLog(serveMemories)(req)
	default:
		return logger.AccessLog(controllers.HandleUnknownEndpoint)(req)
	}
}

func main() {
	lambda.Start(routeRequests)
}
