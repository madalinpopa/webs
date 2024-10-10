package webs

import (
	"encoding/json"
	"encoding/xml"
	"net/http"
	"strings"
)

const (

	// ContentTypeJSON represents the MIME type for JSON data: "application/json".
	ContentTypeJSON = "application/json"

	// ContentTypeXML represents the MIME type for XML data: "application/xml".
	ContentTypeXML = "application/xml"
)

// mergeHeaders merges two sets of HTTP headers into a new http.Header object.
func mergeHeaders(headers http.Header, newHeaders http.Header) http.Header {
	result := make(http.Header)
	addHeaders(result, headers)
	addHeaders(result, newHeaders)
	return result

}

// addHeaders adds key-value pairs from newHeaders to the provided headers.
// headers: The original headers to which new headers will be added.
// newHeaders: The headers to be added to the original headers.
func addHeaders(headers http.Header, newHeaders http.Header) {
	for key, values := range newHeaders {
		for _, value := range values {
			headers.Add(key, value)
		}
	}
}

// getRequestBody marshals the given body into a byte slice based on the provided content type.
// Supports both JSON and XML content types, defaults to JSON if a content type is unrecognized.
func getRequestBody(contentType string, body interface{}) ([]byte, error) {

	if body == nil {
		return nil, nil
	}

	switch strings.ToLower(contentType) {
	case ContentTypeJSON:
		return json.Marshal(body)
	case ContentTypeXML:
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}
