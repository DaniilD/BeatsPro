package services

import (
	"BeatsPro/internal/db"
	"BeatsPro/internal/repositories"
)

var serviceLocator *ServiceLocator

type ServiceLocator struct {
	tagService *TagService
}

func GetServiceLocator() *ServiceLocator {
	if serviceLocator == nil {
		serviceLocator = &ServiceLocator{}
	}

	return serviceLocator
}

func (serviceLocator *ServiceLocator) GetTagService() *TagService {
	if serviceLocator.tagService == nil {
		serviceLocator.tagService = NewTagService(
			repositories.NewTagRepository(db.GetStoreLocator().GetStore()))
	}

	return serviceLocator.tagService
}
