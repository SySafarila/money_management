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
	Date        time.Time `json:"date"`
	CategoryId  *string   `json:"category_id"`
}

type TransactionCategory struct {
	Transaction
	Category *Category `json:"category"`
}

type TransactionQueries struct {
	IsIncome   string `json:"is_income" validate:"omitempty,oneof=true false"`
	CategoryId string `json:"category_id" validate:"omitempty,uuid"`
}
