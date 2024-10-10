package webs

import (
	"net/http"
	"testing"
)

func TestClient_requestHeaders(t *testing.T) {
	client := Client{}
	commonHeaders := make(http.Header)
	commonHeaders.Set("Content-Type", "application/json")
	commonHeaders.Set("User-Agent", "go-client-client")

	client.builder.SetHeaders(commonHeaders)

	requestHeaders := make(http.Header)
	requestHeaders.Set("X-Request-ID", "abc123")

	finalHeaders := mergeHeaders(commonHeaders, requestHeaders)
	if len(finalHeaders) != 3 {
		t.Errorf("expected 3 headers, got %d", len(finalHeaders))
	}
	if finalHeaders.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type to be application/json, got %s", finalHeaders.Get("Content-Type"))
	}
	if finalHeaders.Get("User-Agent") != "go-client-client" {
		t.Errorf("expected User-Agent to be go-client-client, got %s", finalHeaders.Get("User-Agent"))
	}

	if finalHeaders.Get("X-Request-ID") != "abc123" {
		t.Errorf("expected X-Request-ID to be abc123, got %s", finalHeaders.Get("X-Request-ID"))
	}
}
