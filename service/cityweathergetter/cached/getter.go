package cached

import (
	"errors"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
)

type weatherGetter struct {
	getter abstraction.CityWeatherGetter
	cache  abstraction.CacheGetter
}

// New - creates new instance of cached weather getter
func New(getter abstraction.CityWeatherGetter, cache abstraction.CacheGetter) abstraction.CityWeatherGetter {
	return &weatherGetter{ getter: getter, cache: cache}
}	

func (w *weatherGetter) GetWeather(city string) (m *model.CityWeather, err error) {

	defer func() {
		if r:=recover();r!=nil{
			m = nil
			err = errors.New("Wrong item type")
		}
	}()

	x, err := w.cache.Get(city, w.getFunc)
	return x.(*model.CityWeather), err	
}

func (w *weatherGetter) getFunc(key string) (interface{}, error) {
	return w.getter.GetWeather(key)
}