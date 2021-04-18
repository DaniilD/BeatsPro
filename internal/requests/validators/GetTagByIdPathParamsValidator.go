package requests_validators

import (
	requests_validators_rules "BeatsPro/internal/requests/validators/rules"
	"errors"
)

type GetTagByIdPathParamsValidator struct {
	isIntegerRule *requests_validators_rules.IsIntegerRule
}

func NewGetTagByIdPathParamsValidator(isIntegerRule *requests_validators_rules.IsIntegerRule) *GetTagByIdPathParamsValidator {
	return &GetTagByIdPathParamsValidator{
		isIntegerRule: isIntegerRule,
	}
}

func (getTagByIdPathParamsValidator *GetTagByIdPathParamsValidator) Validate(params map[string]string) error {
	if !getTagByIdPathParamsValidator.isIntegerRule.Validate(params["id"]) {
		return errors.New("math: id parameter must be a number")
	}

	return nil
}
