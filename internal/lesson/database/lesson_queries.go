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
    		lessons (code, name, underline, image, image_cover, excerpt, description)
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *`,
		deleteLesson: "UPDATE lessons SET deleted_at = NOW() WHERE uuid = $1",
		getLesson:    "SELECT * FROM lessons WHERE uuid = $1",
		listLesson:   "SELECT * FROM lessons",
		updateLesson: `UPDATE lessons 
			SET code = $1, name = $2, underline = $3, image = $4, image_cover = $5, excerpt = $6, description = $7 
			WHERE uuid = $8 RETURNING *`,
	}
}
