package base

import (
	"encoding/json"
	"memoriesbot/pkg/status"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestServeWhoami(t *testing.T) {
	type data struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	type response struct {
		Meta status.Meta `json:"meta"`
		Data data        `json:"data"`
	}

	result, _ := ServeWhoami(events.APIGatewayProxyRequest{})

	errorMessage := "CASE 1: Should be able to serve whoami request"

	res := response{}
	err := json.Unmarshal([]byte(result.Body), &res)

	if err != nil {
		t.Error(errorMessage)
	}

	if result.StatusCode != status.Ok.Code ||
		res.Meta.Code != status.Ok.Code ||
		res.Meta.Message != status.Ok.Message ||
		res.Meta.Retryable != status.Ok.Retryable ||
		res.Data.Name == "" ||
		res.Data.Version == "" {
		t.Error(errorMessage)
	}
}

func TestServeHealthCheck(t *testing.T) {
	type data struct {
		Status string `json:"status"`
	}

	type response struct {
		Meta status.Meta `json:"meta"`
		Data data        `json:"data"`
	}

	result, _ := ServeHealthCheck(events.APIGatewayProxyRequest{})

	errorMessage := "CASE 1: Should be able to serve health check request"

	res := response{}
	err := json.Unmarshal([]byte(result.Body), &res)

	if err != nil {
		t.Error(errorMessage)
	}

	if result.StatusCode != status.Ok.Code ||
		res.Meta.Code != status.Ok.Code ||
		res.Meta.Message != status.Ok.Message ||
		res.Meta.Retryable != status.Ok.Retryable ||
		res.Data.Status != "HEALTHY" {
		t.Error(errorMessage)
	}
}

func TestHandleUnknownEndpoint(t *testing.T) {
	type data struct{}

	type response struct {
		Meta status.Meta `json:"meta"`
		Data data        `json:"data"`
	}

	result, _ := HandleUnknownEndpoint(events.APIGatewayProxyRequest{})

	errorMessage := "CASE 1: Should be able to handle unknown endpoint request"

	res := response{}
	err := json.Unmarshal([]byte(result.Body), &res)

	if err != nil {
		t.Error(errorMessage)
	}

	if result.StatusCode != status.Unauthorized.Code ||
		res.Meta.Code != status.Unauthorized.Code ||
		res.Meta.Message != status.Unauthorized.Message ||
		res.Meta.Retryable != status.Unauthorized.Retryable {
		t.Error(errorMessage)
	}
}
