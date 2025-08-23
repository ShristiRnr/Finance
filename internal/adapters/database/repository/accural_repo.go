package repository

import (
	"sync"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type InMemoryAccrualRepo struct {
	mu       sync.RWMutex
	accruals map[string]finance.Accrual
}

func NewInMemoryAccrualRepo() *InMemoryAccrualRepo {
	return &InMemoryAccrualRepo{
		accruals: make(map[string]finance.Accrual),
	}
}

func (r *InMemoryAccrualRepo) Save(a finance.Accrual) (finance.Accrual, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.accruals[a.ID] = a
	return a, nil
}

func (r *InMemoryAccrualRepo) Update(a finance.Accrual) (finance.Accrual, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.accruals[a.ID]; !exists {
		return finance.Accrual{}, nil
	}
	r.accruals[a.ID] = a
	return a, nil
}

func (r *InMemoryAccrualRepo) FindAll() ([]finance.Accrual, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]finance.Accrual, 0, len(r.accruals))
	for _, v := range r.accruals {
		list = append(list, v)
	}
	return list, nil
}

func (r *InMemoryAccrualRepo) FindByID(id string) (finance.Accrual, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if accrual, exists := r.accruals[id]; exists {
		return accrual, nil
	}
	return finance.Accrual{}, nil
}

func (r *InMemoryAccrualRepo) DeleteByID(id string) (finance.Accrual, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	accrual, exists := r.accruals[id]
	if !exists {
		return finance.Accrual{}, nil
	}
	delete(r.accruals, id)
	return accrual, nil
}
