package repository

import (
	"context"
	"github.com/engineerXIII/Diploma-server/internal/routers"

	"github.com/jmoiron/sqlx"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"

	"github.com/engineerXIII/Diploma-server/internal/models"
)

// Routers Repository
type routerRepo struct {
	db *sqlx.DB
}

// Routers Repository constructor
func NewRouterRepository(db *sqlx.DB) routers.Repository {
	return &routerRepo{db: db}
}

// Create router
func (r *routerRepo) Create(ctx context.Context, router *models.Router) (*models.Router, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.Create")
	defer span.Finish()

	c := &models.Router{}
	if err := r.db.QueryRowxContext(
		ctx,
		createRouter,
		&router.Type,
		&router.Address,
		&router.Hostname,
		&router.Port,
	).StructScan(c); err != nil {
		return nil, errors.Wrap(err, "routerRepo.Create.StructScan")
	}

	return c, nil
}

// GetByID router
func (r *routerRepo) GetByID(ctx context.Context, routerID int64) (*models.Router, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routerRepo.GetByID")
	defer span.Finish()

	router := &models.Router{}
	if err := r.db.GetContext(ctx, router, getRouterByID, routerID); err != nil {
		return nil, errors.Wrap(err, "routerRepo.GetByID.GetContext")
	}
	return router, nil
}
