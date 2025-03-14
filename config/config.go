package config

import (
	"os"

	"github.com/charmbracelet/log"
)

const (
	KubeconfigEnvName = "KUBECONFIG"
)

var ErrKubeconfigEmpty = "kubeconfig is empty"

type Config struct {
	kubeconfig string
}

func New() (*Config, error) {
	kubeconfig := os.Getenv(KubeconfigEnvName)
	if kubeconfig == "" {
		log.Warn(ErrKubeconfigEmpty)
	}

	return &Config{
		kubeconfig: kubeconfig,
	}, nil
}

func (c *Config) Clientset() string {
	return c.kubeconfig
}
