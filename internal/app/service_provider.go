package app

import (
	"context"
	"log"

	"github.com/Sysleec/auth/internal/api/user"
	"github.com/Sysleec/auth/internal/closer"
	"github.com/Sysleec/auth/internal/config"
	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/repository"
	userRepository "github.com/Sysleec/auth/internal/repository/user"
	"github.com/Sysleec/auth/internal/service"
	userService "github.com/Sysleec/auth/internal/service/user"
	"github.com/jackc/pgx/v4/pgxpool"
)

type serviceProvider struct {
	pgConfig   config.PGConfig
	grpcConfig config.GRPCConfig

	pgPool         *pgxpool.Pool
	userRepository repository.UserRepository

	userService service.UserService

	userServ *user.Server
}

func NewServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to load pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to load grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) PGPool(ctx context.Context) *pgxpool.Pool {
	if s.pgPool == nil {
		pool, err := pgxpool.Connect(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to connect to database: %v", err)
		}

		err = pool.Ping(ctx)
		if err != nil {
			log.Fatalf("failed to ping database: %s", err.Error())
		}
		closer.Add(func() error {
			pool.Close()
			return nil
		})

		s.pgPool = pool
	}

	return s.pgPool
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepo(s.PGPool(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) UserServer(ctx context.Context) *user.Server {
	if s.userServ == nil {
		s.userServ = user.NewServer(s.UserService(ctx))
	}

	return s.userServ
}
