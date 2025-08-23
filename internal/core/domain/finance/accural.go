package finance

import "time"

type Accrual struct {
	ID           string
	Description  string
	Amount       Money
	AccrualDate  time.Time
	Audit        AuditFields
	ExternalRefs []ExternalRef
}

// AccrualRepository is the port (interface) for persistence.
type AccrualRepository interface {
	Save(accrual Accrual) (Accrual, error)
	Update(accrual Accrual) (Accrual, error)
	FindAll() ([]Accrual, error)
	FindByID(id string) (Accrual, error)
	DeleteByID(id string) (Accrual, error)
}
