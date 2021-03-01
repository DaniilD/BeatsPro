package Track_test

import (
	"BeatsPro/internal/models/Track"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewTrack(t *testing.T) {
	//Act
	trackAct := Track.NewTrack()

	//Exp
	trackExp := &Track.Track{}

	//Assert
	assert.Equal(t, trackAct, trackExp)
}
