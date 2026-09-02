package utilities

import (
	"bufio"
	"strconv"

	"github.com/spf13/cobra"
)

// GetFlagOrInput retrieves the flag value if set or prompts the user for input, optionally hiding the input for secret
// values.
//
// Parameters:
//   - cmd: The cobra command instance.
//   - flagName: The name of the flag to retrieve.
//   - reader: The bufio.Reader instance for user input.
//   - prompt: The prompt message to display to the user.
//   - secret: Indicates whether the input should be hidden (true) or not (false).
//
// Returns:
//   - The value from the flag or user input.
//   - An error if any occurs during input retrieval.
func GetFlagOrInput(cmd *cobra.Command, flagName string, reader *bufio.Reader, prompt string, secret bool) (string, error) {
	flagType := cmd.Flag(flagName).Value.Type()
	value := cmd.Flag(flagName).Value.String()

	switch flagType {
	case "string":
		if value != "" {
			return value, nil
		}

		break
	case "int":
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return "", err
		}

		if intValue >= 0 {
			return value, nil
		}

		// the default value for int flags will be -1, so we need to handle it
		value = ""

		break
	default:
		break
	}

	if secret {
		return Prompt(reader, prompt)
	}

	return PromptSecret(reader, prompt)
}
