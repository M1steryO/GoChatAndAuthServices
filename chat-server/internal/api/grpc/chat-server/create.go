package chat_server

import (
	"chat-server/internal/converter"
	desc "chat-server/pkg/chat_v1"
	"context"
)

func (i *Implementation) Create(ctx context.Context, request *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.service.Create(ctx, converter.ToChatServiceFromApi(request.GetChat()))
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
