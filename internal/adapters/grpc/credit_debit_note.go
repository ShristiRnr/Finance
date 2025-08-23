package grpc

import (
	"context"
	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"github.com/ShristiRnr/Finance/internal/core/services"

	"google.golang.org/protobuf/types/known/emptypb"
)

type CreditDebitNoteHandler struct {
	service services.CreditDebitNoteService
	pb.UnimplementedCreditDebitNoteServiceServer
}

func NewCreditDebitNoteHandler(s services.CreditDebitNoteService) *CreditDebitNoteHandler {
	return &CreditDebitNoteHandler{service: s}
}

func (h *CreditDebitNoteHandler) CreateCreditDebitNote(ctx context.Context, req *pb.CreateCreditDebitNoteRequest) (*pb.CreditDebitNote, error) {
	note := finance.CreditDebitNote{
		ID:        req.Note.Id,
		InvoiceID: req.Note.InvoiceId,
		Type:      finance.NoteType(req.Note.Type),
		Amount:    fromProtoMoney(req.Note.Amount),
		Reason:    req.Note.Reason,
		// Audit & ExternalRefs mapping left as exercise
	}

	created, err := h.service.Create(note)
	if err != nil {
		return nil, err
	}
	return toProtoNote(created), nil
}

func (h *CreditDebitNoteHandler) GetCreditDebitNote(ctx context.Context, req *pb.GetCreditDebitNoteRequest) (*pb.CreditDebitNote, error) {
	note, err := h.service.GetByID(req.Id)
	if err != nil {
		return nil, err
	}
	return toProtoNote(note), nil
}

func (h *CreditDebitNoteHandler) ListCreditDebitNotes(ctx context.Context, req *pb.ListCreditDebitNotesRequest) (*pb.ListCreditDebitNotesResponse, error) {
	notes, err := h.service.List(int(req.Page.Offset), int(req.Page.Limit))
	if err != nil {
		return nil, err
	}

	resp := &pb.ListCreditDebitNotesResponse{}
	for _, n := range notes {
		resp.Notes = append(resp.Notes, toProtoNote(n))
	}
	return resp, nil
}

func (h *CreditDebitNoteHandler) UpdateCreditDebitNote(ctx context.Context, req *pb.UpdateCreditDebitNoteRequest) (*pb.CreditDebitNote, error) {
	note := finance.CreditDebitNote{
		ID:        req.Note.Id,
		InvoiceID: req.Note.InvoiceId,
		Type:      finance.NoteType(req.Note.Type),
		Amount:    fromProtoMoney(req.Note.Amount),
		Reason:    req.Note.Reason,
	}

	updated, err := h.service.Update(note)
	if err != nil {
		return nil, err
	}
	return toProtoNote(updated), nil
}

func (h *CreditDebitNoteHandler) DeleteCreditDebitNote(ctx context.Context, req *pb.DeleteCreditDebitNoteRequest) (*emptypb.Empty, error) {
	err := h.service.Delete(req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
