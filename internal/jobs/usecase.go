//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package jobs

import (
	"context"
	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/google/uuid"
)

type UseCase interface {
	Create(ctx context.Context, job *models.Job) (*models.Job, error)
	Status(ctx context.Context, jobID uuid.UUID) (*models.Job, error)
}
