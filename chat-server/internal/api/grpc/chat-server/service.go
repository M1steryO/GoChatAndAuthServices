package chat_server

import (
	"chat-server/internal/service"
	desc "chat-server/pkg/chat_v1"
)

type Implementation struct {
	desc.UnimplementedChatV1Server
	service service.ChatService
}

func NewImplementation(service service.ChatService) *Implementation {
	return &Implementation{
		service: service,
	}
}
