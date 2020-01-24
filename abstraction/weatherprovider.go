package abstraction

import "github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"

// WeatherProvider interface
type WeatherProvider interface {
	GetWeather(cities []string) model.Weather
}
