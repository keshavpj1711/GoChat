package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keshavpj1711/GoChat/pkg/websocket"
)


// defining our websocket endpoint
func serveWs(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Host)

	// upgrade conn to websocket conn
	ws, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}
	// listen for new msgs coming indefinitely thru our websocket conn
	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	// Only handling the websocket connection route
	http.HandleFunc("/ws", serveWs)
}

func main() {
	fmt.Println("Starting Server...")
	setupRoutes()
	fmt.Println("Output at: ", "http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
