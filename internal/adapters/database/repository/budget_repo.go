package repository

import (
	"sync"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type InMemoryBudgetRepo struct {
	mu      sync.RWMutex
	budgets map[string]*finance.Budget
}

func NewInMemoryBudgetRepo() *InMemoryBudgetRepo {
	return &InMemoryBudgetRepo{
		budgets: make(map[string]*finance.Budget),
	}
}

func (r *InMemoryBudgetRepo) Create(b *finance.Budget) (*finance.Budget, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.budgets[b.ID] = b
	return b, nil
}

func (r *InMemoryBudgetRepo) List(offset, limit int) ([]*finance.Budget, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*finance.Budget, 0, len(r.budgets))
	for _, v := range r.budgets {
		list = append(list, v)
	}
	return list, nil
}

type InMemoryBudgetAllocationRepo struct {
	mu          sync.RWMutex
	allocations map[string]*finance.BudgetAllocation
}

func NewInMemoryBudgetAllocationRepo() *InMemoryBudgetAllocationRepo {
	return &InMemoryBudgetAllocationRepo{
		allocations: make(map[string]*finance.BudgetAllocation),
	}
}

func (r *InMemoryBudgetAllocationRepo) Allocate(a *finance.BudgetAllocation) (*finance.BudgetAllocation, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.allocations[a.ID] = a
	return a, nil
}

func (r *InMemoryBudgetAllocationRepo) List(offset, limit int) ([]*finance.BudgetAllocation, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	list := make([]*finance.BudgetAllocation, 0, len(r.allocations))
	for _, v := range r.allocations {
		list = append(list, v)
	}
	return list, nil
}
