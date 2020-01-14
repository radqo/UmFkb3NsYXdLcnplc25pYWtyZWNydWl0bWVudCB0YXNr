package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("app start")

	conf := getConfiguration()

	client := &http.Client{
		Timeout: time.Duration(conf.Weather.ClientTimeoutInSeconds) * time.Second,
	}

	s := &server{service: newService(conf, client)}

	s.Run(conf.API.Port)

	s.waitGroup.Wait()
}
