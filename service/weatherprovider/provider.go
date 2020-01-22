package weatherprovider

import (
	"strings"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
)

type service struct {
	getter abstraction.CityWeatherGetter
}

// New - creates new provider instance
func New(getter abstraction.CityWeatherGetter) abstraction.WeatherProvider {
	return &service{getter: getter}
}

func (s *service) GetWeather(cities []string) (model.Weather) {

	result := model.Weather{
		Errors: []model.WeatherError{},
		Cities: []model.CityWeather{},
	}

	ch := make(chan chanResult)

	for _, city := range cities {

		go s.goGetWeatherForCity(strings.ToLower(city), ch)
	}

	for i := 0; i < len(cities); i++ {

		cres := <-ch

		if cres.Error != nil {
			result.Errors = append(result.Errors,
				model.WeatherError{
					City:         cres.City,
					ErrorMessage: cres.Error.Error(),
				})
		} else {
			result.Cities = append(result.Cities, *cres.Model)
		}
	}

	return result
}

type chanResult struct {
	City  string
	Model *model.CityWeather
	Error error
}

func (s *service) goGetWeatherForCity(city string, c chan chanResult) {
	m, err := s.getter.GetWeather(city)
	c <- chanResult{Model: m, Error: err, City: city}
}


