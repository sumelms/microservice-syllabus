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
    		activities (code, name, underline, image, image_cover, excerpt, description)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		deleteActivity: "UPDATE activities SET deleted_at = NOW() WHERE uuid = $1",
		getActivity:    "SELECT * FROM activities WHERE uuid = $1",
		listActivity:   "SELECT * FROM activities",
		updateActivity: `UPDATE activities 
			SET code = $1, name = $2, underline = $3, image = $4, image_cover = $5, excerpt = $6, description = $7 
			WHERE uuid = $8 RETURNING *`,
	}
}
