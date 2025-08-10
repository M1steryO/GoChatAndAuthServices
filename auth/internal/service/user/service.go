package user

import (
	"auth/internal/repository"
	"auth/internal/service"
)

type serv struct {
	db repository.UserRepository
}

func NewUserService(repo repository.UserRepository) service.UserService {
	return &serv{
		db: repo,
	}
}
