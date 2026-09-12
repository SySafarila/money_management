package dtos

type TransactionCreateDto struct {
	Amount      int64   `json:"amount" validate:"required,number"`
	IsIncome    *bool   `json:"is_income" validate:"required,boolean"`
	Description string  `json:"description" validate:"required"`
	Date        string  `json:"date" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	CategoryId  *string `json:"category_id"`
}

func (t TransactionCreateDto) TableName() string {
	return "transactions"
}
