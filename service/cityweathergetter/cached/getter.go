package cached

import (
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
)

type weatherGetter struct {
	getter abstraction.CityWeatherGetter
	cache  abstraction.CacheOperator
}

// New - creates new instance of cached weather getter
func New(getter abstraction.CityWeatherGetter, cache abstraction.CacheOperator) abstraction.CityWeatherGetter {
	return &weatherGetter{getter: getter, cache: cache}
}

func (w *weatherGetter) GetWeather(city string) (m *model.CityWeather, err error) {

	x, found := w.cache.Get(city)

	if found {
		return x.(*model.CityWeather), nil
	}

	m, err = w.getter.GetWeather(city)

	if err != nil {
		return nil, err
	}

	w.cache.Set(city, m)

	return m, nil
}
