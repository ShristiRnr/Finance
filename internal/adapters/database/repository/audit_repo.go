package repository

import (
	"sync"
	"time"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type InMemoryAuditRepo struct {
	mu     sync.RWMutex
	events []finance.AuditEvent
}

func NewInMemoryAuditRepo() *InMemoryAuditRepo {
	return &InMemoryAuditRepo{}
}

func (r *InMemoryAuditRepo) Save(e finance.AuditEvent) (finance.AuditEvent, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events = append(r.events, e)
	return e, nil
}

func (r *InMemoryAuditRepo) FindAll() ([]finance.AuditEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]finance.AuditEvent, len(r.events))
	copy(list, r.events)
	return list, nil
}

func (r *InMemoryAuditRepo) FindByID(id string) (finance.AuditEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	for _, e := range r.events {
		if e.ID == id {
			return e, nil
		}
	}
	return finance.AuditEvent{}, nil
}

func (r *InMemoryAuditRepo) FindByResource(resourceType, resourceID string) ([]finance.AuditEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []finance.AuditEvent
	for _, e := range r.events {
		if e.ResourceType == resourceType && e.ResourceID == resourceID {
			result = append(result, e)
		}
	}
	return result, nil
}

func (r *InMemoryAuditRepo) FindByUser(userID string) ([]finance.AuditEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []finance.AuditEvent
	for _, e := range r.events {
		if e.UserID == userID {
			result = append(result, e)
		}
	}
	return result, nil
}

func (r *InMemoryAuditRepo) FilterEvents(userID, action, resourceType, resourceID string, from, to *time.Time) ([]finance.AuditEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var result []finance.AuditEvent
	for _, e := range r.events {
		if userID != "" && e.UserID != userID {
			continue
		}
		if action != "" && e.Action != action {
			continue
		}
		if resourceType != "" && e.ResourceType != resourceType {
			continue
		}
		if resourceID != "" && e.ResourceID != resourceID {
			continue
		}
		if from != nil && e.Timestamp.Before(*from) {
			continue
		}
		if to != nil && e.Timestamp.After(*to) {
			continue
		}
		result = append(result, e)
	}
	return result, nil
}
