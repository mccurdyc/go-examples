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
	history *[]Message

	// channel for previous message history
	stream chan Message

	// bi-directional channel
	broadcast chan Message
}

func NewConnectionHub() *ConnectionHub {
	var h = []Message{}
	var s = make(chan Message)

	// goroutine for adding messages to history
	go func() {
		fmt.Println("starting goroutine for listening for history")
		for {
			m := <-s
			fmt.Printf("received message for history: %+v\n", m)
			h = append(h, m)
			fmt.Printf("history now: %+v\n", h)
		}
	}()

	return &ConnectionHub{
		clients:   make(map[*websocket.Conn]bool),
		history:   &h,
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

	chub.WriteHistory(conn)
	fmt.Println("Wrote history")

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
		fmt.Printf("adding message to stream: %+v\n", msg)
		chub.stream <- msg
	}
}

// WriteMessage handles writing a message received
// in the broadcast channel to all of the clients
func (chub *ConnectionHub) WriteMessage() {
	for {
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
}

func (chub *ConnectionHub) WriteHistory(c *websocket.Conn) {
	fmt.Printf("All history: %+v\n", chub.history)
	for _, m := range *chub.history {
		fmt.Printf("History: %+v\n", m)

		out, err := json.Marshal(m)
		if err != nil {
			log.Println(errors.Wrap(err, "error encoding json"))
			return
		}

		if err := c.WriteMessage(websocket.TextMessage, out); err != nil {
			delete(chub.clients, c)
			log.Println(errors.Wrap(err, "error writing message to client"))
			return
		}
	}
}
