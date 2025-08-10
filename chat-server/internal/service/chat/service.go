package chat

import (
	"chat-server/internal/repository"
	"chat-server/internal/service"
)

type serv struct {
	db repository.ChatRepository
}

func NewChatService(db repository.ChatRepository) service.ChatService {
	return &serv{
		db: db,
	}
}
