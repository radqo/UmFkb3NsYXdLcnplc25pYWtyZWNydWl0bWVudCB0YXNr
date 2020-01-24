package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/radqo/UmFkb3NsYXdLcnplc25pYWtyZWNydWl0bWVudCB0YXNr/abstraction"
	"log"
	"net/http"
	"sync"
	"time"
)

type server struct {
	service    abstraction.WeatherProvider
	httpServer *http.Server
	waitGroup  sync.WaitGroup
}

// New - creates new instance of server
func New(provider abstraction.WeatherProvider) abstraction.Server {
	s := server{service: provider}
	return &s
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

	s.waitGroup.Wait()
}

func (s *server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)
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

		resp, err := json.Marshal(s.service.GetWeather(cities))

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(resp)
		}
	}
}
