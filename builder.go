package webs

import (
	"net"
	"net/http"
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

// Builder defines an interface for constructing RequestHandler objects with configurable HTTP client settings.
type Builder interface {
	Build() *Client
}

// Transporter defines an interface for getting an *http.Transport instance.
type Transporter interface {
	getTransport() *http.Transport
}

// DefaultsRetriever defines methods for retrieving default configuration values like connection timeout, response timeout, and max idle connections.
type DefaultsRetriever interface {
	getConnectionTimeout() time.Duration
	getResponseTimeout() time.Duration
	getMaxIdleConnsPerHost() int
}

// ClientBuilder assists in creating customized HTTP clients by configuring headers, timeouts, and connection limits.
type ClientBuilder struct {
	headers             http.Header
	transport           http.Transport
	connectTimeout      time.Duration
	responseTimeout     time.Duration
	disableTimeouts     bool
	maxIdleConnsPerHost int
}

// NewClientBuilder creates a new instance of ClientBuilder for configuring customized HTTP clients.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{}
}

// Build finalizes the ClientBuilder configuration and returns a newly constructed Client instance.
func (cb *ClientBuilder) Build() *Client {

	transport := cb.getTransport()

	baseClient := &http.Client{
		Transport: transport,
		Timeout:   cb.getConnectionTimeout(),
	}
	client := &Client{
		client:  baseClient,
		headers: cb.headers,
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

// getResponseTimeout calculates and returns the appropriate response timeout duration for the HTTP client.
func (cb *ClientBuilder) getResponseTimeout() time.Duration {
	if cb.responseTimeout > 0 {
		return cb.responseTimeout
	}
	if cb.disableTimeouts {
		return 0
	}
	return defaultResponseTimeout
}

// getConnectionTimeout returns the configured connection timeout duration or the default value if not set.
func (cb *ClientBuilder) getConnectionTimeout() time.Duration {
	if cb.connectTimeout > 0 {
		return cb.connectTimeout
	}
	if cb.disableTimeouts {
		return 0
	}
	return defaultConnectionTimeout
}

// getMaxIdleConnsPerHost returns the maximum number of idle connections per host. If not set, defaults to 1.
func (cb *ClientBuilder) getMaxIdleConnsPerHost() int {
	if cb.maxIdleConnsPerHost > 0 {
		return cb.maxIdleConnsPerHost
	}
	return defaultMaxIdleConnsPerHost
}

// getTransport configures and returns an *http.Transport with custom timeout settings and connection limits.
func (cb *ClientBuilder) getTransport() *http.Transport {
	return &http.Transport{
		MaxIdleConnsPerHost:   cb.getMaxIdleConnsPerHost(),
		ResponseHeaderTimeout: cb.getResponseTimeout(),
		DialContext: (&net.Dialer{
			Timeout: cb.getConnectionTimeout(),
		}).DialContext,
	}
}
