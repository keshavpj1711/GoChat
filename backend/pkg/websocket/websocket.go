package websocket

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// defining an upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// checking the origin of our conn
	// this will allow us to make req from our React dev server to here

	// As discussed earlier for now we do no such checking and allow anyone
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}
