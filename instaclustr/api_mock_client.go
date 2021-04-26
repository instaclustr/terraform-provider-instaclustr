package instaclustr

import (
	"net/http"
	"time"
)

// RoundTripFunc .
type RoundTripFunc func(request *http.Request) *http.Response

// RoundTrip .
func (mockFunction RoundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return mockFunction(request), nil
}

func (c *APIClient) InitMockClient(mockFunction RoundTripFunc) {
	c.InitClient("", "", "")
	var client = &http.Client{
		Timeout:   time.Second * 60,
		Transport: RoundTripFunc(mockFunction),
	}
	c.SetClient(client)
}
