package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"

	_utilitieslogging "github.com/kieranroneill/dns-updater/internal/utilities/logging"
	"github.com/spf13/cobra"
)

func NewLogCommand(logger *_utilitieslogging.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "log",
		Short: "Show recent logs",
		Long:  `Print recent log lines from the dns-updater log file stored under "$HOME/.dns-updater/logs".`,
		RunE: func(cmd *cobra.Command, args []string) error {
			log.Print(fmt.Sprintf("%v", cmd.Flag("lineCount").Value))
			var buf []string

			file, err := os.Open(filepath.Clean(logger.FilePath()))
			if err != nil {
				// the no log file not existing is an acceptable error
				if os.IsNotExist(err) {
					return nil
				}

				return err
			}
			defer func(_file *os.File) {
				err = _file.Close()
				if err != nil {
					log.Fatalf("failed to close log file: %s", err)
				}
			}(file)

			// get the last n lines
			lines, err := strconv.ParseInt(cmd.Flag("lineCount").Value.String(), 10, 64)
			if err != nil {
				lines = 50
			}

			if len(args) > 0 {
				if _, err = fmt.Sscanf(args[0], "%d", &lines); err != nil {
					return fmt.Errorf("invalid line count: %w", err)
				}
			}

			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				buf = append(buf, scanner.Text())

				if len(buf) > int(lines) {
					buf = buf[1:]
				}
			}

			if err = scanner.Err(); err != nil {
				return err
			}

			for _, line := range buf {
				fmt.Println(line)
			}

			return nil
		},
	}
}
