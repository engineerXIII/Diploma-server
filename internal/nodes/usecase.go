//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock
package nodes

import (
	"context"

	"github.com/engineerXIII/Diploma-server/internal/models"
)

// Nodes use case
type UseCase interface {
	GetList(ctx context.Context) (*models.NodeList, error)
	GetByID(ctx context.Context, nodeID int64) (*models.Node, error)
}
