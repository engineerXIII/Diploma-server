//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package routers

import (
	"context"

	"github.com/engineerXIII/Diploma-server/internal/models"
)

type Repository interface {
	Create(ctx context.Context, router *models.Router) (*models.Router, error)
	//Update(ctx context.Context, router *models.Router) (*models.Router, error)
	//Delete(ctx context.Context, routerID uuid.UUID) error
	GetByID(ctx context.Context, routerID int64) (*models.Router, error)
}
