package http

import (
	"fmt"
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/engineerXIII/Diploma-server/internal/nodes"
	"github.com/engineerXIII/Diploma-server/pkg/httpErrors"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/engineerXIII/Diploma-server/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type nodesHandlers struct {
	cfg    *config.Config
	nodeUC nodes.UseCase
	logger logger.Logger
}

// NewNodesHandlers Routers handlers constructor
func NewNodesHandlers(cfg *config.Config, nodeUC nodes.UseCase, logger logger.Logger) nodes.Handlers {
	return &nodesHandlers{cfg: cfg, nodeUC: nodeUC, logger: logger}
}

// GetList
// @Summary Get node list
// @Description Get node list
// @Tags Nodes
// @Accept  json
// @Produce  json
// @Success 200 {object} models.NodeList
// @Failure 500 {object} httpErrors.RestErr
// @Router /node [get]
func (h *nodesHandlers) GetList() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "nodesHandlers.GetByID")
		defer span.Finish()

		nodeList, err := h.nodeUC.GetList(ctx)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, nodeList)
	}
}

// GetByID
// @Summary Get node
// @Description Get node by id
// @Tags Nodes
// @Accept  json
// @Produce  json
// @Param id path int true "node_id"
// @Success 200 {object} models.Node
// @Failure 500 {object} httpErrors.RestErr
// @Router /node/{id} [get]
func (h *nodesHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "nodesHandlers.GetByID")
		defer span.Finish()

		var nID int64
		fmt.Sscanf(c.Param("node_id"), "%d", &nID)

		node, err := h.nodeUC.GetByID(ctx, nID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, node)
	}
}

// Status
// @Summary Node Status
// @Description Get node Status by id
// @Tags Nodes
// @Accept  json
// @Produce  json
// @Param id path int true "ndoe_id"
// @Success 200 {object} models.HealthStatus
// @Failure 500 {object} httpErrors.RestErr
// @Router /router/{id} [get]
func (h *nodesHandlers) Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "routersHandlers.GetByID")
		defer span.Finish()

		var rtID int64
		fmt.Sscanf(c.Param("router_id"), "%d", &rtID)

		status := models.HealthStatus{
			Status:     "OK",
			StatusCode: "0",
			Message:    "Running",
		}

		return c.JSON(http.StatusOK, status)
	}
}
