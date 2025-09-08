package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64        `db:"id"`
	Info      UserInfo     `db:""`
	Password  string       `db:"password"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type UserInfo struct {
	Username string `db:"username"`
	Name     string `db:"name"`
	Role     string `db:"role"`
}
