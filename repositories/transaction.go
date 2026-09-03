package repositories

import (
	"money_management/database"
	"money_management/models"
)

type TransactionRepository interface {
	Find(user models.CurrentUser, id string) (models.Transaction, error)
	FindAll(user models.CurrentUser) ([]models.Transaction, error)
	Create(transaction models.Transaction) (models.Transaction, error)
	Update(user models.CurrentUser, id string) (models.Transaction, error)
	Delete(user models.CurrentUser, id string) error
}

type transactionRepository struct{}

func (t transactionRepository) Find(user models.CurrentUser, id string) (models.Transaction, error) {
	var transaction models.Transaction
	err := database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).First(&transaction).Error
	return transaction, err
}

func (t transactionRepository) FindAll(user models.CurrentUser) ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := database.DB.Where("user_id = ?", user.UserId).Order("updated_at DESC").Find(&transactions).Error
	return transactions, err
}

func (t transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	err := database.DB.Create(&transaction).Error
	return transaction, err
}

func (t transactionRepository) Update(user models.CurrentUser, id string) (models.Transaction, error) {
	//TODO implement me
	panic("implement me")
}

func (t transactionRepository) Delete(user models.CurrentUser, id string) error {
	return database.DB.Where("id = ?", id).Where("user_id = ?", user.UserId).Delete(&models.Transaction{}).Error
}

func NewTransactionRepository() TransactionRepository {
	return transactionRepository{}
}
