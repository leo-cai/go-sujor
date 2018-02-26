package authJWT

import (
	"github.com/dgrijalva/jwt-go"
	"time"
	"log"
)

var (
	key = []byte("sujor-api")
)

// 生成jwt
func GenerateJWT() (tokenString string, err error) {
	// token
	token := jwt.New(jwt.SigningMethodHS256)
	// Headers
	token.Header["alg"] = "HS256"
	token.Header["typ"] = "JWT"
	// Claims
	token.Claims = &jwt.StandardClaims{
		NotBefore: time.Now().Unix(),
		ExpiresAt: time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		Issuer:    "admin",
	}
	tokenString, err = token.SignedString(key)
	log.Println(tokenString)
	return
}

// 校验jwt
func CheckJWT(tokenString string) bool {
	_, err := jwt.Parse(tokenString, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Println("parse claim failed: ", err)
		return false
	}
	return true
}

// 刷新jwt
func RefreshJWT (tokenString string) bool {
	if CheckJWT(tokenString) {
		return true
	}
	return false
}
