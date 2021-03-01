package configs_test

import (
	"BeatsPro/internal/configs"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDBConfig(t *testing.T) {
	//Act
	dbConfigAct := configs.NewDBConfig()

	//Exp
	dbConfigExp := &configs.DBConfig{}

	//Assert
	assert.Equal(t, dbConfigAct, dbConfigExp)
}
