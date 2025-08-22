package grpc

import (
	"context"
	"math/big"

	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
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
	domainAccrual := finance.Accrual{
		ID:          req.Accrual.Id,
		Description: req.Accrual.Description,
		Amount:      MoneyFromProto(req.Accrual.Amount),
		AccrualDate: req.Accrual.AccrualDate.AsTime(),
		ExternalRefs: ExternalRefsFromProto(req.Accrual.ExternalRefs),
	}

	saved, err := h.service.CreateAccrual(domainAccrual)
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
	list, err := h.service.ListAccruals()
	if err != nil {
		return nil, err
	}

	var pbAccruals []*pb.Accrual
	for _, a := range list {
		pbAccruals = append(pbAccruals, &pb.Accrual{
			Id:          a.ID,
			Description: a.Description,
			Amount:      MoneyToProto(a.Amount),
			AccrualDate: timestamppb.New(a.AccrualDate),
		})
	}

	return &pb.ListAccrualsResponse{
		Accruals: pbAccruals,
	}, nil
}

// --- Converters ---
func MoneyFromProto(pm *pb.Money) finance.Money {
	if pm == nil {
		return finance.Money{}
	}

	units := big.NewInt(pm.Units)
	nanos := big.NewInt(int64(pm.Nanos))

	total := new(big.Int).Mul(units, big.NewInt(100))
	total.Add(total, new(big.Int).Div(nanos, big.NewInt(1e7)))

	return finance.Money{
		Currency: pm.CurrencyCode,
		Amount:   total.Int64(),
	}
}

func MoneyToProto(m finance.Money) *pb.Money {
	units := m.Amount / 100
	nanos := (m.Amount % 100) * 1e7
	return &pb.Money{
		CurrencyCode: m.Currency,
		Units:        units,
		Nanos:        int32(nanos),
	}
}

func ExternalRefsFromProto(p []*pb.ExternalRef) []finance.ExternalRef {
	if p == nil {
		return nil
	}
	var refs []finance.ExternalRef
	for _, r := range p {
		refs = append(refs, finance.ExternalRef{
			System: r.System,
			RefId:  r.RefId,
		})
	}
	return refs
}

func ExternalRefsToProto(refs []finance.ExternalRef) []*pb.ExternalRef {
	var pbRefs []*pb.ExternalRef
	for _, r := range refs {
		pbRefs = append(pbRefs, &pb.ExternalRef{
			System: r.System,
			RefId:  r.RefId,
		})
	}
	return pbRefs
}
