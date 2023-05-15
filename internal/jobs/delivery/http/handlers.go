package http

import (
	"github.com/engineerXIII/Diploma-server/config"
	jobs "github.com/engineerXIII/Diploma-server/internal/jobs"
	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/engineerXIII/Diploma-server/pkg/httpErrors"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/engineerXIII/Diploma-server/pkg/utils"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type jobsHandlers struct {
	cfg    *config.Config
	jobUC  jobs.UseCase
	logger logger.Logger
}

func NewJobsHandlers(cfg *config.Config, jobUC jobs.UseCase, logger logger.Logger) jobs.Handlers {
	return &jobsHandlers{cfg: cfg, jobUC: jobUC, logger: logger}
}

// Create
// @Summary Create new job
// @Description create new job
// @Tags Jobs
// @Accept  json
// @Produce  json
// @Success 201 {object} models.Job
// @Failure 500 {object} httpErrors.RestErr
// @Router /job [post]
func (h jobsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "jobsHandlers.Create")
		defer span.Finish()

		job := &models.Job{}

		if err := utils.SanitizeRequest(c, job); err != nil {
			return utils.ErrResponseWithLog(c, h.logger, err)
			// return err
		}

		createdJob, err := h.jobUC.Create(ctx, job)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusCreated, createdJob)
	}
}

// Status
// @Summary Job Status
// @Description Get job Status by id
// @Tags Jobs
// @Accept  json
// @Produce  json
// @Param id path int true "job_id"
// @Success 200 {object} models.Jobs
// @Failure 500 {object} httpErrors.RestErr
// @Router /job/{id}/status [get]
func (h jobsHandlers) Status() echo.HandlerFunc {
	return func(c echo.Context) error {
		span, ctx := opentracing.StartSpanFromContext(utils.GetRequestCtx(c), "jobsHandlers.Status")
		defer span.Finish()

		jobID, err := uuid.Parse(c.Param("job_id"))
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		job, err := h.jobUC.Status(ctx, jobID)
		if err != nil {
			utils.LogResponseError(c, h.logger, err)
			return c.JSON(httpErrors.ErrorResponse(err))
		}

		return c.JSON(http.StatusOK, job)
	}
}
