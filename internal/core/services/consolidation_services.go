package services

import (
	"time"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
)

type ConsolidationService struct {
	repo finance.ConsolidationRepository
}

func NewConsolidationService(r finance.ConsolidationRepository) *ConsolidationService {
	return &ConsolidationService{repo: r}
}

func (s *ConsolidationService) ConsolidateEntities(entityIDs []string, period finance.ReportPeriod) (*finance.ConsolidatedReport, error) {
	// Fake consolidation logic for now
	report := &finance.ConsolidatedReport{
		ID:             "rep-" + time.Now().Format("20060102150405"),
		OrganizationID: "org-123",
		PeriodStart:    period.StartDate,
		PeriodEnd:      period.EndDate,
		GeneratedAt:    time.Now(),
	}

	err := s.repo.SaveReport(report)
	if err != nil {
		return nil, err
	}
	return report, nil
}
