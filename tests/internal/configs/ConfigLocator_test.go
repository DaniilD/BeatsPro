package configs_test

import (
	"BeatsPro/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetConfigLocator_returnConfigLocatorInstance(t *testing.T) {
	//Act
	configLocatorAct := configs.GetConfigLocator()

	//Exp
	configLocatorExp := &configs.ConfigLocator{}

	//Assert
	assert.Equal(t, configLocatorAct, configLocatorExp)
}

func TestGetConfigLocator_returnTwoEqualInstance(t *testing.T) {
	//Act
	configLocatorAct := configs.GetConfigLocator()

	//Exp
	configLocatorExp := configs.GetConfigLocator()

	//Assert
	assert.ObjectsAreEqual(configLocatorExp, configLocatorAct)
}

func TestServerConfigInstance_returnServerConfigInstance(t *testing.T) {
	//Act
	serverConfigAct := configs.GetConfigLocator().ServerConfigInstance()

	//Exp
	serverConfigExp := configs.NewServerConfig()

	//Assert
	assert.Equal(t, serverConfigAct, serverConfigExp)
}

func TestDBConfigInstance_returnDBConfigInstance(t *testing.T) {
	//Act
	dbConfigAct := configs.GetConfigLocator().DBConfigInstance()

	//Exp
	dbConfigExp := configs.NewDBConfig()

	//Assert
	assert.Equal(t, dbConfigAct, dbConfigExp)
}

func TestServerConfigInstance_returnTwoEqualInstance(t *testing.T) {
	//Act
	serverConfigAct := configs.GetConfigLocator().ServerConfigInstance()
	serverConfigAct.Host = "127.0.0.1"

	//Exp
	serverConfigExp := configs.GetConfigLocator().ServerConfigInstance()

	//Assert
	assert.ObjectsAreEqual(serverConfigAct, serverConfigExp)
}

func TestDBConfigInstance_returnTwoEqualInstance(t *testing.T) {
	//Act
	dbConfigAct := configs.GetConfigLocator().DBConfigInstance()
	dbConfigAct.Host = "127.0.0.1"

	//Exp
	dbConfigExp := configs.GetConfigLocator().DBConfigInstance()

	//Assert
	assert.ObjectsAreEqual(dbConfigAct, dbConfigExp)
}
