package Track

import "BeatsPro/internal/requests"

type TrackFactory struct {
}

func NewTrackFactory() *TrackFactory {
	return &TrackFactory{}
}

func (factory *TrackFactory) MakeFromRequest(trackRequest *requests.Track) *Track {
	track := new(Track)

	track.Name = trackRequest.Name
	track.Description = trackRequest.Description
	track.Duration = trackRequest.Duration
	track.Bpm = trackRequest.Bpm
	track.IsDeleted = trackRequest.IsDeleted

	return nil
}
