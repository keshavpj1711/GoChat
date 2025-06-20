package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/keshavpj1711/GoChat/pkg/websocket"
)


// defining our websocket endpoint
func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")

	// upgrade conn to websocket conn
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(w, err)
	}
	
	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	// Only handling the websocket connection route
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

func main() {
	fmt.Println("Starting Server...")
	setupRoutes()
	fmt.Println("Output at: ", "http://localhost:3000")
	http.ListenAndServe(":3000", nil)
}
