package model

import (
	"database/sql"
	"time"
)

type User struct {
	Id        int64
	Info      UserInfo
	Password  string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type UserInfo struct {
	Email string
	Name  string
	Role  string
}
