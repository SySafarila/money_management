package models

import (
	"time"
)

type Category struct {
	Id        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
