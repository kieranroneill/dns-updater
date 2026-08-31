package main

import (
	"fmt"
	"log"
	"os"

	_commands "github.com/kieranroneill/dns-updater/internal/commands"
	_constants "github.com/kieranroneill/dns-updater/internal/constants"
	_utilitieslogging "github.com/kieranroneill/dns-updater/internal/utilities/logging"
	"github.com/spf13/cobra"
)

var Version string

func main() {
	var lineCountFlag int64

	logger, err := _utilitieslogging.NewLogger()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	rootCommand := &cobra.Command{
		Use:     "dns-updater",
		Short:   _constants.AppShortDescription,
		Long:    _constants.AppLongDescription,
		Version: Version,
	}

	// create commands
	logCommand := _commands.NewLogCommand(logger)

	// add flags
	logCommand.Flags().Int64VarP(&lineCountFlag, "lineCount", "l", 50, "The amount of lines to display. Defaults to 50.")

	// add commands
	rootCommand.AddCommand(logCommand)

	if err = rootCommand.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
