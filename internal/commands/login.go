package commands

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/digitalocean/godo"
	_adapters "github.com/kieranroneill/dns-updater/internal/adapters"
	_dtos "github.com/kieranroneill/dns-updater/internal/dtos"
	_utilitiesterminal "github.com/kieranroneill/dns-updater/internal/utilities/terminal"
	"github.com/spf13/cobra"
)

func NewLoginCommand(settings *_dtos.Settings, logger *_adapters.LogAdapter) *cobra.Command {
	return &cobra.Command{
		Use:   "login",
		Short: "Configure DigitalOcean API token and DNS record",
		Long:  `Interactively configure your DigitalOcean API token and the A record to update.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			var record *godo.DomainRecord

			reader := bufio.NewReader(os.Stdin)
			apiTokenInput, err := _utilitiesterminal.GetFlagOrInput(cmd, "token", reader, "DigitalOcean API token: ", true)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			apiTokenInput = strings.TrimSpace(apiTokenInput)

			if apiTokenInput == "" {
				return fmt.Errorf("api token cannot be empty")
			}

			domainInput, err := _utilitiesterminal.GetFlagOrInput(cmd, "domain", reader, "Domain (e.g. example.com): ", false)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			domainInput = strings.TrimSpace(domainInput)

			if domainInput == "" {
				return fmt.Errorf("domain cannot be empty")
			}

			recordNameInput, err := _utilitiesterminal.GetFlagOrInput(cmd, "name", reader, `Record name (e.g. "sub" for "sub.example.com", or "@" for "example.com"): `, false)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			recordNameInput = strings.TrimSpace(recordNameInput)

			if recordNameInput == "" {
				return fmt.Errorf("record name cannot be empty")
			}

			recordIDInput, err := _utilitiesterminal.GetFlagOrInput(cmd, "id", reader, "Record ID (leave empty to auto-detect): ", false)
			if err != nil {
				logger.Error(err.Error())
				return err
			}

			recordIDInput = strings.TrimSpace(recordIDInput)

			ctx := context.Background()
			doClient := _adapters.NewDigitalOceanAdapter(apiTokenInput)

			// if no record id was supplied, try and fetch it using the do api
			if recordIDInput == "" {
				record, err = doClient.DomainRecordByNameAndType(ctx, domainInput, recordNameInput, "A")
				if err != nil {
					err = fmt.Errorf("failed to auto-detect record ID: %w", err)
					logger.Error(err.Error())
					return err
				}
			}

			// if a record id was supplied, attempt to get it from the do api
			if recordIDInput != "" {
				id, err := strconv.Atoi(recordIDInput)
				if err != nil {
					err = fmt.Errorf(`invalid record id %s supplied: %w`, recordIDInput, err)
					logger.Error(err.Error())
					return err
				}

				record, err = doClient.DomainRecordByID(ctx, domainInput, id)
				if err != nil {
					err = fmt.Errorf(`failed to get domain record %s: %w`, recordIDInput, err)
					logger.Error(err.Error())
					return err
				}
			}

			// check the record actually exists
			if record == nil {
				return fmt.Errorf("unable to find record %s on domain %s", recordNameInput, domainInput)
			}

			config := &_dtos.Config{
				Auth: _dtos.AuthConfig{
					APIToken: apiTokenInput,
				},
				Domain: domainInput,
				Record: _dtos.RecordConfig{
					ID:   record.ID,
					Name: record.Name,
					Type: record.Type,
				},
			}
			configAdapter := _adapters.NewConfigAdapter(settings.ConfigPath)

			if err := configAdapter.SetConfig(config); err != nil {
				return fmt.Errorf("failed to save config: %w", err)
			}

			logger.Info(`login configuration saved successfully to "%s".`, settings.ConfigPath)

			return nil
		},
	}
}
