package main

import (
	"grafolito/backend/internal/identity/api"
	"log"
	"net/http"
)

func main() {

	api.MapEndpoints()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Found!"))
	})

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
