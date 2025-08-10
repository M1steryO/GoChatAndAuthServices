package service

import (
	"chat-server/internal/model"
	"context"
)

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
}
