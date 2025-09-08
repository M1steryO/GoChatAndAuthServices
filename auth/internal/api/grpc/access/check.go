package access

import (
	jwtUtils "auth/internal/utils/jwt"
	descAccess "auth/pkg/access_v1"
	"context"
	"errors"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"strings"
)

const jwtPrefix = "Bearer "
const accessTokenSecretKey = "W4/X+LLjehdxptt4YgGFCvMpq5ewptpZZYRHY6A72g01"

var accessibleMap = map[string]string{
	"auth/User/v1/Get": "ADMIN",
}

func (i *Implementation) Check(ctx context.Context, req *descAccess.CheckRequest) (*emptypb.Empty, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("metadata is not provided")
	}
	authHeader, ok := md["authorization"]
	if !ok || len(authHeader) == 0 {
		return nil, errors.New("authorization header is not provided")
	}
	if !strings.HasPrefix(authHeader[0], jwtPrefix) {
		return nil, errors.New("invalid authorization header")
	}

	accessToken := strings.TrimPrefix(authHeader[0], jwtPrefix)

	claims, err := jwtUtils.VerifyToken(accessToken, []byte(accessTokenSecretKey))
	if err != nil {
		return nil, errors.New("invalid access token")
	}

	role, ok := accessibleMap[req.GetEndpointAddress()]

	if !ok {
		return &emptypb.Empty{}, nil
	}

	if role != claims.Role {
		return nil, errors.New("invalid access")
	}

	return nil, errors.New("access denied")
}
