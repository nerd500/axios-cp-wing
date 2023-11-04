// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package database

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID
	CreatedBy string
	CreatedAt time.Time
	Title     string
	Link      string
	Tags      []string
	Platform  sql.NullString
}

type User struct {
	ID             uuid.UUID
	Email          string
	HashedPassword string
	Salt           string
	FirstName      string
	LastName       string
	AuthToken      string
	IsAdminUser    bool
}
