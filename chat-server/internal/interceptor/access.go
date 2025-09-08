package interceptor

import (
	"chat-server/internal/client/rpc"
	"context"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const jwtPrefix = "Bearer "
const exampleAccessToken = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3NTY0Nzc4NDksImVtYWlsIjoidWxsYW1jbyBpbmNpZGlkdW50IGlydXJlIHBhcmlhdHVyIGN1cGlkYXRhdCIsInJvbGUiOiJBRE1JTiJ9.oQf-DviEpUMG4ClG9jNrJ1UEeDUFwHGYh2-yiG2m3Mk"

type AccessInterceptor struct {
	client rpc.AuthServiceClient
}

func NewAccessInterceptor(cl rpc.AuthServiceClient) *AccessInterceptor {
	return &AccessInterceptor{client: cl}
}

func (a *AccessInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	accessToken := exampleAccessToken
	md := metadata.New(map[string]string{"Authorization": jwtPrefix + accessToken})
	ctx = metadata.NewOutgoingContext(ctx, md)

	span, ctx := opentracing.StartSpanFromContext(ctx, "check access")
	defer span.Finish()
	span.SetTag("endpointAddress", info.FullMethod)
	err := a.client.Check(ctx, info.FullMethod)

	if err != nil {
		return nil, status.Error(codes.Unauthenticated, "invalid token")
	}
	return handler(ctx, req)
}
