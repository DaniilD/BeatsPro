package operations

import (
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/requests"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
	"encoding/json"
	"net/http"
)

type CreateTagOperation struct {
	validator                *requests_validators.CreateTagRequestValidator
	tagFactory               *Tag.TagFactory
	tagService               *services.TagService
	createTagResponseFactory *responseFactories.CreateTagResponseFactory
}

func NewCreateTagOperation(
	validator *requests_validators.CreateTagRequestValidator,
	tagFactory *Tag.TagFactory,
	tagService *services.TagService,
	createTagResponseFactory *responseFactories.CreateTagResponseFactory) *CreateTagOperation {
	return &CreateTagOperation{
		validator:                validator,
		tagFactory:               tagFactory,
		tagService:               tagService,
		createTagResponseFactory: createTagResponseFactory,
	}
}

func (operation *CreateTagOperation) Handle(r *http.Request) (interface{}, error) {
	var createTagRequest requests.CreateTagRequest

	if err := json.NewDecoder(r.Body).Decode(&createTagRequest); err != nil {
		return nil, err
	}

	if err := operation.validate(&createTagRequest); err != nil {
		return nil, err
	}

	tag := operation.tagFactory.Make(createTagRequest.Tag.Title, false)
	id, err := operation.tagService.CreateTag(tag)

	if err != nil {
		return nil, err
	}

	createTagResponse := operation.createTagResponseFactory.Make(response.STATUS_SUCCESS, id)

	return createTagResponse, nil
}

func (operation *CreateTagOperation) validate(createTagRequest *requests.CreateTagRequest) error {
	if err := operation.validator.Validate(createTagRequest); err != nil {
		return err
	}

	return nil
}
