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
	SourceSystem     string // e.g. "ERP", "Subsidiary A"
	ReferenceType    string // e.g. "Invoice", "LedgerEntry"
	ReferenceID      string
	Audit           AuditFields
	ExternalRefs    []ExternalRef
}

// ConsolidatedReport = snapshot of financials across orgs/entities
type ConsolidatedReport struct {
	ID             string
	OrganizationID string
	PeriodStart    time.Time
	PeriodEnd      time.Time
	Entries        []ConsolidationEntry
	GeneratedAt    time.Time
	Audit          AuditFields
}

type CashFlowForecast struct {
	ID            string
	OrganizationID string
	PeriodStart   time.Time
	PeriodEnd     time.Time
	ForecastLines []ForecastLine
	GeneratedAt   time.Time
	Audit         AuditFields
	ExternalRefs  []ExternalRef
}

// ForecastLine represents one line item of inflow or outflow
type ForecastLine struct {
	ID          string
	Description string
	Category    string // e.g. "Revenue", "Expense", "Loan Repayment"
	Amount      Money
	IsInflow    bool
	ExpectedOn  time.Time
	SourceSystem string
	ReferenceID  string
}



type ConsolidationRepository interface {
	SaveReport(report *ConsolidatedReport) error
	FindByID(id string) (*ConsolidatedReport, error)
	FindAll() ([]*ConsolidatedReport, error)
}


type TreasuryRepository interface {
    SaveForecast(forecast *CashFlowForecast) error
    GetForecastByID(id string) (*CashFlowForecast, error)
    ListForecasts() ([]*CashFlowForecast, error)
}
