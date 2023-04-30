// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: queries.sql

package gendb

import (
	"context"
	"database/sql"
)

const getMyScore = `-- name: GetMyScore :one
SELECT my_score.team_size, my_score.words, my_score.seconds, (
	SELECT COUNT(their_score.ctid) FROM vierkantle.scores their_score
	WHERE their_score.board_name=$1
	AND (their_score.words > my_score.words
	  OR (their_score.words = my_score.words AND their_score.seconds < my_score.seconds)
	  OR (their_score.words = my_score.words AND their_score.seconds = my_score.seconds AND their_score.ctid < my_score.ctid)
	)
) AS rank
FROM vierkantle.scores my_score
WHERE my_score.board_name=$1 AND my_score.anonymous_id=$2
`

type GetMyScoreParams struct {
	BoardName   string
	AnonymousID int64
}

type GetMyScoreRow struct {
	TeamSize int32
	Words    int32
	Seconds  int32
	Rank     int64
}

func (q *Queries) GetMyScore(ctx context.Context, arg GetMyScoreParams) (GetMyScoreRow, error) {
	row := q.db.QueryRow(ctx, getMyScore, arg.BoardName, arg.AnonymousID)
	var i GetMyScoreRow
	err := row.Scan(
		&i.TeamSize,
		&i.Words,
		&i.Seconds,
		&i.Rank,
	)
	return i, err
}

const getScores = `-- name: GetScores :many
SELECT anonymous_id, users.username, team_size, words, seconds, COUNT(*) OVER() AS full_count
FROM vierkantle.scores
LEFT JOIN vierkantle.users ON users.id=scores.user_id
WHERE board_name=$1
ORDER BY words DESC, seconds ASC, scores.ctid ASC
LIMIT $2 OFFSET $3
`

type GetScoresParams struct {
	BoardName string
	Limit     int32
	Offset    int32
}

type GetScoresRow struct {
	AnonymousID int64
	Username    sql.NullString
	TeamSize    int32
	Words       int32
	Seconds     int32
	FullCount   int64
}

func (q *Queries) GetScores(ctx context.Context, arg GetScoresParams) ([]GetScoresRow, error) {
	rows, err := q.db.Query(ctx, getScores, arg.BoardName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetScoresRow
	for rows.Next() {
		var i GetScoresRow
		if err := rows.Scan(
			&i.AnonymousID,
			&i.Username,
			&i.TeamSize,
			&i.Words,
			&i.Seconds,
			&i.FullCount,
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

const getUsersWithEmailOrName = `-- name: GetUsersWithEmailOrName :many
SELECT id, username, email FROM vierkantle.users WHERE username=$1 OR email=$2
`

type GetUsersWithEmailOrNameParams struct {
	Username string
	Email    sql.NullString
}

type GetUsersWithEmailOrNameRow struct {
	ID       int64
	Username string
	Email    sql.NullString
}

func (q *Queries) GetUsersWithEmailOrName(ctx context.Context, arg GetUsersWithEmailOrNameParams) ([]GetUsersWithEmailOrNameRow, error) {
	rows, err := q.db.Query(ctx, getUsersWithEmailOrName, arg.Username, arg.Email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetUsersWithEmailOrNameRow
	for rows.Next() {
		var i GetUsersWithEmailOrNameRow
		if err := rows.Scan(&i.ID, &i.Username, &i.Email); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const loginUser = `-- name: LoginUser :one
UPDATE vierkantle.users SET last_login_at=NOW() WHERE id=$1 RETURNING username
`

func (q *Queries) LoginUser(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRow(ctx, loginUser, id)
	var username string
	err := row.Scan(&username)
	return username, err
}

const registerUser = `-- name: RegisterUser :one
INSERT INTO vierkantle.users (username, email, registered_at, last_login_at)
VALUES ($1, $2, NOW(), NULL) RETURNING id
`

type RegisterUserParams struct {
	Username string
	Email    sql.NullString
}

func (q *Queries) RegisterUser(ctx context.Context, arg RegisterUserParams) (int64, error) {
	row := q.db.QueryRow(ctx, registerUser, arg.Username, arg.Email)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const setScore = `-- name: SetScore :exec
INSERT INTO vierkantle.scores (board_name, anonymous_id, user_id, team_size, words, seconds, started_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW())
ON CONFLICT (board_name, anonymous_id) DO UPDATE SET user_id=$3, team_size=$4, words=$5, seconds=$6, updated_at=NOW()
`

type SetScoreParams struct {
	BoardName   string
	AnonymousID int64
	UserID      sql.NullInt64
	TeamSize    int32
	Words       int32
	Seconds     int32
}

func (q *Queries) SetScore(ctx context.Context, arg SetScoreParams) error {
	_, err := q.db.Exec(ctx, setScore,
		arg.BoardName,
		arg.AnonymousID,
		arg.UserID,
		arg.TeamSize,
		arg.Words,
		arg.Seconds,
	)
	return err
}
