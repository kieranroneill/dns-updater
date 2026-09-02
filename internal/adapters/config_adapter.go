package adapters

import (
	"os"

	_dtos "github.com/kieranroneill/dns-updater/internal/dtos"
	"go.yaml.in/yaml/v3"
)

type ConfigAdapter struct {
	configPath string
}

/**
 * constructors
 */

func NewConfigAdapter(configPath string) *ConfigAdapter {
	return &ConfigAdapter{
		configPath: configPath,
	}
}

/**
 * public methods
 */

func (c *ConfigAdapter) GetConfig() (*_dtos.Config, error) {
	var config _dtos.Config

	file, err := os.ReadFile(c.configPath)
	if err != nil {
		return nil, err
	}

	if err = yaml.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *ConfigAdapter) SetConfig(config *_dtos.Config) error {
	bytes, err := yaml.Marshal(config)
	if err != nil {
		return err
	}

	if err = os.WriteFile(c.configPath, bytes, 0600); err != nil {
		return err
	}

	return nil
}
