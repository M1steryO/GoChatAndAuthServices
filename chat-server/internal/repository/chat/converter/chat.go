package converter

import (
	servModel "chat-server/internal/model"
	"chat-server/internal/repository/chat/model"
)

func ToChatFromRepo(chat *model.Chat) *servModel.Chat {
	return &servModel.Chat{
		Id:        chat.Id,
		Usernames: chat.Usernames,
		CreatedAt: chat.CreatedAt,
		UpdatedAt: chat.UpdatedAt,
	}
}

func ToMessageRepoFromService(msg *servModel.Message) *model.Message {
	return &model.Message{
		From:      msg.From,
		Text:      msg.Text,
		Timestamp: msg.Timestamp,
	}
}
