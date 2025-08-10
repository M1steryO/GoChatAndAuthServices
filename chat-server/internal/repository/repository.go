package repository

import (
	"chat-server/internal/repository/chat/model"
	"context"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
}
