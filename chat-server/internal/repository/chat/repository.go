package chat

import (
	"chat-server/internal/repository"
	"chat-server/internal/repository/chat/model"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

type repo struct {
	db *pgxpool.Pool
}

func NewChatRepository(db *pgxpool.Pool) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (s *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	var lastInsertId int64
	tx, err := s.db.Begin(ctx)

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
