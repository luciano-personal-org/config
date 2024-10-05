// // Package config provides functionality for configuration management.
// package config

// import (
// 	"os"

// 	"github.com/luciano-personal-org/exception"

// 	"github.com/joho/godotenv"
// )

// // Config interface for config.
// type Config interface {
// 	Get(key string) string
// }

// // configImpl struct for configImpl.
// type configImpl struct {
// }

// // Get retrieves the value of the specified environment variable.
// func (config *configImpl) Get(key string) string {
// 	return os.Getenv(key)
// }

// // New returns a new instance of the Config interface, loaded with configuration values from the specified .env files.
// func New(filenames ...string) Config {
// 	err := godotenv.Load(filenames...)
// 	exception.PanicIfNeeded(err)
// 	return &configImpl{}
// }

// Package config provides functionality for configuration management.
package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/luciano-personal-org/exception"
	"github.com/spf13/viper"
)

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
func New() Config {
	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigFile(".env.yaml")
	// for _, filename := range filenames {
	// v.ReadConfig(bytes.NewBuffer([]byte(filename)))
	// // 	// v.SetConfigFile(filename)
	// // 	err := v.MergeInConfig()
	// // 	exception.PanicIfNeeded(err)
	// // }
	if err := v.ReadInConfig(); err != nil {
		var fileNotFoundError *os.PathError
		if errors.As(err, &fileNotFoundError) {
			// Config file not found; ignore error if desired
			fmt.Printf("Local config file not found, using defaults values\n")
		} else {
			exception.PanicIfNeeded(fmt.Errorf("fatal error when reading local config: %w", err))
		}
	}

	fmt.Println(v.AllSettings())
	return &configImpl{v: v}
}
