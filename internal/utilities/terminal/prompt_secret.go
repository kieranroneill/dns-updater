package utilities

import (
	"bufio"
	"fmt"
	"strings"
	"syscall"

  "golang.org/x/term"
)

func PromptSecret(r *bufio.Reader, label string) (string, error) {
	fmt.Print(label)

	b, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		// fallback to normal prompt if terminal doesn't support hidden input
		return Prompt(r, label)
	}

	fmt.Println()

	return strings.TrimSpace(string(b)), nil
}
