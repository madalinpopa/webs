package webs

import (
	"bytes"
	"errors"
	"io"
	"net"
	"net/http"
	"sync"
	"time"
)

const (

	// defaultResponseTimeout is the default duration to wait for a response before timing out in HTTP client requests.
	defaultResponseTimeout = 0

	// defaultConnectionTimeout specifies the default timeout duration for establishing a connection to an HTTP server.
	defaultConnectionTimeout = 0

	// defaultMaxIdleConnsPerHost specifies the default maximum number of idle connections to keep per host in the HTTP client.
	defaultMaxIdleConnsPerHost = 1
)

// Client represents a customizable HTTP client built with the help of ClientBuilder.
type Client struct {
	builder    *ClientBuilder
	client     *http.Client
	clientOnce sync.Once
}

// ExecuteRequest sends an HTTP request with the specified method, URL, headers, and body, then returns the response.
func (c *Client) ExecuteRequest(method, url string, headers http.Header, body interface{}) (*Response, error) {
	allHeaders := mergeHeaders(c.builder.headers, headers)

	requestBody, err := getRequestBody(allHeaders.Get("Content-Type"), body)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("failed to create request")
	}

	request.Header = allHeaders

	client := c.getHttpClient()
	response, err := client.Do(request)
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

// getHttpClient returns a singleton instance of http.Client configured with custom timeouts and transport settings.
func (c *Client) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	c.clientOnce.Do(func() {
		c.client = &http.Client{
			Timeout: c.getConnectionTimeout(),
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   c.getMaxIdleConnsPerHost(),
				ResponseHeaderTimeout: c.getResponseTimeout(),
				DialContext: (&net.Dialer{
					Timeout: c.getConnectionTimeout(),
				}).DialContext,
			},
		}
	})

	return c.client
}

// getResponseTimeout calculates and returns the appropriate response timeout duration for the HTTP client.
func (c *Client) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

// getConnectionTimeout returns the configured connection timeout duration or the default value if not set.
func (c *Client) getConnectionTimeout() time.Duration {
	if c.builder.connectTimeout > 0 {
		return c.builder.connectTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

// getMaxIdleConnsPerHost returns the maximum number of idle connections per host. If not set, defaults to 1.
func (c *Client) getMaxIdleConnsPerHost() int {
	if c.builder.maxIdleConnsPerHost > 0 {
		return c.builder.maxIdleConnsPerHost
	}
	return defaultMaxIdleConnsPerHost
}
