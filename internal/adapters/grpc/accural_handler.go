package grpc

import (
	"context"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/services"
)

type AccrualHandler struct {
	service *services.AccrualService
	pb.UnimplementedAccrualServiceServer
}

func NewAccrualHandler(service *services.AccrualService) *AccrualHandler {
	return &AccrualHandler{service: service}
}

func (h *AccrualHandler) CreateAccrual(ctx context.Context, req *pb.CreateAccrualRequest) (*pb.Accrual, error) {
	accrual := AccrualFromProto(req.Accrual)
	saved, err := h.service.CreateAccrual(accrual)
	if err != nil {
		return nil, err
	}
	return AccrualToProto(saved), nil
}

func (h *AccrualHandler) ListAccruals(ctx context.Context, req *pb.ListAccrualsRequest) (*pb.ListAccrualsResponse, error) {
	list, err := h.service.ListAccruals()
	if err != nil {
		return nil, err
	}

	var pbAccruals []*pb.Accrual
	for _, a := range list {
		pbAccruals = append(pbAccruals, AccrualToProto(a))
	}

	return &pb.ListAccrualsResponse{Accruals: pbAccruals}, nil
}

func (h *AccrualHandler) GetAccrualById(ctx context.Context, req *pb.GetAccrualByIdRequest) (*pb.Accrual, error) {
	a, err := h.service.GetAccrualByID(req.Id)
	if err != nil {
		return nil, err
	}
	return AccrualToProto(a), nil
}

func (h *AccrualHandler) UpdateAccrual(ctx context.Context, req *pb.UpdateAccrualRequest) (*pb.Accrual, error) {
	a := AccrualFromProto(req.Accrual)
	updated, err := h.service.UpdateAccrual(a)
	if err != nil {
		return nil, err
	}
	return AccrualToProto(updated), nil
}

func (h *AccrualHandler) DeleteAccrualById(ctx context.Context, req *pb.DeleteAccrualRequest) (*pb.Accrual, error) {
	deleted, err := h.service.DeleteAccrualByID(req.Id)
	if err != nil {
		return nil, err
	}
	return AccrualToProto(deleted), nil
}