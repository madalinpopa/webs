package webs

import (
	"net/http"
	"testing"
)

func Test_addHeaders(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	newHeaders := make(http.Header)
	newHeaders.Add("Authorization", "Bearer 123")

	addHeaders(headers, newHeaders)

	if headers.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type to be application/json, got %s", headers.Get("Content-Type"))
	}
	if headers.Get("Authorization") != "Bearer 123" {
		t.Errorf("expected Authorization to be Bearer 123, got %s", headers.Get("Authorization"))
	}
}

func Test_mergeHeaders(t *testing.T) {
	headers := make(http.Header)
	headers.Add("Content-Type", "application/json")
	newHeaders := make(http.Header)
	newHeaders.Add("Authorization", "Bearer 123")

	results := mergeHeaders(headers, newHeaders)
	if results.Get("Content-Type") != "application/json" {
		t.Errorf("expected Content-Type to be application/json, got %s", results.Get("Content-Type"))
	}
	if results.Get("Authorization") != "Bearer 123" {
		t.Errorf("expected Authorization to be Bearer 123, got %s", results.Get("Authorization"))
	}
}

func Test_getRequestBody(t *testing.T) {

	t.Run("nobBodyNilResponse", func(t *testing.T) {

		body, err := getRequestBody("", nil)

		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		if body != nil {
			t.Errorf("expected body to be nil, got %s", body)
		}
	})

	t.Run("bodyWithJsonResponse", func(t *testing.T) {
		requestBody := []string{"foo", "bar"}

		body, err := getRequestBody("application/json", requestBody)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		expectedBody := `["foo","bar"]`
		if string(body) != expectedBody {
			t.Errorf("invalid jsob body, expected, %s, got %s", expectedBody, body)
		}
	})

	t.Run("bodyWithXmlResponse", func(t *testing.T) {
		requestBody := []string{"foo", "bar"}

		body, err := getRequestBody("application/xml", requestBody)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		expectedBody := `<string>foo</string><string>bar</string>`
		if string(body) != expectedBody {
			t.Errorf("expected %s, got %s", expectedBody, body)
		}
	})

	t.Run("defaultResponse", func(t *testing.T) {
		requestBody := []string{"foo", "bar"}

		body, err := getRequestBody("", requestBody)
		if err != nil {
			t.Errorf("expected no error, got %s", err)
		}
		expectedBody := `["foo","bar"]`
		if string(body) != expectedBody {
			t.Errorf("invalid jsob body, expected, %s, got %s", expectedBody, body)
		}
	})

}
