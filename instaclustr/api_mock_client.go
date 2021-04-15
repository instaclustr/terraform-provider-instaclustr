package instaclustr

import (
	"net/http"
	"time"
)

type APIMockClient struct {
	APIClient
}

// RoundTripFunc .
type RoundTripFunc func(request *http.Request) *http.Response

// RoundTrip .
func (mockFunction RoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return mockFunction(request), nil
}

func (c *APIMockClient) InitClient(mockFunction RoundTripFunc) {
    c.APIClient.InitClient("", "", "")
	var client = &http.Client{
		Timeout:   time.Second * 60,
		Transport: RoundTripFunc(mockFunction),
    }
    c.APIClient.SetClient(client)
}
