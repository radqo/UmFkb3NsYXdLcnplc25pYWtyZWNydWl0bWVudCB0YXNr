package main

import (
	"strings"
)

type service struct {
	weatherCache *cache
	owm          *owmService
}

type weatherGetter interface {
	getWeather(cities []string) weatherModel
}

func newService(c configuration, client httpClient) weatherGetter {
	s := &service{}
	s.weatherCache = newCache(c.API.CacheTimeoutInSeconds)
	s.owm = newOwmService(c.Weather, client)
	return s
}

func (s *service) getWeather(cities []string) weatherModel {

	result := weatherModel{
		Errors: []weatherErrorModel{},
		Cities: []cityWeatherModel{},
	}

	ch := make(chan chanResult)

	for _, city := range cities {

		go s.goGetWeatherForCity(strings.ToLower(city), ch)
	}

	for i := 0; i < len(cities); i++ {

		cres := <-ch

		if cres.Error != nil {
			result.Errors = append(result.Errors,
				weatherErrorModel{
					City:         cres.City,
					ErrorMessage: cres.Error.Error(),
				})
		} else {
			result.Cities = append(result.Cities, cres.Model)
		}
	}

	return result
}

type chanResult struct {
	City  string
	Model cityWeatherModel
	Error error
}

func (s *service) goGetWeatherForCity(city string, c chan chanResult) {
	model, err := s.weatherCache.Get(city, s.getWeatherForCity)

	m := cityWeatherModel{}

	switch v := model.(type) {
	case cityWeatherModel:
		m = v
	}

	c <- chanResult{Model: m, Error: err, City: city}
}

func (s *service) getWeatherForCity(city string) (interface{}, error) {
	resp, err := s.owm.getOpenWeather(city)
	if err != nil {
		return cityWeatherModel{}, err
	}
	return convert(resp), nil
}

func convert(r owmResponse) cityWeatherModel {
	cw := cityWeatherModel{
		Name:        r.Name,
		Description: r.Weather[0].Description,
		Temperature: temperatureModel{
			Current:    r.Main.Temp,
			Minimal:    r.Main.TempMin,
			Maximal:    r.Main.TempMax,
			FeellsLike: r.Main.FeelsLike,
		},
		Pressure: r.Main.Pressure,
		Humidity: r.Main.Humidity,
		Wind: windModel{
			Speed: r.Wind.Speed,
			Deg:   r.Wind.Deg,
		},
		Cloudiness: r.Clouds.All,
	}
	return cw
}
