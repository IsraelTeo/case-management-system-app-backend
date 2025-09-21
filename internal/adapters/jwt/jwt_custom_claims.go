package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// Claims define la estructura personalizada de los datos (claims) que se almacenan dentro del token JWT.
type Claims struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email"`
	Role   string    `json:"role"`
	jwt.StandardClaims
}
