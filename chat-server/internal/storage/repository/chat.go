package repository

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type Chat struct {
	Id        int64
	Usernames []string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
type Storage struct {
	Pool *pgxpool.Pool
}

func (s *Storage) CreateChat(ctx context.Context, chat *Chat) (int64, error) {
	var lastInsertId int64
	tx, err := s.Pool.Begin(ctx)

	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()

	err = tx.QueryRow(ctx,
		`INSERT INTO "chat" DEFAULT VALUES 
			 RETURNING id;`).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	for _, username := range chat.Usernames {
		_, err = tx.Exec(ctx,
			`INSERT INTO "chat_member" (chat_id, username) VALUES ($1, $2)`,
			lastInsertId, username,
		)
		if err != nil {
			return 0, err
		}
	}

	return lastInsertId, nil

}
