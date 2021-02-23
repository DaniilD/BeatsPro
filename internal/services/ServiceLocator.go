package services

var serviceLocator *ServiceLocator

type ServiceLocator struct {

}

func GetServiceLocator() *ServiceLocator {
	if serviceLocator == nil {
		serviceLocator = &ServiceLocator{}
	}

	return serviceLocator
}


