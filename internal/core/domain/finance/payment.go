package finance

import "time"

// PaymentStatus captures the lifecycle of a payment
type PaymentStatus int

const (
	PaymentStatusUnspecified PaymentStatus = iota
	PaymentStatusDue
	PaymentStatusPartiallyPaid
	PaymentStatusPaid
	PaymentStatusWriteOff
)

// PaymentDue = expected payment linked to an invoice
type PaymentDue struct {
	ID        string
	InvoiceID string
	AmountDue Money
	DueDate   time.Time
	Status    PaymentStatus

	Audit        AuditFields
	ExternalRefs []ExternalRef
}

// BankTransaction = real-world bank movement (imported from bank feed, statement, etc.)
type BankTransaction struct {
	ID              string
	Amount          Money
	TransactionDate time.Time
	Description     string

	Audit        AuditFields
	ExternalRefs []ExternalRef
}
