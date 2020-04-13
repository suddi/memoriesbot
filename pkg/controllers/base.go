package base

import (
	"memoriesbot/pkg/status"

	"github.com/aws/aws-lambda-go/events"
)

type whoami struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type healthCheck struct {
	Status string `json:"status"`
}

type empty struct{}

// ServeWhoami - serves route GET /
func ServeWhoami(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := whoami{
		Name:    "memoriesbot",
		Version: "1.0.0",
	}
	return status.SendResponse(status.Ok, response)
}

// ServeHealthCheck - serves route GET /v1/status
func ServeHealthCheck(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	response := healthCheck{
		Status: "HEALTHY",
	}
	return status.SendResponse(status.Ok, response)
}

// HandleUnknownEndpoint - serves unknown routes
func HandleUnknownEndpoint(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return status.SendResponse(status.Unauthorized, empty{})
}
