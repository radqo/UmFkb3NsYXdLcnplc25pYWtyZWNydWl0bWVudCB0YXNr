package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"sync"
	"time"
)

type server struct {
	service    weatherGetter
	httpServer *http.Server
	waitGroup  sync.WaitGroup
}

func (s *server) About(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"about": "weather api"}`))
}

func (s *server) GetWeather(w http.ResponseWriter, r *http.Request) {
	qkeys := r.URL.Query()
	w.Header().Set("Content-Type", "application/json")

	cities := qkeys["city"]

	if len(cities) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "no cities specified"}`))
	} else {
		w.WriteHeader(http.StatusOK)

		resp, err := json.Marshal(s.service.getWeather(cities))

		if err != nil {
			w.Write([]byte(err.Error()))
		} else {
			w.Write(resp)
		}
	}
}

func (s *server) Run(port string) {
	r := mux.NewRouter()

	r.HandleFunc("/", s.About).Methods(http.MethodGet)
	r.HandleFunc("/weather", s.GetWeather).Methods(http.MethodGet)

	s.waitGroup.Add(1)
	s.httpServer = &http.Server{Addr: fmt.Sprintf(":%s", port), Handler: r}

	go func() {

		if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		s.waitGroup.Done()
	}()
}

func (s *server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
}
