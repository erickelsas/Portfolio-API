package authorization

import (
	"errors"
	"fmt"
	"net/http"
	"portfolio-api/src/config"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func TokenCreate(userId uint64) (string, error) {
	perm := jwt.MapClaims{}
	perm["authorized"] = true
	perm["exp"] = time.Now().Add(time.Hour * 2190).Unix()
	perm["userId"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perm)

	return token.SignedString(config.SecretKey)
}

func TokenValidate(r *http.Request) error {
	tokenString := getToken(r)
	token, erro := jwt.Parse(tokenString, getKey)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("Token inválido")
}

func getToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}

	return ""
}

func getKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}

func GetUserId(r *http.Request) (uint64, error) {
	tokenString := getToken(r)
	token, erro := jwt.Parse(tokenString, getKey)
	if erro != nil {
		return 0, erro
	}

	if perm, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userId, erro := strconv.ParseUint(fmt.Sprintf("%.0f", perm["userId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}

		return userId, nil
	}

	return 0, errors.New("Token inválido")
}
