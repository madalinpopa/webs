package webs

import (
	"net/http"
	"time"
)

type Builder interface {
	Build() RequestHandler
	SetHeaders(headers http.Header) *ClientBuilder
	DisableTimeouts(disable bool) *ClientBuilder
	SetConnectTimeout(timeout time.Duration) *ClientBuilder
	SetResponseTimeout(timeout time.Duration) *ClientBuilder
	SetMaxIdleConnectionsPerHost(maxIdleConnsPerHost int) *ClientBuilder
}

type ClientBuilder struct {
	headers             http.Header
	disableTimeouts     bool
	connectTimeout      time.Duration
	responseTimeout     time.Duration
	maxIdleConnsPerHost int
}

func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}

func (cb *ClientBuilder) Build() *Client {
	client := &Client{
		builder: cb,
	}
	return client
}

func (cb *ClientBuilder) SetHeaders(headers http.Header) *ClientBuilder {
	cb.headers = headers
	return cb
}

func (cb *ClientBuilder) DisableTimeouts(disable bool) *ClientBuilder {
	cb.disableTimeouts = disable
	return cb
}

func (cb *ClientBuilder) SetConnectTimeout(timeout time.Duration) *ClientBuilder {
	cb.connectTimeout = timeout
	return cb
}

func (cb *ClientBuilder) SetResponseTimeout(timeout time.Duration) *ClientBuilder {
	cb.responseTimeout = timeout
	return cb
}

func (cb *ClientBuilder) SetMaxIdleConnectionsPerHost(maxIdleConnsPerHost int) *ClientBuilder {
	cb.maxIdleConnsPerHost = maxIdleConnsPerHost
	return cb
}
