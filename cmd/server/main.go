package main

import (
	"log"
	"net/http"
	"time"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/tool"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/adapter/openweather"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/server/rest"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/weatherprovider"
	cache "github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/cache/inmemory"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/cityweathergetter/cached"
)

func main() {
	log.Println("app start")

	appConf := &model.AppConfiguration{}

	err := tools.Read("../../config/config.json", appConf)

	if err != nil {
		log.Fatal(err)
	}

	clientConf := &openweather.Configuration{}

	err = tools.Read("../../config/openweatherconfig.json", clientConf)

	if err != nil {
		log.Fatal(err)
	}


	client := &http.Client{
		Timeout: time.Duration(appConf.ClientTimeoutInSeconds) * time.Second,
	}

	openweathergetter := openweather.New(*clientConf, client)

	c := cache.New(appConf.CacheTimeoutInSeconds)

	cachedGetter := cached.New(openweathergetter, c)

	provider := weatherprovider.New(cachedGetter)

	s := rest.New(provider)

	s.Run(appConf.Port)

}


