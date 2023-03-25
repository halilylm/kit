package docker_test

import (
	"github.com/halilylm/kit/docker"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStartContainer(t *testing.T) {
	var id string
	t.Run("starts the container", func(t *testing.T) {
		c, err := docker.StartContainer("redis", "6379")
		if err != nil {
			t.Fatal("error creating container", err.Error())
		}
		id = c.ID
		assert.NotNil(t, c)
	})
	t.Run("stops the container", func(t *testing.T) {
		err := docker.StopContainer(id)
		assert.NoError(t, err)
	})
}
