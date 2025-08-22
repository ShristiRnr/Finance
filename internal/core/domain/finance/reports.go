package finance

import "time"

type ReportPeriod struct {
    StartDate time.Time
    EndDate   time.Time
}

// Profit & Loss
type ProfitLossReport struct {
    Period    ReportPeriod
    Revenue   Money
    Expenses  Money
    NetProfit Money
    Audit     AuditFields
}

// Balance Sheet
type BalanceSheetReport struct {
    Period ReportPeriod
    Assets Money
    Liabi  Money
    Equity Money
    Audit  AuditFields
}

// Trial Balance
type TrialBalanceReport struct {
    Period  ReportPeriod
    Debits  Money
    Credits Money
    Audit   AuditFields
}

// Compliance report (simplified placeholder)
type ComplianceReport struct {
    Period  ReportPeriod
    Details string
    Audit   AuditFields
}
