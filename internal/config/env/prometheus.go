package env

import (
	"net"
	"os"

	"github.com/pkg/errors"
)

const (
	promHostEnvName = "PROMETHEUS_HOST"
	promPortEnvName = "PROMETHEUS_PORT"
)

type promConfig struct {
	host string
	port string
}

func NewPrometheusConfig() (*promConfig, error) {
	host := os.Getenv(promHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("prometheus host not found")
	}

	port := os.Getenv(promPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("prometheus port not found")
	}

	return &promConfig{
		host: host,
		port: port,
	}, nil
}

func (cfg *promConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
