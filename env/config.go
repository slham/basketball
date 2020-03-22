package env

import (
	"fmt"
	"github.com/slham/toolbelt/l"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
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

func Load(env string) (Config, bool) {
	//load config file
	var config Config
	config.Env = env
	wd, _ := os.Getwd()
	for !strings.HasSuffix(wd, "basketball") {
		wd = filepath.Dir(wd)
	}
	path := fmt.Sprintf("%s/env/%s.yml", wd, env)
	//envPath, _ := filepath.Abs(path)
	l.Debug(nil, "path:%s", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		l.Error(nil, "could not read env config file %v", err)
		return config, false
	}

	//unmarshal config
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		l.Error(nil, "could not unmarshall yaml file %v", err)
		return config, false
	}

	return config, true
}
