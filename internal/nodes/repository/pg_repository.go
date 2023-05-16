package repository

import (
	"context"
	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/engineerXIII/Diploma-server/internal/nodes"
	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
)

// Nodes Repository
type nodeRepo struct {
	db *sqlx.DB
}

// Nodes Repository constructor
func NewNodeRepository(db *sqlx.DB) nodes.Repository {
	return &nodeRepo{db: db}
}

// GetByID node
func (r *nodeRepo) GetByID(ctx context.Context, nodeID int64) (*models.Node, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "nodeRepo.GetByID")
	defer span.Finish()

	node := &models.Node{}
	if err := r.db.GetContext(ctx, node, getNodeByID, nodeID); err != nil {
		return nil, errors.Wrap(err, "nodeRepo.GetByID.GetContext")
	}
	return node, nil
}

func (r *nodeRepo) GetList(ctx context.Context) (*models.NodeList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "nodeRepo.GetList")
	defer span.Finish()

	var list []models.Node
	routerList := &models.NodeList{}
	if err := r.db.Select(&list, getNodeList); err != nil {
		return nil, errors.Wrap(err, "nodeRepo.GetList.GetContext")
	}
	routerList.Nodes = list
	routerList.Size = len(list)
	return routerList, nil
}
