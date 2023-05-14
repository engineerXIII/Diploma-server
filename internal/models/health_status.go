package models

type HealthStatus struct {
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
	Message    string `json:"message" validate:"omitempty"`
}
