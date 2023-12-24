package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/Sysleec/auth/internal/config"
	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/converter"
	userRepository "github.com/Sysleec/auth/internal/repository/user"
	"github.com/Sysleec/auth/internal/service"
	userService "github.com/Sysleec/auth/internal/service/user"
	desc "github.com/Sysleec/auth/pkg/user_v1"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", ".env", "path to config file")
}

type server struct {
	desc.UnimplementedUserV1Server
	userService service.UserService
}

func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	id, err := s.userService.Create(ctx, converter.ToUserFromDesc(req))
	if err != nil {
		return nil, err
	}
	log.Printf("created user with id %d", id)

	return &desc.CreateResponse{Id: id}, nil
}

func (s *server) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	user, err := s.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	fmt.Printf("got user: %+v\n", user)

	return converter.ToUserFromService(user), nil
}

func main() {
	flag.Parse()

	ctx := context.Background()

	err := config.Load(configPath)
	if err != nil {
		log.Fatal(err)
	}

	grpcConfig, err := env.NewGRPCConfig()
	if err != nil {
		log.Fatal(err)
	}

	pgConfig, err := env.NewPGConfig()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", grpcConfig.Address())
	if err != nil {
		log.Fatal(err)
	}

	pool, err := pgxpool.Connect(ctx, pgConfig.DSN())
	if err != nil {
		log.Fatal(err)
	}

	userRepo := userRepository.NewRepo(pool)
	userSrv := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, &server{userService: userSrv})

	fmt.Printf("starting server at %s\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
