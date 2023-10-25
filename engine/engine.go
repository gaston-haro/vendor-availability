package engine

import (
	"fmt"

	"github.com/pedidosya/vendor-availability/models"
)

// Engine implementation
type Engine struct {
	Env     string
	Version string
}

// New engine instantiation
func New(cfg *models.Config) (Spec, error) {
	// Use cfg here to initialize engine
	return &Engine{
		Env:     cfg.Env,
		Version: cfg.Version,
	}, nil
}

// GetInfo implementation
func (e *Engine) GetInfo() map[string]interface{} {
	return map[string]interface{}{
		"env":     e.Env,
		"version": e.Version,
	}
}

// SayHi implementation
func (e *Engine) SayHi(name string) string {
	return fmt.Sprintf("Hi %s, how are you?", name)
}
