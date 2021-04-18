package response

import "BeatsPro/internal/response/responseModels"

type GetByIdResponse struct {
	Status string `json:"status"`
	Data   struct {
		Tag *responseModels.Tag `json:"tag"`
	} `json:"data"`
}

func NewGetByIdResponse() *GetByIdResponse {
	return &GetByIdResponse{}
}
