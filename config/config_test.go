package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig(t *testing.T) {
	config := New()
	assert.NotNil(t, config)
}

func TestWithDatabase(t *testing.T) {
	config := New(WithDatabase("sqlite", "db.sqlite"))
	assert.Equal(t, "sqlite", config.Database.Driver)
	assert.Equal(t, "db.sqlite", config.Database.DSN)
}

func TestWithPort(t *testing.T) {
	config := New(WithPort(8080))
	assert.Equal(t, 8080, config.Server.Port)
}

func TestWithDebug(t *testing.T) {
	config := New(WithDebug(true))
	assert.True(t, config.Server.Debug)
}

func TestWithGracePeriod(t *testing.T) {
	config := New(WithGracePeriod(30))
	assert.Equal(t, 30, config.Server.GracePeriod)
}
