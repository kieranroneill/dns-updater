package utilities

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptSuccess(t *testing.T) {
	expected := "example.com"
	reader := bufio.NewReader(strings.NewReader(fmt.Sprintf("%s\n", expected)))
	actual, err := Prompt(reader, "Domain: ")

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestPromptReturnsReadError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	actual, err := Prompt(reader, "Domain: ")

	assert.ErrorIs(t, err, io.EOF)
	assert.Empty(t, actual)
}
