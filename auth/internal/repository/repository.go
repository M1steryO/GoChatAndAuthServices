package repository

import (
	modelRepo "auth/internal/repository/user/model"
	"context"
)

type UserRepository interface {
	Create(ctx context.Context, user *modelRepo.User) (int64, error)
	Get(ctx context.Context, id int64) (*modelRepo.User, error)
	CreateLog(ctx context.Context, userId int64, action string) error
}
