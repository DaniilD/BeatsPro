package requests_validators

import (
	requests_validators_rules "BeatsPro/internal/requests/validators/rules"
	"errors"
)

type DeleteTagPathParamsValidator struct {
	isIntegerRule *requests_validators_rules.IsIntegerRule
}

func NewDeleteTagPathParamsValidator(isIntegerRule *requests_validators_rules.IsIntegerRule) *DeleteTagPathParamsValidator {
	return &DeleteTagPathParamsValidator{
		isIntegerRule: isIntegerRule,
	}
}

func (deleteTagPathParamsValidator *DeleteTagPathParamsValidator) Validate(params map[string]string) error {

	if deleteTagPathParamsValidator.isIntegerRule.Validate(params["id"]) {
		return errors.New("math: id parameter must be a number")
	}

	return nil
}
