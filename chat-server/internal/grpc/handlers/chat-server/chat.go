package chat_server

import (
	"chat-server/internal/storage/repository"
	desc "chat-server/pkg/chat_v1"
	"context"
	"errors"
)

var (
	errPasswordNotMatch = errors.New("passwords not match")
)

type Server struct {
	desc.UnimplementedChatV1Server
	Storage *repository.Storage
}

func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	chat := &repository.Chat{
		Usernames: req.Chat.Usernames,
	}
	id, err := s.Storage.CreateChat(ctx, chat)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
