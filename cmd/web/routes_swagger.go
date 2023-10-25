package main

import (
	"net/http"
	"os"

	"github.com/pedidosya/@project_name@/models"
	"github.com/pedidosya/peya-go/logs"
	"github.com/pedidosya/peya-go/server"
)

// Associate route to expose swagger apidoc on non-productive endpoints
func addApidocRoutes(cfg *models.Config, s *server.Server) {
	if cfg.Documentation.Enabled {
		// Expose swagger YAML definition
		s.AddRouteWithOptions(
			"/doc/swagger",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/octet-stream")
				http.ServeFile(w, r, "api"+string(os.PathSeparator)+"swagger.yaml")
			},
			&server.RouteOptions{
				CORSEnabled: true,
				CORSOptions: []server.CORSOption{
					server.AllowedMethods([]string{http.MethodGet}),
					server.AllowedOrigins([]string{"https://jarvis.peya.app"}),
				},
				TimeOutSeconds: 0,
			},
			http.MethodGet)

		// Expose swagger compiled HTML view
		s.Static("/api/", "api")
		s.AddRouteWithOptions(
			"/doc",
			func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "text/html")
				http.ServeFile(w, r, "api"+string(os.PathSeparator)+"api.html")
			},
			&server.RouteOptions{
				CORSEnabled: true,
				CORSOptions: []server.CORSOption{
					server.AllowedMethods([]string{http.MethodGet}),
					server.AllowedOrigins([]string{"https://jarvis.peya.app"}),
				},
				TimeOutSeconds: 0,
			},
			http.MethodGet)
		logs.Info("main: documentation enabled")
	} else {
		logs.Info("main: documentation disabled")
	}
}
