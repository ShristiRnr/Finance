package services

import (
	"time"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type TreasuryService struct {
	repo finance.TreasuryRepository
}

func NewTreasuryService(r finance.TreasuryRepository) *TreasuryService {
	return &TreasuryService{repo: r}
}

func (s *TreasuryService) GenerateCashFlowForecast(period finance.ReportPeriod) (*finance.CashFlowForecast, error) {
	forecast := &finance.CashFlowForecast{
		ID:             "cf-" + time.Now().Format("20060102150405"),
		OrganizationID: "org-123",
		PeriodStart:    period.StartDate,
		PeriodEnd:      period.EndDate,
		GeneratedAt:    time.Now(),
		ForecastLines: []finance.ForecastLine{
			{
				ID:          "fl-1",
				Description: "Projected Sales",
				Category:    "Revenue",
				Amount:      finance.Money{Currency: "INR", Amount: 100000},
				IsInflow:    true,
				ExpectedOn:  period.StartDate.AddDate(0, 0, 15),
				SourceSystem: "ERP",
			},
			{
				ID:          "fl-2",
				Description: "Loan Repayment",
				Category:    "Expense",
				Amount:      finance.Money{Currency: "INR", Amount: 20000},
				IsInflow:    false,
				ExpectedOn:  period.EndDate.AddDate(0, 0, -5),
				SourceSystem: "Bank",
			},
		},
	}

	if err := s.repo.SaveForecast(forecast); err != nil {
		return nil, err
	}
	return forecast, nil
}
