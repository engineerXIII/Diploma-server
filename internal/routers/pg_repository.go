//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package comments

import (
	"context"

	"github.com/google/uuid"

	"github.com/engineerXIII/Diploma-server/internal/models"
)

type Repository interface {
	Create(ctx context.Context, comment *models.Router) (*models.Router, error)
	Update(ctx context.Context, comment *models.Router) (*models.Router, error)
	Delete(ctx context.Context, commentID uuid.UUID) error
	GetByID(ctx context.Context, commentID uuid.UUID) (*models.Router, error)
}
