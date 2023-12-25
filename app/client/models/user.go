package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `json:"id"`
	Email          string    `json:"email"`
	Name           *string   `json:"name"`
	HashedPassword string    `json:"-"`
}

type RegisterRequest struct {
	Email    string  `json:"email"`
	Name     *string `json:"name"`
	Password string  `json:"password"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
