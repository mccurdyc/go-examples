package handlers

import (
	"io"
	"net/http"
)

// One is a handler function that will do whatever
// you tell it to do. It is restricted in that this
// function must only have http.ResponseWriter and *http.Request
// as its parameters.
func One(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello, world!")
}
