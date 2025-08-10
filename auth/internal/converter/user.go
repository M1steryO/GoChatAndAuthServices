package converter

import (
	"auth/internal/model"
	modelRepo "auth/internal/repository/user/model"
	utils "auth/internal/utils/storage"
	desc "auth/pkg/user_v1"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateUserModelFromApi(req *desc.CreateRequest) *model.CreateUserModel {
	return &model.CreateUserModel{
		Info:            ToUserInfoFromApi(req),
		Password:        req.GetPassword(),
		ConfirmPassword: req.GetPasswordConfirm(),
	}
}

func ToUserInfoFromApi(req *desc.CreateRequest) model.UserInfo {
	return model.UserInfo{
		Name:  req.Info.Name,
		Email: req.Info.Email,
		Role:  req.Info.Role.String(),
	}
}

func ToUserCreateRepoFromService(user *model.CreateUserModel) *modelRepo.User {
	return &modelRepo.User{
		Info:     ToUserInfoRepoFromService(user.Info),
		Password: user.Password,
	}
}

func ToUserInfoRepoFromService(user model.UserInfo) modelRepo.UserInfo {
	return modelRepo.UserInfo{
		Name:  user.Name,
		Email: user.Email,
		Role:  user.Role,
	}
}

func ToUserFromService(user *model.User) *desc.User {
	fmt.Println(user.UpdatedAt)
	return &desc.User{
		Id: user.Id,
		Info: &desc.UserInfo{
			Name:  user.Info.Name,
			Email: user.Info.Email,
			Role:  desc.Role(utils.GetRoleIdByStr(user.Info.Role)),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: func() *timestamppb.Timestamp {
			var updatedAt *timestamppb.Timestamp
			if user.UpdatedAt.Valid {
				updatedAt = timestamppb.New(user.UpdatedAt.Time)
			}
			return updatedAt
		}(),
	}
}
