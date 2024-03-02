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
			activities (content_uuid, name, description, content_type, taxonomy)
			VALUES ($1, $2, $3, $4, $5) RETURNING *`,
		deleteActivity: `UPDATE activities SET deleted_at = NOW() WHERE uuid = $1 AND deleted_at IS NULL`,
		getActivity:    `SELECT * FROM activities WHERE uuid = $1 AND deleted_at IS NULL`,
		listActivity:   `SELECT * FROM activities WHERE deleted_at IS NULL`,
		updateActivity: `UPDATE activities
			SET content_uuid = $1, name = $2, description = $3, content_type = $4, taxonomy = $5
			WHERE uuid = $6 AND deleted_at IS NULL RETURNING *`,
	}
}
