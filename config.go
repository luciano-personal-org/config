// Package config provides functionality for configuration management.
package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config interface for config.
type Config interface {
	Get(key string) string
}

// configImpl struct for configImpl.
type configImpl struct {
}

// Get retrieves the value of the specified environment variable.
func (config *configImpl) Get(key string) string {
	return os.Getenv(key)
}

// New returns a new instance of the Config interface, loaded with configuration values from the specified .env files.
func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfNeeded(err)
	return &configImpl{}
}
