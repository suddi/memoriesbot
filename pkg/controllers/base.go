package base

import (
	"log"
	"memoriesbot/pkg/config"
	"memoriesbot/pkg/lambda"
	"memoriesbot/pkg/status"
	"os"

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

var logger = log.New(os.Stderr, "ERROR [memoriesbot/pkg/controllers/base]: ", log.Llongfile)

var whoamiResponse []whoami = []whoami{}

func init() {
	botName := "knightwatcherbot"
	response, err := GetLambdaWhoami(botName)
	if err != nil {
		logger.Fatalf("Could not retrieve app details for %s", botName)
		os.Exit(1)
	}

	c := config.Get()
	whoamiResponse = append(whoamiResponse, whoami{
		Name:    c.App.Name,
		Version: c.App.Version,
	})
	whoamiResponse = append(whoamiResponse, whoami{
		Name:    response["name"].(string),
		Version: response["version"].(string),
	})
}

// GetLambdaWhoami - make request to "knightwatcherbot" lambda get app details
func GetLambdaWhoami(botName string) (map[string]interface{}, error) {
	payload := &lambda.Payload{
		RequestContext: lambda.RequestContext{
			ResourcePath: "/",
			HTTPMethod:   "GET",
		},
	}
	response, err := lambda.Invoke(botName, payload)
	if err != nil {
		return nil, err
	}

	data := response.Data.(map[string]interface{})
	return data, nil
}

// ServeWhoami - serves route GET /
func ServeWhoami(_ events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return status.SendResponse(status.Ok, whoamiResponse)
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
