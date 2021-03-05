package services

import (
	"BeatsPro/internal/db"
	"BeatsPro/internal/repositories"
)

var serviceLocator *ServiceLocator

type ServiceLocator struct {
	tagService   *TagService
	trackService *TrackService
	userService  *UserService
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
			repositories.NewTagRepository(
				db.GetStoreLocator().GetStore()))
	}

	return serviceLocator.tagService
}

func (serviceLocator *ServiceLocator) GetTrackService() *TrackService {
	if serviceLocator.trackService == nil {
		serviceLocator.trackService = NewTrackService(
			repositories.NewTrackRepository(
				db.GetStoreLocator().GetStore()))
	}

	return serviceLocator.trackService
}

func (serviceLocator *ServiceLocator) GetUserService() *UserService {
	if serviceLocator.userService == nil {
		serviceLocator.userService = NewUserService(
			repositories.NewUserRepository(db.GetStoreLocator().GetStore()))
	}

	return serviceLocator.userService
}
