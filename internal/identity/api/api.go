package api

import (
	"encoding/json"
	"fmt"
	"grafolito/backend/internal/identity/application/services"
	"grafolito/backend/internal/identity/contracts/requests"
	"grafolito/backend/internal/identity/contracts/responses"
	"grafolito/backend/internal/shared/api/constants"
	"net/http"
)

var basePath = "/api/identity"

func MapEndpoints() {
	http.HandleFunc(fmt.Sprintf("%s %s%s", http.MethodPost, basePath, "/login"), loginHandler)
	http.HandleFunc(fmt.Sprintf("%s %s%s", http.MethodPost, basePath, "/logout"), logoutHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var decoder = json.NewDecoder(r.Body)
	var request requests.LoginRequest
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var response []byte
	var loginResponse any
	loginResponse, err = services.Login(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		loginResponse = responses.ProblemDetails{
			Type:     constants.ProblemDetailBadRequestType,
			Title:    "Login Error",
			Detail:   err.Error(),
			Instance: r.URL.String(),
		}
	}
	response, err = json.Marshal(loginResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(response)
	return
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Logout"))
}
