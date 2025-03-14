package repositories

import (
	"log"

	"go-todo/db"
	"go-todo/models"
)

func GetAllTasks() ([]models.Task, error) {
	rows, err := db.DB.Query("SELECT id, title, description, status, created_at, updated_at FROM tasks")
	if err != nil {
		log.Println("Query Error:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.CreatedAt, &task.UpdatedAt); err != nil {
			log.Println("Row Scan Error:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}
