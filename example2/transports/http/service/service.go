package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mccurdyc/go-examples/example2/transports/http/handlers"
	"github.com/mccurdyc/go-examples/example2/transports/http/middleware"
)

type Service struct {
	Launched time.Time
	Server   http.Server
}

// NewService intiallizes a new service
// this would comparable to a constructor in PHP
func NewService(host string, port int) *Service {
	addr := fmt.Sprintf("%s:%d", host, port)

	return &Service{
		Launched: time.Now(),
		Server: http.Server{
			Addr:         addr,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
		},
	}
}

func (s *Service) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/one", handlers.One)
	r.HandleFunc("/two/{/name}", handlers.Two)

	// lets actually log this
	r.Handle("/three", middleware.Log(http.HandlerFunc(handlers.Three)))

	http.Handle("/", r)

	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
