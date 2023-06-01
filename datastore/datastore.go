package datastore

import (
	"database/sql"
	"fmt"
	"tasks/config"

	_ "github.com/go-sql-driver/mysql"
)

type Datastore struct {
	TaskDatastore TaskDatastore
	UserDatastore UserDatastore
}

func New(c *config.Config) *Datastore {
	connString := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", c.Env.DbUser, c.Env.DbPassword, c.Env.DbHost, c.Env.DbName)
	db, err := sql.Open("mysql", connString)
	if err != nil {
		// TODO: better error handling
		panic(err)
	}

	return &Datastore{
		TaskDatastore: NewTask(db),
		UserDatastore: NewUser(db),
	}
}
