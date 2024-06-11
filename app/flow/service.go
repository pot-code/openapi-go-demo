package flow

import (
	"context"
	"flow-editor-server/app/account"
	"flow-editor-server/gen/flow"

	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type service struct {
	db *gorm.DB
	c  Converter
	az *Authz
	tp trace.Tracer
}

// CopyFlow implements flow.Service.
func (s *service) CopyFlow(ctx context.Context, copyId string) (err error) {
	var m Flow
	if err := s.db.First(&m, copyId).Error; err != nil {
		return err
	}
	if err := s.az.CheckPermission(ctx, &m, "copy"); err != nil {
		return err
	}
	if err := s.az.CheckCreatePermission(ctx); err != nil {
		return err
	}

	_, span := s.tp.Start(ctx, "CopyFlow")
	defer span.End()

	return s.db.Model(&Flow{}).Omit("id").Create(&m).Error
}

// CreateFlow implements flow.Service.
func (s *service) CreateFlow(ctx context.Context, data *flow.CreateFlowData) (res *flow.FlowDetailData, err error) {
	if err := s.az.CheckCreatePermission(ctx); err != nil {
		return nil, err
	}

	_, span := s.tp.Start(ctx, "CreateFlow")
	defer span.End()

	a := account.Context(ctx)
	m := &Flow{
		Title: *data.Title,
		Nodes: data.Nodes,
		Edges: data.Edges,
		Owner: a.UserID,
	}
	if err := s.db.Create(m).Error; err != nil {
		return nil, err
	}
	return s.c.FlowToFlowDetail(m), nil
}

// DeleteFlow implements flow.Service.
func (s *service) DeleteFlow(ctx context.Context, id string) (err error) {
	var f Flow
	if err := s.db.First(&f, id).Error; err != nil {
		return err
	}
	if err := s.az.CheckPermission(ctx, &f, "delete"); err != nil {
		return err
	}
	return s.db.Delete(&f).Error
}

// GetFlow implements flow.Service.
func (s *service) GetFlow(ctx context.Context, id string) (res *flow.FlowDetailData, err error) {
	a := account.Context(ctx)
	var flow Flow
	if err := s.db.Where("id = ? AND owner = ?", id, a.UserID).First(&flow).Error; err != nil {
		return nil, err
	}
	return s.c.FlowToFlowDetail(&flow), nil
}

// GetFlowList implements flow.Service.
func (s *service) GetFlowList(ctx context.Context) (res []*flow.FlowListItemData, err error) {
	a := account.Context(ctx)
	var flows []*Flow
	if err := s.db.Find(&flows, "owner = ?", a.UserID).Error; err != nil {
		return nil, err
	}
	return s.c.FlowSliceToFlowList(flows), nil
}

// UpdateFlow implements flow.Service.
func (s *service) UpdateFlow(ctx context.Context, payload *flow.UpdateFlowPayload) (res *flow.FlowDetailData, err error) {
	data := payload.Data

	var model Flow
	if err := s.db.First(&model, "id = ?", payload.ID).Error; err != nil {
		return nil, err
	}
	if err := s.az.CheckPermission(ctx, &model, "update"); err != nil {
		return nil, err
	}
	if data.Edges != nil {
		model.Edges = data.Edges
	}
	if data.Nodes != nil {
		model.Nodes = data.Nodes
	}
	if data.Title != nil {
		model.Title = *data.Title
	}
	if err := s.db.Save(&model).Error; err != nil {
		return nil, err
	}
	return s.c.FlowToFlowDetail(&model), nil
}

var _ flow.Service = (*service)(nil)

func NewService(db *gorm.DB, c Converter, a *Authz, tp trace.Tracer) *service {
	return &service{db: db, c: c, az: a, tp: tp}
}
