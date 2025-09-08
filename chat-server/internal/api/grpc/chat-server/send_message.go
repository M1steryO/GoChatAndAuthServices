package chat_server

import (
	"chat-server/internal/converter"
	desc "chat-server/pkg/chat_v1"
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) SendMessage(ctx context.Context, req *desc.SendMessageRequest) (*emptypb.Empty, error) {

	err := i.service.SendMessage(ctx, req.GetChatId(), converter.ToMessageServiceFromApi(req.GetMessage()))
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
