package operations

import (
	"BeatsPro/internal/operations/operationErrors"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response"
	"BeatsPro/internal/services"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type DeleteTagOperation struct {
	deleteTagPathParamsValidator *requests_validators.DeleteTagPathParamsValidator
	tagService                   *services.TagService
}

func NewDeleteTagOperation(
	deleteTagPathParamsValidator *requests_validators.DeleteTagPathParamsValidator,
	tagService *services.TagService,
) *DeleteTagOperation {
	return &DeleteTagOperation{
		deleteTagPathParamsValidator: deleteTagPathParamsValidator,
		tagService:                   tagService,
	}
}

func (operation *DeleteTagOperation) Handle(r *http.Request) (interface{}, error) {
	params := mux.Vars(r)

	if err := operation.deleteTagPathParamsValidator.Validate(params); err != nil {
		return nil, operationErrors.NewValidateError(err.Error())
	}

	id, _ := strconv.Atoi(params["id"])
	tag, err := operation.tagService.GetById(id)

	if err != nil {
		return nil, operationErrors.NewTagNotFoundError()
	}

	tag.IsDeleted = true

	if err := operation.tagService.UpdateTag(tag); err != nil {
		return nil, err
	}

	return response.NewSuccessResponse(), nil
}
