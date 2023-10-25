package engine

import (
	"testing"

	"github.com/pedidosya/@project_name@/models"
	"github.com/stretchr/testify/assert"
)

func TestGetInfo(t *testing.T) {
	e, err := New(&models.Config{
		Version: "test version",
		Env:     "test env",
	})

	assert.Nil(t, err)
	assert.Equal(t, "test version", e.GetInfo()["version"])
	assert.Equal(t, "test env", e.GetInfo()["env"])
}

func TestHello(t *testing.T) {
	e, err := New(&models.Config{
		Version: "test version",
		Env:     "test env",
	})

	assert.Nil(t, err)
	assert.Equal(t, "Hi john, how are you?", e.SayHi("john"))
}
