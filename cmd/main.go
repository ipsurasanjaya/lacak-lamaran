package main

import (
	"log"
	"net/http"
)

func main() {

	h := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		w.Write([]byte("Hello world"))
	}

	http.HandleFunc("/", h)

	log.Println("starting web server at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("error when run the server")
	}
}
