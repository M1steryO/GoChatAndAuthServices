package main

import (
	"auth/internal/config"
	"auth/internal/grpc/handlers/user"
	db "auth/internal/storage"
	desc "auth/pkg/user_v1"
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var configPath = ""

func init() {
	flag.StringVar(&configPath, "config-path", "local.env", "path to config file")
}

func main() {
	flag.Parse()
	ctx := context.Background()
	err := config.Load(configPath)
	if err != nil {
		log.Fatalf("failed to load config %v", err)
	}

	grpcConfig, err := config.NewGRPCConfig()
	if err != nil {
		log.Fatalf("failed to load grpc config %v", err)
	}

	dbConfig, err := config.NewDBConfig()
	if err != nil {
		log.Fatalf("failed to load db config: %v", err)
	}
	_ = dbConfig

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	storage, err := db.NewStorage(ctx, dbConfig)
	if err != nil {
		log.Fatalf("failed to connect to storage: %v", err)
	}

	defer storage.Pool.Close()

	desc.RegisterUserV1Server(s, &user.Server{
		Storage: storage,
	})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
