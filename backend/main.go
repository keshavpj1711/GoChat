package main

import (
	"fmt"
	"net/http"
)

func setupRoutes()  {
	// 
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Simple Server")
	})
}

func main() {
	setupRoutes()
	fmt.Println("Starting Server...")
	fmt.Println("Output at: ", "http://localhost:8080")
	http.ListenAndServe(":8080",nil)
}