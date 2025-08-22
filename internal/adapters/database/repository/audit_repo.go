package repository

import (
	"sync"

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

func (r *InMemoryAuditRepo) FindAll() ([]finance.AuditEvent, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]finance.AuditEvent, 0, len(r.events))
	for _, v := range r.events {
		list = append(list, v)
	}
	return list, nil
}
