package utils

import (
	"fmt"
	"money_management/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

func ValidateAndParseJwt(token string) (models.CurrentUser, error) {
	if token == "" {
		return models.CurrentUser{}, ErrInvalidToken
	}
	claims := &models.CurrentUser{}

	tokenClaims, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetJWTSecretKey()), nil
	}, jwt.WithValidMethods([]string{"HS256"}))
	if err != nil {
		return models.CurrentUser{}, ErrInvalidToken
	}
	if !tokenClaims.Valid {
		return models.CurrentUser{}, ErrInvalidToken
	}
	return *claims, nil
}

func GetJWTSecretKey() string {
	return viper.GetString("jwt_secret_key")
}
