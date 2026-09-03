package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	Id         string     `json:"id"`
	Username   string     `json:"username"`
	FullName   string     `json:"full_name"`
	Password   string     `json:"-"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	VerifiedAt *time.Time `json:"verified_at"`
}

type UserQueryParams struct {
	Id       string `query:"id"`
	Username string `query:"username"`
	FullName string `query:"full_name"`
	Verified string `query:"verified_at"`
}

type CurrentUser struct {
	UserId string `json:"user_id"`
	jwt.RegisteredClaims
}
