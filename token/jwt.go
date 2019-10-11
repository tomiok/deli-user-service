package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

const alg = "HS256"

var Token *JWTToken

type JWTToken struct {
	AuthJWT *jwtauth.JWTAuth
}

func Init(secret string) {
	jjwt := jwtauth.New(alg, secret, nil)

	Token = &JWTToken{
		AuthJWT: jjwt,
	}
}

func Encode() (string, error) {
	_, jsonToken, err := Token.AuthJWT.Encode(jwt.MapClaims{"user": "test"})

	return jsonToken, err
}
