package database

const (
	createLesson = "create lesson"
	deleteLesson = "delete lesson by uuid"
	getLesson    = "get lesson by uuid"
	listLesson   = "list lesson"
	updateLesson = "update lesson by uuid"
)

func queriesLesson() map[string]string {
	return map[string]string{
		createLesson: `INSERT INTO 
    		lessons (name, description, objective, type, module)
			VALUES (:name, :description, :objective, :type, :module) RETURNING *`,
		deleteLesson: "UPDATE lessons SET deleted_at = NOW() WHERE uuid = :uuid",
		getLesson:    "SELECT * FROM lessons WHERE uuid = :uuid",
		listLesson:   "SELECT * FROM lessons",
		updateLesson: `UPDATE lessons 
			SET name = :name, description = :description, objective = :objective, type = :type, module = :module 
			WHERE uuid = :uuid RETURNING *`,
	}
}
