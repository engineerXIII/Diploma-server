package models

type Router struct {
	RouterID int64  `json:"router_id" db:"id" validate:"omitempty,uuid"`
	Type     string `json:"router_type" db:"router_type" validate:"required"`
	Address  string `json:"router_address" db:"router_address" validate:"required"`
	Hostname string `json:"router_hostname" db:"router_hostname" validate:"required"`
	Port     int64  `json:"router_port" db:"router_port"`
}

type RouterList struct {
	Size    int      `json:"size"`
	Routers []Router `json:"routers"`
}
