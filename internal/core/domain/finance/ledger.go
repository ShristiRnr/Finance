package finance

import "time"

type LedgerSide int
const (
	LedgerUnspecified LedgerSide = iota
	LedgerDebit
	LedgerCredit
)

type LedgerEntry struct {
	ID            string
	AccountID     string
	Description   string
	Side          LedgerSide
	Amount        Money
	TransactionAt time.Time
	CostCenterID  string
	ReferenceType string
	ReferenceID   string
	Audit         AuditFields
	ExternalRefs  []ExternalRef
}
