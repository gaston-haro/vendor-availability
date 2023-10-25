package main

import (
	"net/http"

	"github.com/pedidosya/peya-go/server"
	"github.com/pedidosya/vendor-availability/engine"
	"github.com/pedidosya/vendor-availability/models"
)

// Associate routes with handlers here
func routes(s *server.Server, cfg *models.Config, e engine.Spec) {
	s.AddRoute(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			handleHealth(e, w, r)
		},
	)

	s.AddRoute(
		"/v{version}/hello/{name}",
		func(w http.ResponseWriter, r *http.Request) {
			handleHello(e, w, r)
		},
	)

	addApidocRoutes(cfg, s)
}
