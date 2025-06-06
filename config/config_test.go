package config_test

import (
	"testing"

	"github.com/starrysilk/immortal/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadfromFile(t *testing.T) {
	cfg, err := config.Load("./config.yml")
	require.NoError(t, err, "error must be nil.")

	assert.Equal(t, uint16(7777), cfg.WebsocketServer.Port)
	assert.Equal(t, "0.0.0.0", cfg.WebsocketServer.Bind)
}
