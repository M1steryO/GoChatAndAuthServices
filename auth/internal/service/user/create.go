package user

import (
	"auth/internal/converter"
	"auth/internal/model"
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
)

var ErrPasswordNotMatch = errors.New("password does not match")

func (s *serv) Create(ctx context.Context, user *model.CreateUserModel) (int64, error) {
	if user.Password != user.ConfirmPassword {
		return 0, ErrPasswordNotMatch
	}
	password, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, errors.New("failed to generate password: " + err.Error())
	}

	user.Password = string(password)
	var id int64
	err = s.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, err = s.db.Create(ctx, converter.ToUserCreateRepoFromService(user))
		if err != nil {
			return err
		}
		err = s.db.CreateLog(ctx, id, "create_account")
		return err
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}
