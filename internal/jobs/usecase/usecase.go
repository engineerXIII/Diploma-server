package usecase

import (
	"context"
	"github.com/engineerXIII/Diploma-server/config"
	"github.com/engineerXIII/Diploma-server/internal/jobs"
	"github.com/engineerXIII/Diploma-server/internal/models"
	"github.com/engineerXIII/Diploma-server/pkg/logger"
	"github.com/google/uuid"
	"github.com/opentracing/opentracing-go"
	"time"
)

type jobsUC struct {
	cfg    *config.Config
	logger logger.Logger
}

// Jobs UseCase constructor
func NewJobsUseCase(cfg *config.Config, logger logger.Logger) jobs.UseCase {
	return &jobsUC{cfg: cfg, logger: logger}
}

// Create job
func (u *jobsUC) Create(ctx context.Context, job *models.Job) (*models.Job, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "jobsUC.Create")
	defer span.Finish()
	// Moc
	return &models.Job{
		JobID:     uuid.New(),
		Name:      job.Name,
		JobType:   job.JobType,
		Status:    "Running",
		Message:   "Initiating autoheal procedure",
		CreatedAt: time.Now(),
	}, nil
}

// Job status
func (u *jobsUC) Status(ctx context.Context, jobID uuid.UUID) (*models.Job, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "jobsUC.Status")
	defer span.Finish()

	// Moc
	return &models.Job{
		JobID:      jobID,
		Status:     "Finished",
		FinishedAt: time.Now(),
		Message: "Found issue with network interface...Analyzing...\n" +
			"Trying to shutdown and reload network connection... OK\n" +
			"Network interface reloaded\n" +
			"Checking connectivity...OK\n" +
			"Autoheal completed. No user required.",
	}, nil
}
