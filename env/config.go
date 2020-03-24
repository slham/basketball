package env

import (
	"github.com/slham/toolbelt/l"
	"os"
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

func Load() (Config, bool) {
	//load config file
	var config Config
	//env := os.Getenv("ENVIRONMENT")
	//config.Env = env
	//wd, _ := os.Getwd()
	//for !strings.HasSuffix(wd, "basketball") {
	//	wd = filepath.Dir(wd)
	//}
	//path := fmt.Sprintf("./env/%s.yml", env)
	//envPath, _ := filepath.Abs(path)
	//l.Debug(nil, "path:%s :: envPath:%s", path, envPath)
	//
	//data, err := ioutil.ReadFile(envPath)
	//if err != nil {
	//	l.Error(nil, "could not read env config file from %s %v", envPath, err)
		config.Runtime.Port = os.Getenv("RUNTIME_PORT")
		config.Storage.FileName = os.Getenv("STORAGE_FILENAME")
		config.Storage.Bucket = os.Getenv("STORAGE_BUCKET")
		config.Storage.Prefix = os.Getenv("STORAGE_PREFIX")
		config.L.Level = l.Level(os.Getenv("LOG_LEVEL"))
	//	return config, true
	//}
	//
	////unmarshal config
	//err = yaml.Unmarshal(data, &config)
	//if err != nil {
	//	l.Error(nil, "could not unmarshall yaml file %v", err)
	//	return config, false
	//}

	return config, true
}
