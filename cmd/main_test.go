package main

import (
	"encoding/json"
	"fmt"
	"memoriesbot/pkg/status"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func TestRouteRequestsWhoami(t *testing.T) {
	type data struct {
		Name    string `json:"name"`
		Version string `json:"version"`
	}

	type response struct {
		Meta status.Meta `json:"meta"`
		Data data        `json:"data"`
	}

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/",
		Resource:   "/",
	}

	result, _ := routeRequests(req)

	errorMessage := fmt.Sprintf("CASE 1: Should be able to route %s %s", req.HTTPMethod, req.Path)

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

func TestRouteRequestsHealthCheck(t *testing.T) {
	type data struct {
		Status string `json:"status"`
	}

	type response struct {
		Meta status.Meta `json:"meta"`
		Data data        `json:"data"`
	}

	req := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/v1/status",
		Resource:   "/v1/status",
	}

	result, _ := routeRequests(req)

	errorMessage := fmt.Sprintf("CASE 1: Should be able to route %s %s", req.HTTPMethod, req.Path)

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

func TestRouteRequestsUnknown(t *testing.T) {
	type data struct{}

	type response struct {
		Meta status.Meta `json:"meta"`
		Data data        `json:"data"`
	}

	whoami := events.APIGatewayProxyRequest{
		HTTPMethod: "GET",
		Path:       "/unknown",
		Resource:   "/unknown",
	}

	result, _ := routeRequests(whoami)

	errorMessage := fmt.Sprintf("CASE 1: Should be able to route unknown endpoint %s %s", whoami.HTTPMethod, whoami.Path)

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
