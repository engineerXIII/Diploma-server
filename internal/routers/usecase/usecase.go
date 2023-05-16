package usecase

import (
	"context"
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/models"
	int_routers "github.com/engineerXIII/Diploma-server/internal/routers"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/opentracing/opentracing-go"
)

// Routers UseCase
type routersUC struct {
	cfg    *config.Config
	rtRepo int_routers.Repository
	logger logger.Logger
}

// Routers UseCase constructor
func NewRoutersUseCase(cfg *config.Config, rtRepo int_routers.Repository, logger logger.Logger) int_routers.UseCase {
	return &routersUC{cfg: cfg, rtRepo: rtRepo, logger: logger}
}

// Create router
func (u *routersUC) Create(ctx context.Context, router *models.Router) (*models.Router, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routersUC.Create")
	defer span.Finish()
	return u.rtRepo.Create(ctx, router)
}

// GetByID router
func (u *routersUC) GetByID(ctx context.Context, routerID int64) (*models.Router, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routersUC.GetByID")
	defer span.Finish()

	return u.rtRepo.GetByID(ctx, routerID)
}

// Get router list
func (u *routersUC) GetList(ctx context.Context) (*models.RouterList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routersUC.GetList")
	defer span.Finish()

	return u.rtRepo.GetList(ctx)
}

// Status router
func (u *routersUC) Status(ctx context.Context, routerID int64) (*models.HealthStatus, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "routersUC.Status")
	defer span.Finish()

	// Moc
	return &models.HealthStatus{Status: "Healthy", StatusCode: "0"}, nil
}
