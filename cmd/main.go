package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {

	h := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		by, _ := json.Marshal("Hello World")

		w.Write(by)
	}

	http.HandleFunc("/", h)

	log.Println("starting web server at http://localhost:1213")
	if err := http.ListenAndServe(":1213", nil); err != nil {
		log.Fatal("error when run the server")
	}
}
