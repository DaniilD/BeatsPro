package responseFactories

var responseFactoryLocator *ResponseFactoryLocator

type ResponseFactoryLocator struct {
	createTagResponseFactory     *CreateTagResponseFactory
	updateTagResponseFactory     *UpdateTagResponseFactory
	getByIdResponseFactory       *GetByIdResponseFactory
	consumerErrorResponseFactory *ConsumerErrorResponseFactory
}

func GetResponseFactoryLocator() *ResponseFactoryLocator {
	if responseFactoryLocator == nil {
		responseFactoryLocator = &ResponseFactoryLocator{}
	}

	return responseFactoryLocator
}

func (factoryLocator *ResponseFactoryLocator) GetCreateTagResponseFactory() *CreateTagResponseFactory {
	if factoryLocator.createTagResponseFactory == nil {
		factoryLocator.createTagResponseFactory = NewCreateTagResponseFactory()
	}

	return factoryLocator.createTagResponseFactory
}

func (factoryLocator *ResponseFactoryLocator) GetUpdateTagResponseFactory() *UpdateTagResponseFactory {
	if factoryLocator.updateTagResponseFactory == nil {
		factoryLocator.updateTagResponseFactory = NewUpdateTagResponseFactory()
	}

	return factoryLocator.updateTagResponseFactory
}

func (factoryLocator *ResponseFactoryLocator) GetByIdResponseFactory() *GetByIdResponseFactory {
	if factoryLocator.getByIdResponseFactory == nil {
		factoryLocator.getByIdResponseFactory = NewGetByIdResponseFactory()
	}

	return factoryLocator.getByIdResponseFactory
}

func (factoryLocator *ResponseFactoryLocator) GetConsumerErrorResponseFactory() *ConsumerErrorResponseFactory {
	if factoryLocator.consumerErrorResponseFactory == nil {
		factoryLocator.consumerErrorResponseFactory = new(ConsumerErrorResponseFactory)
	}

	return factoryLocator.consumerErrorResponseFactory
}
