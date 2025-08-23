package repository

import (
	"errors"
	"sync"
)

type InMemoryTreasuryRepo struct {
	mu       sync.RWMutex
	forecast string
}

func NewInMemoryTreasuryRepo() *InMemoryTreasuryRepo {
	return &InMemoryTreasuryRepo{}
}

func (r *InMemoryTreasuryRepo) SaveForecast(forecast string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.forecast = forecast
	return nil
}

func (r *InMemoryTreasuryRepo) GetLatestForecast() (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.forecast == "" {
		return "", errors.New("no forecast available")
	}
	return r.forecast, nil
}
