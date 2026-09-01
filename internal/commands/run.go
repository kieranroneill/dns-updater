package commands

import (
	_adapters "github.com/kieranroneill/dns-updater/internal/adapters"
	_dtos "github.com/kieranroneill/dns-updater/internal/dtos"
	"github.com/spf13/cobra"
)

func NewRunCommand(settings *_dtos.Settings, logger *_adapters.LogAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "run",
		Short: "Fetches the public IP address and updates the DigitalOcean DNS record. Requires login.",
		Long:  `Print recent log lines from the dns-updater log file stored under "$HOME/.dns-updater/logs".`,
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}
}
