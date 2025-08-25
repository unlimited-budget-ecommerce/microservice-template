package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewConfig(t *testing.T) {
	cfg := MustNew(".")

	require.NotNil(t, cfg)
	assert.NotEmpty(t, cfg.Service.Name)
	assert.NotEmpty(t, cfg.Service.Port)
	assert.NotEmpty(t, cfg.Service.Env)
	assert.NotEmpty(t, cfg.Logger.Level)
}

func TestGetConfigFromViper(t *testing.T) {
	_ = MustNew(".")

	assert.NotEmpty(t, viper.GetString("service.name"))
	assert.NotEmpty(t, viper.GetString("service.port"))
	assert.NotEmpty(t, viper.GetString("service.env"))
	assert.NotEmpty(t, viper.GetString("logger.level"))
}

func TestReplaceConfigWithEnv(t *testing.T) {
	os.Setenv("SERVICE_NAME", "test-service")
	os.Setenv("SERVICE_PORT", "test-port")
	os.Setenv("SERVICE_ENV", "test-env")
	os.Setenv("LOGGER_LEVEL", "test-level")

	_ = MustNew(".")

	assert.Equal(t, "test-service", viper.GetString("service.name"))
	assert.Equal(t, "test-port", viper.GetString("service.port"))
	assert.Equal(t, "test-env", viper.GetString("service.env"))
	assert.Equal(t, "test-level", viper.GetString("logger.level"))
}
