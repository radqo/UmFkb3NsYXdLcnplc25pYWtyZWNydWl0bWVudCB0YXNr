package main

import (
	"gotest.tools/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServer(t *testing.T) {

	sut := &server{service: newService(configuration{}, clientSrv{
		getResp: func() (*http.Response, error) {
			response := httptest.ResponseRecorder{
				Code: 404,
			}
			return response.Result(), nil
		},
	})}

	sut.Run("50300")
	sut.Shutdown()
	sut.waitGroup.Wait()
}

func TestAbout(t *testing.T) {
	sut := &server{}

	body := strings.NewReader("{}")

	request := httptest.NewRequest("Get", "/", body)

	respRec := httptest.NewRecorder()

	sut.About(respRec, request)

	assert.Equal(t, respRec.Result().StatusCode, http.StatusOK)
}

type getterMock struct {
	cities []string
}

func (g *getterMock) getWeather(cities []string) weatherModel {
	g.cities = cities
	return weatherModel{}
}

func TestWeather(t *testing.T) {

	mock := &getterMock{}
	sut := &server{service: mock}

	body := strings.NewReader("{}")

	request := httptest.NewRequest("Get", "http://localhost:50300/weather?city=aaa&city=bbb", body)

	respRec := httptest.NewRecorder()

	sut.GetWeather(respRec, request)

	assert.Equal(t, respRec.Result().StatusCode, http.StatusOK)
	assert.Equal(t, len(mock.cities), 2)
}

func TestWeatherBadRequest(t *testing.T) {

	mock := &getterMock{}
	sut := &server{service: mock}

	body := strings.NewReader("{}")

	request := httptest.NewRequest("Get", "http://localhost:50300/weather", body)

	respRec := httptest.NewRecorder()

	sut.GetWeather(respRec, request)

	assert.Equal(t, respRec.Result().StatusCode, http.StatusBadRequest)
}
