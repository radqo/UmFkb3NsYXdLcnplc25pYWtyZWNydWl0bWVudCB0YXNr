package main

import "fmt"

type weatherModel struct {
	Cities []cityWeatherModel
	Errors []weatherErrorModel
}

type cityWeatherModel struct {
	Name        string
	Description string
	Temperature temperatureModel
	Pressure    int32
	Humidity    int32
	Wind        windModel
	Cloudiness  int32
}

type temperatureModel struct {
	Current    float32
	Minimal    float32
	Maximal    float32
	FeellsLike float32
}

type windModel struct {
	Speed float32
	Deg   int32
}

type weatherErrorModel struct {
	City         string
	ErrorMessage string
}

type apiError struct {
	Message string
	Code    int
}

func (e apiError) Error() string {
	return fmt.Sprintf("[%v] %s", e.Code, e.Message)
}
