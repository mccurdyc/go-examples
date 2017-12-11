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

type ConnectionHub struct {
	clients map[*websocket.Conn]bool

	// history of previous messages for new clients
	history []Message

	// channel for previous message history
	stream chan Message

	// bi-directional channel
	broadcast chan Message
}

func NewConnectionHub() *ConnectionHub {
	var h = []Message{}
	var s = make(chan Message)

	go func() {
		for {
			m := <-s
			h = append(h, m)
		}
	}()

	return &ConnectionHub{
		clients:   make(map[*websocket.Conn]bool),
		history:   h,
		broadcast: make(chan Message),
		stream:    s,
	}
}

func (chub *ConnectionHub) add(c *websocket.Conn) {
	chub.clients[c] = true
	return
}

// HandleConnection handles upgrading an http
// connection to a websocket connection.
func (chub *ConnectionHub) HandleConnection(w http.ResponseWriter, r *http.Request) {
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
	chub.add(conn)
	// handle reading new messages in its own thread
	go chub.ReadMessages(conn)

	return
}

// ReadMessages runs indefinitely reading in and parsing
// messages and sending them to the connection hub's
// broadcast channel so that other connection can access them.
func (chub *ConnectionHub) ReadMessages(c *websocket.Conn) {
	for {
		var msg Message

		if err := c.ReadJSON(&msg); err != nil {
			log.Println(errors.Wrap(err, "error reading/parsing message from websocket"))
			delete(chub.clients, c)
			return
		}

		fmt.Printf("reading message: %+v\n", msg)
		chub.broadcast <- msg
		chub.stream <- msg
	}
}

// WriteMessage handles writing a message received
// in the broadcast channel to all of the clients
func (chub *ConnectionHub) WriteMessage() {
	msg := <-chub.broadcast
	fmt.Printf("writing message: %+v\n", msg)

	out, err := json.Marshal(msg)
	if err != nil {
		log.Println(errors.Wrap(err, "error encoding json"))
		return
	}

	for c := range chub.clients {
		if err := c.WriteMessage(websocket.TextMessage, out); err != nil {
			delete(chub.clients, c)
			log.Println(errors.Wrap(err, "error writing message to all clients"))
			return
		}
	}
}
