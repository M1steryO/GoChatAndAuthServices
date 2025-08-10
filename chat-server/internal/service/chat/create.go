package chat

import (
	"chat-server/internal/converter"
	"chat-server/internal/model"
	"context"
)

func (s *serv) Create(ctx context.Context, chat *model.Chat) (int64, error) {
	id, err := s.db.Create(ctx, converter.ToChatRepoFromService(chat))
	if err != nil {
		return 0, err
	}
	return id, nil
}
