package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Found!"))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
