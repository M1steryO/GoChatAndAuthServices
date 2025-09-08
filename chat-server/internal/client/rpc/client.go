package rpc

import "context"

type AuthServiceClient interface {
	Check(ctx context.Context, endpointAddress string) error
}
