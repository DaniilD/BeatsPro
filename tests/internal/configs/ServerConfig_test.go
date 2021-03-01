package configs_test

import (
	"BeatsPro/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServerConfig(t *testing.T) {
	//Act
	serverConfigAct := configs.NewServerConfig()

	//Exp
	serverConfigExp := &configs.ServerConfig{}

	//Assert
	assert.Equal(t, serverConfigAct, serverConfigExp)
}
