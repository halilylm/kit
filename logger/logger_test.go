package logger_test

import (
	"github.com/halilylm/kit/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("creates a sugared logger", func(t *testing.T) {
		l, err := logger.New("service")

		assert.NoError(t, err)
		assert.NotNil(t, l)
	})
}
