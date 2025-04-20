package services

import (
	"errors"
	"grafolito/backend/internal/identity/contracts/requests"
	"grafolito/backend/internal/identity/contracts/responses"
	"grafolito/backend/internal/identity/domain/entities"
	"log"
)

func Login(request *requests.LoginRequest) (*responses.LoginResponse, error) {

	if request.Username != "admin" || request.Password != "1234" {
		log.Printf("Invalid credentials for user %s with password %s", request.Username, request.Password)
		return nil, errors.New("Invalid credentials")
	}
	user := entities.User{Username: request.Username}
	user.SetPassword(request.Password)
	token, err := CreateToken(&user)
	if err != nil {
		return nil, err
	}

	response := responses.LoginResponse{Token: token}

	return &response, nil
}
