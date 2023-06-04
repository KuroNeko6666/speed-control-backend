package form

import "github.com/golang-jwt/jwt/v4"

type JwtClaim struct {
	jwt.StandardClaims
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
