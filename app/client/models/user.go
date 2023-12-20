package models

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type CreateUserRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	Token string `json:"token"`
}
