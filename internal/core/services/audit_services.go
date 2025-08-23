package services

import (
	"time"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type AuditService struct {
	repo finance.AuditEventRepository
}

func NewAuditService(repo finance.AuditEventRepository) *AuditService {
	return &AuditService{repo: repo}
}

func (s *AuditService) RecordEvent(event finance.AuditEvent) (finance.AuditEvent, error) {
	if err := event.Validate(); err != nil {
		return finance.AuditEvent{}, err
	}
	return s.repo.Save(event)
}

func (s *AuditService) GetEventByID(id string) (finance.AuditEvent, error) {
	return s.repo.FindByID(id)
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

func (s *AuditService) FilterEvents(userID, action, resourceType, resourceID string, from, to *time.Time) ([]finance.AuditEvent, error) {
	return s.repo.FilterEvents(userID, action, resourceType, resourceID, from, to)
}
