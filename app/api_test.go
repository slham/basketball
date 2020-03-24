package app

import (
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestApp_Initialize(t *testing.T) {
	a := App{}
	assert.Equal(t, true, a.Initialize("dev"))
}

func TestApp_InitializeNoConfigFile(t *testing.T) {
	a := App{}
	assert.Equal(t, false, a.Initialize("scoobydoo"))
}