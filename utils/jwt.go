package utils

import (
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

)

func CreateToken(Phone string) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	tokenexp, _ := strconv.Atoi(GetConfigValue("tokenexp"))
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["phone"] = Phone
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(GetConfigValue("token_secrets")))
	return tokenString
}

func CheckToken(tokenString string) string {
	Phone := ""
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(GetConfigValue("token_secrets")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	Phone = claims["phone"].(string)
	return Phone
}
