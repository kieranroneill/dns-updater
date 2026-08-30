package utilities

import (
	_dtosconfigs "github.com/kieranroneill/dns-updater/internal/dtos/configs"
	_utilitiesos "github.com/kieranroneill/dns-updater/internal/utilities/os"
)

func CreateConfig(version string) _dtosconfigs.ConfigDTO {
	return _dtosconfigs.ConfigDTO{
		LogLevel: _utilitiesos.GetEnvWithDefault("LOG_LEVEL", "error"),
		Version:  version,
	}
}
