package webs

import (
	"errors"
	"sync"
)

var mockServer = MockServer{
	mocks: make(map[string]*Mock),
}

type MockServer struct {
	enabled bool
	mutex   sync.Mutex
	mocks   map[string]*Mock
}

func StartMockServer() {
	mockServer.mutex.Lock()
	defer mockServer.mutex.Unlock()

	mockServer.enabled = true
}

func StopMockServer() {
	mockServer.mutex.Lock()
	defer mockServer.mutex.Unlock()
	mockServer.enabled = false
}

func AddMock(mock Mock) {
	mockServer.mutex.Lock()
	defer mockServer.mutex.Unlock()

	key := mockServer.getMockKey(mock.Method, mock.Url, mock.RequestBody)
	mockServer.mocks[key] = &mock
}

func (ms *MockServer) getMockKey(method, url, body string) string {
	return method + url + body
}

func (ms *MockServer) getMock(method, url, body string) *Mock {
	if !ms.enabled {
		return nil
	}

	if mock := ms.mocks[ms.getMockKey(method, url, body)]; mock != nil {
		return mock
	}
	mock := Mock{
		Error: errors.New("mock not found for " + method + " " + url + " " + body + ""),
	}
	return &mock
}
