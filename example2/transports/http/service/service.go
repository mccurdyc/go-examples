package service

/*
* @author Colton J. McCurdy
*	GitHub: mccurdyc
* Email:  mccurdyc22@gmail.com
* Date: 2018-02-15
 */

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/mccurdyc/go-examples/example2/transports/http/handlers"
	"github.com/mccurdyc/go-examples/example2/transports/http/middleware"
)

// Service will host the server and the time when it started.
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

// Start will register the routes with the appropriate handler functions and
// serve on the port specified by the PORT environment variable and will default
// to 8080 if not set.
func (s *Service) Start() {
	r := mux.NewRouter()

	r.HandleFunc("/one", handlers.One)
	r.HandleFunc("/two/{name}", handlers.Two)

	// lets actually log this
	r.Handle("/three", middleware.Log(http.HandlerFunc(handlers.Three)))
	r.Handle("/four", middleware.Log(http.HandlerFunc(handlers.Four)))

	http.Handle("/", r)

	if err := s.Server.ListenAndServe(); err != nil {
		panic(err)
	}
}
