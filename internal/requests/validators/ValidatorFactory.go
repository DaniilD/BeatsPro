package requests_validators

import requests_validators_rules "BeatsPro/internal/requests/validators/rules"

type ValidatorFactory struct {
}

func NewValidatorFactory() *ValidatorFactory {
	return &ValidatorFactory{}
}

func (validatorFactory *ValidatorFactory) MakeCreateTagRequestValidator() *CreateTagRequestValidator {
	return NewCreateTagRequestValidator(requests_validators_rules.NewNotEmptyRule())
}

func (validatorFactory *ValidatorFactory) MakeUpdateTagPathParamsValidator() *UpdateTagPathParamsValidator {
	return NewUpdateTagPathParamsValidator(requests_validators_rules.NewIsIntegerRule())
}

func (validatorFactory *ValidatorFactory) MakeUpdateTagRequestValidator() *UpdateTagRequestValidator {
	return NewUpdateTagRequestValidator(requests_validators_rules.NewNotEmptyRule())
}

func (validatorFactory *ValidatorFactory) MakeDeleteTagPathParamsValidator() *DeleteTagPathParamsValidator {
	return NewDeleteTagPathParamsValidator(requests_validators_rules.NewIsIntegerRule())
}
