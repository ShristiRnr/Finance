package services

import (
	"errors"
	money "google.golang.org/genproto/googleapis/type/money"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

// -------------------- Request/Response Structs --------------------

type CreateBudgetRequest struct {
	Budget *finance.Budget
}

type ListBudgetsRequest struct {
	Page  int
	Limit int
}

type AllocateBudgetRequest struct {
	Allocation *finance.BudgetAllocation
}

type ListBudgetAllocationsRequest struct {
	Page  int
	Limit int
}

// Response for budget comparison (domain level)
type BudgetComparisonResponse struct {
	BudgetID        string
	TotalBudget     *money.Money
	TotalAllocated  *money.Money
	TotalSpent      *money.Money
	RemainingBudget *money.Money
}

// -------------------- Budget Service --------------------

type BudgetService struct {
	BudgetRepo finance.BudgetRepository
}

func (s *BudgetService) CreateBudget(req *CreateBudgetRequest) (*finance.Budget, error) {
	if req.Budget == nil {
		return nil, errors.New("budget cannot be nil")
	}
	return s.BudgetRepo.Create(req.Budget)
}

func (s *BudgetService) ListBudgets(req *ListBudgetsRequest) ([]*finance.Budget, error) {
	return s.BudgetRepo.List(req.Page, req.Limit)
}

// -------------------- Budget Allocation Service --------------------

type BudgetAllocationService struct {
	AllocationRepo finance.BudgetAllocationRepository
}

func (s *BudgetAllocationService) AllocateBudget(req *AllocateBudgetRequest) (*finance.BudgetAllocation, error) {
	if req.Allocation == nil {
		return nil, errors.New("allocation cannot be nil")
	}
	return s.AllocationRepo.Allocate(req.Allocation)
}

func (s *BudgetAllocationService) ListAllocations(req *ListBudgetAllocationsRequest) ([]*finance.BudgetAllocation, error) {
	return s.AllocationRepo.List(req.Page, req.Limit)
}

// -------------------- Budget Comparison Service --------------------

type BudgetComparisonService struct {
	BudgetRepo     finance.BudgetRepository
	AllocationRepo finance.BudgetAllocationRepository
}

func (s *BudgetComparisonService) CompareBudget(budgetID string) (*BudgetComparisonResponse, error) {
	// Fetch all budgets (for simplicity)
	budgets, err := s.BudgetRepo.List(0, 1000)
	if err != nil {
		return nil, err
	}

	var budget *finance.Budget
	for _, b := range budgets {
		if b.ID == budgetID {
			budget = b
			break
		}
	}
	if budget == nil {
		return nil, errors.New("budget not found")
	}

	// Fetch all allocations
	allocs, err := s.AllocationRepo.List(0, 1000)
	if err != nil {
		return nil, err
	}

	var totalAllocated, totalSpent int64
	for _, a := range allocs {
		if a.BudgetID == budgetID {
			if a.AllocatedAmount != nil {
				totalAllocated += a.AllocatedAmount.Units
			}
			if a.SpentAmount != nil {
				totalSpent += a.SpentAmount.Units
			}
		}
	}

	remaining := budget.Total.Units - totalAllocated

	return &BudgetComparisonResponse{
		BudgetID:    budgetID,
		TotalBudget: budget.Total,
		TotalAllocated: &money.Money{
			CurrencyCode: budget.Total.CurrencyCode,
			Units:        totalAllocated,
			Nanos:        0,
		},
		TotalSpent: &money.Money{
			CurrencyCode: budget.Total.CurrencyCode,
			Units:        totalSpent,
			Nanos:        0,
		},
		RemainingBudget: &money.Money{
			CurrencyCode: budget.Total.CurrencyCode,
			Units:        remaining,
			Nanos:        0,
		},
	}, nil
}
