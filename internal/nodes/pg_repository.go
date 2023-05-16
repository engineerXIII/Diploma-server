//go:generate mockgen -source pg_repository.go -destination mock/pg_repository_mock.go -package mock
package nodes

import (
	"context"

	"github.com/engineerXIII/Diploma-server/internal/models"
)

type Repository interface {
	GetList(ctx context.Context) (*models.NodeList, error)
	GetByID(ctx context.Context, nodeID int64) (*models.Node, error)
}
