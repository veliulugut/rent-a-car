package jwt

import golang_jwt "github.com/golang-jwt/jwt/v5"

type Interface interface {
	Generate(id int) (string, error)
	Parse(token string) (*golang_jwt.RegisteredClaims, error)
}
