package webs

import (
	"encoding/json"
	"net/http"
)

// Response represents an HTTP response, encapsulating status, statusCode, headers, and body.
type Response struct {
	status     string
	statusCode int
	headers    http.Header
	body       []byte
}

// Status returns the HTTP status string of the response.
func (r *Response) Status() string {
	return r.status
}

// StatusCode returns the HTTP status code of the response.
func (r *Response) StatusCode() int {
	return r.statusCode
}

// Headers returns the HTTP headers of the response.
func (r *Response) Headers() http.Header {
	return r.headers
}

// Bytes returns the body of the HTTP response as a slice of bytes.
func (r *Response) Bytes() []byte {
	return r.body
}

// String returns the body of the HTTP response as a string.
func (r *Response) String() string {
	return string(r.body)
}

// UnmarshalJson parses the JSON-encoded body of the response into the target interface.
func (r *Response) UnmarshalJson(target interface{}) error {
	return json.Unmarshal(r.body, target)
}
