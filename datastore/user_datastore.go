package datastore

import (
	"database/sql"
	"fmt"
	"log"
)

type UserDatastore interface {
	CreateUser(User UserEntity) error
	// GetUser() ([]User, error)
}

type User struct {
	db *sql.DB
}

func NewUser(db *sql.DB) *User {
	return &User{
		db,
	}
}

func (t *User) CreateUser(user UserEntity) error {
	fmt.Println("user datastore")
	stmt, err := t.db.Prepare("INSERT INTO user(id, username, user_role) VALUES (? , ? ,?)")
	if err != nil {
		msg := fmt.Sprintf("Failed to prepare insert into user: %s", err.Error())
		log.Fatal(msg)
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID, user.Username, user.Role)
	if err != nil {
		return fmt.Errorf("failed to execute insert into user: %s", err.Error())
	}

	return nil
}
