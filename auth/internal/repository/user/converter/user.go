package converter

import (
	"auth/internal/model"
	modelRepo "auth/internal/repository/user/model"
)

func ToUserInfoFromRepo(user modelRepo.UserInfo) model.UserInfo {
	return model.UserInfo{
		Name:     user.Name,
		Username: user.Username,
		Role:     user.Role,
	}
}

func ToUserFromRepo(user *modelRepo.User) *model.User {
	return &model.User{
		Id:        user.Id,
		Info:      ToUserInfoFromRepo(user.Info),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
