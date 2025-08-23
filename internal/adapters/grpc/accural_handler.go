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

// --- Converters ---
func AccrualFromProto(p *pb.Accrual) finance.Accrual {
	if p == nil {
		return finance.Accrual{}
	}
	return finance.Accrual{
		ID:           p.Id,
		Description:  p.Description,
		Amount:       MoneyFromProto(p.Amount),
		AccrualDate:  p.AccrualDate.AsTime(),
		ExternalRefs: ExternalRefsFromProto(p.ExternalRefs),
	}
}

func AccrualToProto(a finance.Accrual) *pb.Accrual {
	return &pb.Accrual{
		Id:           a.ID,
		Description:  a.Description,
		Amount:       MoneyToProto(a.Amount),
		AccrualDate:  timestamppb.New(a.AccrualDate),
		ExternalRefs: ExternalRefsToProto(a.ExternalRefs),
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