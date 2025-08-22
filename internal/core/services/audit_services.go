package services

import "github.com/ShristiRnr/Finance/internal/core/domain/finance"

type AuditService struct {
	repo finance.AuditEventRepository
}

func NewAuditService(repo finance.AuditEventRepository) *AuditService {
	return &AuditService{repo: repo}
}

func (s *AuditService) RecordEvent(event finance.AuditEvent) (finance.AuditEvent, error) {
	// ✅ apply validation before saving
	if err := event.Validate(); err != nil {
		return finance.AuditEvent{}, err
	}
	return s.repo.Save(event)
}

func (s *AuditService) GetEventsByResource(resourceType, resourceID string) ([]finance.AuditEvent, error) {
	return s.repo.FindByResource(resourceType, resourceID)
}

func (s *AuditService) GetEventsByUser(userID string) ([]finance.AuditEvent, error) {
	return s.repo.FindByUser(userID)
}

func (s *AuditService) ListEvents() ([]finance.AuditEvent, error) {
	return s.repo.FindAll()
}
