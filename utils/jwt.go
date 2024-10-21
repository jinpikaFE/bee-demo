package utils

import (
	"bee-demo/models"
	"errors"
	"strconv"
	"time"

	"github.com/beego/beego/v2/core/logs"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// MyCustomClaims
// This struct is the payload
// 此结构是有效负载
type MyCustomClaims struct {
	UserID int `json:"userID"`
	jwt.StandardClaims
}

// JwtPayload
// This struct is the parsing of token payload
// 此结构是对token有效负载的解析
type JwtPayload struct {
	UserName  string `json:"userName"`
	UserID    int    `json:"userID"`
	IssuedAt  int64  `json:"iat"` // 发布日期
	ExpiresAt int64  `json:"exp"` // 过期时间
}

func CreateToken(user *models.User) string {
	tokenexp, _ := strconv.Atoi(GetConfigValue("tokenexp"))
	expireAt := time.Now().Add(time.Second * time.Duration(tokenexp)).Unix()
	logs.Info("Token 将到期于：", time.Unix(expireAt, 0))

	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)

	claims["exp"] = expireAt
	claims["iat"] = time.Now().Unix()
	claims["userID"] = user.Id
	claims["userName"] = user.UserName
	token.Claims = claims
	tokenString, _ := token.SignedString([]byte(GetConfigValue("token_secrets")))
	return tokenString
}

func CheckToken(tokenString string) (*JwtPayload, error) {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("错误: token验证失败")
		}
		return []byte(GetConfigValue("token_secrets")), nil
	})
	claims, _ := token.Claims.(jwt.MapClaims)
	return &JwtPayload{
		UserName:  claims["userName"].(string),     // 用户名：发行者
		UserID:    int(claims["userID"].(float64)), // 转换为int
		IssuedAt:  int64(claims["iat"].(float64)),  // 转换为int64
		ExpiresAt: int64(claims["exp"].(float64)),  // 转换为int64
	}, nil
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
