package operations

import (
	"BeatsPro/internal/operations/operationErrors"
	"BeatsPro/internal/requests"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UpdateTagOperation struct {
	validator                *requests_validators.UpdateTagRequestValidator
	pathValidator            *requests_validators.UpdateTagPathParamsValidator
	tagService               *services.TagService
	updateTagResponseFactory *responseFactories.UpdateTagResponseFactory
}

func NewUpdateTagOperation(
	validator *requests_validators.UpdateTagRequestValidator,
	pathValidator *requests_validators.UpdateTagPathParamsValidator,
	tagService *services.TagService,
	updateTagResponseFactory *responseFactories.UpdateTagResponseFactory,
) *UpdateTagOperation {
	return &UpdateTagOperation{
		validator:                validator,
		pathValidator:            pathValidator,
		tagService:               tagService,
		updateTagResponseFactory: updateTagResponseFactory,
	}
}

func (operation *UpdateTagOperation) Handle(r *http.Request) (interface{}, error) {
	var updateTagRequest *requests.UpdateTagRequest

	if err := json.NewDecoder(r.Body).Decode(&updateTagRequest); err != nil {
		return nil, new(operationErrors.JsonInvalidStructError)
	}

	params := mux.Vars(r)

	if err := operation.pathValidator.Validate(params); err != nil {
		return nil, operationErrors.NewValidateError(err.Error())
	}

	if err := operation.validator.Validate(updateTagRequest); err != nil {
		return nil, operationErrors.NewValidateError(err.Error())
	}

	tagId, err := strconv.Atoi(params["id"])

	if err != nil {
		sentry.CaptureException(err)
	}

	tag, err := operation.tagService.GetById(tagId)

	if err != nil {
		return nil, operationErrors.NewTagNotFoundError()
	}

	tag.Title = updateTagRequest.Tag.Title

	if err := operation.tagService.UpdateTag(tag); err != nil {
		return nil, err
	}

	updateTagResponse := operation.updateTagResponseFactory.Make(response.STATUS_SUCCESS)

	return updateTagResponse, nil
}
