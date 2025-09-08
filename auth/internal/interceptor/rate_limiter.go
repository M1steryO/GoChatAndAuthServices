package interceptor

import (
	"auth/internal/utils/rate_limiter"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RateLimiterInterceptor struct {
	rateLimiter *rate_limiter.TokenBucketLimiter
}

func NewRateLimiterInterceptor(rateLimiter *rate_limiter.TokenBucketLimiter) *RateLimiterInterceptor {
	return &RateLimiterInterceptor{rateLimiter: rateLimiter}
}
func (r *RateLimiterInterceptor) Unary(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	if !r.rateLimiter.Allow() {
		return nil, status.Error(codes.ResourceExhausted, "rate limit exceeded")
	}
	return handler(ctx, req)
}
