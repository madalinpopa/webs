package webs

import (
	"net/http"
	"testing"
)

// TestResponse_Status verifies that the Status method of the Response struct correctly returns the expected status string.
func TestResponse_Status(t *testing.T) {
	resp := Response{status: "200 OK"}
	if resp.Status() != "200 OK" {
		t.Errorf("Status() = %v, want %v", resp.Status(), "200 OK")
	}
}

// TestResponse_StatusCode tests if the StatusCode method of the Response struct returns the correct HTTP status code.
func TestResponse_StatusCode(t *testing.T) {
	resp := Response{statusCode: 200}
	if resp.StatusCode() != 200 {
		t.Errorf("StatusCode() = %v, want %v", resp.StatusCode(), 200)
	}
}

// TestResponse_Headers checks if the Headers() function of the Response struct returns the correct HTTP headers.
func TestResponse_Headers(t *testing.T) {
	headers := http.Header{"Content-Type": []string{"application/json"}}
	resp := &Response{headers: headers}
	if got := resp.Headers(); len(got) != 1 || got.Get("Content-Type") != "application/json" {
		t.Errorf("Headers() = %v, want %v", got, headers)
	}
}

// TestResponse_Bytes tests the Bytes method of the Response type to ensure it correctly returns the body byte slice.
func TestResponse_Bytes(t *testing.T) {
	body := []byte("hello, world")
	resp := &Response{body: body}
	if got := resp.Bytes(); string(got) != "hello, world" {
		t.Errorf("Bytes() = %v, want %v", got, body)
	}
}

// TestResponse_String tests the String() method of Response to ensure it correctly returns the body as a string.
func TestResponse_String(t *testing.T) {
	body := []byte("hello, world")
	resp := &Response{body: body}
	if got := resp.String(); got != "hello, world" {
		t.Errorf("String() = %v, want %v", got, "hello, world")
	}
}

// TestResponse_UnmarshalJson tests the UnmarshalJson method of the Response type to ensure correct JSON unmarshalling.
func TestResponse_UnmarshalJson(t *testing.T) {
	body := []byte(`{"key":"value"}`)
	resp := &Response{body: body}

	var target map[string]string
	if err := resp.UnmarshalJson(&target); err != nil {
		t.Errorf("UnmarshalJson() error = %v", err)
		return
	}

	if target["key"] != "value" {
		t.Errorf("UnmarshalJson() target = %v, want %v", target, map[string]string{"key": "value"})
	}
}
