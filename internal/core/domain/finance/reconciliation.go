package finance

import "time"

type ReconciliationStatus string
const (
	ReconcileMatched   ReconciliationStatus = "MATCHED"
	ReconcileUnmatched ReconciliationStatus = "UNMATCHED"
	ReconcilePending   ReconciliationStatus = "PENDING"
)

type ReconciliationResult struct {
	ID            string
	PaymentID     string
	BankTxnID     string
	Status        ReconciliationStatus
	MatchedAmount Money
	ReconciledAt  time.Time
	Audit         AuditFields
}
