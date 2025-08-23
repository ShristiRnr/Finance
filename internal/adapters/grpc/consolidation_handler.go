package grpc

import (
	"context"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"github.com/ShristiRnr/Finance/internal/core/services"
)

type ConsolidationHandler struct {
	service *services.ConsolidationService
	pb.UnimplementedConsolidationServiceServer
}

func NewConsolidationHandler(s *services.ConsolidationService) *ConsolidationHandler {
	return &ConsolidationHandler{service: s}
}

func (h *ConsolidationHandler) ConsolidateEntities(ctx context.Context, req *pb.ConsolidationRequest) (*pb.ConsolidationResponse, error) {
	period := finance.ReportPeriod{StartDate: req.Period.Start.AsTime(), EndDate: req.Period.End.AsTime()}

	report, err := h.service.ConsolidateEntities(req.EntityIds, period)
	if err != nil {
		return nil, err
	}

	return &pb.ConsolidationResponse{Report: report.ID}, nil
}
