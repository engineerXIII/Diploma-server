package models

type NodeAuth struct {
	HeaderName  string `json:"header_name" db:"header_name"`
	SecretName  string `json:"secret_name" db:"secret_name"`
	SecretValue string `json:"secret_value" db:"secret_value`
}

type Node struct {
	NodeID  int64    `json:"node_id,omit_empty" db:"id"`
	Address string   `json:"address" db:"node_address"`
	Name    string   `json:"name" db:"node_name"`
	Type    string   `json:"type" db:"hv_type"`
	Auth    NodeAuth `json:"auth"`
}

type NodeList struct {
	Size  int    `json:"size"`
	Nodes []Node `json:"nodes"`
}
