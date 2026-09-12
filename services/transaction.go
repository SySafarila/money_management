package services

import (
	"money_management/dtos"
	"money_management/models"
	"money_management/repositories"
	"time"
)

type TransactionService interface {
	GetAll(user models.CurrentUser, q models.TransactionQueries) (map[string][]models.TransactionCategory, error)
	GetDetail(user models.CurrentUser, id string) (models.TransactionCategory, error)
	NewTransaction(user models.CurrentUser, data dtos.TransactionCreateDto) (models.Transaction, error)
	UpdateTransaction(user models.CurrentUser, id string, data dtos.TransactionCreateDto) (models.Transaction, error)
	DeleteTransaction(user models.CurrentUser, id string) error
}

type transactionService struct {
	transactionRepository repositories.TransactionRepository
	categoryService       CategoryService
}

func (t transactionService) UpdateTransaction(user models.CurrentUser, id string, data dtos.TransactionCreateDto) (models.Transaction, error) {
	if data.CategoryId != nil {
		_, err := t.categoryService.GetDetail(user, *data.CategoryId)
		if err != nil {
			return models.Transaction{}, err
		}
	}
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
	transaction.CategoryId = data.CategoryId
	result, err := t.transactionRepository.Update(user, transaction.Id, transaction.Transaction)
	return result, err
}

func (t transactionService) GetDetail(user models.CurrentUser, id string) (models.TransactionCategory, error) {
	transaction, err := t.transactionRepository.Find(user, id)
	if err != nil {
		return models.TransactionCategory{}, err
	}
	if transaction.CategoryId != nil {
		category, err := t.categoryService.GetDetail(user, *transaction.CategoryId)
		if err != nil {
			return models.TransactionCategory{}, err
		}
		return models.TransactionCategory{
			Transaction: transaction,
			Category:    &category,
		}, nil
	}
	return models.TransactionCategory{
		Transaction: transaction,
		Category:    nil,
	}, nil
}

func (t transactionService) GetAll(user models.CurrentUser, q models.TransactionQueries) (map[string][]models.TransactionCategory, error) {
	transactions, err := t.transactionRepository.FindAll(user, q)
	if err != nil {
		return nil, err
	}
	categories, _ := t.categoryService.GetAll(user)
	categoriesMap := make(map[string]*models.Category)
	for _, category := range categories {
		categoriesMap[category.Id] = &category
	}

	groupByDate := make(map[string][]models.TransactionCategory)
	for _, transaction := range transactions {
		date := transaction.UpdatedAt.Format("2006-01-02")
		if transaction.CategoryId != nil {
			groupByDate[date] = append(groupByDate[date], models.TransactionCategory{
				Transaction: transaction,
				Category:    categoriesMap[*transaction.CategoryId],
			})
		} else {
			groupByDate[date] = append(groupByDate[date], models.TransactionCategory{
				Transaction: transaction,
				Category:    nil,
			})

		}
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
	if data.CategoryId != nil {
		_, err := t.categoryService.GetDetail(user, *data.CategoryId)
		if err != nil {
			return models.Transaction{}, err
		}
	}
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
		CategoryId:  data.CategoryId,
	}
	result, err := t.transactionRepository.Create(transaction)
	return result, err
}

func NewTransactionService(repo repositories.TransactionRepository, categoryService CategoryService) TransactionService {
	return &transactionService{transactionRepository: repo, categoryService: categoryService}
}
