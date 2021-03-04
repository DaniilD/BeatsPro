package Track

import "time"

type Track struct {
	Id               int
	UserId           int
	Name             string
	Description      string
	DateTimeCreation time.Time
	Duration         int
	Bpm              int
	IsDeleted        bool
}

func NewTrack() *Track {
	return &Track{}
}
