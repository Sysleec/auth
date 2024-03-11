package env

import (
	"os"

	"github.com/pkg/errors"
)

const (
	logLevel = "LOG_LEVEL"
)

type loggerConfig struct {
	logLevel string
}

func NewLoggerConfig() (*loggerConfig, error) {
	logLevel := os.Getenv(logLevel)
	if len(logLevel) == 0 {
		return nil, errors.New("log level not found")
	}

	return &loggerConfig{
		logLevel: logLevel,
	}, nil
}

func (cfg *loggerConfig) LogLevel() string {
	return cfg.logLevel
}
