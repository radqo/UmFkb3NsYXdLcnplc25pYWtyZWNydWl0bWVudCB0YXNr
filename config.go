package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

const confFile = "config.json"

type apiConfiguration struct {
	Port                  string `json:"port"`
	CacheTimeoutInSeconds int    `json:"cacheTimeoutInSeconds"`
}

type weatherConfiguration struct {
	URL                    string `json:"url"`
	APIKey                 string `json:"apikey"`
	Lang                   string `json:"lang"`
	Units                  string `json:"units"`
	ClientTimeoutInSeconds int    `json:"clientTimeoutInSeconds"`
}

type configuration struct {
	API     apiConfiguration     `json:"api"`
	Weather weatherConfiguration `json:"weather"`
}

var _conf configuration

func getConfiguration() configuration {
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
