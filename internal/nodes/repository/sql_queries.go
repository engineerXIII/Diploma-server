package repository

const (
	getNodeByID = `SELECT id, node_address, node_name, hv_type	
						FROM hv_node h
						WHERE h.id = $1`
	getNodeList = `SELECT id, node_address, node_name, hv_type	
						FROM hv_node h
						LIMIT 100`
)
