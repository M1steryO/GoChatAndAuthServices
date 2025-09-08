package service

import (
	"chat-server/internal/model"
	"context"
)

type Stream interface {
	Send(*model.Message) error
	Context() context.Context
}

type ChatService interface {
	Create(ctx context.Context, chat *model.Chat) (int64, error)
	SendMessage(ctx context.Context, chatId int64, message *model.Message) error
	Get(ctx context.Context, chatId int64) (*model.Chat, error)
	ConnectChat(ctx context.Context, chatId int64, username string, stream Stream) error
}
