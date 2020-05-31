package lambda

import (
	"encoding/json"
	"memoriesbot/pkg/config"
	"memoriesbot/pkg/status"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
)

// RequestContext - AWS Lambda requestContext object
type RequestContext struct {
	ResourcePath string `json:"resourcePath"`
	HTTPMethod   string `json:"httpMethod"`
}

// Payload - AWS Lambda input for InvokeInput.Payload
type Payload struct {
	RequestContext        RequestContext `json:"requestContext"`
	Headers               interface{}    `json:"headers"`
	QueryStringParameters interface{}    `json:"queryStringParameters"`
	Body                  interface{}    `json:"body"`
	PathParameters        interface{}    `json:"pathParameters"`
}

// InvokeResponse - AWS Lambda outpout
type InvokeResponse struct {
	StatusCode int         `json:"statusCode"`
	Headers    interface{} `json:"headers"`
	Body       string      `json:"body"`
}

// Invoke another lambda function
func Invoke(functionName string, payload *Payload) (status.Response, error) {
	obj, err := json.Marshal(payload)

	if err != nil {
		return status.Response{}, err
	}

	c := config.Get()

	service := lambda.New(session.New(&aws.Config{
		Region: &c.Aws.Region,
	}))
	input := &lambda.InvokeInput{
		FunctionName: aws.String(functionName),
		Payload:      obj,
	}

	result, err := service.Invoke(input)

	if err != nil {
		return status.Response{}, err
	}

	output := InvokeResponse{}
	err = json.Unmarshal(result.Payload, &output)

	if err != nil {
		return status.Response{}, err
	}

	response := status.Response{}
	err = json.Unmarshal([]byte(output.Body), &response)

	if err != nil {
		return status.Response{}, err
	}

	return response, nil
}
