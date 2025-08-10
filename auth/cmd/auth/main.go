package main

import (
	"auth/internal/api/grpc/user"
	"auth/internal/config"
	db "auth/internal/repository/user"
	serv "auth/internal/service/user"
	desc "auth/pkg/user_v1"
	"context"
	"flag"
	"github.com/jackc/pgx/v4/pgxpool"
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
	defer lis.Close()

	pool, err := pgxpool.Connect(ctx, dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer pool.Close()

	repo := db.NewUserRepository(pool)
	service := serv.NewUserService(repo)

	s := grpc.NewServer()
	reflection.Register(s)

	desc.RegisterUserV1Server(s, user.NewUserImplementation(service))

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
