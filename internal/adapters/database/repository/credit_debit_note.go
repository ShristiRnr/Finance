package repository

import (
	"errors"
	"sync"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type InMemoryCreditDebitNoteRepo struct {
	mu    sync.RWMutex
	notes map[string]finance.CreditDebitNote
}

func NewInMemoryCreditDebitNoteRepo() *InMemoryCreditDebitNoteRepo {
	return &InMemoryCreditDebitNoteRepo{
		notes: make(map[string]finance.CreditDebitNote),
	}
}

func (r *InMemoryCreditDebitNoteRepo) Save(note finance.CreditDebitNote) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.notes[note.ID] = note
	return nil
}

func (r *InMemoryCreditDebitNoteRepo) FindByID(id string) (finance.CreditDebitNote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	note, exists := r.notes[id]
	if !exists {
		return finance.CreditDebitNote{}, errors.New("note not found")
	}
	return note, nil
}

func (r *InMemoryCreditDebitNoteRepo) FindAll() ([]finance.CreditDebitNote, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	list := make([]finance.CreditDebitNote, 0, len(r.notes))
	for _, note := range r.notes {
		list = append(list, note)
	}
	return list, nil
}

func (r *InMemoryCreditDebitNoteRepo) Update(note finance.CreditDebitNote) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.notes[note.ID]; !exists {
		return errors.New("note not found")
	}
	r.notes[note.ID] = note
	return nil
}

func (r *InMemoryCreditDebitNoteRepo) Delete(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, exists := r.notes[id]; !exists {
		return errors.New("note not found")
	}
	delete(r.notes, id)
	return nil
}
