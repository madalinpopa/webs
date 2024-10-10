package webs

import (
	"net/http"
	"testing"
	"time"
)

// TestClientBuilderDefaultOptions validates that the ClientBuilder initializes with the correct default configuration values.
func TestClientBuilderDefaultOptions(t *testing.T) {

	builder := NewClientBuilder()
	if builder == nil {
		t.Error("expected builder to be not nil")
	}

	t.Run("defaultConnectionTimeout", func(t *testing.T) {
		timeout := builder.getConnectionTimeout()
		if timeout != defaultConnectionTimeout {
			t.Errorf("expected timeout to be %d, got %d", defaultConnectionTimeout, timeout)
		}
	})

	t.Run("defaultResponseTimeout", func(t *testing.T) {
		timeout := builder.getResponseTimeout()
		if timeout != defaultResponseTimeout {
			t.Errorf("expected timeout to be %d, got %d", defaultResponseTimeout, timeout)
		}
	})

	t.Run("defaultMaxIdleConnections", func(t *testing.T) {
		conns := builder.getMaxIdleConnsPerHost()
		if conns != defaultMaxIdleConnsPerHost {
			t.Errorf("expected conns to be %d, got %d", defaultMaxIdleConnsPerHost, conns)
		}
	})
}

// TestClientBuilderWithDefaultOptions tests that the ClientBuilder correctly applies default options for connection and response timeouts, as well as the max number of idle connections per host.
func TestClientBuilderWithDefaultOptions(t *testing.T) {

	builder := NewClientBuilder().
		SetConnectTimeout(3 * time.Second).
		SetResponseTimeout(4 * time.Second).
		SetMaxIdleConnectionsPerHost(5)

	if builder == nil {
		t.Error("expected client to be not nil")
		t.FailNow()
	}
	timeout := builder.getConnectionTimeout()
	if timeout != 3*time.Second {
		t.Errorf("expected timeout to be %d, got %d", 3*time.Second, timeout)
	}

	timeout = builder.getResponseTimeout()
	if timeout != 4*time.Second {
		t.Errorf("expected timeout to be %d, got %d", 4*time.Second, timeout)
	}

	conns := builder.getMaxIdleConnsPerHost()
	if conns != 5 {
		t.Errorf("expected conns to be %d, got %d", 5, conns)
	}

}

// TestClientCreationWithCustomHeaders verifies the creation of an HTTP client with custom headers using ClientBuilder.
func TestClientCreationWithCustomHeaders(t *testing.T) {

	headers := make(http.Header)
	headers.Add("X-Custom-Header", "custom-value")

	builder := NewClientBuilder().SetHeaders(headers)
	if builder == nil {
		t.Error("expected client to be not nil")
		t.FailNow()
	}

	if builder.headers == nil {
		t.Error("expected headers to be not nil")
	}

	if builder.headers.Get("X-Custom-Header") != "custom-value" {
		t.Errorf("expected header value to be 'custom-value', got %s", builder.headers.Get("X-Custom-Header"))
	}

}

// TestClientBuilder_DisabledTimeouts verifies the behavior of ClientBuilder when timeouts are disabled.
func TestClientBuilder_DisabledTimeouts(t *testing.T) {
	builder := NewClientBuilder().DisableTimeouts(true)
	if builder == nil {
		t.Error("expected client to be not nil")
		t.FailNow()
	}

	timeout := builder.getConnectionTimeout()
	if timeout != 0 {
		t.Errorf("expected timeout to be 0, got %d", timeout)
	}

	timeout = builder.getResponseTimeout()
	if timeout != 0 {
		t.Errorf("expected timeout to be 0, got %d", timeout)
	}

}
