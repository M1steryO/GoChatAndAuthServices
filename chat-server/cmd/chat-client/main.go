package main

import (
	"bufio"
	"chat-server/internal/logger"
	desc "chat-server/pkg/chat_v1"
	"context"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.NewClient(":50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.Printf("failed to close conn: %v", err)
		}
	}()

	client := desc.NewChatV1Client(conn)
	chat, err := client.Create(ctx, &desc.CreateRequest{
		Chat: &desc.Chat{
			Usernames: []string{
				"m1stery18", "oleg28",
			},
		},
	})
	if err != nil {
		logger.Error("Failed to create chat: ", err.Error())
		return
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		err := connectServer(ctx, client, chat.GetId(), "m1stery18", time.Second*5)
		if err != nil {
			log.Fatalf("Failed to connect chat: %s", err.Error())
		}
	}()
	go func() {
		defer wg.Done()
		err := connectServer(ctx, client, chat.GetId(), "oleg28", time.Second*5)
		if err != nil {
			log.Fatalf("Failed to connect chat: %s", err.Error())
		}
	}()
	wg.Wait()

}

func connectServer(ctx context.Context, client desc.ChatV1Client, chatID int64, username string, period time.Duration) error {
	stream, err := client.ConnectChat(ctx, &desc.ConnectChatRequest{
		ChatId:   chatID,
		Username: username,
	})
	if err != nil {
		log.Fatalf("Failed to connect chat: %s", err.Error())
	}

	go func() {
		for {
			res, err := stream.Recv()
			if err != nil {
				if err == io.EOF {
					log.Fatalf("Connection closed.")
					return
				}
				log.Fatalf("Failed to receive message: %s", err.Error())
			}
			if res.GetFrom() != username {
				log.Printf("[%v] - [from: %s]: %s\n",
					color.YellowString(res.GetTimestamp().AsTime().Format(time.RFC3339)),
					color.BlueString(res.GetFrom()),
					res.GetText(),
				)
			}

		}

	}()
	scanner := bufio.NewScanner(os.Stdin)
	var lines strings.Builder

	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}

		lines.WriteString(line)
		lines.WriteString("\n")
		_, err = client.SendMessage(ctx, &desc.SendMessageRequest{
			ChatId: chatID,
			Message: &desc.Message{
				From:      username,
				Text:      lines.String(),
				Timestamp: timestamppb.Now(),
			},
		})
		if err != nil {
			log.Println("Failed to send message: ", err.Error())
			return err
		}
	}

	err = scanner.Err()
	if err != nil {
		log.Println("failed to scan message: ", err)
	}
	return nil
}
