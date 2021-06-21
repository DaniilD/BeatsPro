package operations

import (
	"BeatsPro/internal/operations/operationErrors"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type GetByIdTagOperation struct {
	paramsValidator *requests_validators.GetTagByIdPathParamsValidator
	tagService      *services.TagService
	responseFactory *responseFactories.GetByIdResponseFactory
}

func NewGetByIdTagOperation(
	paramsValidator *requests_validators.GetTagByIdPathParamsValidator,
	tagService *services.TagService,
	responseFactory *responseFactories.GetByIdResponseFactory,
) *GetByIdTagOperation {
	return &GetByIdTagOperation{
		paramsValidator: paramsValidator,
		tagService:      tagService,
		responseFactory: responseFactory,
	}
}

func (operation *GetByIdTagOperation) Handle(r *http.Request) (interface{}, error) {
	params := mux.Vars(r)

	if err := operation.paramsValidator.Validate(params); err != nil {
		return nil, operationErrors.NewValidateError(err.Error())
	}

	tagId, err := strconv.Atoi(params["id"])
	if err != nil {
		return nil, err
	}

	tag, err := operation.tagService.GetById(tagId)

	if err != nil {
		return nil, operationErrors.NewTagNotFoundError()
	}

	return operation.responseFactory.Make(tag), nil
}
