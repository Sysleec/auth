package env

import (
	"os"

	"github.com/pkg/errors"
)

const (
	refreshTokenSecretKey = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenSecretKey  = "ACCESS_TOKEN_SECRET_KEY"
)

type jwtConfig struct {
	refreshTokenSecretKey string
	accessTokenSecretKey  string
}

// NewJWTConfig creates new JWT configuration
func NewJWTConfig() (*jwtConfig, error) {
	access := os.Getenv(accessTokenSecretKey)
	if len(access) == 0 {
		return nil, errors.New("access token secret key not found")
	}
	refresh := os.Getenv(refreshTokenSecretKey)
	if len(refresh) == 0 {
		return nil, errors.New("refresh token secret key not found")
	}

	return &jwtConfig{
		refreshTokenSecretKey: refresh,
		accessTokenSecretKey:  access,
	}, nil
}

func (cfg *jwtConfig) RefreshTokenSecretKey() string {
	return cfg.refreshTokenSecretKey
}

func (cfg *jwtConfig) AccessTokenSecretKey() string {
	return cfg.accessTokenSecretKey
}
