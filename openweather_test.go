package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestOwmResponseDeserialization(t *testing.T) {

	raw := `{"coord":{"lon":2.35,"lat":48.85},"weather":[{"id":804,"main":"Clouds","description":"całkowite zachmurzenie","icon":"04n"}],"base":"stations","main":{"temp":7.1,"feels_like":3.47,"temp_min":6.11,"temp_max":8.33,"pressure":1029,"humidity":87},"visibility":10000,"wind":{"speed":3.6,"deg":300},"clouds":{"all":90},"dt":1578674817,"sys":{"type":1,"id":6550,"country":"FR","sunrise":1578642114,"sunset":1578672803},"timezone":3600,"id":2988507,"name":"Paryż","cod":200}`

	resp := owmResponse{}
	err := json.Unmarshal([]byte(raw), &resp)

	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, resp.Name, "Paryż")
	assert.Equal(t, resp.Main.Temp, float32(7.1))
	assert.Equal(t, resp.Wind.Speed, float32(3.6))
}

type client struct {
}

func (c client) Get(url string) (resp *http.Response, err error) {

	b := bytes.NewBufferString(`{"coord":{"lon":2.35,"lat":48.85},"weather":[{"id":804,"main":"Clouds","description":"całkowite zachmurzenie","icon":"04n"}],"base":"stations","main":{"temp":7.1,"feels_like":3.47,"temp_min":6.11,"temp_max":8.33,"pressure":1029,"humidity":87},"visibility":10000,"wind":{"speed":3.6,"deg":300},"clouds":{"all":90},"dt":1578674817,"sys":{"type":1,"id":6550,"country":"FR","sunrise":1578642114,"sunset":1578672803},"timezone":3600,"id":2988507,"name":"Paryż","cod":200}`)

	response := httptest.ResponseRecorder{
		Code: 200,
		Body: b,
	}

	return response.Result(), nil
}

func TestGetOpenWeatherOK(t *testing.T) {

	conf := weatherConfiguration{
		URL:                    "abc",
		APIKey:                 "x",
		Lang:                   "pl",
		Units:                  "metric",
		ClientTimeoutInSeconds: 30,
	}

	sut := newOwmService(conf, client{})

	r, err := sut.getOpenWeather("city name")

	assert.NilError(t, err)
	assert.Equal(t, r.Name, "Paryż")

}

type client404 struct {
}

func (c client404) Get(url string) (resp *http.Response, err error) {

	b := bytes.NewBufferString(`{"coord":{"lon":2.35,"lat":48.85},"weather":[{"id":804,"main":"Clouds","description":"całkowite zachmurzenie","icon":"04n"}],"base":"stations","main":{"temp":7.1,"feels_like":3.47,"temp_min":6.11,"temp_max":8.33,"pressure":1029,"humidity":87},"visibility":10000,"wind":{"speed":3.6,"deg":300},"clouds":{"all":90},"dt":1578674817,"sys":{"type":1,"id":6550,"country":"FR","sunrise":1578642114,"sunset":1578672803},"timezone":3600,"id":2988507,"name":"Paryż","cod":404}`)

	response := httptest.ResponseRecorder{
		Code: 404,
		Body: b,
	}

	return response.Result(), nil
}

func TestGetOpenWeatherNotFound(t *testing.T) {

	conf := weatherConfiguration{
		URL:                    "abc",
		APIKey:                 "x",
		Lang:                   "pl",
		Units:                  "metric",
		ClientTimeoutInSeconds: 30,
	}

	sut := newOwmService(conf, client404{})

	_, err := sut.getOpenWeather("city name")

	assert.Check(t, strings.Contains(err.Error(), "404"))
}
