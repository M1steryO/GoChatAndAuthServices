package adapters

import (
	"chat-server/internal/converter"
	"chat-server/internal/model"
	"chat-server/internal/service"
	desc "chat-server/pkg/chat_v1"
	"context"
)

type grpcStream struct {
	stream desc.ChatV1_ConnectChatServer
}

func NewGRPCStream(stream desc.ChatV1_ConnectChatServer) service.Stream {
	return &grpcStream{
		stream: stream,
	}
}

func (s *grpcStream) Send(msg *model.Message) error {
	return s.stream.Send(converter.ToMessageApiFromService(msg))
}

func (s *grpcStream) Context() context.Context {
	return s.stream.Context()
}
