package controllers

import (
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/requests"
	requests_validators "BeatsPro/internal/requests/validators"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TagController struct {
	createTagRequestValidator *requests_validators.CreateTagRequestValidator
	tagFactory                *Tag.TagFactory
}

func NewTagController(
	createTagRequestValidator *requests_validators.CreateTagRequestValidator,
	tagFactory *Tag.TagFactory) *TagController {
	return &TagController{
		createTagRequestValidator: createTagRequestValidator,
		tagFactory:                tagFactory,
	}
}

func (tagController *TagController) CreateTag(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var CreateTagRequest requests.CreateTagRequest

	if err := json.NewDecoder(r.Body).Decode(&CreateTagRequest); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tagController.createTagRequestValidator.Validate(&CreateTagRequest); err != nil {
		// ошибка 400
		fmt.Fprintln(os.Stdout, err)
		return
	}

	tag := tagController.tagFactory.Make(CreateTagRequest.Tag.Title)

	fmt.Fprintln(os.Stdout, tag.Title)
}
