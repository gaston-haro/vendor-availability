package main

import (
	"testing"

	"github.com/pedidosya/@project_name@/models"
	"github.com/pedidosya/peya-go/server"
	"github.com/stretchr/testify/assert"
)

func TestHello(t *testing.T) {
	// Mock engine so we can fix the expected response
	e := &mockedEngine{}
	e.On("SayHi", "John").Return("hello john")

	// Create a server and define routes
	cfg := &models.Config{Server: &server.Config{}}
	s := createServer(cfg)
	routes(s, cfg, e)

	// Create a test request and ask server to serve it
	// this allows to test the route definition
	req, rr := createGETTest(t, "/v1/hello/John")
	s.ServeHTTP(rr, req)

	// Assert response based on mocked engine behaviour
	expected := `{"message": "hello john"}`
	assert.JSONEq(t, expected, rr.Body.String())
}
