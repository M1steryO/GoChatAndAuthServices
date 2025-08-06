package user

import (
	"auth/internal/storage/repository"
	desc "auth/pkg/user_v1"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

var (
	errPasswordNotMatch = errors.New("passwords not match")
)

type Server struct {
	desc.UnimplementedUserV1Server
	Storage *repository.Storage
}

func (s *Server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	log.Printf("Received id %+v", req.GetId())
	user, err := s.Storage.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &desc.GetResponse{
		User: &desc.User{
			Id: req.GetId(),
			Info: &desc.UserInfo{
				Name:  user.Name,
				Email: user.Email,
				Role:  desc.Role(user.Role),
			},
			CreatedAt: timestamppb.New(user.CreatedAt),
			UpdatedAt: func() *timestamppb.Timestamp {
				if user.UpdatedAt != nil {
					return timestamppb.New(*user.UpdatedAt)
				}
				return nil
			}(),
		},
	}, nil
}
func (s *Server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	if req.Password != req.GetPassword() {
		return nil, errPasswordNotMatch
	}
	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to generate password: " + err.Error())
	}

	user := repository.User{
		Name:     req.User.Info.Name,
		Email:    req.User.Info.Email,
		Password: string(password),
		Role:     int(req.User.Info.Role.Number()),
	}
	id, err := s.Storage.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
