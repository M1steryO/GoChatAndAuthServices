package chat

import (
	"chat-server/internal/converter"
	"chat-server/internal/model"
	"context"
)

func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {

	var chatId int64
	err := s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, err := s.db.Create(ctx, converter.ToChatRepoFromService(chat))
		chatId = id

		return err
	})
	s.mxChannel.Lock()
	
	if _, ok := s.channels[chatId]; !ok {
		s.channels[chatId] = make(chan *model.Message, 100)
	}
	s.mxChannel.Unlock()
	if err != nil {
		return 0, err
	}
	return chatId, nil
}
