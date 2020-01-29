package openweather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"errors"

	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
)

// Configuration - open weather client configuration
type Configuration struct {
	URL    string `json:"url"`
	APIKey string `json:"apikey"`
	Lang   string `json:"lang"`
	Units  string `json:"units"`
}

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
	client  httpClient
	counter int
	conf    Configuration
}

//New - creates new instance of client
func New(c Configuration, client httpClient) abstraction.CityWeatherGetter {
	return &owmService{conf: c, client: client}
}

func (s *owmService) GetWeather(city string) (m *model.CityWeather, err error) {
	log.Println("Call for city: " + city)

	defer func() {
		if r:=recover();r!=nil{
			log.Println(r)
			err = errors.New("Internal server error")
			m = nil
		}
	}()

	url := fmt.Sprintf("%s?q=%s&lang=%s&appid=%s&units=%s", s.conf.URL, city, s.conf.Lang, s.conf.APIKey, s.conf.Units)

	response, err := s.client.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		e := owmError{Code: response.Status}
		return nil, e
	}

	b, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Printf("Error reading body [%s]\n", err.Error())
		return nil, owmError{Code: "500"}
	}

	resp := owmResponse{}
	err = json.Unmarshal(b, &resp)

	if err != nil {
		log.Printf("Error deserializing body [%s]\n", err.Error())
		return nil, owmError{Code: "500"}
	}

	return convert(resp), nil
}

func convert(r owmResponse) *model.CityWeather {
	cw := &model.CityWeather{
		Name:        r.Name,
		Description: r.Weather[0].Description,
		Temperature: model.Temperature{
			Current:    r.Main.Temp,
			Minimal:    r.Main.TempMin,
			Maximal:    r.Main.TempMax,
			FeellsLike: r.Main.FeelsLike,
		},
		Pressure: r.Main.Pressure,
		Humidity: r.Main.Humidity,
		Wind: model.Wind{
			Speed: r.Wind.Speed,
			Deg:   r.Wind.Deg,
		},
		Cloudiness: r.Clouds.All,
	}
	return cw
}
