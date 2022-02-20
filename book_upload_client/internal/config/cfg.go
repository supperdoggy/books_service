package cfg

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

type appConfig struct {
	PathToTxt string `json:"path_to_txt"`
	SecurityToken string `json:"security_token"`

	Url string `json:"url"`
	Port string `json:"port"`
}

var config appConfig

func init() {
	path := flag.String("config", "", "for config path")
	flag.Parse()
	if path == nil || *path == "" {
		panic("no config path specified")
	}

	data, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err.Error())
	}

	if err := json.Unmarshal(data, &config); err != nil {
		panic(err.Error())
	}

	if config.PathToTxt == "" ||
		config.SecurityToken == "" ||
		config.Url == "" ||
		config.Port == "" {
		panic("fill all the config values!!!!")
	}
}

func GetConfig() appConfig {
	return config
}

