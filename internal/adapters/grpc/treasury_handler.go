package grpc

import (
	"context"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"github.com/ShristiRnr/Finance/internal/core/services"
)

type TreasuryHandler struct {
	service *services.TreasuryService
	pb.UnimplementedTreasuryServiceServer
}

func NewTreasuryHandler(s *services.TreasuryService) *TreasuryHandler {
	return &TreasuryHandler{service: s}
}

func (h *TreasuryHandler) GenerateCashFlowForecast(ctx context.Context, req *pb.CashFlowForecastRequest) (*pb.CashFlowForecastResponse, error) {
	period := finance.ReportPeriod{StartDate: req.Period.Start.AsTime(), EndDate: req.Period.End.AsTime()}

	forecast, err := h.service.GenerateCashFlowForecast(period)
	if err != nil {
		return nil, err
	}

	return &pb.CashFlowForecastResponse{ForecastDetails: forecast}, nil
}
