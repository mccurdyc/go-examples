package connections

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

type ConnectionPool struct {
	clients map[*websocket.Conn]bool

	// bi-directional channel
	broadcast chan Message
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		clients:   make(map[*websocket.Conn]bool),
		broadcast: make(chan Message, 3),
	}
}

func (cp *ConnectionPool) add(c *websocket.Conn) {
	cp.clients[c] = true
	return
}

// HandleConnection handles upgrading an http
// connection to a websocket connection.
func (cp *ConnectionPool) HandleConnection(w http.ResponseWriter, r *http.Request) {
	var u = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	// upgrade http connection to websocket connection
	conn, err := u.Upgrade(w, r, nil)

	if err != nil {
		log.Println(errors.Wrap(err, "error upgrading to websocket connection"))
		return
	}

	// add the websocket connection to the pool of connections
	cp.add(conn)
	// handle reading new messages in its own thread
	go cp.ReadMessages(conn)

	return
}

// ReadMessages runs indefinitely reading in and parsing
// messages and sending them to the connection pool's
// broadcast channel so that other connection can access them.
func (cp *ConnectionPool) ReadMessages(c *websocket.Conn) {
	for {
		var msg Message

		if err := c.ReadJSON(&msg); err != nil {
			log.Println(errors.Wrap(err, "error reading/parsing message from websocket"))
			delete(cp.clients, c)
			return
		}

		fmt.Printf("reading message: %+v\n", msg)
		cp.broadcast <- msg
	}
}

// WriteMessage handles writing a message received
// in the broadcast channel to all of the clients
func (cp *ConnectionPool) WriteMessage() {
	msg := <-cp.broadcast
	fmt.Printf("writing message: %+v\n", msg)

	out, err := json.Marshal(msg)

	if err != nil {
		log.Println(errors.Wrap(err, "error encoding json"))
		return
	}

	for c := range cp.clients {
		if err := c.WriteMessage(websocket.TextMessage, out); err != nil {
			delete(cp.clients, c)
			log.Println(errors.Wrap(err, "error writing message to all clients"))
			return
		}
	}
}
