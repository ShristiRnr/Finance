package finance

import "time"

// ConsolidationEntry = one record during financial consolidation
type ConsolidationEntry struct {
	ID              string
	OrganizationID  string
	AccountID       string
	Amount          Money
	PeriodStart     time.Time
	PeriodEnd       time.Time
	SourceSystem    string // e.g. "ERP", "Subsidiary A"
	ReferenceType   string // e.g. "Invoice", "LedgerEntry"
	ReferenceID     string

	Audit        AuditFields
	ExternalRefs []ExternalRef
}

// ConsolidatedReport = snapshot of financials across orgs/entities
type ConsolidatedReport struct {
	ID             string
	OrganizationID string
	PeriodStart    time.Time
	PeriodEnd      time.Time
	Entries        []ConsolidationEntry

	GeneratedAt time.Time
	Audit       AuditFields
}
