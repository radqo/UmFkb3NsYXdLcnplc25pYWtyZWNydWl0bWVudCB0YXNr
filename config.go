package main

import (
	"log"
	"io/ioutil"
	"encoding/json"
)

const confFile = "config.json"

type apiConfiguration struct {
	Port           string `json:"port"`
	CacheInMinutes int    `json:"cacheMinutes"`
}

type watherConfiguration struct {
	APIKey string `json:"apikey"`
	Lang   string `json:"lang"`
}

type configuration struct {
	API     apiConfiguration    `json:"api"`
	Weather watherConfiguration `json:"weather"`
}

var _conf apiConfiguration

func getConfiguration() apiConfiguration {
	return _conf
}

func init() {
	log.Println("config initialization")
	
	file, err := ioutil.ReadFile(confFile)

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &_conf)

	if err != nil {
		log.Fatal(err)
	}
}
