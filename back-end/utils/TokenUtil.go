package utils

import (
	"back-end/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
)

func ExtractTokenClaimsFromCookie(cookie http.Cookie) (jwt.MapClaims, error) {
	token, err := jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, errors.New("unauthorized")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("something went wrong in parsing token")
}

func GenerateToken(userData model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["usr"] = userData.Username
	tokenSigned, err := token.SignedString([]byte("mh_secret"))
	return tokenSigned, err
}
