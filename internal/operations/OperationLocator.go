package operations

import (
	"BeatsPro/internal/models/Tag"
	requests_validators "BeatsPro/internal/requests/validators"
	"BeatsPro/internal/response/responseFactories"
	"BeatsPro/internal/services"
)

var operationLocator *OperationLocator

type OperationLocator struct {
	createTagOperation  *CreateTagOperation
	updateTagOperation  *UpdateTagOperation
	deleteTagOperation  *DeleteTagOperation
	GetByIdTagOperation *GetByIdTagOperation
}

func GetOperationLocator() *OperationLocator {
	if operationLocator == nil {
		operationLocator = &OperationLocator{}
	}

	return operationLocator
}

func (locator OperationLocator) GetUpdateTagOperation() *UpdateTagOperation {
	if locator.updateTagOperation == nil {
		locator.updateTagOperation = NewUpdateTagOperation(
			requests_validators.NewValidatorFactory().MakeUpdateTagRequestValidator(),
			requests_validators.NewValidatorFactory().MakeUpdateTagPathParamsValidator(),
			services.GetServiceLocator().GetTagService(),
			responseFactories.GetResponseFactoryLocator().GetUpdateTagResponseFactory(),
		)
	}

	return locator.updateTagOperation
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

func (locator OperationLocator) GetDeleteTagOperation() *DeleteTagOperation {
	if locator.deleteTagOperation == nil {
		locator.deleteTagOperation = NewDeleteTagOperation(
			requests_validators.NewValidatorFactory().MakeDeleteTagPathParamsValidator(),
			services.GetServiceLocator().GetTagService(),
		)
	}

	return locator.deleteTagOperation
}

func (locator OperationLocator) GetGetByIdTagOperation() *GetByIdTagOperation {
	if locator.GetByIdTagOperation == nil {
		locator.GetByIdTagOperation = NewGetByIdTagOperation(
			requests_validators.NewValidatorFactory().MakeGetByIdPathParamsValidator(),
			services.GetServiceLocator().GetTagService(),
			responseFactories.GetResponseFactoryLocator().GetByIdResponseFactory(),
		)
	}

	return locator.GetByIdTagOperation
}
