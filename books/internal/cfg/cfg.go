package cfg

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config appConfig

type appConfig struct {
	DBName string
	CollectionName string
	Port string

	IsProd bool
}

func init() {
	path := os.Getenv("config")
	if path == "" {
		panic("no config path specified")
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}
}

func GetConfig() appConfig {
	return config
}
