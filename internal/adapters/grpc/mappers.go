package grpc
 import(
	"math/big"
	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"google.golang.org/protobuf/types/known/timestamppb"
 )

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

func AuditEventFromProto(p *pb.AuditEvent) finance.AuditEvent {
	if p == nil {
		return finance.AuditEvent{}
	}
	return finance.AuditEvent{
		ID:           p.Id,
		UserID:       p.UserId,
		Action:       p.Action,
		Timestamp:    p.Timestamp.AsTime(),
		Details:      p.Details,
		ResourceType: p.ResourceType,
		ResourceID:   p.ResourceId,
	}
}

func AuditEventToProto(e finance.AuditEvent) *pb.AuditEvent {
	return &pb.AuditEvent{
		Id:           e.ID,
		UserId:       e.UserID,
		Action:       e.Action,
		Timestamp:    timestamppb.New(e.Timestamp),
		Details:      e.Details,
		ResourceType: e.ResourceType,
		ResourceId:   e.ResourceID,
	}
}

func AuditEventsToProto(events []finance.AuditEvent) []*pb.AuditEvent {
	var result []*pb.AuditEvent
	for _, e := range events {
		result = append(result, AuditEventToProto(e))
	}
	return result
}

func mapDomainAllocationToProto(a *finance.BudgetAllocation) *pb.BudgetAllocation {
	return &pb.BudgetAllocation{
		Id:              a.ID,
		BudgetId:        a.BudgetID,
		DepartmentId:    a.DepartmentID,
		AllocatedAmount: a.AllocatedAmount,
		SpentAmount:     a.SpentAmount,
		RemainingAmount: a.RemainingAmount,
	}
}

func mapDomainBudgetToProto(b *finance.Budget) *pb.Budget {
	return &pb.Budget{
		Id:           b.ID,
		Name:         b.Name,
		TotalAmount:  b.Total,
		Status:       b.Status,
		Audit:        &pb.AuditFields{CreatedBy: b.Audit.CreatedBy, CreatedAt: b.Audit.CreatedAt},
		ExternalRefs: mapDomainExternalRefsToProto(b.ExternalRefs),
	}
}

func mapDomainExternalRefsToProto(refs []finance.ExternalRef) []*pb.ExternalRef {
	protoRefs := make([]*pb.ExternalRef, len(refs))
	for i, r := range refs {
		protoRefs[i] = &pb.ExternalRef{
			System: r.System,
			RefId:     r.RefId,
		}
	}
	return protoRefs
}

func mapProtoExternalRefsToDomain(refs []*pb.ExternalRef) []finance.ExternalRef {
	domainRefs := make([]finance.ExternalRef, len(refs))
	for i, r := range refs {
		domainRefs[i] = finance.ExternalRef{
			System: r.System,
			RefId:     r.RefId,
		}
	}
	return domainRefs
}

func toProtoNote(note finance.CreditDebitNote) *pb.CreditDebitNote {
	return &pb.CreditDebitNote{
		Id:        note.ID,
		InvoiceId: note.InvoiceID,
		Type:      pb.NoteType(note.Type),
		Amount:    toProtoMoney(note.Amount),
		Reason:    note.Reason,
	}
}

func toProtoMoney(m finance.Money) *pb.Money {
	return &pb.Money{
		CurrencyCode: m.Currency,
		Units:        m.Amount,
		Nanos:        0,
	}
}

func fromProtoMoney(m *pb.Money) finance.Money {
	if m == nil {
		return finance.Money{}
	}
	return finance.Money{
		Currency: m.CurrencyCode,
		Amount:   m.Units, // proto Units -> domain Amount
	}
}

// Proto → Domain
func mapProtoPeriodToDomain(p *pb.ReportPeriod) finance.ReportPeriod {
    return finance.ReportPeriod{
        StartDate: p.PeriodStart.AsTime(),
        EndDate:   p.PeriodEnd.AsTime(),
    }
}

// Domain → Proto
func mapDomainPeriodToProto(p finance.ReportPeriod) *pb.ReportPeriod {
    return &pb.ReportPeriod{
        StartDate: timestamppb.New(p.StartDate),
        EndDate:   timestamppb.New(p.EndDate),
    }
}

