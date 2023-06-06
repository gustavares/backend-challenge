package datastore

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type TaskDatastore interface {
	CreateTask(TaskEntity) error
	GetTasksFromUser(userId string) ([]TaskEntity, error)
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

func (t *Task) GetTasksFromUser(userId string) ([]TaskEntity, error) {
	rows, err := t.db.Query("SELECT id, user_id, summary, performed_in FROM task WHERE user_id=? AND deleted=FALSE", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []TaskEntity
	for rows.Next() {
		var task TaskEntity

		if err := rows.Scan(&task.ID, &task.UserId, &task.Summary, &task.PerformedIn); err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}
