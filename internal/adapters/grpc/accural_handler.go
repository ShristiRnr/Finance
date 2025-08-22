package grpc

import (
	"context"
	"math/big"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	
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
		Amount:      MoneyFromProto(req.Accrual.Amount),
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
		Amount:      MoneyToProto(saved.Amount),
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
			Amount:     MoneyToProto(a.Amount),
			AccrualDate: timestamppb.New(a.AccrualDate),
		})
	}

	return &pb.ListAccrualsResponse{
		Accruals: pbAccruals,
	}, nil
}

// Convert proto Money → domain Money
func MoneyFromProto(pm *pb.Money) finance.Money {
    if pm == nil {
        return finance.Money{}
    }

    units := big.NewInt(pm.Units)                // Units = int64
    nanos := big.NewInt(int64(pm.Nanos))        // Nanos = int32, cast to int64

    total := new(big.Int).Mul(units, big.NewInt(100)) // units → minor units
    total.Add(total, new(big.Int).Div(nanos, big.NewInt(1e7)))

    return finance.Money{
        Currency: pm.CurrencyCode,
        Amount:   total.Int64(),
    }
}

// Convert domain Money → proto Money
func MoneyToProto(m finance.Money) *pb.Money {
    units := m.Amount / 100
    nanos := (m.Amount % 100) * 1e7
    return &pb.Money{
        CurrencyCode: m.Currency,
        Units:        units,
        Nanos:        int32(nanos),
    }
}
