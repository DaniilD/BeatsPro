package responseFactories

var responseFactoryLocator *ResponseFactoryLocator

type ResponseFactoryLocator struct {
	createTagResponseFactory *CreateTagResponseFactory
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
