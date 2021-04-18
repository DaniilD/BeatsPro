package responseFactories

import (
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/response"
	"BeatsPro/internal/response/responseModels"
)

type GetByIdResponseFactory struct {
}

func NewGetByIdResponseFactory() *GetByIdResponseFactory {
	return &GetByIdResponseFactory{}
}

func (GetByIdResponseFactory *GetByIdResponseFactory) Make(tag *Tag.Tag) *response.GetByIdResponse {
	getByIdResponse := response.NewGetByIdResponse()
	getByIdResponse.Status = response.STATUS_SUCCESS
	getByIdResponse.Data.Tag = responseModels.NewTag()
	getByIdResponse.Data.Tag.Id = tag.Id
	getByIdResponse.Data.Tag.Title = tag.Title

	return getByIdResponse
}
