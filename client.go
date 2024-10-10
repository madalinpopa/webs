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
	defaultResponseTimeout     = 0
	defaultConnectionTimeout   = 0
	defaultMaxIdleConnsPerHost = 1
)

type Client struct {
	builder    *ClientBuilder
	client     *http.Client
	clientOnce sync.Once
}

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

func (c *Client) Get(url string, headers http.Header) (*Response, error) {
	return c.ExecuteRequest(http.MethodGet, url, headers, nil)
}

func (c *Client) Post(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.ExecuteRequest(http.MethodPost, url, headers, body)
}

func (c *Client) Put(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.ExecuteRequest(http.MethodPut, url, headers, body)
}

func (c *Client) Patch(url string, headers http.Header, body interface{}) (*Response, error) {
	return c.ExecuteRequest(http.MethodPatch, url, headers, body)
}

func (c *Client) Delete(url string, headers http.Header) (*Response, error) {
	return c.ExecuteRequest(http.MethodDelete, url, headers, nil)
}

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

func (c *Client) getResponseTimeout() time.Duration {
	if c.builder.responseTimeout > 0 {
		return c.builder.responseTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}
func (c *Client) getConnectionTimeout() time.Duration {
	if c.builder.connectTimeout > 0 {
		return c.builder.connectTimeout
	}
	if c.builder.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

func (c *Client) getMaxIdleConnsPerHost() int {
	if c.builder.maxIdleConnsPerHost > 0 {
		return c.builder.maxIdleConnsPerHost
	}
	return defaultMaxIdleConnsPerHost
}
