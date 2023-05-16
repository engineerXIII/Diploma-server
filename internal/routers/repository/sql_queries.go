package repository

const (
	createRouter = `INSERT INTO hv_router (router_type, router_address, router_hostname, router_port) VALUES ($1, $2, $3, $4) RETURNING *`

	getRouterByID = `SELECT router_type, router_address, router_hostname, router_port	
						FROM hv_router h
						WHERE h.id = $1`
	getRouterList = `SELECT router_type, router_address, router_hostname, router_port	
						FROM hv_router h
						LIMIT 100`
)
