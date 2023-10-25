package main

import (
	"testing"

	"github.com/pedidosya/@project_name@/models"
	"github.com/pedidosya/peya-go/server"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	// Mock engine so we can fix the expected response
	e := &mockedEngine{}
	e.On("GetInfo").Return(map[string]interface{}{"hello": "world"})

	// Create a server and define routes
	cfg := &models.Config{Server: &server.Config{}}
	s := createServer(cfg)
	routes(s, cfg, e)

	// Create a test request and ask server to serve it
	// this allows to test the route definition
	req, rr := createGETTest(t, "/")
	s.ServeHTTP(rr, req)

	// Assert response based on mocked engine behaviour
	expected := `{"name":"@project_name@","info":{"hello":"world"}}`
	assert.JSONEq(t, expected, rr.Body.String())
}
