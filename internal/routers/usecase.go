//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package routers

import (
	"context"

	"github.com/engineerXIII/Diploma-server/internal/models"
)

// Routers use case
type UseCase interface {
	Create(ctx context.Context, router *models.Router) (*models.Router, error)
	//Update(ctx context.Context, router *models.Router) (*models.Router, error)
	//Delete(ctx context.Context, routerID uuid.UUID) error
	GetByID(ctx context.Context, routerID int64) (*models.Router, error)
	Status(ctx context.Context, routerID int64) (*models.HealthStatus, error)
}
