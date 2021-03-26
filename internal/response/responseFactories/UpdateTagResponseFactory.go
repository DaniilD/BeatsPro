package responseFactories

import "BeatsPro/internal/response"

type UpdateTagResponseFactory struct {
}

func NewUpdateTagResponseFactory() *UpdateTagResponseFactory {
	return &UpdateTagResponseFactory{}
}

func (updateTagResponseFactory *UpdateTagResponseFactory) Make(status string) *response.UpdateTagResponse {
	updateTagResponse := response.NewUpdateTagResponse()
	updateTagResponse.Status = status

	return updateTagResponse
}
