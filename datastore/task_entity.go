package datastore

import (
	"time"

	"github.com/google/uuid"
)

type TaskEntity struct {
	ID          uuid.UUID  `json:"id"`
	UserId      uuid.UUID  `json:"userId"`
	Deleted     bool       `json:"deleted"`
	Summary     string     `json:"summary"`
	PerformedIn *time.Time `json:"performed_in"`
}
