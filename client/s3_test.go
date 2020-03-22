package client

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestInitializeSession(t *testing.T) {
	InitializeSession()
	assert.Equal(t, sess.Config.Region, "us-west-2")
}
