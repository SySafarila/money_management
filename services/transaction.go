package services

import (
	"money_management/models"
	"money_management/repositories"
)

type TransactionService interface {
	GetAll(user models.CurrentUser) (map[string][]models.Transaction, error)
	GetDetail(user models.CurrentUser, id string) (models.Transaction, error)
	AddIncome(user models.CurrentUser, amount int64, desc string) (models.Transaction, error)
	AddExpense(user models.CurrentUser, amount int64, desc string) (models.Transaction, error)
	DeleteTransaction(user models.CurrentUser, id string) error
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
}

func (t transactionService) GetDetail(user models.CurrentUser, id string) (models.Transaction, error) {
	return t.transactionRepository.Find(user, id)
}

func (t transactionService) GetAll(user models.CurrentUser) (map[string][]models.Transaction, error) {
	transactions, err := t.transactionRepository.FindAll(user)
	if err != nil {
		return nil, err
	}
	groupByDate := make(map[string][]models.Transaction)
	for _, transaction := range transactions {
		date := transaction.UpdatedAt.Format("2006-01-02")
		groupByDate[date] = append(groupByDate[date], transaction)
	}
	return groupByDate, err
}

func (t transactionService) DeleteTransaction(user models.CurrentUser, id string) error {
	return t.transactionRepository.Delete(user, id)
}

func (t transactionService) AddIncome(user models.CurrentUser, amount int64, desc string) (models.Transaction, error) {
	transaction := models.Transaction{
		UserId:      user.UserId,
		Amount:      amount,
		IsIncome:    true,
		Description: desc,
	}
	result, err := t.transactionRepository.Create(transaction)
	if err != nil {
		return transaction, err
	}
	return result, nil
}

func (t transactionService) AddExpense(user models.CurrentUser, amount int64, desc string) (models.Transaction, error) {
	transaction := models.Transaction{
		UserId:      user.UserId,
		Amount:      amount,
		IsIncome:    false,
		Description: desc,
	}
	result, err := t.transactionRepository.Create(transaction)
	if err != nil {
		return transaction, err
	}
	return result, nil
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: repo}
}
