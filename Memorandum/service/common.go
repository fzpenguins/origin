package service

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

const secretKey = "penQee"

type Claims struct {
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(id uint, userName string, password string) (myTokenString string, err error) {
	var claims Claims
	claims = Claims{
		Id:       id,
		UserName: userName,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			Issuer:    "Memorandum",
		},
	}
	myToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
	myTokenString, err = myToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return myTokenString, nil
}

func ParseToken(myTokenString string) (*Claims, error) {
	myToken, err := jwt.ParseWithClaims(myTokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if myToken != nil {
		if c, ok := myToken.Claims.(*Claims); ok && myToken.Valid {
			return c, nil
		}
	}
	return nil, err
}
