package instaclustr

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

/*
SetupMock used to mock request/response calls to the provisioning api

Usage example:

client := SetupMock(t, id, fmt.Sprintf(`{"id":"%s"}`, id), 202)

Where:
 t: is the testing library,
 request: is the resource path in this case the id of the cluster,
 response: is the api response in this case a string representing a json object with an id the same as the request
 responseCode: is the api response code in this case 202

*/
func SetupMock(t *testing.T, request string, response string, responseCode int) *APIMockClient {
	requestString := fmt.Sprintf("/provisioning/v1/%s", request)
	client := new(APIMockClient)
	client.InitClient(func(req *http.Request) *http.Response {
		// Test request parameters
		if req.URL.String() != requestString {
			t.Fatalf("Unexpected request, expected '%s', but was '%s'", requestString, req.URL.String())
		}
		return &http.Response{
			StatusCode: responseCode,
			// Send response to be tested
			Body: ioutil.NopCloser(bytes.NewBufferString(response)),
			// Must be set to non-nil value or it panics
			Header: make(http.Header),
		}
	})
	return client
}
