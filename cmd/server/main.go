package main

import (
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/adapter/openweather"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/model"
	cache "github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/cache/inmemory"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/cityweathergetter/cached"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/server/rest"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/service/weatherprovider"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/tool"
	"log"
	"net/http"
	"path"
	"runtime"
	"time"
)

func main() {
	log.Println("app start")

	appConf := &model.AppConfiguration{}

	_, filename, _, _ := runtime.Caller(0)

	confPath := path.Join(path.Dir(filename), "../../config")

	err := tools.Read(path.Join(confPath, "config.json"), appConf)

	if err != nil {
		log.Fatal(err)
	}

	clientConf := &openweather.Configuration{}

	err = tools.Read(path.Join(confPath, "openweatherconfig.json"), clientConf)

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
