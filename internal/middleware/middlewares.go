package middleware

import (
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/auth"
	"github.com/engineerXIII/Diploma-server/internal/session"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
)

// Middleware manager
type MiddlewareManager struct {
	sessUC  session.UCSession
	authUC  auth.UseCase
	cfg     *config.Config
	origins []string
	logger  logger.Logger
}

// Middleware manager constructor
func NewMiddlewareManager(sessUC session.UCSession, authUC auth.UseCase, cfg *config.Config, origins []string, logger logger.Logger) *MiddlewareManager {
	return &MiddlewareManager{sessUC: sessUC, authUC: authUC, cfg: cfg, origins: origins, logger: logger}
}
