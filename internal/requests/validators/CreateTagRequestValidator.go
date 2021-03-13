package requests_validators

import (
	"BeatsPro/internal/requests"
	requests_validators_rules "BeatsPro/internal/requests/validators/rules"
	"errors"
)

type CreateTagRequestValidator struct {
	notEmptyRule *requests_validators_rules.NotEmptyRule
}

func NewCreateTagRequestValidator(
	notEmptyRule *requests_validators_rules.NotEmptyRule) *CreateTagRequestValidator {

	return &CreateTagRequestValidator{
		notEmptyRule: notEmptyRule,
	}
}

func (createTagRequestValidator *CreateTagRequestValidator) Validate(
	createTagRequest *requests.CreateTagRequest) error {

	if !createTagRequestValidator.notEmptyRule.Validate(createTagRequest) {
		return errors.New("math: Empty Request")
	}

	if !createTagRequestValidator.notEmptyRule.Validate(createTagRequest.Tag) {
		return errors.New("math: Parameter \"tag\" is required")
	}

	if !createTagRequestValidator.notEmptyRule.Validate(createTagRequest.Tag.Title) {
		return errors.New("math: Parameter \"title\" is required")
	}

	return nil
}
