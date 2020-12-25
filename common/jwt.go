package common

import (
	jwtGo "github.com/dgrijalva/jwt-go"
	"gopkg.in/errgo.v2/fmt/errors"
)

func ParseToken(token, secret string) (*jwtGo.Token, error) {
	return jwtGo.Parse(token, func(token *jwtGo.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwtGo.SigningMethodHMAC); !ok {
			return nil, errors.Newf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
}

func ParseTokenClaims(tokenStr, secret string) (jwtGo.MapClaims, error) {
	token, err := ParseToken(tokenStr, secret)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwtGo.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}