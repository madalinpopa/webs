package webs

import "testing"

func TestClientBuilder_Build(t *testing.T) {

	client := NewClientBuilder().Build()
	if client == nil {
		t.Errorf("expected client, got nil")
	}
}
