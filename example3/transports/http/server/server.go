package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mccurdyc/go-examples/example3/middleware"
	"github.com/mccurdyc/go-examples/example3/transports/http/connections"
	"github.com/mccurdyc/go-examples/example3/transports/http/handlers"
)

// Service has a launched time and contains a Server.
type Service struct {
	Launched time.Time
	Server   http.Server
}

// NewService creates a new Service with the launched time set and a server configured.
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

// Start registers the necessary routes with their handler functions, then starts
// listening for requests.
func (s *Service) Start() {
	chub := connections.NewConnectionHub()
	r := mux.NewRouter()

	// use middleware for handling connections - http://www.alexedwards.net/blog/making-and-using-middleware
	r.Handle("/chat", middleware.Connect(handlers.Chat(chub), chub))
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("public")))
	http.Handle("/", r)

	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
