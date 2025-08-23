package finance

type NoteType int

const (
	NoteUnspecified NoteType = iota
	NoteCredit
	NoteDebit
)

type CreditDebitNote struct {
	ID          string
	InvoiceID   string
	Type        NoteType
	Amount      Money
	Reason      string
	Audit       AuditFields
	ExternalRefs []ExternalRef
}

// Repository interface
type CreditDebitNoteRepository interface {
	Save(note CreditDebitNote) (CreditDebitNote, error)
	FindByID(id string) (CreditDebitNote, error)
	FindAll(offset, limit int) ([]CreditDebitNote, error)
	Update(note CreditDebitNote) (CreditDebitNote, error)
	Delete(id string) error
}
