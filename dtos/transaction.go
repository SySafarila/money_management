package dtos

type TransactionCreateDto struct {
	Amount      int64  `json:"amount" validate:"required"`
	IsIncome    *bool  `json:"is_income" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (t TransactionCreateDto) TableName() string {
	return "transactions"
}
