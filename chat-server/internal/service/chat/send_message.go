package chat

import (
	"chat-server/internal/model"
	"chat-server/internal/repository/chat/converter"
	"context"
)

func (s *serv) SendMessage(ctx context.Context, chatId int64, message *model.Message) error {
	_, err := s.Get(ctx, chatId)
	if err != nil {
		return err
	}

	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := s.db.CreateMessage(ctx, chatId, converter.ToMessageRepoFromService(message))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	s.channels[chatId] <- message
	return nil
}
