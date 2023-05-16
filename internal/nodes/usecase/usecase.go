package usecase

import (
	"context"
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/engineerXIII/Diploma-server/internal/nodes"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

// Nodes UseCase
type nodesUC struct {
	cfg    *config.Config
	nRepo  nodes.Repository
	logger logger.Logger
}

// Nodes UseCase constructor
func NewNodesUseCase(cfg *config.Config, nRepo nodes.Repository, logger logger.Logger) nodes.UseCase {
	return &nodesUC{cfg: cfg, nRepo: nRepo, logger: logger}
}

// GetByID nodes
func (u *nodesUC) GetByID(ctx context.Context, nodeID int64) (*models.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "nodesUC.GetByID")
	defer span.Finish()

	return u.nRepo.GetByID(ctx, nodeID)
}

// Get node list
func (u *nodesUC) GetList(ctx context.Context) (*models.NodeList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "nodesUC.GetList")
	defer span.Finish()

	return u.nRepo.GetList(ctx)
}
