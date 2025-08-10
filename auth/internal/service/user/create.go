package user

import (
	"auth/internal/converter"
	"auth/internal/model"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var errPasswordNotMatch = errors.New("password does not match")

func (s *serv) Create(ctx context.Context, user *model.CreateUserModel) (int64, error) {
	if user.Password != user.ConfirmPassword {
		return 0, errPasswordNotMatch
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to generate password: " + err.Error())
	}

	user.Password = string(password)

	id, err := s.db.Create(ctx, converter.ToUserCreateRepoFromService(user))
	if err != nil {
		return 0, err
	}
	return id, nil
}
