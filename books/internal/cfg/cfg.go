package cfg

import (
	"encoding/json"
	"flag"
	"io/ioutil"
)

var config appConfig

type appConfig struct {
	DBUrl                     string `json:"db_url"`
	DBName                    string `json:"db_name"`
	BooksCollectionName       string `json:"books_collection_name"`
	ScoresCollectionName      string `json:"scores_collection_name"`
	BooksDataCollectionName   string `json:"books_data_collection_name"`
	CurrentPageCollectionName string `json:"current_page_collection_name"`

	SecurityToken string `json:"security_token"`

	Url  string `json:"url"`
	Port string `json:"port"`

	IsProd bool `json:"is_prod"`
}

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

	if config.DBName == "" ||
		config.BooksCollectionName == "" ||
		config.BooksDataCollectionName == "" ||
		config.ScoresCollectionName == "" ||
		config.CurrentPageCollectionName == "" ||
		config.Port == "" ||
		config.SecurityToken == "" {
		panic("fill all the config values!!!!")
	}
}

func GetConfig() appConfig {
	return config
}
