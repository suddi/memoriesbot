package status

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
)

// Meta - struct used for "meta" properties in response
type Meta struct {
	Code      int      `json:"code"`
	Message   string   `json:"message"`
	Details   []string `json:"details"`
	Retryable bool     `json:"retryable"`
}

// Pagination - struct used for "pagination" properties in response
type Pagination struct {
	NextPage string `json:"nextPage,omitempty"`
}

// Response - response struct with "meta" and "data"
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// ListResponse - complete response struct with "meta", "data" and "pagination"
type ListResponse struct {
	Meta       Meta        `json:"meta"`
	Data       interface{} `json:"data"`
	Pagination *Pagination `json:"pagination,omitempty"`
}

// Ok - use status.Ok for 200 - OK response
var Ok Meta = Meta{
	Code:      200,
	Message:   "OK",
	Details:   make([]string, 0),
	Retryable: true,
}

// Created - use status.Created for 201 - CREATED response
var Created Meta = Meta{
	Code:      201,
	Message:   "Created",
	Details:   make([]string, 0),
	Retryable: true,
}

// NoContent - use status.NoContent for 204 NO CONTENT response
var NoContent Meta = Meta{
	Code:      204,
	Message:   "No Content",
	Details:   make([]string, 0),
	Retryable: true,
}

// MovedPermanently - use status.MovedPermanently for 301 MOVED PERMANENTLY response
var MovedPermanently Meta = Meta{
	Code:      301,
	Message:   "Moved Permanently",
	Details:   make([]string, 0),
	Retryable: true,
}

// BadRequest - use status.BadRequest for 400 BAD REQUEST response
var BadRequest Meta = Meta{
	Code:      400,
	Message:   "Bad Request",
	Details:   make([]string, 0),
	Retryable: false,
}

// Unauthorized - use status.Unauthorized for 401 UNAUTHORIZED response
var Unauthorized Meta = Meta{
	Code:      401,
	Message:   "Unauthorized",
	Details:   make([]string, 0),
	Retryable: false,
}

// Forbidden - use status.Forbidden for 403 FORBIDDEN response
var Forbidden Meta = Meta{
	Code:      403,
	Message:   "Invalid User",
	Details:   make([]string, 0),
	Retryable: false,
}

// NotFound - use status.NotFound for 404 NOT FOUND response
var NotFound Meta = Meta{
	Code:      404,
	Message:   "Not Found",
	Details:   make([]string, 0),
	Retryable: false,
}

// NotAcceptable - use status.NotAcceptable for 406 NOT ACCEPTABLE response
var NotAcceptable Meta = Meta{
	Code:      406,
	Message:   "Not Acceptable",
	Details:   make([]string, 0),
	Retryable: false,
}

// RequestTimeout - use status.RequestTimeout for 408 REQUEST TIMEOUT response
var RequestTimeout Meta = Meta{
	Code:      408,
	Message:   "Request Timeout",
	Details:   make([]string, 0),
	Retryable: true,
}

// Conflict - use status.Conflict for 409 CONFLICT response
var Conflict Meta = Meta{
	Code:      409,
	Message:   "Conflict",
	Details:   make([]string, 0),
	Retryable: false,
}

// TooManyRequests - use status.TooManyRequests for 429 TOO MANY REQUESTS response
var TooManyRequests Meta = Meta{
	Code:      429,
	Message:   "Too Many Requests",
	Details:   make([]string, 0),
	Retryable: true,
}

// InternalError - use status.InternalError for 500 INTERNAL ERROR response
var InternalError Meta = Meta{
	Code:      500,
	Message:   "Internal Error",
	Details:   make([]string, 0),
	Retryable: false,
}

// NotImplemented - use status.NotImplemented for 501 NOT IMPLEMENTED response
var NotImplemented Meta = Meta{
	Code:      501,
	Message:   "Not Implemented",
	Details:   make([]string, 0),
	Retryable: false,
}

// BadGateway - use status.BadGateway for 502 BAD GATEWAT response
var BadGateway Meta = Meta{
	Code:      502,
	Message:   "Bad Gateway",
	Details:   make([]string, 0),
	Retryable: false,
}

var errorLogger = log.New(os.Stderr, "ERROR: ", log.Llongfile)

// SendResponse - used to provide non-error response to user
func SendResponse(status Meta, body interface{}) (events.APIGatewayProxyResponse, error) {
	resp := Response{
		Meta: status,
		Data: body,
	}
	obj, err := json.Marshal(resp)

	if err != nil {
		return sendInternalError()
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status.Code,
		Body:       string(obj),
	}, nil
}

// SendError - used to provide error response to user
func SendError(status Meta, body interface{}) (events.APIGatewayProxyResponse, error) {
	if status.Code >= 200 && status.Code < 300 {
		errorLogger.Println("Cannot sendError with a success status code")
		return sendInternalError()
	}

	resp := Response{
		Meta: status,
		Data: body,
	}

	obj, err := json.Marshal(resp)

	if err != nil {
		return sendInternalError()
	}

	return events.APIGatewayProxyResponse{
		StatusCode: status.Code,
		Body:       string(obj),
	}, nil
}

// SendRedirect - send redirect to specified URL
func SendRedirect(url string) (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: MovedPermanently.Code,
		Headers: map[string]string{
			"Location": url,
		},
	}, nil
}

func sendInternalError() (events.APIGatewayProxyResponse, error) {
	response := `{"meta":{"code":500,"message":"Internal Error","details":[],"retryable":false},"data":{}}`
	return events.APIGatewayProxyResponse{
		StatusCode: InternalError.Code,
		Body:       response,
	}, nil
}
