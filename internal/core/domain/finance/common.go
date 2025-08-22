package finance

import "time"

// Value Object
type Money struct {
	Currency string
	Amount   int64 // use minor units (paise, cents)
}

type RequestMetadata struct {
	RequestID      string
	OrganizationID string
	TenantID       string
}

type AuditFields struct {
	CreatedAt time.Time
	CreatedBy string
	UpdatedAt time.Time
	UpdatedBy string
	Revision  string
}

type ExternalRef struct {
	System string
	RefId    string
}

type PartyKind int

const (
	PartyUnspecified PartyKind = iota
	PartyCustomer
	PartyVendor
)

type PartyRef struct {
	Kind        PartyKind
	Ref         ExternalRef
	DisplayName string
}
