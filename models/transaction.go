package models

import (
	"time"
)

type Transaction struct {
	Id          string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserId      string    `json:"user_id"`
	Amount      int64     `json:"amount"`
	IsIncome    bool      `json:"is_income"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
