package env

import (
	"github.com/slham/toolbelt/l"
	"gopkg.in/go-playground/assert.v1"
	"testing"
)

func TestLoadDev(t *testing.T) {
	config, ok := Load("dev")
	assert.Equal(t, ok, true)
	assert.Equal(t, config.L.Level, l.DEBUG)
	assert.Equal(t, config.Storage.Prefix, "")
	assert.Equal(t, config.Storage.Bucket, "")
	assert.Equal(t, config.Storage.FileName, "1583510437.yaml")
	assert.Equal(t, config.Env, "dev")
	assert.Equal(t, config.Runtime.Port, "8090")
}

func TestLoadProd(t *testing.T) {
	config, ok := Load("prod")
	assert.Equal(t, ok, true)
	assert.Equal(t, config.L.Level, l.INFO)
	assert.Equal(t, config.Storage.Prefix, "player-stats/2020")
	assert.Equal(t, config.Storage.Bucket, "sheldonsandbox-basketball")
	assert.Equal(t, config.Storage.FileName, "")
	assert.Equal(t, config.Env, "prod")
	assert.Equal(t, config.Runtime.Port, "80")
}

func TestLoadNoFile(t *testing.T) {
	config, ok := Load("blah")
	assert.Equal(t, ok, false)
	assert.Equal(t, Config{Env: "blah"}, config)
}
