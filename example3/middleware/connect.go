package middleware

import (
	"log"
	"net/http"

	"github.com/mccurdyc/go-examples/example3/transports/http/connections"
)

// Connect handles upgrading and adding the connection to the connection pool
func Connect(next http.Handler, cp *connections.ConnectionPool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("handling connection in middlware")
		cp.HandleConnection(w, r)
		log.Println("finished handling connection in middlware")
		next.ServeHTTP(w, r)
	})
}
