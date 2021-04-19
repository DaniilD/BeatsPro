package controllers

import (
	"BeatsPro/internal/Enums"
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/operations"
	"BeatsPro/internal/operations/operationErrors"
	"BeatsPro/internal/requests"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response"
	"BeatsPro/internal/response/consumerError"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
	"encoding/json"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type TagController struct {
	createTagOperation     *operations.CreateTagOperation
	consumerErrorFactory   *responseFactories.ConsumerErrorResponseFactory
	validatorFactory       *requests_validators.ValidatorFactory
	tagFactory             *Tag.TagFactory
	tagService             *services.TagService
	responseFactoryLocator *responseFactories.ResponseFactoryLocator
}

func NewTagController(
	createTagOperation *operations.CreateTagOperation,
	consumerErrorFactory *responseFactories.ConsumerErrorResponseFactory,
	validatorFactory *requests_validators.ValidatorFactory,
	tagFactory *Tag.TagFactory,
	tagService *services.TagService,
	responseFactoryLocator *responseFactories.ResponseFactoryLocator) *TagController {

	return &TagController{
		createTagOperation:     createTagOperation,
		consumerErrorFactory:   consumerErrorFactory,
		validatorFactory:       validatorFactory,
		tagFactory:             tagFactory,
		tagService:             tagService,
		responseFactoryLocator: responseFactoryLocator,
	}
}

func (controller *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	createTagResponse, err := controller.createTagOperation.Handle(r)

	if err != nil {
		if e, ok := err.(*operationErrors.JsonInvalidStructError); ok {
			createTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), Enums.JSON_INVALID_STRUCT)
		} else if e, ok := err.(*operationErrors.ValidateError); ok {
			createTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), Enums.JSON_INVALID_STRUCT)
		} else {
			sentry.CaptureException(err)
		}
	}

	err = json.NewEncoder(w).Encode(createTagResponse)

	if err != nil {
		sentry.CaptureException(err)
	}
}

func (controller *TagController) UpdateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateTagRequest requests.UpdateTagRequest
	params := mux.Vars(r)

	paramsValidator := controller.validatorFactory.MakeUpdateTagPathParamsValidator()

	if err := paramsValidator.Validate(params); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&updateTagRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	requestValidator := controller.validatorFactory.MakeUpdateTagRequestValidator()

	if err := requestValidator.Validate(&updateTagRequest); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	tagId, _ := strconv.Atoi(params["id"])
	tag, err := controller.tagService.GetById(tagId)

	if err != nil {
		errorResponse := consumerError.NewError(1, "Тэг не найден")
		responseError := consumerError.NewNotFoundErrorResponse(response.STATUS_ERROR, errorResponse)

		json.NewEncoder(w).Encode(responseError)
		return
	}

	tag.Title = updateTagRequest.Tag.Title

	if err := controller.tagService.UpdateTag(tag); err != nil {
		// ошибка 500
		fmt.Fprintln(os.Stdout, err)
		return
	}

	updateTagResponse := controller.responseFactoryLocator.GetUpdateTagResponseFactory().Make(response.STATUS_SUCCESS)

	err = json.NewEncoder(w).Encode(updateTagResponse)

	if err != nil {
		sentry.CaptureException(err)
	}
}

func (controller *TagController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	paramsValidator := controller.validatorFactory.MakeDeleteTagPathParamsValidator()

	if err := paramsValidator.Validate(params); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	id, _ := strconv.Atoi(params["id"])
	tag, err := controller.tagService.GetById(id)

	if err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}
	tag.IsDeleted = true
	err = controller.tagService.UpdateTag(tag)

	if err != nil {
		// ошибка 500
		fmt.Fprintln(os.Stdout, err)
		return
	}

	successResponse := response.NewSuccessResponse()

	err = json.NewEncoder(w).Encode(successResponse)

	if err != nil {
		sentry.CaptureException(err)
	}

}

func (controller *TagController) GetTagById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	paramsValidator := controller.validatorFactory.MakeGetByIdPathParamsValidator()

	if err := paramsValidator.Validate(params); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	tagId, err := strconv.Atoi(params["id"])
	if err != nil {
		sentry.CaptureException(err)
	}

	tag, err := controller.tagService.GetById(tagId)

	if err != nil {
		errorResponse := consumerError.NewError(1, "Тэг не найден")
		responseError := consumerError.NewNotFoundErrorResponse(response.STATUS_ERROR, errorResponse)

		err := json.NewEncoder(w).Encode(responseError)
		if err != nil {
			sentry.CaptureException(err)
		}
		return
	}

	getTagResponse := controller.responseFactoryLocator.GetByIdResponseFactory().Make(tag)

	err = json.NewEncoder(w).Encode(getTagResponse)

	if err != nil {
		sentry.CaptureException(err)
	}
}
