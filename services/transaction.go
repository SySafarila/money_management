package services

import (
	"money_management/dtos"
	"money_management/models"
	"money_management/repositories"
	"time"
)

type TransactionService interface {
	GetAll(user models.CurrentUser) (map[string][]models.Transaction, error)
	GetDetail(user models.CurrentUser, id string) (models.Transaction, error)
	NewTransaction(user models.CurrentUser, data dtos.TransactionCreateDto) (models.Transaction, error)
	UpdateTransaction(user models.CurrentUser, id string, data dtos.TransactionCreateDto) (models.Transaction, error)
	DeleteTransaction(user models.CurrentUser, id string) error
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
}

func (t transactionService) UpdateTransaction(user models.CurrentUser, id string, data dtos.TransactionCreateDto) (models.Transaction, error) {
	transaction, errDetailTransaction := t.GetDetail(user, id)
	if errDetailTransaction != nil {
		return models.Transaction{}, errDetailTransaction
	}
	date, errDate := time.Parse("2006-01-02T15:04:05Z07:00", data.Date)
	if errDate != nil {
		return models.Transaction{}, errDate
	}
	transaction.Amount = data.Amount
	transaction.IsIncome = *data.IsIncome
	transaction.Description = data.Description
	transaction.Date = date
	result, err := t.transactionRepository.Update(user, transaction.Id, transaction)
	return result, err
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
	transaction, err := t.GetDetail(user, id)
	if err != nil {
		return err
	}
	return t.transactionRepository.Delete(user, transaction.Id)
}

func (t transactionService) NewTransaction(user models.CurrentUser, data dtos.TransactionCreateDto) (models.Transaction, error) {
	date, errDate := time.Parse("2006-01-02T15:04:05Z07:00", data.Date)
	if errDate != nil {
		return models.Transaction{}, errDate
	}
	transaction := models.Transaction{
		UserId:      user.UserId,
		Amount:      data.Amount,
		IsIncome:    *data.IsIncome,
		Description: data.Description,
		Date:        date,
	}
	result, err := t.transactionRepository.Create(transaction)
	return result, err
}

func NewTransactionService(repo repositories.TransactionRepository) TransactionService {
	return &transactionService{transactionRepository: repo}
}
