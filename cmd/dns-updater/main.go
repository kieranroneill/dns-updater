package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	_adapters "github.com/kieranroneill/dns-updater/internal/adapters"
	_commands "github.com/kieranroneill/dns-updater/internal/commands"
	_constants "github.com/kieranroneill/dns-updater/internal/constants"
	"github.com/kieranroneill/dns-updater/internal/dtos"
	_utilitiesapplication "github.com/kieranroneill/dns-updater/internal/utilities/application"
	"github.com/spf13/cobra"
)

var Version string

func main() {
	var lineCountFlag int64

	configDirectory, err := _utilitiesapplication.ConfigDirectory()
	if err != nil {
		log.Fatalf("failed to get config path: %v", err)
	}

	logDirectory, err := _utilitiesapplication.LogDirectory()
	if err != nil {
		log.Fatalf("failed to get log path: %v", err)
	}

	settings := &dtos.Settings{
		ConfigPath: filepath.Join(configDirectory, "config.yml"),
		LogPath:    filepath.Join(logDirectory, "dns-updater.log"),
		Version:    Version,
	}
	logger := _adapters.NewLogAdapter(settings.LogPath)
	rootCommand := &cobra.Command{
		Use:     "dns-updater",
		Short:   _constants.AppShortDescription,
		Long:    _constants.AppLongDescription,
		Version: Version,
	}

	// create commands
	logCommand := _commands.NewLogCommand(settings)
	loginCommand := _commands.NewLoginCommand(settings, logger)
	runCommand := _commands.NewRunCommand(settings, logger)

	// add flags
	logCommand.Flags().Int64VarP(&lineCountFlag, "lineCount", "l", 50, "The amount of lines to display. Defaults to 50.")

	// add commands
	rootCommand.AddCommand(
		logCommand,
		loginCommand,
		runCommand,
	)

	if err = rootCommand.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
