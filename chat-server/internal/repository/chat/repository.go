package chat

import (
	"chat-server/internal/repository"
	"chat-server/internal/repository/chat/model"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/M1steryO/platform_common/pkg/db"
	"strings"
)

var ChatNotFoundError = errors.New("chat not found")

type repo struct {
	db db.Client
}

func NewChatRepository(db db.Client) repository.ChatRepository {
	return &repo{
		db: db,
	}
}

func (s *repo) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	var lastInsertId int64

	err := s.db.DB().QueryRowContext(ctx, db.Query{
		Title: "Create Chat",
		Query: `INSERT INTO "chat" DEFAULT VALUES
			 RETURNING id;`,
	}).Scan(&lastInsertId)
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO chat_member (chat_id, username) VALUES `
	args := []interface{}{lastInsertId}
	var placeholders []string

	for i, username := range chat.Usernames {
		args = append(args, username)
		placeholders = append(placeholders, fmt.Sprintf("($1, $%d)", i+2))
	}

	_, err = s.db.DB().ExecContext(ctx, db.Query{
		Title: "Create chat memebers dependency",
		Query: query + strings.Join(placeholders, ","),
	}, args...)
	if err != nil {
		return 0, err
	}

	return lastInsertId, nil

}

func (s *repo) Get(ctx context.Context, chatId int64) (*model.Chat, error) {
	var msg model.Chat
	err := s.db.DB().ScanOneContext(ctx, &msg, db.Query{
		Title: "Get chat",
		Query: "SELECT id, created_at, updated_at FROM chat WHERE id = $1 ",
	}, chatId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ChatNotFoundError
		}
		return nil, err
	}
	return &msg, nil
}

func (s *repo) CreateMessage(ctx context.Context, chatId int64, msg *model.Message) error {
	_, err := s.db.DB().ExecContext(ctx, db.Query{
		Title: "Create message",
		Query: `INSERT INTO message (chat_id, "from", text, timestamp) VALUES ($1, $2, $3, $4)`,
	}, chatId, msg.From, msg.Text, msg.Timestamp)
	if err != nil {
		return err
	}
	return nil
}
