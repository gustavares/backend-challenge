package datastore

import (
	"github.com/google/uuid"
)

type UserEntity struct {
	ID       uuid.UUID `json:"id,omitempty"`
	Username string    `json:"username"`
	Role     string    `json:"userRole"`
	Deleted  bool      `json:"deleted,omitempty"`
}
