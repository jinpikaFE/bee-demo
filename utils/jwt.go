package utils

import (
	"bee-demo/models"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(user *models.User) string {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	tokenexp, _ := strconv.Atoi(GetConfigValue("tokenexp"))
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenexp)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = user.Id
	claims["userName"] = user.UserName
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

// CheckPasswordHash 检查输入的密码与哈希值是否匹配
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
