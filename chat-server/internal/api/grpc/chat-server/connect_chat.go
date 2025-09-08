package chat_server

import (
	"chat-server/internal/api/grpc/adapters"
	desc "chat-server/pkg/chat_v1"
)

func (i *Implementation) ConnectChat(req *desc.ConnectChatRequest, stream desc.ChatV1_ConnectChatServer) error {
	err := i.service.ConnectChat(stream.Context(), req.GetChatId(), req.GetUsername(), adapters.NewGRPCStream(stream))
	if err != nil {
		return err
	}
	return nil
}
