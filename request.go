package webs

import (
	"net/http"
)

// RequestHandler represents an interface for handling HTTP requests and responses.
// ExecuteRequest sends an HTTP request with the specified method, URL, headers, and body, returning a Response and an error.
type RequestHandler interface {
	ExecuteRequest(method, url string, headers http.Header, body interface{}) (*Response, error)
}

// Requester is an interface for making HTTP requests including Do, Get, Post, Put, Patch, and Delete methods.
type Requester interface {
	Do(req *http.Request) (*Response, error)
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Put(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}
