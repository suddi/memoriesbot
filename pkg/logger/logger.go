package logger

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// Debug - debug log level
var Debug *log.Logger = log.New(os.Stdout, "", log.Llongfile)

// Info - info log level
var Info *log.Logger = log.New(os.Stdout, "", log.Llongfile)

// Warn - warn log level
var Warn *log.Logger = log.New(os.Stdout, "", log.Llongfile)

// Error - error log level
var Error *log.Logger = log.New(os.Stderr, "", log.Llongfile)

// Fatal - fatal log level
var Fatal *log.Logger = log.New(os.Stderr, "", log.Llongfile)

type accessLog struct {
	Method   string `json:"method"`
	Path     string `json:"path"`
	RouteKey string `json:"routeKey"`
}

// RoutingFunction - function declaration for routing functions
type RoutingFunction func(events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error)

// AccessLog - access logger that wraps a routing function
func AccessLog(fn RoutingFunction) RoutingFunction {
	return func(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		response, err := fn(req)

		Info.Println(&accessLog{
			Method:   req.HTTPMethod,
			Path:     req.Resource,
			RouteKey: fmt.Sprintf("%s %s", req.HTTPMethod, req.Resource),
		})
		return response, err
	}

}
