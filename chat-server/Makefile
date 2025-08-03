LOCAL_BIN:=$(CURDIR)/bin

install-golangci-lint:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.3.0

lint:
	$(LOCAL_BIN)/golangci-lint run ./... --config .golangci.pipeline.yaml

install-deps:
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest # Для структур
	GOBIN=$(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1 # Для gRPC сервера и клиента gRPC

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc


generate:
	make generate-chat-api

generate-chat-api:
	mkdir -p pkg/chat_v1
	# Указываем путь, где искать .proto файлы и их импорты
 	# --go_out: выходная директория для сгенерированных структур Protobuf
 	# --go_opt: опции для Go-генератора
	# paths=source_relative - сохраняет структуру директорий как в исходниках
	# Указываем путь к плагину protoc-gen-go (генератор Go-кода)
 	# Генерируем Go-код для gRPC сервисов
   # --go-grpc_out: выходная директория для gRPC-кода
   # --go-grpc_opt: опции для gRPC-генератора
   # paths=source_relative - аналогично для gRPC-кода
   # Указываем путь к плагину protoc-gen-go-grpc (генератор gRPC-кода)
   # Исходный .proto файл для компиляции
	protoc --proto_path api/chat_v1 \
        --go_out pkg/chat_v1 \
        --go_opt=paths=source_relative \
        --plugin=protoc-gen-go=bin/protoc-gen-go \
        --go-grpc_out=pkg/chat_v1 \
        --go-grpc_opt=paths=source_relative \
        --plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
        api/chat_v1/chat.proto