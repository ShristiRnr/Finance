package finance

type NoteType int
const (
	NoteUnspecified NoteType = iota
	NoteCredit
	NoteDebit
)

type CreditDebitNote struct {
	ID        string
	InvoiceID string
	Type      NoteType
	Amount    Money
	Reason    string
	Audit     AuditFields
	ExternalRefs []ExternalRef
}
