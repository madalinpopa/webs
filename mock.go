package webs

import (
	"fmt"
	"net/http"
)

// Mock represents a mock HTTP request and response, useful for testing and simulating HTTP interactions.
type Mock struct {
	Method      string
	Url         string
	RequestBody string

	Error              error
	ResponseBody       string
	ResponseStatusCode int
}

// GetResponse returns a simulated HTTP response based on the mock's state, or an error if one is set.
func (m *Mock) GetResponse() (*Response, error) {
	if m.Error != nil {
		return nil, m.Error
	}
	response := Response{
		status: fmt.Sprintf("%d %s", m.ResponseStatusCode, http.StatusText(m.ResponseStatusCode)),
	}
	return &response, nil
}

// RoundTripFunc defines a function type that takes an HTTP request and returns an HTTP response.
type RoundTripFunc func(req *http.Request) *Response

// RoundTrip executes the RoundTripFunc with the provided *http.Request and returns the resulting *Response and error.
func (f RoundTripFunc) RoundTrip(req *http.Request) (*Response, error) {
	return f(req), nil
}

type MockBuilder struct{}

func (mb *MockBuilder) Build() *Mock {
	return &Mock{}
}
