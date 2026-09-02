package utilities

import (
	"bufio"
	"fmt"
	"strings"
)

func Prompt(r *bufio.Reader, label string) (string, error) {
	fmt.Print(label)

	s, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(s), nil
}
