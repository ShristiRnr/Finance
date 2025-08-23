package grpc

import (
	"context"

	pb "github.com/ShristiRnr/Finance/api/proto/financepb"
	"github.com/ShristiRnr/Finance/internal/core/domain/finance"
	"github.com/ShristiRnr/Finance/internal/core/services"
)

// -------------------- Budget Handler --------------------

type BudgetHandler struct {
	BudgetService *services.BudgetService
}

func (h *BudgetHandler) CreateBudget(ctx context.Context, req *pb.CreateBudgetRequest) (*pb.Budget, error) {
	budget := &finance.Budget{
		ID:     req.Budget.Id,
		Name:   req.Budget.Name,
		Total:  req.Budget.TotalAmount,
		Status: req.Budget.Status,
		Audit: finance.AuditFields{
			CreatedBy:  req.Budget.Audit.CreatedBy,
			CreatedAt:  req.Budget.Audit.CreatedAt,
		},
		ExternalRefs: mapProtoExternalRefsToDomain(req.Budget.ExternalRefs),
	}

	created, err := h.BudgetService.CreateBudget(&services.CreateBudgetRequest{Budget: budget})
	if err != nil {
		return nil, err
	}

	return mapDomainBudgetToProto(created), nil
}

func (h *BudgetHandler) ListBudgets(ctx context.Context, req *pb.ListBudgetsRequest) (*pb.ListBudgetsResponse, error) {
	budgets, err := h.BudgetService.ListBudgets(&services.ListBudgetsRequest{
		Page:  int(req.Page.PageNumber),
		Limit: int(req.Page.PageSize),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.ListBudgetsResponse{}
	for _, b := range budgets {
		resp.Budgets = append(resp.Budgets, mapDomainBudgetToProto(b))
	}
	return resp, nil
}

// -------------------- Budget Allocation Handler --------------------

type BudgetAllocationHandler struct {
	AllocationService *services.BudgetAllocationService
}

func (h *BudgetAllocationHandler) AllocateBudget(ctx context.Context, req *pb.AllocateBudgetRequest) (*pb.BudgetAllocation, error) {
	allocation := &finance.BudgetAllocation{
		ID:              req.Allocation.Id,
		BudgetID:        req.Allocation.BudgetId,
		DepartmentID:    req.Allocation.DepartmentId,
		AllocatedAmount: req.Allocation.AllocatedAmount,
		SpentAmount:     req.Allocation.SpentAmount,
		RemainingAmount: req.Allocation.RemainingAmount,
	}

	created, err := h.AllocationService.AllocateBudget(&services.AllocateBudgetRequest{Allocation: allocation})
	if err != nil {
		return nil, err
	}

	return mapDomainAllocationToProto(created), nil
}

func (h *BudgetAllocationHandler) ListBudgetAllocations(ctx context.Context, req *pb.ListBudgetAllocationsRequest) (*pb.ListBudgetAllocationsResponse, error) {
	allocs, err := h.AllocationService.ListAllocations(&services.ListBudgetAllocationsRequest{
		Page:  int(req.Page.PageNumber),
		Limit: int(req.Page.PageSize),
	})
	if err != nil {
		return nil, err
	}

	resp := &pb.ListBudgetAllocationsResponse{}
	for _, a := range allocs {
		resp.Allocations = append(resp.Allocations, mapDomainAllocationToProto(a))
	}
	return resp, nil
}

// -------------------- Budget Comparison Handler --------------------

type BudgetComparisonHandler struct {
	ComparisonService *services.BudgetComparisonService
}

func (h *BudgetComparisonHandler) GetBudgetComparisonReport(ctx context.Context, req *pb.BudgetComparisonRequest) (*pb.BudgetComparisonResponse, error) {
	report, err := h.ComparisonService.CompareBudget(req.BudgetId)
	if err != nil {
		return nil, err
	}

	return &pb.BudgetComparisonResponse{
		BudgetId:        report.BudgetID,
		TotalBudget:     report.TotalBudget,
		TotalAllocated:  report.TotalAllocated,
		TotalSpent:      report.TotalSpent,
		RemainingBudget: report.RemainingBudget,
	}, nil
}
