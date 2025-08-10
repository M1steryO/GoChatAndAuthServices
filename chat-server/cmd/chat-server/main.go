package main

import (
	"chat-server/internal/api/grpc/chat-server"
	"chat-server/internal/config"
	"chat-server/internal/repository/chat"
	chat2 "chat-server/internal/service/chat"
	desc "chat-server/pkg/chat_v1"
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

	pool, err := pgxpool.Connect(ctx, dbConfig.GetDSN())
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer pool.Close()
	chatRepo := chat.NewChatRepository(pool)
	chatService := chat2.NewChatService(chatRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, chat_server.NewImplementation(chatService))

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
