package grpc

import (
	"context"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"github.com/ShristiRnr/Finance/internal/core/services"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AuditHandler struct {
	pb.UnimplementedAuditTrailServiceServer
	service *services.AuditService
}

func NewAuditHandler(s *services.AuditService) *AuditHandler {
	return &AuditHandler{service: s}
}

func (h *AuditHandler) RecordAuditEvent(ctx context.Context, req *pb.RecordAuditEventRequest) (*pb.AuditEvent, error) {
	domainEvent := AuditEventFromProto(req.Event)

	saved, err := h.service.RecordEvent(domainEvent)
	if err != nil {
		return nil, err
	}

	return AuditEventToProto(saved), nil
}

func (h *AuditHandler) ListAuditEvents(ctx context.Context, req *pb.ListAuditEventsRequest) (*pb.ListAuditEventsResponse, error) {
	events, err := h.service.ListEvents()
	if err != nil {
		return nil, err
	}

	return &pb.ListAuditEventsResponse{
		Events: AuditEventsToProto(events),
		// Page: TODO: Implement real pagination
	}, nil
}

// -------- converters ----------

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
