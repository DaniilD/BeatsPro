package responseFactories

import "BeatsPro/internal/response"

type CreateTagResponseFactory struct {
}

func NewCreateTagResponseFactory() *CreateTagResponseFactory {
	return &CreateTagResponseFactory{}
}

func (createTagResponseFactory *CreateTagResponseFactory) Make(status string, tagId int) *response.CreateTagResponse {
	createTagResponse := response.NewCreateTagResponse()
	createTagResponse.Status = status
	createTagResponse.Data.Id = tagId

	return createTagResponse
}
