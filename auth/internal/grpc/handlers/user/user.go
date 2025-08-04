package user

import (
	desc "auth/pkg/user_v1"
	"context"
	"github.com/brianvoe/gofakeit/v7"
	"log"
)

type Server struct {
	desc.UnimplementedUserV1Server
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received id %+v", req.GetId())
	return &desc.GetResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
				Role:  desc.Role_ADMIN,
			},
		},
	}, nil
}
