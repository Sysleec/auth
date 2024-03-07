package config

import (
	"github.com/joho/godotenv"
)

// SwaggerConfig is an interface for Swagger configuration
type SwaggerConfig interface {
	Address() string
}

// HTTPConfig is an interface for HTTP configuration
type HTTPConfig interface {
	Address() string
}

// GRPCConfig is an interface for gRPC configuration
type GRPCConfig interface {
	Address() string
}

// PGConfig is an interface for PostgreSQL configuration
type PGConfig interface {
	DSN() string
}

// JWTConfig is an interface for JWT configuration
type JWTConfig interface {
	AccessTokenSecretKey() string
	RefreshTokenSecretKey() string
}

// Load loads configuration from .env file
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
