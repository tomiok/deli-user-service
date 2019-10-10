package token

import "github.com/go-chi/jwtauth"

const alg  = "HS256"

func Init(secret string)  {
	jwtauth.New(alg, secret, nil)
}