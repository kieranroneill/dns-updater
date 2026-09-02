package utilities

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromptSecretSuccess(t *testing.T) {
	expected := "super-secret"
	reader := bufio.NewReader(strings.NewReader(fmt.Sprintf("%s\n", expected)))
	actual, err := PromptSecret(reader, "API Token: ")

	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestPromptSecretReturnsReadError(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader(""))
	actual, err := PromptSecret(reader, "API Token: ")

	assert.ErrorIs(t, err, io.EOF)
	assert.Empty(t, actual)
}
