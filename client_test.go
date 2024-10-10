package webs

import "testing"

func TestClient_Build(t *testing.T) {

	client := NewClientBuilder().Build()
	if client == nil {
		t.Errorf("expected client, got nil")
	}
}

//func TestClient_Get(t *testing.T) {
//	builder := NewClientBuilder()
//
//}
