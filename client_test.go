package webs

import (
	"github.com/h2non/gock"
	"net/http"
	"testing"
)

// TestClient_Build verifies that the ClientBuilder successfully creates a non-nil Client instance when Build is invoked.
func TestClient_Build(t *testing.T) {

	client := NewClientBuilder().Build()
	if client == nil {
		t.Errorf("expected client, got nil")
	}
}

// TestClient_Get validates that the Get method of the custom HTTP client returns an expected response without errors.
func TestClient_Get(t *testing.T) {
	defer gock.Off()

	gock.New("https://server.com").
		Get("/").
		Reply(200).
		JSON(map[string]string{"hello": "world"})

	client := NewClientBuilder().Build()
	gock.InterceptClient(client.client)

	res, err := client.Get("https://server.com", nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}
	if res.statusCode != 200 {
		t.Errorf("expected status code 200, got %d", res.statusCode)
	}

}

// TestClient_Post is a test function that verifies the behavior of the Client's Post method.
// It uses the gock library to mock HTTP responses and ensures that the client correctly sends a POST request,
// processes the response, and unmarshals JSON content.
func TestClient_Post(t *testing.T) {
	defer gock.Off()

	gock.New("https://server.com").
		Post("/").
		MatchType("json").
		JSON(map[string]string{"foo": "bar"}).
		Reply(201).
		JSON(map[string]string{"foo": "bar"})

	client := NewClientBuilder().Build()
	gock.InterceptClient(client.client)

	body := map[string]string{"foo": "bar"}

	header := http.Header{}
	header.Add("Content-Type", "application/json")

	res, err := client.Post("https://server.com", header, body)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}
	if res.statusCode != 201 {
		t.Errorf("expected status code 201, got %d", res.statusCode)
	}

	var content = make(map[string]string)
	err = res.UnmarshalJson(&content)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if content["foo"] != "bar" {
		t.Errorf("expected foo to be bar, got %s", content["foo"])
	}
}

// TestClient_Put tests the Put method of the Client by sending a JSON request to a mock server and verifying the response.
func TestClient_Put(t *testing.T) {
	defer gock.Off()

	gock.New("https://server.com").
		Put("/").
		MatchType("json").
		JSON(map[string]string{"foo": "bar"}).
		Reply(200).
		JSON(map[string]string{"foo": "bar"})

	client := NewClientBuilder().Build()
	gock.InterceptClient(client.client)

	body := map[string]string{"foo": "bar"}
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	res, err := client.Put("https://server.com", header, body)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}

	if res.statusCode != 200 {
		t.Errorf("expected status code 200, got %d", res.statusCode)
	}
}

// TestClient_Patch tests the Client's ability to send an HTTP PATCH request and handle the response correctly.
func TestClient_Patch(t *testing.T) {
	defer gock.Off()

	gock.New("https://server.com").
		Patch("/").
		MatchType("json").
		JSON(map[string]string{"foo": "bar"}).
		Reply(204).
		JSON(map[string]string{"foo": "bar"})

	client := NewClientBuilder().Build()
	gock.InterceptClient(client.client)

	body := map[string]string{"foo": "bar"}
	header := http.Header{}
	header.Add("Content-Type", "application/json")

	res, err := client.Patch("https://server.com", header, body)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}
	if res.statusCode != 204 {
		t.Errorf("expected status code 204, got %d", res.statusCode)
	}
}

// TestClient_Delete validates the HTTP DELETE request functionality of the Client.
// It mocks a DELETE request to "https://server.com/bar" and expects a 204 No Content status in the response.
// The test ensures no errors occur during the request and the correct status code is returned.
func TestClient_Delete(t *testing.T) {
	defer gock.Off()
	gock.New("https://server.com").
		Delete("/bar").
		Reply(204)

	client := NewClientBuilder().Build()
	gock.InterceptClient(client.client)

	res, err := client.Delete("https://server.com/bar", nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}
	if res.statusCode != 204 {
		t.Errorf("expected status code 204, got %d", res.statusCode)
	}
}

// TestClient_Do tests the Do function of the Client type to ensure it correctly sends a GET request and handles the response.
func TestClient_Do(t *testing.T) {
	defer gock.Off()

	gock.New("https://server.com").
		Get("/").
		Reply(200).
		JSON(map[string]string{"hello": "world"})

	client := NewClientBuilder().Build()
	gock.InterceptClient(client.client)

	req, err := http.NewRequest("GET", "https://server.com", nil)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	res, err := client.Do(req)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if res == nil {
		t.Fatal("expected response, got nil")
	}

	if res.statusCode != 200 {
		t.Errorf("expected status code 200, got %d", res.statusCode)
	}

}
