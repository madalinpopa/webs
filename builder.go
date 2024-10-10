package webs

import (
	"net/http"
	"time"
)

// Builder defines an interface for constructing RequestHandler objects with configurable HTTP client settings.
type Builder interface {
	Build() RequestHandler
	SetHeaders(headers http.Header) *ClientBuilder
	DisableTimeouts(disable bool) *ClientBuilder
	SetConnectTimeout(timeout time.Duration) *ClientBuilder
	SetResponseTimeout(timeout time.Duration) *ClientBuilder
	SetMaxIdleConnectionsPerHost(maxIdleConnsPerHost int) *ClientBuilder
}

// ClientBuilder assists in creating customized HTTP clients by configuring headers, timeouts, and connection limits.
type ClientBuilder struct {
	headers             http.Header
	disableTimeouts     bool
	connectTimeout      time.Duration
	responseTimeout     time.Duration
	maxIdleConnsPerHost int
}

// NewClientBuilder creates a new instance of ClientBuilder for configuring customized HTTP clients.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}

// Build finalizes the ClientBuilder configuration and returns a newly constructed Client instance.
func (cb *ClientBuilder) Build() *Client {
	client := &Client{
		builder: cb,
	}
	return client
}

// SetHeaders sets custom HTTP headers to be used in the requests.
func (cb *ClientBuilder) SetHeaders(headers http.Header) *ClientBuilder {
	cb.headers = headers
	return cb
}

// DisableTimeouts configures the ClientBuilder to enable or disable timeout settings.
func (cb *ClientBuilder) DisableTimeouts(disable bool) *ClientBuilder {
	cb.disableTimeouts = disable
	return cb
}

// SetConnectTimeout sets the connection timeout duration for the client.
func (cb *ClientBuilder) SetConnectTimeout(timeout time.Duration) *ClientBuilder {
	cb.connectTimeout = timeout
	return cb
}

// SetResponseTimeout sets the response timeout duration for the client.
func (cb *ClientBuilder) SetResponseTimeout(timeout time.Duration) *ClientBuilder {
	cb.responseTimeout = timeout
	return cb
}

// SetMaxIdleConnectionsPerHost sets the maximum number of idle connections to keep per-host for the HTTP client.
func (cb *ClientBuilder) SetMaxIdleConnectionsPerHost(maxIdleConnsPerHost int) *ClientBuilder {
	cb.maxIdleConnsPerHost = maxIdleConnsPerHost
	return cb
}
