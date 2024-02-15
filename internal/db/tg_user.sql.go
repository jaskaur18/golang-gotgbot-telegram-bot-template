// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tg_user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkUserExist = `-- name: CheckUserExist :one
SELECT EXISTS(
  SELECT 1 FROM users
  WHERE telegram_id = $1
) AS exist
`

func (q *Queries) CheckUserExist(ctx context.Context, telegramID pgtype.Int8) (bool, error) {
	row := q.db.QueryRow(ctx, checkUserExist, telegramID)
	var exist bool
	err := row.Scan(&exist)
	return exist, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  telegram_id, first_name, last_name, username, language, user_type
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id
`

type CreateUserParams struct {
	TelegramID pgtype.Int8 `json:"telegram_id"`
	FirstName  string      `json:"first_name"`
	LastName   pgtype.Text `json:"last_name"`
	Username   pgtype.Text `json:"username"`
	Language   pgtype.Text `json:"language"`
	UserType   Usertype    `json:"user_type"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (pgtype.UUID, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.TelegramID,
		arg.FirstName,
		arg.LastName,
		arg.Username,
		arg.Language,
		arg.UserType,
	)
	var id pgtype.UUID
	err := row.Scan(&id)
	return id, err
}

const getUserByTelegramID = `-- name: GetUserByTelegramID :one
SELECT id, telegram_id, first_name, last_name, username, language, user_type, created_at, updated_at FROM users
WHERE telegram_id = $1 LIMIT 1
`

func (q *Queries) GetUserByTelegramID(ctx context.Context, telegramID pgtype.Int8) (User, error) {
	row := q.db.QueryRow(ctx, getUserByTelegramID, telegramID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.TelegramID,
		&i.FirstName,
		&i.LastName,
		&i.Username,
		&i.Language,
		&i.UserType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, telegram_id, first_name, last_name, username, language, user_type, created_at, updated_at FROM users
ORDER BY created_at DESC
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.TelegramID,
			&i.FirstName,
			&i.LastName,
			&i.Username,
			&i.Language,
			&i.UserType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserType = `-- name: UpdateUserType :exec
UPDATE users
SET user_type = $2
WHERE telegram_id = $1
RETURNING id
`

type UpdateUserTypeParams struct {
	TelegramID pgtype.Int8 `json:"telegram_id"`
	UserType   Usertype    `json:"user_type"`
}

func (q *Queries) UpdateUserType(ctx context.Context, arg UpdateUserTypeParams) error {
	_, err := q.db.Exec(ctx, updateUserType, arg.TelegramID, arg.UserType)
	return err
}
