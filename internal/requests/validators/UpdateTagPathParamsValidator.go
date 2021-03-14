package requests_validators

import (
	requests_validators_rules "BeatsPro/internal/requests/validators/rules"
	"errors"
)

type UpdateTagPathParamsValidator struct {
	isIntegerRule *requests_validators_rules.IsIntegerRule
}

func NewUpdateTagPathParamsValidator(isIntegerRule *requests_validators_rules.IsIntegerRule) *UpdateTagPathParamsValidator {
	return &UpdateTagPathParamsValidator{
		isIntegerRule: isIntegerRule,
	}
}

func (updateTagPathParamsValidator *UpdateTagPathParamsValidator) Validate(params map[string]string) error {

	if !updateTagPathParamsValidator.isIntegerRule.Validate(params["id"]) {
		return errors.New("math: id parameter must be a number")
	}

	return nil
}
