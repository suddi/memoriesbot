package controllers

import (
	"memoriesbot/pkg/auth"
	"memoriesbot/pkg/logger"
	"memoriesbot/pkg/status"

	"github.com/aws/aws-lambda-go/events"
)

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
		return status.SendResponse(status.Ok, nil)
	}

	code := req.QueryStringParameters["code"]
	if code != "" {
		token, err := auth.ExchangeToken(code)
		if err != nil {
			logger.LogError(err)
			return status.SendResponse(status.Ok, nil)
		}

		logger.Log(token.AccessToken)
		return status.SendResponse(status.Ok, nil)
	}
	return status.SendResponse(status.Ok, nil)
}