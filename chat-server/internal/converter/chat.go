package converter

import repoModel "chat-server/internal/repository/chat/model"
import "chat-server/internal/model"
import desc "chat-server/pkg/chat_v1"

func ToChatApiFromService(chat *model.Chat) *desc.Chat {
	return &desc.Chat{
		Usernames: chat.Usernames,
	}
}

func ToChatRepoFromService(chat *model.Chat) *repoModel.Chat {
	return &repoModel.Chat{
		Usernames: chat.Usernames,
		UpdatedAt: chat.UpdatedAt,
		CreatedAt: chat.CreatedAt,
	}
}

func ToChatServiceFromApi(chat *desc.Chat) *model.Chat {
	return &model.Chat{
		Usernames: chat.Usernames,
	}
}
