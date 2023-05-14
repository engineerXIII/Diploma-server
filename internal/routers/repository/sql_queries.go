package repository

const (
	createRouter = `INSERT INTO hv_router (router_type, router_address, router_hostname, router_port) VALUES ($1, $2, $3, $4) RETURNING *`

	updateComment = `UPDATE comments SET message = $1, updated_at = CURRENT_TIMESTAMP WHERE comment_id = $2 RETURNING *`

	deleteComment = `DELETE FROM comments WHERE comment_id = $1`

	getRouterByID = `SELECT router_type, router_address, router_hostname, router_port	
						FROM hv_router h
						WHERE h.id = $1`

	getTotalCountByNewsID = `SELECT COUNT(comment_id) FROM comments WHERE news_id = $1`
)
