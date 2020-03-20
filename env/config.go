package env

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path/filepath"
)

type Config struct {
	Env string
	L   struct {
		Mode string `yaml:"mode"`
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
	path := fmt.Sprintf("./env/%s.yml", env)
	envPath, _ := filepath.Abs(path)
	log.Printf("path:%s :: envPath:%s\n", path, envPath)

	data, err := ioutil.ReadFile(envPath)
	if err != nil {
		log.Fatalf("could not read env config file %v", err)
		return config, false
	}

	//unmarshal config
	err = yaml.Unmarshal([]byte(data), &config)
	if err != nil {
		log.Fatalf("could not unmarshall yaml file %v", err)
		return config, false
	}

	return config, true
}
