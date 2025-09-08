package model

import (
	"database/sql"
	"time"
)

type Chat struct {
	Id        int64
	Usernames []string
	CreatedAt time.Time
	UpdatedAt sql.NullTime
}

type Message struct {
	From      string
	Text      string
	Timestamp time.Time
}
