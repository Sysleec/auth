FROM golang:1.21.4-alpine AS builder

COPY . /github.com/Sysleec/chat-server/source/
WORKDIR /github.com/Sysleec/chat-server/source/

RUN go mod download
RUN go build -o ./bin/chat_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Sysleec/chat-server/source/bin/chat_server .
COPY --from=builder /github.com/Sysleec/chat-server/source/.env .

CMD ["./chat_server"]