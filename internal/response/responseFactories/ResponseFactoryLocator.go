package responseFactories

var responseFactoryLocator *ResponseFactoryLocator

type ResponseFactoryLocator struct {
	createTagResponseFactory *CreateTagResponseFactory
	updateTagResponseFactory *UpdateTagResponseFactory
}

func GetResponseFactoryLocator() *ResponseFactoryLocator {
	if responseFactoryLocator == nil {
		responseFactoryLocator = &ResponseFactoryLocator{}
	}

	return responseFactoryLocator
}

func (responseFactoryLocator *ResponseFactoryLocator) GetCreateTagResponseFactory() *CreateTagResponseFactory {
	if responseFactoryLocator.createTagResponseFactory == nil {
		responseFactoryLocator.createTagResponseFactory = NewCreateTagResponseFactory()
	}

	return responseFactoryLocator.createTagResponseFactory
}

func (responseFactoryLocator *ResponseFactoryLocator) GetUpdateTagResponseFactory() *UpdateTagResponseFactory {
	if responseFactoryLocator.updateTagResponseFactory == nil {
		responseFactoryLocator.updateTagResponseFactory = NewUpdateTagResponseFactory()
	}

	return responseFactoryLocator.updateTagResponseFactory
}
