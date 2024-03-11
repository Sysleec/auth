package app

import (
	"context"
	"log"

	"github.com/Sysleec/auth/internal/api/access"
	"github.com/Sysleec/auth/internal/api/auth"
	"github.com/Sysleec/auth/internal/api/user"
	"github.com/Sysleec/auth/internal/client/db"
	"github.com/Sysleec/auth/internal/client/db/pg"
	"github.com/Sysleec/auth/internal/client/db/transaction"
	"github.com/Sysleec/auth/internal/closer"
	"github.com/Sysleec/auth/internal/config"
	"github.com/Sysleec/auth/internal/config/env"
	"github.com/Sysleec/auth/internal/repository"
	accessRepository "github.com/Sysleec/auth/internal/repository/access"
	authRepository "github.com/Sysleec/auth/internal/repository/auth"
	userRepository "github.com/Sysleec/auth/internal/repository/user"
	"github.com/Sysleec/auth/internal/service"
	accessService "github.com/Sysleec/auth/internal/service/access"
	authService "github.com/Sysleec/auth/internal/service/auth"
	userService "github.com/Sysleec/auth/internal/service/user"
)

type serviceProvider struct {
	pgConfig      config.PGConfig
	grpcConfig    config.GRPCConfig
	httpConfig    config.HTTPConfig
	swaggerConfig config.SwaggerConfig
	jwtConfig     config.JWTConfig
	loggerConfig  config.LoggerConfig
	promConfig    config.PromConfig

	dbClient         db.Client
	txManager        db.TxManager
	userRepository   repository.UserRepository
	authRepository   repository.AuthRepository
	accessRepository repository.AccessRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	userImpl   *user.Implementation
	authImpl   *auth.Implementation
	accessImpl *access.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := env.NewPGConfig()
		if err != nil {
			log.Fatalf("failed to get pg config: %s", err.Error())
		}

		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) GRPCConfig() config.GRPCConfig {
	if s.grpcConfig == nil {
		cfg, err := env.NewGRPCConfig()
		if err != nil {
			log.Fatalf("failed to get grpc config: %s", err.Error())
		}

		s.grpcConfig = cfg
	}

	return s.grpcConfig
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := env.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get http config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) JWTConfig() config.JWTConfig {
	if s.jwtConfig == nil {
		_, err := env.NewJWTConfig()
		if err != nil {
			log.Fatalf("failed to get jwt config: %s", err.Error())
		}
	}

	return s.jwtConfig
}

func (s *serviceProvider) SwaggerConfig() config.SwaggerConfig {
	if s.swaggerConfig == nil {
		cfg, err := env.NewSwaggerConfig()
		if err != nil {
			log.Fatalf("failed to get swagger config: %s", err.Error())
		}

		s.swaggerConfig = cfg
	}

	return s.swaggerConfig
}

func (s *serviceProvider) LoggerConfig() config.LoggerConfig {
	if s.loggerConfig == nil {
		cfg, err := env.NewLoggerConfig()
		if err != nil {
			log.Fatalf("failed to get logger config: %s", err.Error())
		}

		s.loggerConfig = cfg
	}

	return s.loggerConfig
}

func (s *serviceProvider) PromConfig() config.PromConfig {
	if s.promConfig == nil {
		cfg, err := env.NewPrometheusConfig()
		if err != nil {
			log.Fatalf("failed to get prometheus config: %s", err.Error())
		}

		s.promConfig = cfg
	}

	return s.promConfig
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		cl, err := pg.New(ctx, s.PGConfig().DSN())
		if err != nil {
			log.Fatalf("failed to create db client: %v", err)
		}

		err = cl.DB().Ping(ctx)
		if err != nil {
			log.Fatalf("ping error: %s", err.Error())
		}
		closer.Add(cl.Close)

		s.dbClient = cl
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		s.txManager = transaction.NewTransactionManager(s.DBClient(ctx).DB())
	}

	return s.txManager
}

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = userRepository.NewRepo(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = userService.NewService(
			s.UserRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.userService
}

func (s *serviceProvider) UserImpl(ctx context.Context) *user.Implementation {
	if s.userImpl == nil {
		s.userImpl = user.NewImplementation(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AuthRepository(ctx context.Context) repository.AuthRepository {
	if s.authRepository == nil {
		s.authRepository = authRepository.NewRepo(s.DBClient(ctx))
	}

	return s.authRepository
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = authService.NewService(
			s.AuthRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.authService
}

func (s *serviceProvider) AuthImpl(ctx context.Context) *auth.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth.NewImplementation(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) AccessRepository(ctx context.Context) repository.AccessRepository {
	if s.accessRepository == nil {
		s.accessRepository = accessRepository.NewRepo(s.DBClient(ctx))
	}

	return s.accessRepository
}

func (s *serviceProvider) AccessService(ctx context.Context) service.AccessService {
	if s.accessService == nil {
		s.accessService = accessService.NewService(
			s.AccessRepository(ctx),
			s.TxManager(ctx),
		)
	}

	return s.accessService
}
func (s *serviceProvider) AccessImpl(ctx context.Context) *access.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access.NewImplementation(s.AccessService(ctx))
	}

	return s.accessImpl
}
