package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type owmWeather struct {
	Main        string
	Description string
}

type owmMain struct {
	Temp      float32
	FeelsLike float32 `json:"feels_like"`
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	Pressure  int32
	Humidity  int32
}

type ownWind struct {
	Speed float32
	Deg   int32
}

type owmClouds struct {
	All int32
}

type owmResponse struct {
	Weather    []owmWeather
	Name       string
	Main       owmMain
	Visibility int32
	Wind       ownWind
	Clouds     owmClouds
}

type owmError struct {
	Code string
}

func (m owmError) Error() string {
	return m.Code
}

type httpClient interface {
	Get(url string) (resp *http.Response, err error)
}

type owmService struct {
	client           httpClient
	counter          int
	incrementChannel chan string
	conf             weatherConfiguration
}

func newOwmService(c weatherConfiguration, client httpClient) *owmService {
	s := &owmService{}
	s.conf = c
	s.client = client
	s.counter = 0
	s.incrementChannel = make(chan string)
	go s.increment()
	return s
}

func (s *owmService) increment() {
	for {
		x := <-s.incrementChannel
		s.counter = s.counter + 1
		log.Printf("%v %s\n", s.counter, x)
	}
}

func (s *owmService) getOpenWeather(city string) (owmResponse, error) {

	url := fmt.Sprintf("%s?q=%s&lang=%s&appid=%s&units=%s", s.conf.URL, city, s.conf.Lang, s.conf.APIKey, s.conf.Units)

	response, err := s.client.Get(url)

	s.incrementChannel <- city

	if err != nil {
		return owmResponse{}, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		e := owmError{Code: response.Status}
		return owmResponse{}, e
	}

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error reading body [%s]\n", err.Error())
		return owmResponse{}, owmError{Code: "500"}
	}

	resp := owmResponse{}
	err = json.Unmarshal(b, &resp)

	if err != nil {
		log.Printf("Error deserializing body [%s]\n", err.Error())
		return owmResponse{}, owmError{Code: "500"}
	}

	return resp, nil
}
