package database

const (
	createActivity = "create activity"
	deleteActivity = "delete activity by uuid"
	getActivity    = "get activity by uuid"
	listActivity   = "list activities"
	updateActivity = "update activity by uuid"
)

func queriesActivity() map[string]string {
	return map[string]string{
		createActivity: `INSERT INTO
			activities (fork_id, content_id, name, description, content_type, taxonomy)
			VALUES ($1, $2, $3, $4, $5, $6) RETURNING *`,
		deleteActivity: `UPDATE activities SET deleted_at = NOW() WHERE uuid = $1 AND deleted_at IS NULL`,
		getActivity:    `SELECT * FROM activities WHERE uuid = $1 AND deleted_at IS NULL`,
		listActivity:   `SELECT * FROM activities WHERE deleted_at IS NULL`,
		updateActivity: `UPDATE activities
			SET fork_id = $1, content_id = $2, name = $3, description = $4, content_type = $5, taxonomy = $6
			WHERE uuid = $7 AND deleted_at IS NULL RETURNING *`,
	}
}
