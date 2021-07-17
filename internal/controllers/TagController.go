package controllers

import (
	"BeatsPro/internal/Enums"
	"BeatsPro/internal/operations"
	"BeatsPro/internal/operations/operationErrors"
	"BeatsPro/internal/response/responseFactories"
	"encoding/json"
	"github.com/getsentry/sentry-go"
	"net/http"
)

type TagController struct {
	createTagOperation   *operations.CreateTagOperation
	updateTagOperation   *operations.UpdateTagOperation
	deleteTagOperation   *operations.DeleteTagOperation
	getByIdTagOperation  *operations.GetByIdTagOperation
	consumerErrorFactory *responseFactories.ConsumerErrorResponseFactory
}

func NewTagController(
	createTagOperation *operations.CreateTagOperation,
	updateTagOperation *operations.UpdateTagOperation,
	deleteTagOperation *operations.DeleteTagOperation,
	getByIdTagOperation *operations.GetByIdTagOperation,
	consumerErrorFactory *responseFactories.ConsumerErrorResponseFactory) *TagController {

	return &TagController{
		createTagOperation:   createTagOperation,
		updateTagOperation:   updateTagOperation,
		deleteTagOperation:   deleteTagOperation,
		getByIdTagOperation:  getByIdTagOperation,
		consumerErrorFactory: consumerErrorFactory,
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

	updateTagResponse, err := controller.updateTagOperation.Handle(r)

	if err != nil {
		if e, ok := err.(*operationErrors.JsonInvalidStructError); ok {
			updateTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), Enums.JSON_INVALID_STRUCT)
		} else if e, ok := err.(*operationErrors.ValidateError); ok {
			updateTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), Enums.JSON_INVALID_STRUCT)
		} else if e, ok := err.(*operationErrors.TagNotFoundError); ok {
			updateTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), e.Code)
		} else {
			sentry.CaptureException(err)
		}
	}

	err = json.NewEncoder(w).Encode(updateTagResponse)

	if err != nil {
		sentry.CaptureException(err)
	}
}

func (controller *TagController) DeleteTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	deleteTagResponse, err := controller.deleteTagOperation.Handle(r)

	if err != nil {
		if e, ok := err.(*operationErrors.ValidateError); ok {
			deleteTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), Enums.JSON_INVALID_STRUCT)
		} else if e, ok := err.(*operationErrors.TagNotFoundError); ok {
			deleteTagResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), e.Code)
		} else {
			sentry.CaptureException(err)
		}
	}

	err = json.NewEncoder(w).Encode(deleteTagResponse)

	if err != nil {
		sentry.CaptureException(err)
	}
}

func (controller *TagController) GetTagById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	getByIdResponse, err := controller.getByIdTagOperation.Handle(r)

	if err != nil {
		if e, ok := err.(*operationErrors.ValidateError); ok {
			getByIdResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), Enums.JSON_INVALID_STRUCT)
		} else if e, ok := err.(*operationErrors.TagNotFoundError); ok {
			getByIdResponse = controller.consumerErrorFactory.MakeBadRequestResponse(e.Error(), e.Code)
		} else {
			sentry.CaptureException(err)
		}
	}

	err = json.NewEncoder(w).Encode(getByIdResponse)

	if err != nil {
		sentry.CaptureException(err)
	}
}
