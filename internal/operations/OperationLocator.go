package operations

import (
	"BeatsPro/internal/models/Tag"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
)

var operationLocator *OperationLocator

type OperationLocator struct {
	createTagOperation *CreateTagOperation
}

func GetOperationLocator() *OperationLocator {
	if operationLocator == nil {
		operationLocator = &OperationLocator{}
	}

	return operationLocator
}

func (locator OperationLocator) GetCreateTagOperation() *CreateTagOperation {
	if locator.createTagOperation == nil {
		locator.createTagOperation = NewCreateTagOperation(
			requests_validators.NewValidatorFactory().MakeCreateTagRequestValidator(),
			Tag.NewTagFactory(),
			services.GetServiceLocator().GetTagService(),
			responseFactories.GetResponseFactoryLocator().GetCreateTagResponseFactory(),
		)
	}

	return locator.createTagOperation
}
