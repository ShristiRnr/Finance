package finance

import "time"

type Accrual struct {
	ID          string
	Description string
	Amount      Money
	AccrualDate time.Time
	Audit       AuditFields
	ExternalRefs []ExternalRef
}
