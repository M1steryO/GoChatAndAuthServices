package service

import (
	"auth/internal/model"
	"context"
)

type UserService interface {
	Get(ctx context.Context, id int64) (*model.User, error)
	Create(ctx context.Context, user *model.CreateUserModel) (int64, error)
}
