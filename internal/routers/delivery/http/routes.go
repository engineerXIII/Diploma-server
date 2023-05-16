package http

import (
	"github.com/engineerXIII/Diploma-server/internal/routers"
	"github.com/labstack/echo/v4"

	"github.com/engineerXIII/Diploma-server/internal/middleware"
)

// Map routers routes
func MapRoutersRoutes(routerGroup *echo.Group, h routers.Handlers, mw *middleware.MiddlewareManager) {
	routerGroup.POST("", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	//routerGroup.DELETE("/:router_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	//routerGroup.PUT("/:router_id", h.Update())
	routerGroup.GET("/", h.GetList())
	routerGroup.GET("/:router_id", h.GetByID())
	routerGroup.GET("/:router_id/status", h.Status())
}
