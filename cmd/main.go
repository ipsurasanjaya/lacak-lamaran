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

	log.Println("starting web server at http://localhost:1213")
	if err := http.ListenAndServe(":1213", nil); err != nil {
		log.Fatal("error when run the server")
	}
}
