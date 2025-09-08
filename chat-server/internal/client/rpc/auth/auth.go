package auth

import (
	"chat-server/internal/client/rpc"
	descAccess "chat-server/pkg/access_v1"
	"context"
)

type authServiceClient struct {
	authClient descAccess.AccessV1Client
}

func New(authClient descAccess.AccessV1Client) rpc.AuthServiceClient {
	return &authServiceClient{
		authClient: authClient,
	}
}

func (c *authServiceClient) Check(ctx context.Context, endpointAddress string) error {
	_, err := c.authClient.Check(ctx, &descAccess.CheckRequest{
		EndpointAddress: endpointAddress,
	})
	if err != nil {
		return err
	}
	return nil
}
