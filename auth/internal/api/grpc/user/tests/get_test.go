package tests

import (
	"auth/internal/api/grpc/user"
	"auth/internal/model"
	"auth/internal/model/auth"
	"auth/internal/service"
	"auth/internal/service/mocks"
	desc "auth/pkg/user_v1"
	"context"
	"database/sql"
	"fmt"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/timestamppb"
	"testing"
)

func TestGet(t *testing.T) {
	type userServiceMockFunc func(mc *minimock.Controller) service.UserService

	type args struct {
		ctx context.Context
		req *desc.GetRequest
	}
	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id        = gofakeit.Int64()
		name      = gofakeit.Name()
		email     = gofakeit.Email()
		createdAt = gofakeit.Date()
		updatedAt = sql.NullTime{Valid: true, Time: createdAt}
		role      = "ADMIN"

		serviceErr = fmt.Errorf("service error")

		req = &desc.GetRequest{
			Id: id,
		}

		serviceResp = &model.User{
			Id: id,
			Info: auth.UserInfo{
				Name:  name,
				Email: email,
				Role:  role,
			},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		res = &desc.GetResponse{
			User: &desc.User{
				Id: id,
				Info: &desc.UserInfo{
					Name:  name,
					Email: email,
					Role:  desc.Role_ADMIN,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt.Time),
			},
		}
	)
	defer t.Cleanup(mc.Finish)

	tests := []struct {
		name            string
		args            args
		want            *desc.GetResponse
		err             error
		userServiceMock userServiceMockFunc
	}{{
		name: "success case",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: res,
		err:  nil,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			mock := mocks.NewUserServiceMock(mc)
			mock.GetMock.Expect(ctx, id).Return(serviceResp, nil)
			return mock
		},
	}, {
		name: " ",
		args: args{
			ctx: ctx,
			req: req,
		},
		want: nil,
		err:  serviceErr,
		userServiceMock: func(mc *minimock.Controller) service.UserService {
			mock := mocks.NewUserServiceMock(mc)
			mock.GetMock.Expect(ctx, id).Return(nil, serviceErr)
			return mock
		},
	},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			userServiceMock := tt.userServiceMock(mc)
			api := user.NewUserImplementation(userServiceMock)
			resp, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, resp)
		})
	}

}
