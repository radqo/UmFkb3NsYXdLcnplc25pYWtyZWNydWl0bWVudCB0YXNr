package model

// Weather -  contains weather result by city or errors if result cannot be retreived
type Weather struct {
	Cities []CityWeather
	Errors []WeatherError
}

// CityWeather - weather information for city
type CityWeather struct {
	Name        string
	Description string
	Temperature Temperature
	Pressure    int32
	Humidity    int32
	Wind        Wind
	Cloudiness  int32
}

// Temperature - temperature information
type Temperature struct {
	Current    float32
	Minimal    float32
	Maximal    float32
	FeellsLike float32
}

// Wind - wind information
type Wind struct {
	Speed float32
	Deg   int32
}

// WeatherError - error information
type WeatherError struct {
	City         string
	ErrorMessage string
}

