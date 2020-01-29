package openweather

import (
	"reflect"
	"testing"

	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
)

func Test_convert(t *testing.T) {
	type args struct {
		r owmResponse
	}
	tests := []struct {
		name string
		args args
		want *model.CityWeather
	}{
		{
			name: "Conversion",
			args: args{
				r: owmResponse{
					Weather: []owmWeather{{
						Main:        "Clouds",
						Description: "całkowite zachmurzenie",
					}},
					Main: owmMain{
						Temp:      7.1,
						FeelsLike: 3.47,
						TempMin:   6.11,
						TempMax:   8.33,
						Pressure:  1029,
						Humidity:  87,
					},
					Visibility: 10000,
					Wind: ownWind{
						Speed: 3.6,
						Deg:   300,
					},
					Clouds: owmClouds{
						All: 90,
					},
					Name: "Paryż",
				},
			},
			want: &model.CityWeather{
				Name:        "Paryż",
				Description: "całkowite zachmurzenie",
				Temperature: model.Temperature{
					Current:    7.1,
					Minimal:    6.11,
					Maximal:    8.33,
					FeellsLike: 3.47,
				},
				Pressure: 1029,
				Humidity: 87,
				Wind: model.Wind{
					Speed: 3.6,
					Deg:   300,
				},
				Cloudiness: 90,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}
		})
	}
}
