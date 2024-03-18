package cmd

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewApplication(t *testing.T) {
	cmd := NewApplication()
	assert.NotNil(t, cmd)
	assert.True(t, cmd.SilenceUsage)
	assert.True(t, cmd.SilenceErrors)
}
