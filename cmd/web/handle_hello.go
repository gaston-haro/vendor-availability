package main

import (
	"net/http"

	"github.com/pedidosya/peya-go/logs"
	"github.com/pedidosya/peya-go/server"
	"github.com/pedidosya/vendor-availability/engine"
)

func handleHello(e engine.Spec, w http.ResponseWriter, r *http.Request) {
	logs.Debug("handling hello request")

	name := server.GetStringFromPath(r, "name", "no name")
	server.OK(w, r, helloResponse{Message: e.SayHi(name)})
}

type helloResponse struct {
	Message string `json:"message"`
}
