package models

import "github.com/google/uuid"

type Router struct {
	RouterID uuid.UUID `json:"router_id" db:"id" validate:"required,uuid"`
	Type     string    `json:"router_type" db:"router_type" validate:"required"`
	Address  string    `json:"router_address" db:"router_address" validate:"required"`
	Hostname string    `json:"router_hostname" db:"router_hostname" validate:"required"`
	Port     string    `json:"router_port" db:"router_port"`
}

type RouterList struct {
	TotalCount int       `json:"total_count"`
	TotalPages int       `json:"total_pages"`
	Page       int       `json:"page"`
	Size       int       `json:"size"`
	HasMore    bool      `json:"has_more"`
	Routers    []*Router `json:"routers"`
}
