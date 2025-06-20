package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)


// defining an upgrader
var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,

	// checking the origin of our conn
	// this will allow us to make req from our React dev server to here

	// As discussed earlier for now we do no such checking and allow anyone
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}


// A reader to listen for new messages being sent to our websocket endpoint
func reader(conn *websocket.Conn)  {
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

// defining our websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade conn to websocket conn
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	// listen for new msgs coming indefinitely thru our websocket conn
	reader(ws)
}

func setupRoutes()  {
	// 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})

	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Starting Server...")
	setupRoutes()
	fmt.Println("Output at: ", "http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}