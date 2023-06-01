package datastore

import (
	"database/sql"
	"log"
)

type TaskDatastore interface {
	CreateTask(TaskEntity) error
	// GetTasksByUser() ([]TaskEntity, error)
}

type Task struct {
	db *sql.DB
}

func NewTask(db *sql.DB) *Task {
	return &Task{
		db,
	}
}

func (t *Task) CreateTask(task TaskEntity) error {
	stmt, err := t.db.Prepare("INSERT INTO task(id, user_id, summary) VALUES (? , ? ,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(task.ID, task.UserId, task.Summary)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
