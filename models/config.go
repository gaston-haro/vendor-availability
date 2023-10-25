package models

import (
	"github.com/pedidosya/peya-go/server"
)

// Config for service
// Extend this struct to parse config values from env_* files
type Config struct {
	Env           string            `json:"-"`
	Server        *server.Config    `json:"server"`
	Version       string            `json:"string"`
	Diagnostics   DiagnosticsConfig `json:"diagnostics"`
	Documentation struct {
		Enabled bool `json:"enabled"`
	} `json:"documentation"`
}

type DiagnosticsConfig struct {
	LogLevel    string `json:"log_level"`
	EnablePprof bool   `json:"enable_pprof"`
}
