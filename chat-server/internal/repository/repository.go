package repository

import (
	"chat-server/internal/repository/chat/model"
	"context"
)

type ChatRepository interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	CreateMessage(ctx context.Context, chatID int64, message *model.Message) error
	Get(ctx context.Context, chatId int64) (*model.Chat, error)
}
