FROM golang:1.21.4-alpine AS builder

COPY . /github.com/Sysleec/auth/source/
WORKDIR /github.com/Sysleec/auth/source/

RUN go mod download
RUN go build -o ./bin/auth cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Sysleec/auth/source/bin/auth .
COPY --from=builder /github.com/Sysleec/auth/source/.env .

RUN chmod +x auth
CMD ["./auth"]