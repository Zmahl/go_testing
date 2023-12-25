package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubStore struct {
	response  string
	cancelled bool
}

func (s *StubStore) Fetch() string {
	return s.response
}

func (s *StubStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	data := "hello, world"
	store := &StubStore{response: data}
	svr := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}
