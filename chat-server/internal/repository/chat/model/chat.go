package model

import (
	"database/sql"
	"time"
)

type Chat struct {
	Id        int64        `db:"id"`
	Usernames []string     `db:"usernames"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt sql.NullTime `db:"updated_at"`
}

type Message struct {
	From      string    `db:"from"`
	Text      string    `db:"text"`
	Timestamp time.Time `db:"timestamp"`
}
