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
    		activities (name, description, content_id, content_type)
			VALUES (:name, :description, :content_id, :content_type) RETURNING *`,
		deleteActivity: "UPDATE activities SET deleted_at = NOW() WHERE uuid = :uuid",
		getActivity:    "SELECT * FROM activities WHERE uuid = :uuid",
		listActivity:   "SELECT * FROM activities",
		updateActivity: `UPDATE activities 
			SET name = :name, description = :description, content_id = :content_id, content_type = :content_type 
			WHERE uuid = :uuid RETURNING *`,
	}
}
