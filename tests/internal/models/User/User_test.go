package User_test

import (
	"BeatsPro/internal/models/User"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	//Act
	userAct := User.NewUser()

	//Exp
	userExp := &User.User{}

	//Assert
	assert.Equal(t, userAct, userExp)
}
