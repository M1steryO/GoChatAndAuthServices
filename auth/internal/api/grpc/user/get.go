package user

import (
	"auth/internal/converter"
	desc "auth/pkg/user_v1"
	"context"
	"log"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received id %+v", req.GetId())
	user, err := i.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
