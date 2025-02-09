package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	os.Setenv("TEST_MODE", "true")
	err := LoadConfig()
	assert.NoError(t, err, "LoadConfig should not return an error")

	expectedPort := ":8081"
	actualPort := viper.GetString("SERVER_PORT")
	assert.Equal(t, expectedPort, actualPort, "Expected port server to be :8081")

}
