package requests_validators

import (
	"BeatsPro/internal/requests"
	requests_validators_rules "BeatsPro/internal/requests/validators/rules"
	"errors"
)

type UpdateTagRequestValidator struct {
	notEmptyRule *requests_validators_rules.NotEmptyRule
}

func NewUpdateTagRequestValidator(notEmptyRule *requests_validators_rules.NotEmptyRule) *UpdateTagRequestValidator {
	return &UpdateTagRequestValidator{
		notEmptyRule: notEmptyRule,
	}
}

func (updateTagRequestValidator *UpdateTagRequestValidator) Validate(updateTagRequest *requests.UpdateTagRequest) error {

	if !updateTagRequestValidator.notEmptyRule.Validate(updateTagRequest) {
		return errors.New("math: Empty Request")
	}

	if !updateTagRequestValidator.notEmptyRule.Validate(updateTagRequest.Tag) {
		return errors.New("math: Parameter \"tag\" is required")
	}

	if !updateTagRequestValidator.notEmptyRule.Validate(updateTagRequest.Tag.Title) {
		return errors.New("math: Parameter \"title\" is required")
	}
	return nil
}
