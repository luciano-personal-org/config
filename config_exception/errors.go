package config_exception

import (
	exceptioncore "github.com/luciano-personal-org/exception"
)

var (
	LocalConfigError = exceptioncore.NewTradingError("CFG001", "Unable to read local config")
)
