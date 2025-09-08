package app

import (
	chatServer "chat-server/internal/api/grpc/chat-server"
	dbClient "chat-server/internal/client/db"
	"chat-server/internal/client/db/pg"
	"chat-server/internal/client/db/transaction"
	"chat-server/internal/client/rpc"
	"chat-server/internal/client/rpc/auth"
	"chat-server/internal/closer"
	"chat-server/internal/config"
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"chat-server/pkg/access_v1"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	db "chat-server/internal/repository/chat"
	serv "chat-server/internal/service/chat"
	"context"
	"log"
)

type serviceProvider struct {
	dbConfig     config.DBConfig
	grpcConfig   config.GRPCConfig
	loggerConfig config.LoggerConfig

	authGRPCClient rpc.AuthServiceClient

	chatRepository repository.ChatRepository
	dbClient       dbClient.Client
	txManager      dbClient.TxManager

	chatService service.ChatService

	chatImpl *chatServer.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := config.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to load logger config: %s", err.Error())
		}
		s.loggerConfig = cfg
	}
	return s.loggerConfig
}

func (s *serviceProvider) DBConfig() config.DBConfig {
	if s.dbConfig == nil {
		cfg, err := config.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.dbConfig = cfg
	}

	return s.dbConfig
}
func (s *serviceProvider) DBCClient(ctx context.Context) dbClient.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.DBConfig().GetDSN())
		if err != nil {
			log.Fatalf("failed to connect to db: %s", err.Error())
		}
		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping db: %s", err.Error())
		}
		s.dbClient = cl
		closer.Add(cl.Close)
	}
	return s.dbClient

}

func (s *serviceProvider) TxManager(ctx context.Context) dbClient.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTxManager(s.DBCClient(ctx).DB())
	}
	return s.txManager
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) ChatRepository(ctx context.Context) repository.ChatRepository {
	if s.chatRepository == nil {
		s.chatRepository = db.NewChatRepository(s.DBCClient(ctx))
	}

	return s.chatRepository
}

func (s *serviceProvider) ChatService(ctx context.Context) service.ChatService {
	if s.chatService == nil {
		s.chatService = serv.NewChatService(
			s.ChatRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.chatService
}

func (s *serviceProvider) ChatImpl(ctx context.Context) *chatServer.Implementation {
	if s.chatImpl == nil {
		s.chatImpl = chatServer.NewImplementation(s.ChatService(ctx))
	}

	return s.chatImpl
}

func (s *serviceProvider) AuthGRPCClient() rpc.AuthServiceClient {
	if s.authGRPCClient == nil {
		conn, err := grpc.NewClient(
			s.GRPCConfig().AuthAddress(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())))
		if err != nil {
			log.Fatalf("failed to connect to auth service: %s", err.Error())
		}
		s.authGRPCClient = auth.New(access_v1.NewAccessV1Client(conn))
	}
	return s.authGRPCClient

}
