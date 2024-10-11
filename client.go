package webs

import (
	"bytes"
	"errors"
	"io"
	"net/http"
)

// Client represents a customizable HTTP client built with the help of ClientBuilder.
type Client struct {
	client  *http.Client
	headers http.Header
}

// ExecuteRequest sends an HTTP request with the specified method, URL, headers, and body, then returns the response.
func (c *Client) ExecuteRequest(method, url string, headers http.Header, body interface{}) (*Response, error) {
	allHeaders := mergeHeaders(c.headers, headers)

	requestBody, err := getRequestBody(allHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("failed to create request")
	}

	request.Header = allHeaders

	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)
	customResponse := Response{
		status:     response.Status,
		statusCode: response.StatusCode,
		headers:    response.Header,
		body:       responseBody,
	}
	return &customResponse, nil
}

func (c *Client) Do(req *http.Request) (*Response, error) {
	return c.ExecuteRequest(req.Method, req.URL.String(), req.Header, nil)
}

// Get sends an HTTP GET request to the specified URL with optional headers and returns the response.
func (c *Client) Get(url string, headers http.Header) (*Response, error) {
	return c.ExecuteRequest(http.MethodGet, url, headers, nil)
}

// Post sends an HTTP POST request to the specified URL with provided headers and body, and returns the response.
func (c *Client) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.ExecuteRequest(http.MethodPost, url, headers, body)
}

// Put sends an HTTP PUT request to the specified URL with provided headers and body, and returns the response.
func (c *Client) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.ExecuteRequest(http.MethodPut, url, headers, body)
}

// Patch sends an HTTP PATCH request to the specified URL with provided headers and body, then returns the response.
func (c *Client) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.ExecuteRequest(http.MethodPatch, url, headers, body)
}

// Delete sends an HTTP DELETE request to the specified URL with optional headers and returns the response.
func (c *Client) Delete(url string, headers http.Header) (*Response, error) {
	return c.ExecuteRequest(http.MethodDelete, url, headers, nil)
}
