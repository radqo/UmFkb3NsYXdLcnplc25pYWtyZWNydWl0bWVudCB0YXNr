package main

import (
	"bytes"
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type clientSrv struct {
	getResp func() (*http.Response, error)
}

func (c clientSrv) Get(url string) (resp *http.Response, err error) {
	return c.getResp()
}
func TestServiceCitiesFound(t *testing.T) {

	sut := newService(configuration{}, clientSrv{
		getResp: func() (*http.Response, error) {
			b := bytes.NewBufferString(`{"coord":{"lon":2.35,"lat":48.85},"weather":[{"id":804,"main":"Clouds","description":"całkowite zachmurzenie","icon":"04n"}],"base":"stations","main":{"temp":7.1,"feels_like":3.47,"temp_min":6.11,"temp_max":8.33,"pressure":1029,"humidity":87},"visibility":10000,"wind":{"speed":3.6,"deg":300},"clouds":{"all":90},"dt":1578674817,"sys":{"type":1,"id":6550,"country":"FR","sunrise":1578642114,"sunset":1578672803},"timezone":3600,"id":2988507,"name":"Paryż","cod":404}`)

			response := httptest.ResponseRecorder{
				Code: 200,
				Body: b,
			}
			return response.Result(), nil
		},
	})

	resp := sut.getWeather([]string{"a", "b"})

	assert.Equal(t, len(resp.Cities), 2)
	assert.Equal(t, len(resp.Errors), 0)
}

func TestServiceCitiesNotFound(t *testing.T) {

	sut := newService(configuration{}, clientSrv{
		getResp: func() (*http.Response, error) {
			response := httptest.ResponseRecorder{
				Code: 404,
			}
			return response.Result(), nil
		},
	})

	resp := sut.getWeather([]string{"a", "b"})

	assert.Equal(t, len(resp.Cities), 0)
	assert.Equal(t, len(resp.Errors), 2)
}
