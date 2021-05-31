// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (first_name, last_name, user_name,  email, password, phone_no)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING id, first_name, last_name, user_name, email, password, phone_no, ts
`

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	UserName  string `json:"user_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	PhoneNo   int32  `json:"phone_no"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (Users, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.FirstName,
		arg.LastName,
		arg.UserName,
		arg.Email,
		arg.Password,
		arg.PhoneNo,
	)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.PhoneNo,
		&i.Ts,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, user_name, email, password, phone_no, ts FROM users
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.PhoneNo,
		&i.Ts,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, first_name, last_name, user_name, email, password, phone_no, ts FROM users
WHERE email = $1
LIMIT 1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (Users, error) {
	row := q.db.QueryRowContext(ctx, getUserByEmail, email)
	var i Users
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.UserName,
		&i.Email,
		&i.Password,
		&i.PhoneNo,
		&i.Ts,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, first_name, last_name, user_name, email, password, phone_no, ts FROM users
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]Users, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Users{}
	for rows.Next() {
		var i Users
		if err := rows.Scan(
			&i.ID,
			&i.FirstName,
			&i.LastName,
			&i.UserName,
			&i.Email,
			&i.Password,
			&i.PhoneNo,
			&i.Ts,
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

const updateEmail = `-- name: UpdateEmail :exec
UPDATE users
SET email = $2
WHERE id = $1
`

type UpdateEmailParams struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
}

func (q *Queries) UpdateEmail(ctx context.Context, arg UpdateEmailParams) error {
	_, err := q.db.ExecContext(ctx, updateEmail, arg.ID, arg.Email)
	return err
}

const updateFirstName = `-- name: UpdateFirstName :exec
UPDATE users
SET first_name = $2
WHERE id = $1
`

type UpdateFirstNameParams struct {
	ID        int64  `json:"id"`
	FirstName string `json:"first_name"`
}

func (q *Queries) UpdateFirstName(ctx context.Context, arg UpdateFirstNameParams) error {
	_, err := q.db.ExecContext(ctx, updateFirstName, arg.ID, arg.FirstName)
	return err
}

const updateLastName = `-- name: UpdateLastName :exec
UPDATE users
SET last_name = $2
WHERE id = $1
`

type UpdateLastNameParams struct {
	ID       int64  `json:"id"`
	LastName string `json:"last_name"`
}

func (q *Queries) UpdateLastName(ctx context.Context, arg UpdateLastNameParams) error {
	_, err := q.db.ExecContext(ctx, updateLastName, arg.ID, arg.LastName)
	return err
}

const updatePassword = `-- name: UpdatePassword :exec
UPDATE users
SET password = $2
WHERE id = $1
`

type UpdatePasswordParams struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.ID, arg.Password)
	return err
}

const updatePhoneNo = `-- name: UpdatePhoneNo :exec
UPDATE users
SET phone_no = $2
WHERE id = $1
`

type UpdatePhoneNoParams struct {
	ID      int64 `json:"id"`
	PhoneNo int32 `json:"phone_no"`
}

func (q *Queries) UpdatePhoneNo(ctx context.Context, arg UpdatePhoneNoParams) error {
	_, err := q.db.ExecContext(ctx, updatePhoneNo, arg.ID, arg.PhoneNo)
	return err
}

const updateUserName = `-- name: UpdateUserName :exec
UPDATE users
SET user_name = $2
WHERE id = $1
`

type UpdateUserNameParams struct {
	ID       int64  `json:"id"`
	UserName string `json:"user_name"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) error {
	_, err := q.db.ExecContext(ctx, updateUserName, arg.ID, arg.UserName)
	return err
}
