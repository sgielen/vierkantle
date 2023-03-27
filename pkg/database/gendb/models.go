// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package gendb

import (
	"database/sql"
	"time"
)

type VierkantleScore struct {
	BoardName   string
	AnonymousID int64
	UserID      sql.NullInt64
	TeamSize    int32
	Words       int32
	Seconds     int32
	StartedAt   time.Time
	UpdatedAt   time.Time
}

type VierkantleUser struct {
	ID          int64
	Username    string
	Email       sql.NullString
	LastLoginAt sql.NullTime
}
