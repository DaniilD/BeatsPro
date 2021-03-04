package services

import (
	"BeatsPro/internal/models/Track"
	"BeatsPro/internal/repositories"
)

type TrackService struct {
	trackRepository *repositories.TrackRepository
}

func NewTrackService(trackRepository *repositories.TrackRepository) *TrackService {
	return &TrackService{
		trackRepository: trackRepository,
	}
}

func (trackService *TrackService) CreateTrack(track *Track.Track) (int, error) {
	return trackService.trackRepository.CreateTrack(track)
}

func (trackService *TrackService) UpdateTrack(track *Track.Track) error {
	return trackService.trackRepository.UpdateTrack(track)
}

func (trackService *TrackService) GetById(id int) (*Track.Track, error) {
	return trackService.trackRepository.GetById(id)
}
