package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"memoriesbot/pkg/config"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// InfoLogger - to log out info to Cloudwatch Logs
var infoLogger *log.Logger = log.New(os.Stdout, "", 0)

// ErrorLogger - to log out errors to Cloudwatch Logs
var errorLogger *log.Logger = log.New(os.Stdout, "", 0)

type accessLog struct {
	Method   string `json:"method"`
	Path     string `json:"path"`
	RouteKey string `json:"routeKey"`
}

// RoutingFunction - function declaration for routing functions
type RoutingFunction func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// AccessLog - access logger that wraps a routing function
func AccessLog(fn RoutingFunction) RoutingFunction {
	c := config.Get()

	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		response, err := fn(req)

		infoLogger.SetPrefix(fmt.Sprintf("EXECUTING\tRequestId: %s\tVersion: %s\tINFO\t", req.RequestContext.RequestID, c.Aws.LambdaVersion))
		errorLogger.SetPrefix(fmt.Sprintf("EXECUTING\tRequestId: %s\tVersion: %s\tERROR\t", req.RequestContext.RequestID, c.Aws.LambdaVersion))

		routeKey := fmt.Sprintf("%s %s", req.HTTPMethod, req.Resource)
		logJSON, err := json.Marshal(accessLog{
			Method:   req.HTTPMethod,
			Path:     req.Resource,
			RouteKey: routeKey,
		})

		if err != nil {
			errorMessage := fmt.Sprintf("RequestId: %s\tVersion: %s\tERROR\tFailed to marshal JSON for routeKey = %s", req.RequestContext.RequestID, c.Aws.LambdaVersion, routeKey)
			errorLogger.Println(errorMessage)
			return response, err
		}

		message := fmt.Sprintf("RequestId: %s\tVersion: %s\tINFO\t%s", req.RequestContext.RequestID, c.Aws.LambdaVersion, logJSON)
		infoLogger.Println(message)
		return response, err
	}
}

// Log - used to log messages to stdout
func Log(message string) {
	infoLogger.Println(message)
}

// LogError - used to log errors to stderr
func LogError(errorMessage interface{}) {
	errorLogger.Println(errorMessage)
}
