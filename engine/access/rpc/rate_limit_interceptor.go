package rpc

import (
	"context"
	"time"

	"github.com/rs/zerolog"
	"google.golang.org/grpc"
)

type rateLimiter struct {
	log zerolog.Logger
}

func (interceptor *rateLimiter) Limit() bool {
	return true
}

// TODO: not used currently. May use it later if we want to refine the error message
func (interceptor *rateLimiter) unaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()

	// call the handler
	h, err := handler(ctx, req)

	interceptor.log.Debug().
		Str("method", info.FullMethod).
		Interface("request", req).
		Interface("response", resp).
		Err(err).
		Dur("request_duration", time.Since(start))

	return h, err
}
