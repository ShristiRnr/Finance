package services

import (
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type AccrualService struct {
	repo finance.AccrualRepository
}

func NewAccrualService(repo finance.AccrualRepository) *AccrualService {
	return &AccrualService{repo: repo}
}

func (s *AccrualService) CreateAccrual(a finance.Accrual) (finance.Accrual, error) {
	// Place business rules here if any (e.g., validations, compliance checks)
	return s.repo.Save(a)
}

func (s *AccrualService) ListAccruals() ([]finance.Accrual, error) {
	return s.repo.FindAll()
}
