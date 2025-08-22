package finance

import "time"

type Expense struct {
	ID          string
	Category    string
	Amount      Money
	ExpenseDate time.Time
	CostCenterID string
	Audit       AuditFields
	ExternalRefs []ExternalRef
}

type CostCenter struct {
	ID          string
	Name        string
	Description string
	Audit       AuditFields
}

type CostAllocation struct {
	ID           string
	CostCenterID string
	Amount       Money
	ReferenceType string
	ReferenceID   string
	Audit        AuditFields
}
