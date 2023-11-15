// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: tasks.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
    id,
    created_by,
    created_at,
    last_edited_by,
    last_edited_at,
    title,
    link,
    platform

) VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8
) RETURNING id, created_by, created_at, last_edited_by, last_edited_at, title, link, platform
`

type CreateTaskParams struct {
	ID           uuid.UUID
	CreatedBy    uuid.UUID
	CreatedAt    time.Time
	LastEditedBy uuid.UUID
	LastEditedAt time.Time
	Title        string
	Link         string
	Platform     uuid.UUID
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ID,
		arg.CreatedBy,
		arg.CreatedAt,
		arg.LastEditedBy,
		arg.LastEditedAt,
		arg.Title,
		arg.Link,
		arg.Platform,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedBy,
		&i.CreatedAt,
		&i.LastEditedBy,
		&i.LastEditedAt,
		&i.Title,
		&i.Link,
		&i.Platform,
	)
	return i, err
}

const getAllTasks = `-- name: GetAllTasks :many
SELECT id, created_by, created_at, last_edited_by, last_edited_at, title, link, platform FROM tasks
`

func (q *Queries) GetAllTasks(ctx context.Context) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, getAllTasks)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.CreatedBy,
			&i.CreatedAt,
			&i.LastEditedBy,
			&i.LastEditedAt,
			&i.Title,
			&i.Link,
			&i.Platform,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
