package main

import (
	"fmt"
	"grafolito/backend/internal/identity/api"
	"log"
	"net/http"
)

func main() {

	port := 9000

	api.MapEndpoints()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Not Found!"))
	})

	log.Printf("Listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprint("localhost:", port), nil))
}
