package user

import (
	"auth/internal/converter"
	"auth/internal/logger"
	desc "auth/pkg/user_v1"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (i *Implementation) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	if req.Id == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument")
	}
	logger.Info("Received", slog.Int64("id:", req.GetId()))
	user, err := i.service.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &desc.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
