package chat

import (
	"chat-server/internal/model"
	"chat-server/internal/repository"
	"chat-server/internal/service"
	"github.com/M1steryO/platform_common/pkg/db"
	"sync"
)

type Streams struct {
	streams map[string]service.Stream //nolint:unused
	m       sync.RWMutex              //nolint:unused
}
type serv struct {
	db        repository.ChatRepository
	txManager db.TxManager

	chats  map[int64]*Streams
	mxChat sync.RWMutex

	channels  map[int64]chan *model.Message
	mxChannel sync.RWMutex
}

func NewChatService(db repository.ChatRepository, txManager db.TxManager) service.ChatService {
	return &serv{
		db:        db,
		txManager: txManager,
		chats:     make(map[int64]*Streams),
		channels:  make(map[int64]chan *model.Message),
	}
}
