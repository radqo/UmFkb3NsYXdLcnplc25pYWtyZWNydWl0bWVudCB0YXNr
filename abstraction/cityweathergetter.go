package abstraction

import "github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"

// CityWeatherGetter - interface for weather service
type CityWeatherGetter interface {
	GetWeather(city string) (*model.CityWeather, error)
}
