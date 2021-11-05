package operations

import (
	"BeatsPro/internal/models/Track"
	"BeatsPro/internal/operations/operationErrors"
	"BeatsPro/internal/requests"
	response2 "BeatsPro/internal/response"
	"BeatsPro/internal/services"
	"encoding/json"
	"net/http"
)

type CreateTrackOperation struct {
	trackFactory *Track.TrackFactory
	trackService *services.TrackService
}

func NewCreateTrackOperation(trackFactory *Track.TrackFactory) *CreateTrackOperation {
	return &CreateTrackOperation{
		trackFactory: trackFactory,
	}
}

func (operation *CreateTrackOperation) Handle(r *http.Request) (interface{}, error) {
	var request = &struct {
		Track *requests.Track `json:"track"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, new(operationErrors.JsonInvalidStructError)
	}

	track := operation.trackFactory.MakeFromRequest(request.Track)
	id, err := operation.trackService.CreateTrack(track)

	if err != nil {
		return nil, err
	}

	track.Id = id

	var response = &struct {
		Status string `json:"status"`
		Data   struct {
			Id int `json:"id"`
		} `json:"data"`
	}{}

	response.Status = response2.STATUS_SUCCESS
	response.Data.Id = id

	return response, nil
}
