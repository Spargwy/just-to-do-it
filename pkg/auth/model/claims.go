package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type CustomClaims struct {
	UserID uuid.UUID `json:"user_id"`
	jwt.StandardClaims
}
