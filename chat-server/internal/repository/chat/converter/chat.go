package converter

import (
	"chat-server/internal/repository/chat/model"
	desc "chat-server/pkg/chat_v1"
)

func ToChatFromRepo(chat *model.Chat) *desc.Chat {
	return &desc.Chat{
		Usernames: chat.Usernames,
	}
}
