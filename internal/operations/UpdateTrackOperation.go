package operations

import (
	"BeatsPro/internal/models/Track"
	"BeatsPro/internal/operations/operationErrors"
	"BeatsPro/internal/requests"
	"BeatsPro/internal/services"
	"encoding/json"
	"net/http"
)

type UpdateTrackOperation struct {
	trackFactory *Track.TrackFactory
	trackService *services.TrackService
}

func NewUpdateTrackOperation(trackFactory *Track.TrackFactory, trackService *services.TrackService) *UpdateTrackOperation {
	return &UpdateTrackOperation{
		trackFactory: trackFactory,
		trackService: trackService,
	}
}

func (operation *UpdateTrackOperation) Handle(r *http.Request) (interface{}, error) {
	var request = &struct {
		Id    int             `json:"id" binding:"required"`
		Track *requests.Track `json:"track"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, new(operationErrors.JsonInvalidStructError)
	}

	track, err := operation.trackService.GetById(request.Id)

	if err != nil {
		return nil, err
	}

	return nil, nil
}
