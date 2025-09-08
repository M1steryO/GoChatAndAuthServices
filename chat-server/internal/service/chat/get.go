package chat

import (
	"chat-server/internal/model"
	"chat-server/internal/repository/chat/converter"
	"context"
)

func (s *serv) Get(ctx context.Context, chatId int64) (*model.Chat, error) {
	chat, err := s.db.Get(ctx, chatId)
	if err != nil {
		return nil, err
	}
	return converter.ToChatFromRepo(chat), nil
}
