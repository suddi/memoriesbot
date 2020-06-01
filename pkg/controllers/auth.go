package controllers

import (
	"fmt"
	"memoriesbot/pkg/auth"
	"memoriesbot/pkg/logger"
	"memoriesbot/pkg/status"

	"github.com/aws/aws-lambda-go/events"
)

type empty struct{}

var authState string

func init() {
	state, err := auth.GenerateID()
	if err != nil {
		logger.LogError(err)
	}
	authState = state
}

// ServeRequestToken - serves route GET /v1/auth/google
func ServeRequestToken(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	url := auth.RequestToken(authState)
	return status.SendRedirect(url)
}

// ServeExchangeAuthCode - serves route GET /v1/auth/google/code
func ServeExchangeAuthCode(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	err := req.QueryStringParameters["error"]
	if err != "" {
		logger.LogError(err)
		return status.SendResponse(status.Ok, empty{})
	}

	state := req.QueryStringParameters["state"]
	code := req.QueryStringParameters["code"]
	if code != "" {
		if state != authState {
			logger.LogError("state did not match")
			return status.SendResponse(status.Ok, empty{})
		}

		token, err := auth.ExchangeToken(code)
		if err != nil {
			logger.LogError(err)
			return status.SendResponse(status.Ok, empty{})
		}

		logger.Log(fmt.Sprintf("%+v", token))

		return status.SendResponse(status.Ok, empty{})
	}
	return status.SendResponse(status.Ok, empty{})
}
