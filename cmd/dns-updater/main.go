package main

import (
	"fmt"
	"log/slog"

	_utilitiesconfigs "github.com/kieranroneill/dns-updater/internal/utilities/configs"
	_utilitieslogging "github.com/kieranroneill/dns-updater/internal/utilities/logging"
)

var Version string

func main() {
	config := _utilitiesconfigs.CreateConfig(Version)

	// config logger
	slog.SetLogLoggerLevel(_utilitieslogging.ParseLogLevel(config.LogLevel))

	slog.Info(fmt.Sprintf("Log level: %s", config.LogLevel))
	slog.Info(fmt.Sprintf("Version: %s", config.Version))
}
