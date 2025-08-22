package grpc

import (
	"context"

	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"github.com/ShristiRnr/Finance/internal/adapters/database/repository"
	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
)

type AccrualHandler struct {
	repo finance.AccrualRepository
	pb.UnimplementedAccrualServiceServer
}

func NewAccrualHandler(repo finance.AccrualRepository) *AccrualHandler {
	return &AccrualHandler{repo: repo}
}

func (h *AccrualHandler) CreateAccrual(ctx context.Context, req *pb.CreateAccrualRequest) (*pb.Accrual, error) {
	domainAccrual := finance.Accrual{
		ID:          req.Accrual.Id,
		Description: req.Accrual.Description,
		Amount:      finance.MoneyFromProto(req.Accrual.Amount),
		AccrualDate: req.Accrual.AccrualDate.AsTime(),
		// TODO: map Audit & ExternalRefs also
	}

	saved, err := h.repo.Save(domainAccrual)
	if err != nil {
		return nil, err
	}

	return &pb.Accrual{
		Id:          saved.ID,
		Description: saved.Description,
		Amount:      saved.Amount.ToProto(),
		AccrualDate: timestamppb.New(saved.AccrualDate),
	}, nil
}

func (h *AccrualHandler) ListAccruals(ctx context.Context, req *pb.ListAccrualsRequest) (*pb.ListAccrualsResponse, error) {
	list, err := h.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var pbAccruals []*pb.Accrual
	for _, a := range list {
		pbAccruals = append(pbAccruals, &pb.Accrual{
			Id:          a.ID,
			Description: a.Description,
			Amount:      a.Amount.ToProto(),
			AccrualDate: timestamppb.New(a.AccrualDate),
		})
	}

	return &pb.ListAccrualsResponse{
		Accruals: pbAccruals,
	}, nil
}
