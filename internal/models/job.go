package models

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	JobID      uuid.UUID `json:"job_id,omitempty" db:"job_id"`
	Name       string    `json:"name" db:"name" validate:"required"`
	JobType    string    `json:"job_type" db:"job_type"`
	Status     string    `json:"status" db:"status"`
	Message    string    `json:"message" db:"message"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
	FinishedAt time.Time `json:"finished_at,omitempty" db:"finished_at"`
}
