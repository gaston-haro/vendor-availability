package main

import (
	"net/http"

	"github.com/pedidosya/peya-go/server"
	"github.com/pedidosya/vendor-availability/engine"
)

func handleHealth(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	server.OK(
		w, r, map[string]interface{}{
			"name": "project_name",
			"info": e.GetInfo(),
		},
	)
}
