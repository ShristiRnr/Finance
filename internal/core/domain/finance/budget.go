package finance

import money "google.golang.org/genproto/googleapis/type/money"

type Budget struct {
	ID           string
	Name         string
	Total        *money.Money
	Status       string
	Audit        AuditFields
	ExternalRefs []ExternalRef
}

type BudgetAllocation struct {
	ID              string
	BudgetID        string
	DepartmentID    string
	AllocatedAmount *money.Money
	SpentAmount     *money.Money
	RemainingAmount *money.Money
	Audit           AuditFields
	ExternalRefs    []ExternalRef
}

type BudgetComparisonResponse struct {
	BudgetID        string
	TotalBudget     *money.Money
	TotalAllocated  *money.Money
	TotalSpent      *money.Money
	RemainingBudget *money.Money
}

type BudgetRepository interface {
	Create(budget *Budget) (*Budget, error)
	List(offset, limit int) ([]*Budget, error)
}

type BudgetAllocationRepository interface {
	Allocate(allocation *BudgetAllocation) (*BudgetAllocation, error)
	List(offset, limit int) ([]*BudgetAllocation, error)
}

