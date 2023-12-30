package config

import (
	"github.com/joho/godotenv"
)

// GRPCConfig is an interface for gRPC configuration
type GRPCConfig interface {
	Address() string
}

// PGConfig is an interface for PostgreSQL configuration
type PGConfig interface {
	DSN() string
}

// Load loads configuration from .env file
func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
