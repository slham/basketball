package env

import (
	"github.com/slham/toolbelt/l"
)

type Config struct {
	Env string
	L   struct {
		Level l.Level `yaml:"level"`
	} `yaml:"l,omitempty"`
	Runtime struct {
		Port string `yaml:"port"`
	} `yaml:"runtime,omitempty"`
	Storage struct {
		FileName string `yaml:"fileName"`
		Bucket   string `yaml:"bucket"`
		Prefix   string `yaml:"prefix"`
	} `yaml:"storage"`
}
