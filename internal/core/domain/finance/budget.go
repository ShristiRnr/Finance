package finance

type Budget struct {
	ID        string
	Name      string
	Total     Money
	Status    string
	Audit     AuditFields
	ExternalRefs []ExternalRef
}

type BudgetAllocation struct {
	ID              string
	BudgetID        string
	DepartmentID    string
	AllocatedAmount Money
	SpentAmount     Money
	RemainingAmount Money
	Audit           AuditFields
	ExternalRefs    []ExternalRef
}
