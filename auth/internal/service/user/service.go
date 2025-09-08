package user

import (
	"auth/internal/client/db"
	"auth/internal/repository"
	"auth/internal/service"
)

type serv struct {
	db        repository.UserRepository
	txManager db.TxManager
}

func NewUserService(repo repository.UserRepository, txManager db.TxManager) service.UserService {
	return &serv{
		db:        repo,
		txManager: txManager,
	}
}
