package env

import (
	"log"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	refreshTokenSecretKey = "REFRESH_TOKEN_SECRET_KEY"
	accessTokenSecretKey  = "ACCESS_TOKEN_SECRET_KEY"

	refreshTokenExpiration = "REFRESH_EXPIRATION"
	accessTokenExpiration  = "ACCESS_EXPIRATION"
)

type jwtConfig struct {
	refreshTokenSecretKey string
	accessTokenSecretKey  string

	refreshTokenExpiration string
	accessTokenExpiration  string
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

	refreshEXP := os.Getenv(refreshTokenExpiration)
	if len(refreshEXP) == 0 {
		return nil, errors.New("refresh token expiration not found")
	}

	accessEXP := os.Getenv(accessTokenExpiration)
	if len(accessEXP) == 0 {
		return nil, errors.New("access token expiration not found")
	}

	return &jwtConfig{
		refreshTokenSecretKey: refresh,
		accessTokenSecretKey:  access,

		refreshTokenExpiration: refreshEXP,
		accessTokenExpiration:  accessEXP,
	}, nil
}

func (cfg *jwtConfig) RefreshTokenSecretKey() string {
	return cfg.refreshTokenSecretKey
}

func (cfg *jwtConfig) AccessTokenSecretKey() string {
	return cfg.accessTokenSecretKey
}

func (cfg *jwtConfig) RefreshTokenExpiration() time.Duration {
	duration, err := time.ParseDuration(cfg.refreshTokenExpiration)
	if err != nil {
		log.Fatalf("failed to parse refresh token expiration: %s", err.Error())
	}

	return duration
}

func (cfg *jwtConfig) AccessTokenExpiration() time.Duration {
	duration, err := time.ParseDuration(cfg.accessTokenExpiration)
	if err != nil {
		log.Fatalf("failed to parse access token expiration: %s", err.Error())
	}

	return duration

}
