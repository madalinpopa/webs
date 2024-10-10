package webs

import (
	"net/http"
)

type RequestHandler interface {
	ExecuteRequest(method, url string, headers http.Header, body interface{}) (*Response, error)
}

type Requester interface {
	Do(req *http.Request) (*http.Response, error)
	Get(url string, headers http.Header) (*Response, error)
	Post(url string, headers http.Header, body interface{}) (*Response, error)
	Put(url string, headers http.Header, body interface{}) (*Response, error)
	Patch(url string, headers http.Header, body interface{}) (*Response, error)
	Delete(url string, headers http.Header) (*Response, error)
}
