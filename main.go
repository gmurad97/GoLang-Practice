package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	route := http.NewServeMux()
	route.HandleFunc("/", homeHandler)

	log.Println("Run server: http://localhost:8080")
	serverListener := http.ListenAndServe(":8080", route)
	if serverListener != nil {
		log.Fatal(serverListener)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to home page!")
}
