package config

import (
	"errors"
	"net"
	"os"
)

const (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
	authHostEnvName = "AUTH_GRPC_HOST"
	authPortEnvName = "AUTH_GRPC_PORT"
)

type GRPCConfig interface {
	Address() string
	AuthAddress() string
}
type grpcConfig struct {
	host     string
	port     string
	authHost string
	authPort string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("grpc host not found")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("grpc port not found")
	}

	authHost := os.Getenv(authHostEnvName)
	if len(authHost) == 0 {
		return nil, errors.New("grpc auth host not found")
	}
	authPort := os.Getenv(authPortEnvName)
	if len(authPort) == 0 {
		return nil, errors.New("grpc auth port not found")
	}

	return &grpcConfig{
		host:     host,
		port:     port,
		authHost: authHost,
		authPort: authPort,
	}, nil
}

func (c *grpcConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}

func (c *grpcConfig) AuthAddress() string {
	return net.JoinHostPort(c.authHost, c.authPort)
}
