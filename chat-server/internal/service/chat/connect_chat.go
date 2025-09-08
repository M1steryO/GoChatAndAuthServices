package chat

import (
	"chat-server/internal/repository/chat"
	"chat-server/internal/service"
	"context"
)

func (s *serv) ConnectChat(ctx context.Context, chatId int64, username string, stream service.Stream) error {
	s.mxChannel.RLock()
	chatChannel, ok := s.channels[chatId]
	s.mxChannel.RUnlock()
	if !ok {
		return chat.ChatNotFoundError
	}

	s.mxChat.Lock()
	if _, ok := s.chats[chatId]; !ok {
		s.chats[chatId] = &Streams{
			streams: make(map[string]service.Stream),
		}
	}
	s.mxChat.Unlock()

	s.chats[chatId].m.Lock()
	s.chats[chatId].streams[username] = stream
	s.chats[chatId].m.Unlock()

	for {
		select {
		case msg, okCh := <-chatChannel:
			if !okCh {
				return nil
			}
			for _, st := range s.chats[chatId].streams {
				if err := st.Send(msg); err != nil {
					return err
				}
			}

		case <-ctx.Done():
			s.mxChat.Lock()
			delete(s.chats[chatId].streams, username)
			s.mxChat.Unlock()
		}
	}

}
