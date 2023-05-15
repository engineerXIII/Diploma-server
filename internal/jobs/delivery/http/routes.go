package http

import (
	"github.com/engineerXIII/Diploma-server/internal/jobs"
	"github.com/engineerXIII/Diploma-server/internal/middleware"
	"github.com/labstack/echo/v4"
)

// Map jobs routes
func MapJobsRoutes(jobGroup *echo.Group, h jobs.Handlers, mw *middleware.MiddlewareManager) {
	//jobGroup.POST("", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	jobGroup.POST("", h.Create())
	jobGroup.GET("/:job_id/status", h.Status())
}
