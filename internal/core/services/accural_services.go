package services

import "github.com/ShristiRnr/Finance/internal/core/domain/finance"

type AccrualService struct {
	repo finance.AccrualRepository
}

func NewAccrualService(repo finance.AccrualRepository) *AccrualService {
	return &AccrualService{repo: repo}
}

func (s *AccrualService) CreateAccrual(a finance.Accrual) (finance.Accrual, error) {
	return s.repo.Save(a)
}

func (s *AccrualService) ListAccruals() ([]finance.Accrual, error) {
	return s.repo.FindAll()
}

func (s *AccrualService) GetAccrualByID(id string) (finance.Accrual, error) {
	return s.repo.FindByID(id)
}

func (s *AccrualService) UpdateAccrual(a finance.Accrual) (finance.Accrual, error) {
	return s.repo.Update(a)
}

func (s *AccrualService) DeleteAccrualByID(id string) (finance.Accrual, error) {
	return s.repo.DeleteByID(id)
}
