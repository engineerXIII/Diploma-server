package http

import (
	"github.com/engineerXIII/Diploma-server/internal/middleware"
	"github.com/engineerXIII/Diploma-server/internal/nodes"
	"github.com/labstack/echo/v4"
)

// Map nodes routes
func MapNodesRoutes(nodeGroup *echo.Group, h nodes.Handlers, mw *middleware.MiddlewareManager) {
	nodeGroup.GET("/", h.GetList())
	nodeGroup.GET("/:node_id", h.GetByID())
	nodeGroup.GET("/:node_id/status", h.Status())
}
