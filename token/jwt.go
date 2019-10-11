package token

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
	"time"
)

const alg = "HS256"

var Token *JWTToken

type JWTToken struct {
	AuthJWT *jwtauth.JWTAuth
}

func Init(secret string) {
	jjwt := jwtauth.New(alg, []byte(secret), nil)
	Token = &JWTToken{
		AuthJWT: jjwt,
	}
}

func Encode(userUid string) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userUid
	jwtauth.SetIssuedAt(claims, time.Now())
	jwtauth.SetExpiry(claims, time.Now().Add(time.Hour*24*10))
	_, jsonToken, err := Token.AuthJWT.Encode(claims)

	return jsonToken, err
}
