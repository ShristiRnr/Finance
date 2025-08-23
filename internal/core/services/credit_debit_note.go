package services

import "github.com/ShristiRnr/Finance/internal/core/domain/finance"

type CreditDebitNoteService interface {
	Create(note finance.CreditDebitNote) (finance.CreditDebitNote, error)
	GetByID(id string) (finance.CreditDebitNote, error)
	List(offset, limit int) ([]finance.CreditDebitNote, error)
	Update(note finance.CreditDebitNote) (finance.CreditDebitNote, error)
	Delete(id string) error
}

type creditDebitNoteService struct {
	repo finance.CreditDebitNoteRepository
}

func NewCreditDebitNoteService(repo finance.CreditDebitNoteRepository) CreditDebitNoteService {
	return &creditDebitNoteService{repo: repo}
}

func (s *creditDebitNoteService) Create(note finance.CreditDebitNote) (finance.CreditDebitNote, error) {
	return s.repo.Save(note)
}

func (s *creditDebitNoteService) GetByID(id string) (finance.CreditDebitNote, error) {
	return s.repo.FindByID(id)
}

func (s *creditDebitNoteService) List(offset, limit int) ([]finance.CreditDebitNote, error) {
	return s.repo.FindAll(offset, limit)
}

func (s *creditDebitNoteService) Update(note finance.CreditDebitNote) (finance.CreditDebitNote, error) {
	return s.repo.Update(note)
}

func (s *creditDebitNoteService) Delete(id string) error {
	return s.repo.Delete(id)
}
