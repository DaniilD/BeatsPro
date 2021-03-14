package controllers

import (
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/requests"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/services"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"strconv"
)

type TagController struct {
	validatorFactory *requests_validators.ValidatorFactory
	tagFactory       *Tag.TagFactory
	tagService       *services.TagService
}

func NewTagController(
	validatorFactory *requests_validators.ValidatorFactory,
	tagFactory *Tag.TagFactory,
	tagService *services.TagService) *TagController {
	return &TagController{
		validatorFactory: validatorFactory,
		tagFactory:       tagFactory,
		tagService:       tagService,
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

	tag := tagController.tagFactory.Make(CreateTagRequest.Tag.Title)
	id, err := tagController.tagService.CreateTag(tag)

	if err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
	}

	fmt.Fprintln(os.Stdout, id)
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

	//сначала получить тег потом обновить !!!!
	tag := tagController.tagFactory.Make(updateTagRequest.Tag.Title)
	tagId, _ := strconv.Atoi(params["id"])
	tag.Id = tagId

	if err := tagController.tagService.UpdateTag(tag); err != nil {
		// ошибка 500
		fmt.Fprintln(os.Stdout, err)
		return
	}

	fmt.Fprintln(os.Stdout, "ok")
}

func (tagController *TagController) DeleteTag(w http.ResponseWriter, r *http.Request) {

}
