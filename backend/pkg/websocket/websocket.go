package websocket

import (
	"fmt"
	"io"
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

// A reader to listen for new messages being sent to our websocket endpoint
func Reader(conn *websocket.Conn) {
	for {
		// read in a message
		msgType, p, err := conn.ReadMessage()
		// checking for any errors
		if err != nil {
			log.Println(err)
			return
		}

		// print out the message
		fmt.Println(string(p))

		if err := conn.WriteMessage(msgType, p); err != nil {
			log.Println(err)
			return
		}
	}
}


func Writer(conn *websocket.Conn)  {
	for {
		fmt.Println("Sending")
		messageType, r, err := conn.NextReader()
		if err != nil {
			fmt.Println(err)
			return
		}

		w, err := conn.NextWriter(messageType)
		if err != nil {
			fmt.Println(err)
			return
		}

		if _, err := io.Copy(w, r); err != nil {
			fmt.Println(err)
			return
		}
		
		if err := w.Close(); err != nil {
			fmt.Println(err)
			return
		}
	}
}
