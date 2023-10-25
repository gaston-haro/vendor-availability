package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/mock"
)

func createGETTest(t *testing.T, url string) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal("failed to build request")
	}

	return req, httptest.NewRecorder()
}

func createPUTTest(t *testing.T, url string, bdy interface{}) (*http.Request, *httptest.ResponseRecorder) {
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		t.Fatal("failed to build request")
	}

	return req, httptest.NewRecorder()
}

type mockedEngine struct {
	mock.Mock
}

func (m *mockedEngine) GetInfo() map[string]interface{} {
	args := m.Called()
	return args.Get(0).(map[string]interface{})
}

func (m *mockedEngine) SayHi(name string) string {
	args := m.Called(mock.Anything)
	return args.String(0)
}
