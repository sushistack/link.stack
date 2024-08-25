package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadEnvironment(t *testing.T) {
	envFilePath4Test := "../configs/test.env"

	envContent := "MONGODB_URI=mongodb://localhost:27017\n" +
		"MONGODB_USERNAME=testuser\n" +
		"MONGODB_PASSWORD=testpass\n"

	err := os.WriteFile(envFilePath4Test, []byte(envContent), 0644)
	require.NoError(t, err, "Failed to create .env file")

	options := &EnvironmentOptions{
		EnvFilePath: envFilePath4Test,
	}
	env := LoadEnvironment(options)
	require.NoError(t, err, "LoadEnvironment returned an error")

	assert.Equal(t, "mongodb://localhost:27017", env["MONGODB_URI"], "Expected MONGODB_URI value")
	assert.Equal(t, "testuser", env["MONGODB_USERNAME"], "Expected MONGODB_USERNAME value")
	assert.Equal(t, "testpass", env["MONGODB_PASSWORD"], "Expected MONGODB_PASSWORD value")

	err = os.Remove(envFilePath4Test)
	require.NoError(t, err, "Failed to remove .env file")
}
