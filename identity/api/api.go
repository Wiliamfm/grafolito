package identity

import (
	"fmt"
	"net/http"
)

var basePath = "/api/identity"

func MapEndpoints() {
	http.HandleFunc(fmt.Sprintf("%s %s%s", http.MethodPost, basePath, "/login/"), loginHandler)
	http.HandleFunc(fmt.Sprintf("%s %s%s", http.MethodPost, basePath, "/logout/"), logoutHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Login"))
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
