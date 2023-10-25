package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pedidosya/@project_name@/engine"
	"github.com/pedidosya/@project_name@/models"
	"github.com/pedidosya/peya-go/config"
	"github.com/pedidosya/peya-go/logs"
	"github.com/pedidosya/peya-go/server"
	"github.com/pedidosya/peya-go/vault"
)

func main() {
	logger := logs.InitDefault()
	// logs.SetLevel("debug")

	// Parse dev executable flag
	env := flag.String("E", "dev", "Execution environment")
	flag.Parse()

	// Get ENV environment variable
	osValue := os.Getenv("ENV")
	if osValue != "" {
		env = &osValue
	}

	// Read configuration file for environment
	cfg, err := readConfig(*env)
	if err != nil {
		panic(fmt.Sprintf("main: can't read config: %s", err.Error()))
	}
	setLogLevel(cfg)

	// Set up engine
	e, err := createEngine(cfg)
	if err != nil {
		panic(fmt.Sprintf("main: can't start engine: %s", err.Error()))
	}

	// Set up server instance
	if cfg.Server != nil {
		cfg.Server.Logger = logger
	}
	s := createServer(cfg)
	startProfiler(cfg, s)

	// Bind application handlers to endpoints
	routes(s, cfg, e)

	// Start server.ServeHTTP() loop
	server.MainLoop(s)
}

func readConfig(env string) (*models.Config, error) {
	cfg := &models.Config{}
	config.ReadConfiguration(env, cfg)
	cfg.Env = env
	cfg.Version = os.Getenv("VERSION")

	vaultValues, err := vault.Read("https://vault.peya.co/v1/secret/@project_name@", env)
	if err != nil {
		panic(fmt.Sprintf("main: error from fault. no vault values found for env %s | %s", env, err.Error()))
	}

	if len(vaultValues) == 0 {
		logs.Warn("main: no values found in vault")
	}

	return cfg, nil
}

func createEngine(cfg *models.Config) (engine.Spec, error) {
	e, err := engine.New(cfg)
	return e, err
}

func createServer(cfg *models.Config) *server.Server {
	s := server.New(cfg.Server)
	return s
}

func startProfiler(cfg *models.Config, s *server.Server) {
	if cfg.Diagnostics.EnablePprof {
		s.EnablePprof(func() bool { return cfg.Diagnostics.EnablePprof })
		logs.Info("main: enabled pprof")
	}
}

func setLogLevel(cfg *models.Config) {
	if cfg.Diagnostics.LogLevel != "" {
		err := logs.SetLevel(cfg.Diagnostics.LogLevel)
		if err != nil {
			logs.Infof("main: main: error setting log level to %s - %s", cfg.Diagnostics.LogLevel, err.Error())
		} else {
			logs.Infof("main: main: log level set to %s", cfg.Diagnostics.LogLevel)
		}
	}
}
