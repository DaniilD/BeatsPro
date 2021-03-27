package controllers

import (
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/requests"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response"
	"BeatsPro/internal/response/consumerError"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type TagController struct {
	validatorFactory       *requests_validators.ValidatorFactory
	tagFactory             *Tag.TagFactory
	tagService             *services.TagService
	responseFactoryLocator *responseFactories.ResponseFactoryLocator
}

func NewTagController(
	validatorFactory *requests_validators.ValidatorFactory,
	tagFactory *Tag.TagFactory,
	tagService *services.TagService,
	responseFactoryLocator *responseFactories.ResponseFactoryLocator) *TagController {

	return &TagController{
		validatorFactory:       validatorFactory,
		tagFactory:             tagFactory,
		tagService:             tagService,
		responseFactoryLocator: responseFactoryLocator,
	}
}

func (tagController *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var CreateTagRequest requests.CreateTagRequest

	if err := json.NewDecoder(r.Body).Decode(&CreateTagRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	validator := tagController.validatorFactory.MakeCreateTagRequestValidator()
	if err := validator.Validate(&CreateTagRequest); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	tag := tagController.tagFactory.Make(CreateTagRequest.Tag.Title, false)
	id, err := tagController.tagService.CreateTag(tag)

	if err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
	}

	createTagResponse := tagController.responseFactoryLocator.GetCreateTagResponseFactory().Make(response.STATUS_SUCCESS, id)

	err = json.NewEncoder(w).Encode(createTagResponse)

	if err != nil {
		//sentry capture error
	}
}

func (tagController *TagController) UpdateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateTagRequest requests.UpdateTagRequest
	params := mux.Vars(r)

	paramsValidator := tagController.validatorFactory.MakeUpdateTagPathParamsValidator()

	if err := paramsValidator.Validate(params); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updateTagRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	requestValidator := tagController.validatorFactory.MakeUpdateTagRequestValidator()

	if err := requestValidator.Validate(&updateTagRequest); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	tagId, _ := strconv.Atoi(params["id"])
	tag, err := tagController.tagService.GetById(tagId)

	if err != nil {
		errorResponse := consumerError.NewError(1, "Тэг не найден")
		responseError := consumerError.NewNotFoundErrorResponse(response.STATUS_ERROR, errorResponse)

		json.NewEncoder(w).Encode(responseError)
		return
	}

	tag.Title = updateTagRequest.Tag.Title

	if err := tagController.tagService.UpdateTag(tag); err != nil {
		// ошибка 500
		fmt.Fprintln(os.Stdout, err)
		return
	}

	updateTagResponse := tagController.responseFactoryLocator.GetUpdateTagResponseFactory().Make(response.STATUS_SUCCESS)

	err = json.NewEncoder(w).Encode(updateTagResponse)

	if err != nil {
		//sentry capture error
	}
}

func (tagController *TagController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	paramsValidator := tagController.validatorFactory.MakeDeleteTagPathParamsValidator()

	if err := paramsValidator.Validate(params); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	id, _ := strconv.Atoi(params["id"])
	tag, err := tagController.tagService.GetById(id)

	if err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}
	tag.IsDeleted = true
	err = tagController.tagService.UpdateTag(tag)

	if err != nil {
		// ошибка 500
		fmt.Fprintln(os.Stdout, err)
		return
	}

	successResponse := response.NewSuccessResponse()

	err = json.NewEncoder(w).Encode(successResponse)

	if err != nil {
		// sentry caputure error
	}

}
