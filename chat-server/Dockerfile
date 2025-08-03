FROM golang:1.23-alpine AS builder

COPY . /m1stery18/github.com/TestGoProject/chat-server/source
WORKDIR /m1stery18/github.com/TestGoProject/chat-server/source

RUN go mod download
RUN go build -o ./bin/chat-server cmd/chat-server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /m1stery18/github.com/TestGoProject/chat-server/source/bin/chat-server .

CMD ["./chat-server"]