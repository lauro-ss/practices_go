package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	jwt.StandardClaims
}

func NewAccessToken(user *User) (string, error) {
	claims := UserClaims{}
	claims.Name = user.Name
	claims.Id = user.Id
	claims.ExpiresAt = time.Now().Add(time.Hour).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken() (string, error) {
	claims := jwt.StandardClaims{}
	claims.ExpiresAt = time.Now().Add(time.Hour * 24 * 7).Unix()

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) (*UserClaims, error) {
	token, err := ParseToken(accessToken, &UserClaims{})
	if err != nil {
		return nil, err
	}

	return token.(*UserClaims), nil
}

func ParseRefreshToken(refreshToken string) (*jwt.StandardClaims, error) {
	token, err := ParseToken(refreshToken, &jwt.StandardClaims{})
	if err != nil {
		return nil, err
	}
	return token.(*jwt.StandardClaims), nil
}

func ParseToken(stringToken string, claims jwt.Claims) (any, error) {
	token, err := jwt.ParseWithClaims(stringToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt method")
		}
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims, nil
}
