package handlers

import (
	"log"
	"net/http"

	"github.com/mccurdyc/websocket-example/internal/connections"
)

func Chat(cp *connections.ConnectionPool) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("writing message")

		cp.WriteMessage()
	})
}
