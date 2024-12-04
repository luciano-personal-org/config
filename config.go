// // Package config provides functionality for configuration management.
package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/luciano-personal-org/config/exception"
	exceptioncore "github.com/luciano-personal-org/exception"
	"github.com/spf13/viper"
)

// Initialize the custom error
var custom_error exceptioncore.TradingError

// Config interface for config.
type Config interface {
	Get(key string) string
}

// configImpl struct for configImpl.
type configImpl struct {
	v *viper.Viper
}

// Get retrieves the value of the specified configuration key.
func (config *configImpl) Get(key string) string {
	return config.v.GetString(key)
}

// New returns a new instance of the Config interface, loaded with configuration values from the specified files.
func NewConfig() Config {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(".env.yaml")
	if err := v.ReadInConfig(); err != nil {
		var fileNotFoundError *os.PathError
		if errors.As(err, &fileNotFoundError) {
			// Config file not found; ignore error if desired
			fmt.Printf("Local config file not found, using defaults values\n")
		} else {
			// exception.PanicIfNeeded(fmt.Errorf("fatal error when reading local config: %w", err))
			custom_error = exception.LocalConfigError
			custom_error.SetOriginalError(err)
			custom_error.SetDetails("When trying to read local config from config pkg")
			exceptioncore.DoPanic(custom_error)
		}
	}
	return &configImpl{v: v}
}
