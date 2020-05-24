package status

import (
	"encoding/json"
	"fmt"
	"testing"
)

type exampleBodyStruct struct {
	Status string `json:"status"`
}

type exampleResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type emptyStruct struct{}

var exampleBody exampleBodyStruct = exampleBodyStruct{
	Status: "HEALTHY",
}

func getExpectedResult(response exampleResponse, t *testing.T) string {
	expectedResult, err := json.Marshal(response)
	if err != nil {
		t.Fatal("Failed to marshal JSON from input")
	}

	return string(expectedResult)
}

func TestSendResponse(t *testing.T) {
	type testCase struct {
		description        string
		inputStatus        Meta
		inputBody          interface{}
		expectedStatusCode int
		expectedBody       string
	}
	testCases := []testCase{
		testCase{
			description:        "Should be able to send 200 - OK response",
			inputStatus:        Ok,
			inputBody:          exampleBody,
			expectedStatusCode: Ok.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: Ok,
				Data: exampleBody,
			}, t),
		},
		testCase{
			description:        "Should be able to send 201 - CREATED response",
			inputStatus:        Created,
			inputBody:          exampleBody,
			expectedStatusCode: Created.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: Created,
				Data: exampleBody,
			}, t),
		},
		testCase{
			description:        "Should be able to send 501 - NOT IMPLEMTNTED response",
			inputStatus:        NotImplemented,
			inputBody:          exampleBody,
			expectedStatusCode: NotImplemented.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: NotImplemented,
				Data: exampleBody,
			}, t),
		},
		testCase{
			description:        "Should not be able to pass invalid JSON",
			inputStatus:        TooManyRequests,
			inputBody:          make(chan int),
			expectedStatusCode: InternalError.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: InternalError,
				Data: emptyStruct{},
			}, t),
		},
	}

	for i, test := range testCases {
		errorMessage := fmt.Sprintf("CASE %d: %s", i, test.description)
		result, err := SendResponse(test.inputStatus, test.inputBody)

		if err != nil {
			t.Error(errorMessage)
			return
		}

		if result.StatusCode != test.expectedStatusCode || result.Body != test.expectedBody {
			t.Logf("result.StatusCode: %+v", result.StatusCode)
			t.Logf("test.expectedStatusCode: %+v", test.expectedStatusCode)
			t.Logf("result.Body: %+v", result.Body)
			t.Logf("test.expectedBody: %+v", test.expectedBody)
			t.Error(errorMessage)
		}
	}
}

func TestSendError(t *testing.T) {
	type testCase struct {
		description        string
		inputStatus        Meta
		inputBody          interface{}
		expectedStatusCode int
		expectedBody       string
	}
	testCases := []testCase{
		testCase{
			description:        "Should be able to send 400 - BAD REQUEST response",
			inputStatus:        BadRequest,
			inputBody:          exampleBody,
			expectedStatusCode: BadRequest.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: BadRequest,
				Data: exampleBody,
			}, t),
		},
		testCase{
			description:        "Should not be able to send 201 - CREATED response",
			inputStatus:        Created,
			inputBody:          exampleBody,
			expectedStatusCode: InternalError.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: InternalError,
				Data: emptyStruct{},
			}, t),
		},
		testCase{
			description:        "Should not be able to pass invalid JSON",
			inputStatus:        TooManyRequests,
			inputBody:          make(chan int),
			expectedStatusCode: InternalError.Code,
			expectedBody: getExpectedResult(exampleResponse{
				Meta: InternalError,
				Data: emptyStruct{},
			}, t),
		},
	}

	for i, test := range testCases {
		errorMessage := fmt.Sprintf("CASE %d: %s", i, test.description)
		result, err := SendError(test.inputStatus, test.inputBody)

		if err != nil {
			t.Error(errorMessage)
			return
		}

		if result.StatusCode != test.expectedStatusCode || result.Body != test.expectedBody {
			t.Logf("result.StatusCode: %+v", result.StatusCode)
			t.Logf("test.expectedStatusCode: %+v", test.expectedStatusCode)
			t.Logf("result.Body: %+v", result.Body)
			t.Logf("test.expectedBody: %+v", test.expectedBody)
			t.Error(errorMessage)
		}
	}
}
