package http

import (
	"fmt"
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/models"
	int_routers "github.com/engineerXIII/Diploma-server/internal/routers"
	"github.com/engineerXIII/Diploma-server/pkg/httpErrors"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/engineerXIII/Diploma-server/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type routersHandlers struct {
	cfg    *config.Config
	rtUC   int_routers.UseCase
	logger logger.Logger
}

// NewCommentsHandlers Routers handlers constructor
func NewRoutersHandlers(cfg *config.Config, rtUC int_routers.UseCase, logger logger.Logger) int_routers.Handlers {
	return &routersHandlers{cfg: cfg, rtUC: rtUC, logger: logger}
}

// Create
// @Summary Create new router
// @Description create new router
// @Tags Routers
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Router
// @Failure 500 {object} httpErrors.RestErr
// @Router /router [post]
func (h *routersHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "routersHandlers.Create")
		defer span.Finish()

		router := &models.Router{}

		if err := utils.SanitizeRequest(c, router); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
			// return err
		}

		createdRouter, err := h.rtUC.Create(ctx, router)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdRouter)
	}
}

// GetByID
// @Summary Get router
// @Description Get router by id
// @Tags Routers
// @Accept  json
// @Produce  json
// @Param id path int true "router_id"
// @Success 200 {object} models.Router
// @Failure 500 {object} httpErrors.RestErr
// @Router /router/{id} [get]
func (h *routersHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "routersHandlers.GetByID")
		defer span.Finish()

		var rtID int64
		fmt.Sscanf(c.Param("router_id"), "%d", &rtID)

		router, err := h.rtUC.GetByID(ctx, rtID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, router)
	}
}

// Status
// @Summary Router Status
// @Description Get router Status by id
// @Tags Routers
// @Accept  json
// @Produce  json
// @Param id path int true "router_id"
// @Success 200 {object} models.HealthStatus
// @Failure 500 {object} httpErrors.RestErr
// @Router /router/{id} [get]
func (h *routersHandlers) Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "routersHandlers.GetByID")
		defer span.Finish()

		var rtID int64
		fmt.Sscanf(c.Param("router_id"), "%d", &rtID)

		status, err := h.rtUC.Status(ctx, rtID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, status)
	}
}
