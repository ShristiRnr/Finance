package grpc

import (
	"context"
	"time"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/services"
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
	}, nil
}

func (h *AuditHandler) GetAuditEventById(ctx context.Context, req *pb.GetAuditEventByIdRequest) (*pb.AuditEvent, error) {
	event, err := h.service.GetEventByID(req.Id)
	if err != nil {
		return nil, err
	}
	return AuditEventToProto(event), nil
}

func (h *AuditHandler) FilterAuditEvents(ctx context.Context, req *pb.FilterAuditEventsRequest) (*pb.FilterAuditEventsResponse, error) {
	var from, to *time.Time
	if req.FromDate != nil {
		t := req.FromDate.AsTime()
		from = &t
	}
	if req.ToDate != nil {
		t := req.ToDate.AsTime()
		to = &t
	}

	events, err := h.service.FilterEvents(req.UserId, req.Action, req.ResourceType, req.ResourceId, from, to)
	if err != nil {
		return nil, err
	}
	return &pb.FilterAuditEventsResponse{
		Events: AuditEventsToProto(events),
	}, nil
}