package Tag_test

import (
	"BeatsPro/internal/models/Tag"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTag(t *testing.T) {
	//Act
	tagAct := Tag.NewTag()

	//Exp
	tagExp := &Tag.Tag{}

	//Assert
	assert.Equal(t, tagAct, tagExp)
}
