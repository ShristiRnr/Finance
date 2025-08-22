package finance

import "time"

type InvoiceType int
const (
	InvoiceUnspecified InvoiceType = iota
	InvoiceSales
	InvoiceProforma
	InvoiceChallan
	InvoicePurchase
)

type InvoiceStatus int
const (
	InvoiceStatusUnspecified InvoiceStatus = iota
	InvoiceDraft
	InvoiceIssued
	InvoicePartiallyPaid
	InvoicePaid
	InvoiceVoid
	InvoiceOverdue
)

type TaxType int
const (
	TaxUnspecified TaxType = iota
	TaxCGST
	TaxSGST
	TaxIGST
	TaxVAT
	TaxOther = 10
)

type TaxLine struct {
	Type        TaxType
	RatePercent float64
	Amount      Money
}

type InvoiceItem struct {
	ID           string
	Name         string
	Description  string
	HSN          string
	Quantity     int
	UnitPrice    Money
	LineSubtotal Money
	Taxes        []TaxLine
	LineTotal    Money
	CostCenterID string
}

type Invoice struct {
	ID              string
	InvoiceNumber   string
	Type            InvoiceType
	InvoiceDate     time.Time
	DueDate         time.Time
	DeliveryDate    time.Time
	Party           PartyRef
	OrganizationID  string
	PONumber        string
	Status          InvoiceStatus
	PaymentRef      string
	Items           []InvoiceItem
	Subtotal        Money
	Taxes           []TaxLine
	Total           Money
	Audit           AuditFields
	GSTIN           string
	ExternalRefs    []ExternalRef
}
