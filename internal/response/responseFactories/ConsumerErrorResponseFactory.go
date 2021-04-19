package responseFactories

import (
	"BeatsPro/internal/response"
	"BeatsPro/internal/response/consumerError"
)

type ConsumerErrorResponseFactory struct {
}

func (factory *ConsumerErrorResponseFactory) MakeBadRequestResponse(message string, code int) *consumerError.BadRequestResponse {
	error := consumerError.NewError(code, message)
	return consumerError.NewBadRequestResponse(response.STATUS_ERROR, error)
}
