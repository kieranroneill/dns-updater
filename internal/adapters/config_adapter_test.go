package adapters

import (
	"io/fs"
	"path/filepath"
	"testing"

	_dtos "github.com/kieranroneill/dns-updater/internal/dtos"
	"github.com/stretchr/testify/assert"
)

var config = _dtos.Config{
	Auth: _dtos.AuthConfig{
		APIToken: "super-secret",
	},
	Domain: "example.com",
	Record: _dtos.RecordConfig{
		ID:   1337,
		Name: "sub",
		Type: "A",
	},
}

func TestGetConfigNoFileExists(t *testing.T) {
	var expectedError *fs.PathError

	adapter := NewConfigAdapter(filepath.Join("testdata", "no_config.yaml"))
	actual, err := adapter.GetConfig()

	assert.ErrorAs(t, err, &expectedError)
	assert.Nil(t, actual)
}

func TestGetConfigSuccess(t *testing.T) {
	adapter := NewConfigAdapter(filepath.Join("testdata", "mock_config.yaml"))
	actual, err := adapter.GetConfig()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, &config, actual)
}

func TestSetConfigSuccess(t *testing.T) {
	tempDirectory := t.TempDir()
	adapter := NewConfigAdapter(filepath.Join(tempDirectory, "config.yaml"))
	err := adapter.SetConfig(&config)
	if err != nil {
		t.Error(err)
	}

	actual, err := adapter.GetConfig()
	if err != nil {
		t.Error(err)
	}

	assert.Equal(t, &config, actual)
}
