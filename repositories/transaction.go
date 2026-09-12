package repositories

import (
	"money_management/database"
	"money_management/models"
)

type TransactionRepository interface {
	Find(user models.CurrentUser, id string) (models.Transaction, error)
	FindAll(user models.CurrentUser, q models.TransactionQueries) ([]models.Transaction, error)
	Create(transaction models.Transaction) (models.Transaction, error)
	Update(user models.CurrentUser, id string, transaction models.Transaction) (models.Transaction, error)
	Delete(user models.CurrentUser, id string) error
}

type transactionRepository struct{}

func (t transactionRepository) Find(user models.CurrentUser, id string) (models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).First(&transaction).Error
	return transaction, err
}

func (t transactionRepository) FindAll(user models.CurrentUser, q models.TransactionQueries) ([]models.Transaction, error) {
	var transactions []models.Transaction
	query := database.DB.Where("user_id = ?", user.UserId).Order("updated_at DESC")

	if q.IsIncome == "true" {
		query = query.Where("is_income = ?", true)
	} else if q.IsIncome == "false" {
		query = query.Where("is_income = ?", false)
	}

	if q.CategoryId != "" {
		query = query.Where("category_id = ?", q.CategoryId)
	}

	err := query.Find(&transactions).Error
	return transactions, err
}

func (t transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	err := database.DB.Create(&transaction).Error
	return transaction, err
}

func (t transactionRepository) Update(user models.CurrentUser, id string, transaction models.Transaction) (models.Transaction, error) {
	err := database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).Updates(&transaction).Error
	return transaction, err
}

func (t transactionRepository) Delete(user models.CurrentUser, id string) error {
	return database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).Delete(&models.Transaction{}).Error
}

func NewTransactionRepository() TransactionRepository {
	return transactionRepository{}
}
