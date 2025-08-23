package repository

import (
	"errors"
	"sync"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type InMemoryConsolidationRepo struct {
	mu      sync.RWMutex
	reports map[string]*finance.ConsolidatedReport
}

func NewInMemoryConsolidationRepo() *InMemoryConsolidationRepo {
	return &InMemoryConsolidationRepo{reports: make(map[string]*finance.ConsolidatedReport)}
}

func (r *InMemoryConsolidationRepo) SaveReport(report *finance.ConsolidatedReport) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.reports[report.ID] = report
	return nil
}

func (r *InMemoryConsolidationRepo) FindByID(id string) (*finance.ConsolidatedReport, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if rep, ok := r.reports[id]; ok {
		return rep, nil
	}
	return nil, errors.New("report not found")
}

func (r *InMemoryConsolidationRepo) FindAll() ([]*finance.ConsolidatedReport, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*finance.ConsolidatedReport, 0, len(r.reports))
	for _, rep := range r.reports {
		list = append(list, rep)
	}
	return list, nil
}
