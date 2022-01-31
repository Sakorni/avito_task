package service

import (
	"avito_task/models"
	"avito_task/repository"
)

type BalanceService struct {
	repo *repository.BalanceRepository
}

func NewBalanceService(repo *repository.BalanceRepository) *BalanceService {
	return &BalanceService{
		repo,
	}
}

// ChangeAmount Used for withdraw and deposit operations
func (b *BalanceService) ChangeAmount(id uint, amount int) (*models.BalanceInfo, error) {
	return b.repo.ChangeAmount(id, amount)
}

func (b *BalanceService) Transfer(transaction *models.Transaction) (error error) {
	return b.repo.Transfer(transaction)
}

func (b *BalanceService) GetBalance(userId uint) (*models.BalanceInfo, error) {
	return b.repo.GetBalance(userId)
}
